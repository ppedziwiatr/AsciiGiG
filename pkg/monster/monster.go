package monster

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/common"
	"math/rand"

	"github.com/ppedziwiatr/ascii-gig/pkg/dungeon"
)

type Monster struct {
	Name       string
	Symbol     rune
	Position   dungeon.Position
	Attributes Attributes
	Health     int
	MaxHealth  int
	Mana       int
	MaxMana    int
	Level      int
	Gold       int
	Type       MonsterType
	Behavior   MonsterBehavior
	DropRate   float64
	AbilitySet []Ability
	ExpValue   int
}

func GenerateMonster(level int, monsterType MonsterType) *Monster {
	switch monsterType {
	case MonsterTypeUndead:
		return GenerateUndead(level)
	case MonsterTypeDemon:
		return GenerateDemon(level)
	case MonsterTypeAnimal:
		return GenerateAnimal(level)
	case MonsterTypeElement:
		return GenerateElemental(level)
	case MonsterTypeAberation:
		return GenerateAberration(level)
	default:

		return GenerateUndead(level)
	}
}

func MoveMonster(monster *Monster, playerX, playerY int, dungeon *dungeon.Dungeon, monsters []*Monster) {

	dx := playerX - monster.Position.X
	dy := playerY - monster.Position.Y

	moveX, moveY := 0, 0
	if dx > 0 {
		moveX = 1
	} else if dx < 0 {
		moveX = -1
	}

	if dy > 0 {
		moveY = 1
	} else if dy < 0 {
		moveY = -1
	}

	tryMonsterMove(monster, moveX, moveY, dungeon, monsters)
}

func MoveMonsterAway(monster *Monster, playerX, playerY int, dungeon *dungeon.Dungeon, monsters []*Monster) {

	dx := monster.Position.X - playerX
	dy := monster.Position.Y - playerY

	moveX, moveY := 0, 0
	if dx > 0 {
		moveX = 1
	} else if dx < 0 {
		moveX = -1
	}

	if dy > 0 {
		moveY = 1
	} else if dy < 0 {
		moveY = -1
	}

	tryMonsterMove(monster, moveX, moveY, dungeon, monsters)
}

func MoveMonsterRandomly(monster *Monster, dungeon *dungeon.Dungeon, monsters []*Monster) {
	dirs := [][2]int{
		{0, -1}, // Up
		{0, 1},  // Down
		{-1, 0}, // Left
		{1, 0},  // Right
	}

	dir := dirs[rand.Intn(len(dirs))]
	tryMonsterMove(monster, dir[0], dir[1], dungeon, monsters)
}

func tryMonsterMove(
	monster *Monster, dx, dy int,
	dungeon *dungeon.Dungeon,
	monsters []*Monster) {

	newX := monster.Position.X + dx
	newY := monster.Position.Y + dy

	entities := make([]common.Entity, len(monsters))
	for i, m := range monsters {
		if m != monster { // Skip self
			entities[i] = m
		}
	}

	if dungeon.IsPositionEmpty(newX, newY, entities) {
		monster.Position.X = newX
		monster.Position.Y = newY
	}
}

func IsAdjacent(monster *Monster, playerX, playerY int) bool {
	dx := abs(playerX - monster.Position.X)
	dy := abs(playerY - monster.Position.Y)
	return dx <= 1 && dy <= 1
}

func CalculateDistanceSquared(monster *Monster, playerX, playerY int) int {
	dx := playerX - monster.Position.X
	dy := playerY - monster.Position.Y
	return dx*dx + dy*dy
}

func ReduceCooldowns(monster *Monster) {
	for i := range monster.AbilitySet {
		if monster.AbilitySet[i].CurrentCD > 0 {
			monster.AbilitySet[i].CurrentCD--
		}
	}
}

func (m *Monster) ChooseAbility() Ability {

	if len(m.AbilitySet) == 0 {
		return GetBasicAttack(m.Level)
	}

	availableAbilities := make([]int, 0)
	for i, ability := range m.AbilitySet {
		if ability.CurrentCD <= 0 && m.Mana >= ability.ManaCost {
			availableAbilities = append(availableAbilities, i)
		}
	}

	if len(availableAbilities) == 0 {
		return m.AbilitySet[0] // Basic attack should always be first and have no cost
	}

	abilityIndex := availableAbilities[rand.Intn(len(availableAbilities))]

	selectedAbility := m.AbilitySet[abilityIndex]
	m.AbilitySet[abilityIndex].CurrentCD = selectedAbility.CoolDown
	m.Mana -= selectedAbility.ManaCost

	return selectedAbility
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
