package game

import (
	"fmt"
	"math/rand"

	"github.com/ppedziwiatr/ascii-gig/pkg/character"
	"github.com/ppedziwiatr/ascii-gig/pkg/dungeon"
	"github.com/ppedziwiatr/ascii-gig/pkg/input"
	"github.com/ppedziwiatr/ascii-gig/pkg/item"
	"github.com/ppedziwiatr/ascii-gig/pkg/monster"
	"github.com/ppedziwiatr/ascii-gig/pkg/ui"
)

const (
	MaxDungeonLevels = 8
)

func NewGame() *Game {
	player := character.NewPlayer("Hero", character.ClassWarrior)

	game := &Game{
		Player:       player,
		CurrentLevel: 0,
		Dungeons:     make([]*DungeonState, MaxDungeonLevels),
		GameOver:     false,
		Victory:      false,
		TurnCount:    0,
		Mode:         ModeExploring,
		Messages:     []string{"Welcome to the dungeon! Find the stairs to descend deeper."},
	}

	for i := 0; i < MaxDungeonLevels; i++ {
		dungeonLevel := dungeon.GenerateDungeon(60, 30, i+1)
		monsters := generateMonstersForLevel(dungeonLevel, i+1)
		items := generateItemsForLevel(dungeonLevel, i+1)

		game.Dungeons[i] = &DungeonState{
			Dungeon:  dungeonLevel,
			Monsters: monsters,
			Items:    items,
		}
	}

	startPos := dungeon.GetStartingPosition(game.Dungeons[0].Dungeon)
	player.SetPosition(startPos.X, startPos.Y)

	return game
}

func (g *Game) Run() {

	for !g.GameOver {

		g.calculateFOV()

		g.renderGame()

		g.handleInput()

		if g.Mode == ModeExploring {
			g.processMonsterTurns()
		}

		if g.Player.Health <= 0 {
			g.GameOver = true
			g.Victory = false
			g.AddMessage("You have died. Game over!")
		}

		if g.CurrentLevel == MaxDungeonLevels-1 {

			if len(g.Dungeons[g.CurrentLevel].Monsters) == 0 {
				g.GameOver = true
				g.Victory = true
				g.AddMessage("You have defeated Diablo and saved the world! Victory!")
			}
		}

		g.TurnCount++
	}

	g.showGameOverScreen()
}

func (g *Game) calculateFOV() {
	currentDungeon := g.Dungeons[g.CurrentLevel].Dungeon
	playerX, playerY := g.Player.GetPosition()

	currentDungeon.GetVisibleTiles(playerX, playerY)
}

func (g *Game) renderGame() {
	currentDungeon := g.Dungeons[g.CurrentLevel].Dungeon
	player := g.Player

	visibleMonsters := g.GetVisibleMonsters()
	visibleItems := g.GetVisibleItems()

	monsterInfos := make([]ui.MonsterInfo, len(visibleMonsters))
	for i, m := range visibleMonsters {
		monsterInfos[i] = ui.MonsterInfo{
			X:         m.Position.X,
			Y:         m.Position.Y,
			Symbol:    m.Symbol,
			Type:      string(m.Type),
			Name:      m.Name,
			Health:    m.Health,
			MaxHealth: m.MaxHealth,
		}
	}

	itemInfos := make([]ui.ItemInfo, len(visibleItems))
	for i, item := range visibleItems {
		itemInfos[i] = ui.ItemInfo{
			X:      item.Position.X,
			Y:      item.Position.Y,
			Symbol: item.Symbol,
			Type:   string(item.Type),
			Name:   item.Name,
		}
	}

	abilityInfos := make([]ui.AbilityDisplayInfo, len(player.Abilities))
	for i, ability := range player.Abilities {
		abilityInfos[i] = ui.AbilityDisplayInfo{
			AbilityInfo: ui.AbilityInfo{
				Name:        ability.Name,
				Description: ability.Description,
				Type:        string(ability.Type),
				Power:       ability.Power,
				ManaCost:    ability.ManaCost,
			},
			Cooldown: ability.CurrentCD,
		}
	}

	var currentTarget *ui.MonsterInfo
	if g.Mode == ModeCombat && g.CurrentTarget != nil {
		currentTarget = &ui.MonsterInfo{
			X:         g.CurrentTarget.Position.X,
			Y:         g.CurrentTarget.Position.Y,
			Symbol:    g.CurrentTarget.Symbol,
			Type:      string(g.CurrentTarget.Type),
			Name:      g.CurrentTarget.Name,
			Health:    g.CurrentTarget.Health,
			MaxHealth: g.CurrentTarget.MaxHealth,
		}
	}

	ui.RenderGameScreen(
		currentDungeon.Width, currentDungeon.Height,
		g.CurrentLevel, MaxDungeonLevels,
		currentDungeon.Tiles,
		currentDungeon.Visible,
		currentDungeon.Visited,
		player.Position.X, player.Position.Y,
		player.Health, player.MaxHealth,
		player.Mana, player.MaxMana,
		player.Experience, player.LevelUpExp,
		player.Gold,
		player.Attributes.Strength, player.Attributes.Agility, player.Attributes.Charisma, player.Attributes.Intelligence,
		monsterInfos,
		itemInfos,
		g.Messages,
		g.GetGameModeString(),
		currentTarget,
		abilityInfos,
	)
}

func (g *Game) handleInput() {

	mode := g.GetGameModeString()

	action, value := input.HandlePlayerInput(mode)

	switch action {
	case "move":
		direction := value
		g.movePlayer(direction)
	case "pickup":
		g.pickupItem()
	case "inventory":
		g.showInventory()
	case "character":
		g.showCharacter()
	case "ability":
		abilityIndex := value
		g.UseAbility(abilityIndex)
	case "run":
		g.AttemptToFlee()
	case "close_inventory":
		g.Mode = ModeExploring
	case "quit":
		g.GameOver = true
	case "none":

	}
}

func (g *Game) movePlayer(direction int) {

	if g.Mode != ModeExploring {
		return
	}

	player := g.Player
	currentDungeon := g.Dungeons[g.CurrentLevel].Dungeon

	var dx, dy int
	switch direction {
	case 0: // North
		dy = -1
	case 1: // South
		dy = 1
	case 2: // West
		dx = -1
	case 3: // East
		dx = 1
	}

	newX := player.Position.X + dx
	newY := player.Position.Y + dy

	if !currentDungeon.IsWalkable(newX, newY) {
		return
	}

	monster := g.GetMonsterAt(newX, newY)
	if monster != nil {
		g.StartCombat(monster)
		return
	}

	if currentDungeon.Tiles[newY][newX] == dungeon.TileStairs {
		if g.CurrentLevel < MaxDungeonLevels-1 {
			g.descendToNextLevel()
			return
		}
	}

	player.SetPosition(newX, newY)

	item := g.GetItemAt(newX, newY)
	if item != nil {
		g.AddMessage(fmt.Sprintf("You see %s. Press [g] to pick it up.", item.Name))
	}
}

func (g *Game) pickupItem() {
	player := g.Player

	itemAtPlayer := g.GetItemAt(player.Position.X, player.Position.Y)
	if itemAtPlayer != nil {
		if player.PickUpItem(itemAtPlayer) {
			g.AddMessage(fmt.Sprintf("You picked up %s.", itemAtPlayer.Name))
			g.RemoveItem(itemAtPlayer)
		} else {
			g.AddMessage("Your inventory is full!")
		}
	} else {
		g.AddMessage("There's nothing here to pick up.")
	}
}

func (g *Game) showInventory() {
	g.Mode = ModeInventory

}

func (g *Game) showCharacter() {
	player := g.Player

	equipment := make(map[string]string)
	equipmentTypes := make(map[string]string)
	for slot, item := range player.Equipment {
		equipment[character.GetEquipmentSlotName(slot)] = item.Name
		equipmentTypes[character.GetEquipmentSlotName(slot)] = string(item.Type)
	}

	abilities := make([]ui.AbilityInfo, len(player.Abilities))
	for i, ability := range player.Abilities {
		abilities[i] = ui.AbilityInfo{
			Name:        ability.Name,
			Description: ability.Description,
			Type:        string(ability.Type),
			Power:       ability.Power,
			ManaCost:    ability.ManaCost,
		}
	}

	strBonus := character.GetAttributeBonus(player, "strength")
	agiBonus := character.GetAttributeBonus(player, "agility")
	chaBonus := character.GetAttributeBonus(player, "charisma")
	intBonus := character.GetAttributeBonus(player, "intelligence")

	ui.DisplayCharacterSheet(
		player.Name,
		player.Level,
		string(player.Class),
		player.Health, player.MaxHealth,
		player.Mana, player.MaxMana,
		player.Experience, player.LevelUpExp,
		player.Gold,
		player.Attributes.Strength, player.Attributes.Strength-strBonus, strBonus,
		player.Attributes.Agility, player.Attributes.Agility-agiBonus, agiBonus,
		player.Attributes.Charisma, player.Attributes.Charisma-chaBonus, chaBonus,
		player.Attributes.Intelligence, player.Attributes.Intelligence-intBonus, intBonus,
		equipment,
		equipmentTypes,
		abilities,
	)
}

func (g *Game) processMonsterTurns() {
	playerX, playerY := g.Player.GetPosition()
	currentDungeon := g.Dungeons[g.CurrentLevel]

	monsterArr := make([]*monster.Monster, len(currentDungeon.Monsters))
	copy(monsterArr, currentDungeon.Monsters)

	for _, m := range currentDungeon.Monsters {

		monsterX, monsterY := m.Position.X, m.Position.Y
		if !currentDungeon.Dungeon.Visible[monsterY][monsterX] {
			continue
		}

		monster.ReduceCooldowns(m)

		distSquared := monster.CalculateDistanceSquared(m, playerX, playerY)

		if monster.IsAdjacent(m, playerX, playerY) {

			g.StartCombat(m)
			return
		}

		switch m.Behavior {
		case monster.BehaviorAggressive:

			if distSquared <= 100 { // Distance of 10 tiles
				monster.MoveMonster(m, playerX, playerY, currentDungeon.Dungeon, monsterArr)
			}
		case monster.BehaviorPatrolling:

			if distSquared <= 25 { // Distance of 5 tiles
				if rand.Float64() < 0.7 {
					monster.MoveMonster(m, playerX, playerY, currentDungeon.Dungeon, monsterArr)
				} else {
					monster.MoveMonsterRandomly(m, currentDungeon.Dungeon, monsterArr)
				}
			} else {
				monster.MoveMonsterRandomly(m, currentDungeon.Dungeon, monsterArr)
			}
		case monster.BehaviorCowardly:

			if distSquared <= 25 { // Distance of 5 tiles
				monster.MoveMonsterAway(m, playerX, playerY, currentDungeon.Dungeon, monsterArr)
			} else {
				monster.MoveMonsterRandomly(m, currentDungeon.Dungeon, monsterArr)
			}
		case monster.BehaviorSmart:

			if distSquared <= 100 { // Distance of 10 tiles

				if m.Health > m.MaxHealth/2 {
					monster.MoveMonster(m, playerX, playerY, currentDungeon.Dungeon, monsterArr)
				} else {

					monster.MoveMonsterAway(m, playerX, playerY, currentDungeon.Dungeon, monsterArr)
				}
			}
		}
	}
}

func (g *Game) descendToNextLevel() {
	g.CurrentLevel++

	theme := g.Dungeons[g.CurrentLevel].Dungeon.Theme
	g.AddMessage(fmt.Sprintf("You descend deeper into %s...", theme.Name))

	startPos := dungeon.GetStartingPosition(g.Dungeons[g.CurrentLevel].Dungeon)
	g.Player.SetPosition(startPos.X, startPos.Y)
}

func (g *Game) showGameOverScreen() {
	ui.DisplayGameOverScreen(
		g.Victory,
		g.Player.Level,
		g.CurrentLevel,
		MaxDungeonLevels,
		g.Player.Gold,
	)
}

func generateMonstersForLevel(d *dungeon.Dungeon, level int) []*monster.Monster {
	monsters := make([]*monster.Monster, 0)

	numMonsters := rand.Intn(5) + 5*level // More monsters in deeper levels
	if numMonsters > 20 {
		numMonsters = 20 // Cap at 20 monsters per level
	}

	for i := 0; i < numMonsters; i++ {

		if len(d.Rooms) <= 2 {
			continue
		}
		roomIdx := rand.Intn(len(d.Rooms)-2) + 1
		room := d.Rooms[roomIdx]

		monsterX := room.X + rand.Intn(room.Width-2) + 1
		monsterY := room.Y + rand.Intn(room.Height-2) + 1

		if d.Tiles[monsterY][monsterX] == dungeon.TileFloor {

			monsterType := d.Theme.MonsterTypes[rand.Intn(len(d.Theme.MonsterTypes))]

			m := monster.GenerateMonster(level, monster.MonsterType(monsterType))
			m.Position.X = monsterX
			m.Position.Y = monsterY

			monsters = append(monsters, m)
		}
	}

	return monsters
}

func generateItemsForLevel(d *dungeon.Dungeon, level int) []*item.Item {
	items := make([]*item.Item, 0)

	numItems := rand.Intn(3) + 2

	for i := 0; i < numItems; i++ {

		roomIdx := rand.Intn(len(d.Rooms))
		room := d.Rooms[roomIdx]

		itemX := room.X + rand.Intn(room.Width-2) + 1
		itemY := room.Y + rand.Intn(room.Height-2) + 1

		if d.Tiles[itemY][itemX] == dungeon.TileFloor {

			newItem := item.GenerateItem(level)
			newItem.Position.X = itemX
			newItem.Position.Y = itemY

			items = append(items, newItem)
		}
	}

	return items
}
