package monster

import (
	"math/rand"
)

func GetAnimalTemplates() []MonsterTemplate {
	return []MonsterTemplate{
		{
			Name:   "Rat",
			Type:   MonsterTypeAnimal,
			Symbol: 'r',
			BaseAttributes: Attributes{
				Strength:     3,
				Agility:      10,
				Intelligence: 2,
				Charisma:     1,
			},
			ExpValue: 5,
			Level:    1,
		},
		{
			Name:   "Spider",
			Type:   MonsterTypeAnimal,
			Symbol: 's',
			BaseAttributes: Attributes{
				Strength:     4,
				Agility:      12,
				Intelligence: 3,
				Charisma:     1,
			},
			ExpValue: 10,
			Level:    2,
		},
		{
			Name:   "Snake",
			Type:   MonsterTypeAnimal,
			Symbol: 'S',
			BaseAttributes: Attributes{
				Strength:     5,
				Agility:      14,
				Intelligence: 4,
				Charisma:     2,
			},
			ExpValue: 12,
			Level:    2,
		},
		{
			Name:   "Cave Bear",
			Type:   MonsterTypeAnimal,
			Symbol: 'b',
			BaseAttributes: Attributes{
				Strength:     14,
				Agility:      8,
				Intelligence: 5,
				Charisma:     4,
			},
			ExpValue: 20,
			Level:    3,
		},
		{
			Name:   "Giant Bat",
			Type:   MonsterTypeAnimal,
			Symbol: 'B',
			BaseAttributes: Attributes{
				Strength:     6,
				Agility:      16,
				Intelligence: 4,
				Charisma:     3,
			},
			ExpValue: 18,
			Level:    3,
		},
	}
}

func CreateAnimalAbilities(monster *Monster) {

	monster.AbilitySet = append(monster.AbilitySet, GetBasicAttack(monster.Level))

	switch monster.Name {
	case "Spider":
		webAbility := Ability{
			Name:        "Web",
			Description: "Ensnares the enemy, reducing their speed",
			ManaCost:    3,
			CoolDown:    4,
			Type:        AbilityTypeDebuff,
			Power:       1,
			Effect: &Effect{
				Name:      "Webbed",
				Duration:  2,
				Magnitude: 3,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, webAbility)
	case "Snake":
		venomAbility := Ability{
			Name:        "Venom Strike",
			Description: "Poisonous bite that causes damage over time",
			ManaCost:    4,
			CoolDown:    3,
			Type:        AbilityTypeAttack,
			Power:       2,
			Effect: &Effect{
				Name:      "Poisoned",
				Duration:  3,
				Magnitude: 2,
				Type:      EffectDamageOverTime,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, venomAbility)
	case "Cave Bear":
		roarAbility := Ability{
			Name:        "Ferocious Roar",
			Description: "Intimidating roar that weakens enemies",
			ManaCost:    5,
			CoolDown:    5,
			Type:        AbilityTypeDebuff,
			Power:       0,
			Effect: &Effect{
				Name:      "Intimidated",
				Duration:  2,
				Magnitude: 2,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, roarAbility)
	case "Giant Bat":
		sonarAbility := Ability{
			Name:        "Echolocation",
			Description: "Uses sound to locate weak spots",
			ManaCost:    3,
			CoolDown:    3,
			Type:        AbilityTypeBuff,
			Power:       0,
			Effect: &Effect{
				Name:      "Enhanced Senses",
				Duration:  3,
				Magnitude: 3,
				Type:      EffectBuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, sonarAbility)
	}
}

func GenerateAnimal(level int) *Monster {
	templates := GetAnimalTemplates()

	var validTemplates []MonsterTemplate
	for _, template := range templates {
		if template.Level <= level+1 && template.Level >= level-1 {
			validTemplates = append(validTemplates, template)
		}
	}

	if len(validTemplates) == 0 {
		validTemplates = templates
	}

	template := validTemplates[rand.Intn(len(validTemplates))]

	levelScaling := float64(level) * 0.5
	baseHealth := 15 + int(levelScaling*8) // Animals have less health
	baseMana := 5 + int(levelScaling*3)    // And less mana

	monster := &Monster{
		Name:   template.Name,
		Symbol: template.Symbol,
		Attributes: Attributes{
			Strength:     template.BaseAttributes.Strength + int(levelScaling),
			Agility:      template.BaseAttributes.Agility + int(levelScaling), // Animals are agile
			Intelligence: template.BaseAttributes.Intelligence + int(levelScaling/2),
			Charisma:     template.BaseAttributes.Charisma + int(levelScaling/3),
		},
		Health:     baseHealth,
		MaxHealth:  baseHealth,
		Mana:       baseMana,
		MaxMana:    baseMana,
		Level:      level,
		Gold:       rand.Intn(3*level) + 1, // Animals don't carry much gold
		Type:       MonsterTypeAnimal,
		Behavior:   MonsterBehavior(rand.Intn(3)), // Animals are never smart
		DropRate:   0.2 + float64(level)*0.03,     // Lower drop rate
		AbilitySet: []Ability{},
		ExpValue:   template.ExpValue * level,
	}

	CreateAnimalAbilities(monster)

	return monster
}
