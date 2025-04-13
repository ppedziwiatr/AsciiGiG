package character

import "github.com/ppedziwiatr/ascii-gig/pkg/common"

type Ability struct {
	Name        string
	Description string
	ManaCost    int
	CoolDown    int
	CurrentCD   int
	Type        AbilityType
	Power       int
	Effect      *Effect
}

type AbilityType string

const (
	AbilityTypeAttack  AbilityType = "attack"
	AbilityTypeHeal    AbilityType = "heal"
	AbilityTypeBuff    AbilityType = "buff"
	AbilityTypeDebuff  AbilityType = "debuff"
	AbilityTypeSpecial AbilityType = "special"
)

type Effect struct {
	Name      string
	Duration  int
	Magnitude int
	Type      common.EffectType
}

func GetBasicAttack(class PlayerClass) Ability {
	switch class {
	case ClassWarrior:
		return Ability{
			Name:        "Slash",
			Description: "A basic attack with your weapon",
			ManaCost:    0,
			CoolDown:    0,
			Type:        AbilityTypeAttack,
			Power:       5,
		}
	case ClassRanger:
		return Ability{
			Name:        "Shoot",
			Description: "Attack an enemy from distance",
			ManaCost:    0,
			CoolDown:    0,
			Type:        AbilityTypeAttack,
			Power:       4,
		}
	case ClassMage:
		return Ability{
			Name:        "Magic Missile",
			Description: "A basic magical attack",
			ManaCost:    5,
			CoolDown:    0,
			Type:        AbilityTypeAttack,
			Power:       6,
		}
	case ClassRogue:
		return Ability{
			Name:        "Backstab",
			Description: "A sneaky attack dealing extra damage",
			ManaCost:    3,
			CoolDown:    2,
			Type:        AbilityTypeAttack,
			Power:       7,
		}
	default:
		return Ability{
			Name:        "Attack",
			Description: "A basic attack",
			ManaCost:    0,
			CoolDown:    0,
			Type:        AbilityTypeAttack,
			Power:       3,
		}
	}
}

func GetNewAbilityForLevel(class PlayerClass, level int) Ability {
	switch class {
	case ClassWarrior:
		switch level {
		case 3:
			return Ability{
				Name:        "Cleave",
				Description: "Attack that hits multiple enemies",
				ManaCost:    5,
				CoolDown:    3,
				Type:        AbilityTypeAttack,
				Power:       7,
			}
		case 6:
			return Ability{
				Name:        "Berserk",
				Description: "Increases strength temporarily",
				ManaCost:    8,
				CoolDown:    5,
				Type:        AbilityTypeBuff,
				Power:       0,
				Effect: &Effect{
					Name:      "Berserker Rage",
					Duration:  5,
					Magnitude: 5,
					Type:      common.EffectBuffAttribute,
				},
			}
		case 9:
			return Ability{
				Name:        "Whirlwind",
				Description: "Massive attack that hits all enemies",
				ManaCost:    12,
				CoolDown:    7,
				Type:        AbilityTypeAttack,
				Power:       15,
			}
		}
	case ClassRanger:
		switch level {
		case 3:
			return Ability{
				Name:        "Quick Shot",
				Description: "Fast attack that doesn't use a turn",
				ManaCost:    5,
				CoolDown:    4,
				Type:        AbilityTypeAttack,
				Power:       6,
			}
		case 6:
			return Ability{
				Name:        "Snare",
				Description: "Trap enemy, reducing their agility",
				ManaCost:    7,
				CoolDown:    5,
				Type:        AbilityTypeDebuff,
				Power:       0,
				Effect: &Effect{
					Name:      "Snared",
					Duration:  3,
					Magnitude: 5,
					Type:      common.EffectDebuffAttribute,
				},
			}
		case 9:
			return Ability{
				Name:        "Rain of Arrows",
				Description: "Powerful attack with high damage",
				ManaCost:    10,
				CoolDown:    6,
				Type:        AbilityTypeAttack,
				Power:       18,
			}
		}
	case ClassMage:
		switch level {
		case 3:
			return Ability{
				Name:        "Fireball",
				Description: "Explosive magic attack with area damage",
				ManaCost:    8,
				CoolDown:    3,
				Type:        AbilityTypeAttack,
				Power:       10,
			}
		case 6:
			return Ability{
				Name:        "Frost Nova",
				Description: "Freezes enemies, reducing their speed",
				ManaCost:    10,
				CoolDown:    5,
				Type:        AbilityTypeDebuff,
				Power:       5,
				Effect: &Effect{
					Name:      "Frozen",
					Duration:  2,
					Magnitude: 3,
					Type:      common.EffectDebuffAttribute,
				},
			}
		case 9:
			return Ability{
				Name:        "Arcane Explosion",
				Description: "Massive magic damage to all nearby enemies",
				ManaCost:    15,
				CoolDown:    8,
				Type:        AbilityTypeAttack,
				Power:       20,
			}
		}
	case ClassRogue:
		switch level {
		case 3:
			return Ability{
				Name:        "Poison Strike",
				Description: "Attack that poisons the enemy",
				ManaCost:    5,
				CoolDown:    4,
				Type:        AbilityTypeAttack,
				Power:       5,
				Effect: &Effect{
					Name:      "Poisoned",
					Duration:  4,
					Magnitude: 2,
					Type:      common.EffectDamageOverTime,
				},
			}
		case 6:
			return Ability{
				Name:        "Shadow Step",
				Description: "Teleport behind enemy for extra damage",
				ManaCost:    8,
				CoolDown:    5,
				Type:        AbilityTypeAttack,
				Power:       12,
			}
		case 9:
			return Ability{
				Name:        "Death Mark",
				Description: "Mark target for instant death after 3 turns",
				ManaCost:    15,
				CoolDown:    10,
				Type:        AbilityTypeDebuff,
				Power:       0,
				Effect: &Effect{
					Name:      "Death Mark",
					Duration:  3,
					Magnitude: 999,
					Type:      common.EffectDamageOverTime,
				},
			}
		}
	}

	return Ability{
		Name:        "Minor Heal",
		Description: "Heal a small amount of health",
		ManaCost:    5,
		CoolDown:    3,
		Type:        AbilityTypeHeal,
		Power:       10,
	}
}

func ReduceCooldowns(abilities []Ability) []Ability {
	for i := range abilities {
		if abilities[i].CurrentCD > 0 {
			abilities[i].CurrentCD--
		}
	}
	return abilities
}
