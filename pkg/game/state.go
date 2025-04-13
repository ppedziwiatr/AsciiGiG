package game

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/character"
	"github.com/ppedziwiatr/ascii-gig/pkg/dungeon"
	"github.com/ppedziwiatr/ascii-gig/pkg/item"
	"github.com/ppedziwiatr/ascii-gig/pkg/monster"
)

type Mode int

const (
	ModeExploring Mode = iota
	ModeCombat
	ModeInventory
	ModeDialog
	ModeGameOver
)

func ModeToString(mode Mode) string {
	switch mode {
	case ModeExploring:
		return "exploring"
	case ModeCombat:
		return "combat"
	case ModeInventory:
		return "inventory"
	case ModeDialog:
		return "dialog"
	case ModeGameOver:
		return "gameover"
	default:
		return "unknown"
	}
}

type Game struct {
	Player        *character.Player
	CurrentLevel  int
	Dungeons      []*DungeonState
	GameOver      bool
	Victory       bool
	TurnCount     int
	Mode          Mode
	CurrentTarget *monster.Monster
	Messages      []string
}

type DungeonState struct {
	Dungeon  *dungeon.Dungeon
	Monsters []*monster.Monster
	Items    []*item.Item
}

func (g *Game) AddMessage(msg string) {
	g.Messages = append(g.Messages, msg)

	if len(g.Messages) > 100 {
		g.Messages = g.Messages[len(g.Messages)-100:]
	}
}

func (g *Game) GetCurrentDungeon() *DungeonState {
	return g.Dungeons[g.CurrentLevel]
}

func (g *Game) GetGameModeString() string {
	return ModeToString(g.Mode)
}

func (g *Game) IsPlayerTurn() bool {

	if g.Mode == ModeExploring {
		return true
	}

	if g.Mode == ModeCombat && g.CurrentTarget != nil {
		playerAgility := g.Player.Attributes.Agility
		monsterAgility := g.CurrentTarget.Attributes.Agility

		if playerAgility > monsterAgility {
			return true
		}

		if monsterAgility > playerAgility {

			return g.TurnCount%2 == 0
		}

		return g.TurnCount%2 == 0
	}

	return true
}

func (g *Game) GetVisibleMonsters() []*monster.Monster {
	currentDungeon := g.GetCurrentDungeon()
	var visible []*monster.Monster

	for _, m := range currentDungeon.Monsters {
		x, y := m.GetPosition()
		if currentDungeon.Dungeon.Visible[y][x] {
			visible = append(visible, m)
		}
	}

	return visible
}

func (g *Game) GetVisibleItems() []*item.Item {
	currentDungeon := g.GetCurrentDungeon()
	var visible []*item.Item

	for _, i := range currentDungeon.Items {
		x, y := i.GetPosition()
		if currentDungeon.Dungeon.Visible[y][x] {
			visible = append(visible, i)
		}
	}

	return visible
}

func (g *Game) IsPositionOccupied(x, y int) bool {
	currentDungeon := g.GetCurrentDungeon()

	for _, m := range currentDungeon.Monsters {
		mx, my := m.GetPosition()
		if mx == x && my == y {
			return true
		}
	}

	return false
}

func (g *Game) GetMonsterAt(x, y int) *monster.Monster {
	currentDungeon := g.GetCurrentDungeon()

	for _, m := range currentDungeon.Monsters {
		mx, my := m.GetPosition()
		if mx == x && my == y {
			return m
		}
	}

	return nil
}

func (g *Game) GetItemAt(x, y int) *item.Item {
	currentDungeon := g.GetCurrentDungeon()

	for _, i := range currentDungeon.Items {
		ix, iy := i.GetPosition()
		if ix == x && iy == y {
			return i
		}
	}

	return nil
}

func (g *Game) RemoveMonster(monster *monster.Monster) {
	currentDungeon := g.GetCurrentDungeon()

	for i, m := range currentDungeon.Monsters {
		if m == monster {
			currentDungeon.Monsters = append(currentDungeon.Monsters[:i], currentDungeon.Monsters[i+1:]...)
			break
		}
	}
}

func (g *Game) RemoveItem(item *item.Item) {
	currentDungeon := g.GetCurrentDungeon()

	for i, it := range currentDungeon.Items {
		if it == item {
			currentDungeon.Items = append(currentDungeon.Items[:i], currentDungeon.Items[i+1:]...)
			break
		}
	}
}

func (g *Game) AddItem(item *item.Item) {
	currentDungeon := g.GetCurrentDungeon()
	currentDungeon.Items = append(currentDungeon.Items, item)
}
