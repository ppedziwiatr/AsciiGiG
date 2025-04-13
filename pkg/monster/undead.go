package monster

import (
	"math/rand"
)

func GetUndeadTemplates() []MonsterTemplate {
	return []MonsterTemplate{
		{
			Name:   "Skeleton",
			Type:   MonsterTypeUndead,
			Symbol: 'S',
			BaseAttributes: Attributes{
				Strength:     6,
				Agility:      8,
				Intelligence: 4,
				Charisma:     2,
			},
			ExpValue: 10,
			Level:    1,
		},
		{
			Name:   "Zombie",
			Type:   MonsterTypeUndead,
			Symbol: 'Z',
			BaseAttributes: Attributes{
				Strength:     10,
				Agility:      4,
				Intelligence: 2,
				Charisma:     2,
			},
			ExpValue: 15,
			Level:    2,
		},
		{
			Name:   "Ghost",
			Type:   MonsterTypeUndead,
			Symbol: 'G',
			BaseAttributes: Attributes{
				Strength:     4,
				Agility:      10,
				Intelligence: 8,
				Charisma:     6,
			},
			ExpValue: 20,
			Level:    3,
		},
		{
			Name:   "Wraith",
			Type:   MonsterTypeUndead,
			Symbol: 'W',
			BaseAttributes: Attributes{
				Strength:     8,
				Agility:      12,
				Intelligence: 10,
				Charisma:     8,
			},
			ExpValue: 30,
			Level:    4,
		},
		{
			Name:   "Lich",
			Type:   MonsterTypeUndead,
			Symbol: 'L',
			BaseAttributes: Attributes{
				Strength:     12,
				Agility:      10,
				Intelligence: 15,
				Charisma:     12,
			},
			ExpValue: 100,
			Level:    7,
		},
	}
}

func CreateUndeadAbilities(monster *Monster) {

	monster.AbilitySet = append(monster.AbilitySet, GetBasicAttack(monster.Level))

	if monster.Level >= 3 {
		lifedrainAbility := Ability{
			Name:        "Life Drain",
			Description: "Drains health from the target",
			ManaCost:    4,
			CoolDown:    4,
			Type:        AbilityTypeAttack,
			Power:       3 + monster.Level/2,
			Effect: &Effect{
				Name:      "Drained",
				Duration:  2,
				Magnitude: 1,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, lifedrainAbility)
	}

	if monster.Level >= 5 {
		if monster.Name == "Lich" {
			necroticBlastAbility := Ability{
				Name:        "Necrotic Blast",
				Description: "Powerful undead magic attack",
				ManaCost:    8,
				CoolDown:    5,
				Type:        AbilityTypeAttack,
				Power:       8 + monster.Level,
			}
			monster.AbilitySet = append(monster.AbilitySet, necroticBlastAbility)
		} else if monster.Name == "Wraith" {
			terrorAbility := Ability{
				Name:        "Terror",
				Description: "Terrifies the target, reducing their stats",
				ManaCost:    6,
				CoolDown:    6,
				Type:        AbilityTypeDebuff,
				Effect: &Effect{
					Name:      "Terrified",
					Duration:  3,
					Magnitude: 3,
					Type:      EffectDebuffAttribute,
				},
			}
			monster.AbilitySet = append(monster.AbilitySet, terrorAbility)
		}
	}
}

func GenerateUndead(level int) *Monster {
	templates := GetUndeadTemplates()

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
	baseHealth := 20 + int(levelScaling*10)
	baseMana := 10 + int(levelScaling*5)

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
		Gold:       rand.Intn(10*level) + level,
		Type:       MonsterTypeUndead,
		Behavior:   MonsterBehavior(rand.Intn(4)),
		DropRate:   0.3 + float64(level)*0.05,
		AbilitySet: []Ability{},
		ExpValue:   template.ExpValue * level,
	}

	CreateUndeadAbilities(monster)

	return monster
}
