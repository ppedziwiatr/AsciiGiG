package monster

import (
	"math/rand"
)

func GetDemonTemplates() []MonsterTemplate {
	return []MonsterTemplate{
		{
			Name:   "Imp",
			Type:   MonsterTypeDemon,
			Symbol: 'i',
			BaseAttributes: Attributes{
				Strength:     5,
				Agility:      10,
				Intelligence: 8,
				Charisma:     6,
			},
			ExpValue: 15,
			Level:    2,
		},
		{
			Name:   "Fallen One",
			Type:   MonsterTypeDemon,
			Symbol: 'f',
			BaseAttributes: Attributes{
				Strength:     8,
				Agility:      8,
				Intelligence: 6,
				Charisma:     5,
			},
			ExpValue: 20,
			Level:    3,
		},
		{
			Name:   "Hellhound",
			Type:   MonsterTypeDemon,
			Symbol: 'h',
			BaseAttributes: Attributes{
				Strength:     12,
				Agility:      14,
				Intelligence: 5,
				Charisma:     4,
			},
			ExpValue: 25,
			Level:    4,
		},
		{
			Name:   "Vortex",
			Type:   MonsterTypeDemon,
			Symbol: 'v',
			BaseAttributes: Attributes{
				Strength:     10,
				Agility:      12,
				Intelligence: 12,
				Charisma:     10,
			},
			ExpValue: 40,
			Level:    5,
		},
		{
			Name:   "Balrog",
			Type:   MonsterTypeDemon,
			Symbol: 'B',
			BaseAttributes: Attributes{
				Strength:     18,
				Agility:      12,
				Intelligence: 14,
				Charisma:     14,
			},
			ExpValue: 120,
			Level:    8,
		},
	}
}

func CreateDemonAbilities(monster *Monster) {

	monster.AbilitySet = append(monster.AbilitySet, GetBasicAttack(monster.Level))

	if monster.Level >= 3 {
		hellfireAbility := Ability{
			Name:        "Hellfire",
			Description: "Burning attack that deals damage over time",
			ManaCost:    5,
			CoolDown:    3,
			Type:        AbilityTypeAttack,
			Power:       2 + monster.Level,
			Effect: &Effect{
				Name:      "Burning",
				Duration:  3,
				Magnitude: 2,
				Type:      EffectDamageOverTime,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, hellfireAbility)
	}

	if monster.Level >= 5 {
		if monster.Name == "Balrog" {
			infernoAbility := Ability{
				Name:        "Inferno",
				Description: "Massive fire attack that engulfs the area",
				ManaCost:    10,
				CoolDown:    6,
				Type:        AbilityTypeAttack,
				Power:       10 + monster.Level,
			}
			monster.AbilitySet = append(monster.AbilitySet, infernoAbility)
		} else if monster.Name == "Vortex" {
			chaosNovaAbility := Ability{
				Name:        "Chaos Nova",
				Description: "Explosion of demonic energy",
				ManaCost:    8,
				CoolDown:    5,
				Type:        AbilityTypeAttack,
				Power:       8 + monster.Level/2,
				Effect: &Effect{
					Name:      "Chaos",
					Duration:  2,
					Magnitude: 2,
					Type:      EffectDebuffAttribute,
				},
			}
			monster.AbilitySet = append(monster.AbilitySet, chaosNovaAbility)
		}
	}
}

func GenerateDemon(level int) *Monster {
	templates := GetDemonTemplates()

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
	baseHealth := 25 + int(levelScaling*12) // Demons have more health
	baseMana := 15 + int(levelScaling*6)    // And more mana for abilities

	monster := &Monster{
		Name:   template.Name,
		Symbol: template.Symbol,
		Attributes: Attributes{
			Strength:     template.BaseAttributes.Strength + int(levelScaling),
			Agility:      template.BaseAttributes.Agility + int(levelScaling),
			Intelligence: template.BaseAttributes.Intelligence + int(levelScaling),
			Charisma:     template.BaseAttributes.Charisma + int(levelScaling/2),
		},
		Health:     baseHealth,
		MaxHealth:  baseHealth,
		Mana:       baseMana,
		MaxMana:    baseMana,
		Level:      level,
		Gold:       rand.Intn(15*level) + level, // Demons carry more gold
		Type:       MonsterTypeDemon,
		Behavior:   MonsterBehavior(rand.Intn(4)),
		DropRate:   0.4 + float64(level)*0.05, // Better drop rate
		AbilitySet: []Ability{},
		ExpValue:   template.ExpValue * level,
	}

	CreateDemonAbilities(monster)

	return monster
}
