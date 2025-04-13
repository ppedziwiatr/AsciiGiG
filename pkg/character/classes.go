package character

import "github.com/ppedziwiatr/ascii-gig/pkg/common"

type PlayerClass string

const (
	ClassWarrior PlayerClass = "warrior"
	ClassRanger  PlayerClass = "ranger"
	ClassMage    PlayerClass = "mage"
	ClassRogue   PlayerClass = "rogue"
)

type ClassInfo struct {
	Name        string
	Description string
	BaseStats   common.Attributes
	StartingHP  int
	StartingMP  int
}

func GetClassInfo(class PlayerClass) ClassInfo {
	switch class {
	case ClassWarrior:
		return ClassInfo{
			Name:        "Warrior",
			Description: "Masters of physical combat, warriors excel at dealing and absorbing damage.",
			BaseStats: common.Attributes{
				Strength:     12,
				Agility:      8,
				Charisma:     6,
				Intelligence: 4,
			},
			StartingHP: 120,
			StartingMP: 30,
		}
	case ClassRanger:
		return ClassInfo{
			Name:        "Ranger",
			Description: "Skilled hunters who excel with ranged weapons and can track enemies.",
			BaseStats: common.Attributes{
				Strength:     8,
				Agility:      12,
				Charisma:     7,
				Intelligence: 7,
			},
			StartingHP: 90,
			StartingMP: 50,
		}
	case ClassMage:
		return ClassInfo{
			Name:        "Mage",
			Description: "Wielders of arcane energy who sacrifice physical prowess for magical power.",
			BaseStats: common.Attributes{
				Strength:     4,
				Agility:      6,
				Charisma:     8,
				Intelligence: 14,
			},
			StartingHP: 70,
			StartingMP: 120,
		}
	case ClassRogue:
		return ClassInfo{
			Name:        "Rogue",
			Description: "Nimble tricksters who excel at stealth, traps, and exploiting enemy weaknesses.",
			BaseStats: common.Attributes{
				Strength:     6,
				Agility:      14,
				Charisma:     10,
				Intelligence: 8,
			},
			StartingHP: 80,
			StartingMP: 40,
		}
	default:
		return GetClassInfo(ClassWarrior)
	}
}

func GetLevelUpStats(class PlayerClass) common.Attributes {
	switch class {
	case ClassWarrior:
		return common.Attributes{
			Strength:     2,
			Agility:      1,
			Charisma:     1,
			Intelligence: 0,
		}
	case ClassRanger:
		return common.Attributes{
			Strength:     1,
			Agility:      2,
			Charisma:     1,
			Intelligence: 0,
		}
	case ClassMage:
		return common.Attributes{
			Strength:     0,
			Agility:      1,
			Charisma:     1,
			Intelligence: 2,
		}
	case ClassRogue:
		return common.Attributes{
			Strength:     0,
			Agility:      2,
			Charisma:     1,
			Intelligence: 1,
		}
	default:
		return common.Attributes{
			Strength:     1,
			Agility:      1,
			Charisma:     1,
			Intelligence: 1,
		}
	}
}

func GetHealthAndManaIncrease(class PlayerClass, attributes common.Attributes) (int, int) {
	healthInc := 10
	manaInc := 5

	switch class {
	case ClassWarrior:
		healthInc += attributes.Strength / 2
		manaInc += attributes.Intelligence / 4
	case ClassRanger:
		healthInc += attributes.Strength / 3
		manaInc += attributes.Intelligence / 3
	case ClassMage:
		healthInc += attributes.Strength / 4
		manaInc += attributes.Intelligence / 2
	case ClassRogue:
		healthInc += attributes.Strength / 3
		manaInc += attributes.Intelligence / 3
	}

	return healthInc, manaInc
}
