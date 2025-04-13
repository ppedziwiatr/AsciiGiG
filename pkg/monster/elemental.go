package monster

import (
	"math/rand"
)

func GetElementalTemplates() []MonsterTemplate {
	return []MonsterTemplate{
		{
			Name:   "Fire Sprite",
			Type:   MonsterTypeElement,
			Symbol: 'f',
			BaseAttributes: Attributes{
				Strength:     6,
				Agility:      12,
				Intelligence: 10,
				Charisma:     8,
			},
			ExpValue: 20,
			Level:    3,
		},
		{
			Name:   "Earth Golem",
			Type:   MonsterTypeElement,
			Symbol: 'E',
			BaseAttributes: Attributes{
				Strength:     15,
				Agility:      5,
				Intelligence: 6,
				Charisma:     4,
			},
			ExpValue: 30,
			Level:    4,
		},
		{
			Name:   "Water Elemental",
			Type:   MonsterTypeElement,
			Symbol: 'W',
			BaseAttributes: Attributes{
				Strength:     10,
				Agility:      10,
				Intelligence: 12,
				Charisma:     8,
			},
			ExpValue: 35,
			Level:    5,
		},
		{
			Name:   "Air Spirit",
			Type:   MonsterTypeElement,
			Symbol: 'A',
			BaseAttributes: Attributes{
				Strength:     8,
				Agility:      16,
				Intelligence: 12,
				Charisma:     10,
			},
			ExpValue: 35,
			Level:    5,
		},
		{
			Name:   "Magma Giant",
			Type:   MonsterTypeElement,
			Symbol: 'M',
			BaseAttributes: Attributes{
				Strength:     18,
				Agility:      8,
				Intelligence: 10,
				Charisma:     6,
			},
			ExpValue: 80,
			Level:    7,
		},
	}
}

func CreateElementalAbilities(monster *Monster) {

	monster.AbilitySet = append(monster.AbilitySet, GetBasicAttack(monster.Level))

	switch monster.Name {
	case "Fire Sprite":
		fireballAbility := Ability{
			Name:        "Fireball",
			Description: "Hurls a ball of fire",
			ManaCost:    5,
			CoolDown:    3,
			Type:        AbilityTypeAttack,
			Power:       4 + monster.Level,
			Effect: &Effect{
				Name:      "Burning",
				Duration:  3,
				Magnitude: 2,
				Type:      EffectDamageOverTime,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, fireballAbility)
	case "Earth Golem":
		quakeAbility := Ability{
			Name:        "Tremor",
			Description: "Shakes the ground to unbalance enemies",
			ManaCost:    6,
			CoolDown:    4,
			Type:        AbilityTypeAttack,
			Power:       3 + monster.Level,
			Effect: &Effect{
				Name:      "Unbalanced",
				Duration:  2,
				Magnitude: 3,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, quakeAbility)
	case "Water Elemental":
		freezeAbility := Ability{
			Name:        "Freeze",
			Description: "Encases the target in ice",
			ManaCost:    7,
			CoolDown:    5,
			Type:        AbilityTypeDebuff,
			Power:       2,
			Effect: &Effect{
				Name:      "Frozen",
				Duration:  2,
				Magnitude: 4,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, freezeAbility)
	case "Air Spirit":
		lightningAbility := Ability{
			Name:        "Lightning Strike",
			Description: "Calls down a bolt of lightning",
			ManaCost:    8,
			CoolDown:    6,
			Type:        AbilityTypeAttack,
			Power:       6 + monster.Level,
		}
		monster.AbilitySet = append(monster.AbilitySet, lightningAbility)
	case "Magma Giant":
		eruptionAbility := Ability{
			Name:        "Eruption",
			Description: "Causes magma to erupt from the ground",
			ManaCost:    10,
			CoolDown:    5,
			Type:        AbilityTypeAttack,
			Power:       8 + monster.Level,
			Effect: &Effect{
				Name:      "Melting",
				Duration:  3,
				Magnitude: 3,
				Type:      EffectDamageOverTime,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, eruptionAbility)
	}

	if monster.Level >= 5 {
		surgeAbility := Ability{
			Name:        "Elemental Surge",
			Description: "Powerful elemental attack",
			ManaCost:    6,
			CoolDown:    5,
			Type:        AbilityTypeAttack,
			Power:       5 + monster.Level,
		}
		monster.AbilitySet = append(monster.AbilitySet, surgeAbility)
	}
}

func GenerateElemental(level int) *Monster {
	templates := GetElementalTemplates()

	var validTemplates []MonsterTemplate
	for _, template := range templates {
		if template.Level <= level+2 && template.Level >= level-2 {
			validTemplates = append(validTemplates, template)
		}
	}

	if len(validTemplates) == 0 {
		validTemplates = templates
	}

	template := validTemplates[rand.Intn(len(validTemplates))]

	levelScaling := float64(level) * 0.5
	baseHealth := 30 + int(levelScaling*10) // Elementals are hardy
	baseMana := 40 + int(levelScaling*8)    // And have lots of mana for abilities

	monster := &Monster{
		Name:   template.Name,
		Symbol: template.Symbol,
		Attributes: Attributes{
			Strength:     template.BaseAttributes.Strength + int(levelScaling),
			Agility:      template.BaseAttributes.Agility + int(levelScaling),
			Intelligence: template.BaseAttributes.Intelligence + int(levelScaling*1.5), // Elementals are smart
			Charisma:     template.BaseAttributes.Charisma + int(levelScaling/2),
		},
		Health:     baseHealth,
		MaxHealth:  baseHealth,
		Mana:       baseMana,
		MaxMana:    baseMana,
		Level:      level,
		Gold:       rand.Intn(5*level) + level*2,
		Type:       MonsterTypeElement,
		Behavior:   MonsterBehavior(rand.Intn(2) + 2), // Elementals are smarter
		DropRate:   0.5 + float64(level)*0.05,         // Better drop rate
		AbilitySet: []Ability{},
		ExpValue:   template.ExpValue * level,
	}

	CreateElementalAbilities(monster)

	return monster
}
