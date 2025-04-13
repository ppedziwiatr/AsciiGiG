package monster

import (
	"math/rand"
)

func GetAberrationTemplates() []MonsterTemplate {
	return []MonsterTemplate{
		{
			Name:   "Crawler",
			Type:   MonsterTypeAberation,
			Symbol: 'c',
			BaseAttributes: Attributes{
				Strength:     12,
				Agility:      10,
				Intelligence: 3,
				Charisma:     2,
			},
			ExpValue: 30,
			Level:    4,
		},
		{
			Name:   "Tentacled Horror",
			Type:   MonsterTypeAberation,
			Symbol: 'T',
			BaseAttributes: Attributes{
				Strength:     14,
				Agility:      8,
				Intelligence: 10,
				Charisma:     4,
			},
			ExpValue: 40,
			Level:    5,
		},
		{
			Name:   "Mind Flayer",
			Type:   MonsterTypeAberation,
			Symbol: 'M',
			BaseAttributes: Attributes{
				Strength:     10,
				Agility:      10,
				Intelligence: 16,
				Charisma:     14,
			},
			ExpValue: 60,
			Level:    6,
		},
		{
			Name:   "Beholder",
			Type:   MonsterTypeAberation,
			Symbol: 'B',
			BaseAttributes: Attributes{
				Strength:     12,
				Agility:      8,
				Intelligence: 18,
				Charisma:     12,
			},
			ExpValue: 80,
			Level:    7,
		},
		{
			Name:   "Elder Thing",
			Type:   MonsterTypeAberation,
			Symbol: 'E',
			BaseAttributes: Attributes{
				Strength:     16,
				Agility:      12,
				Intelligence: 20,
				Charisma:     16,
			},
			ExpValue: 150,
			Level:    8,
		},
	}
}

func CreateAberrationAbilities(monster *Monster) {

	monster.AbilitySet = append(monster.AbilitySet, GetBasicAttack(monster.Level))

	switch monster.Name {
	case "Crawler":
		toxinAbility := Ability{
			Name:        "Toxic Spit",
			Description: "Spits a corrosive toxin",
			ManaCost:    4,
			CoolDown:    3,
			Type:        AbilityTypeAttack,
			Power:       3 + monster.Level,
			Effect: &Effect{
				Name:      "Corroded",
				Duration:  3,
				Magnitude: 2,
				Type:      EffectDamageOverTime,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, toxinAbility)
	case "Tentacled Horror":
		grabAbility := Ability{
			Name:        "Tentacle Grab",
			Description: "Grabs and constricts the target",
			ManaCost:    5,
			CoolDown:    4,
			Type:        AbilityTypeAttack,
			Power:       4 + monster.Level,
			Effect: &Effect{
				Name:      "Constricted",
				Duration:  2,
				Magnitude: 3,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, grabAbility)
	case "Mind Flayer":
		mindBlastAbility := Ability{
			Name:        "Mind Blast",
			Description: "Psychic attack that damages the mind",
			ManaCost:    8,
			CoolDown:    5,
			Type:        AbilityTypeAttack,
			Power:       6 + monster.Level,
			Effect: &Effect{
				Name:      "Disoriented",
				Duration:  2,
				Magnitude: 4,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, mindBlastAbility)
	case "Beholder":
		eyeRayAbility := Ability{
			Name:        "Eye Ray",
			Description: "Shoots a beam of energy from one of its eyes",
			ManaCost:    10,
			CoolDown:    4,
			Type:        AbilityTypeAttack,
			Power:       7 + monster.Level,
		}
		monster.AbilitySet = append(monster.AbilitySet, eyeRayAbility)
	case "Elder Thing":
		cosmicHorrorAbility := Ability{
			Name:        "Cosmic Horror",
			Description: "Reveals glimpses of cosmic horror that damage sanity",
			ManaCost:    12,
			CoolDown:    6,
			Type:        AbilityTypeAttack,
			Power:       10 + monster.Level,
			Effect: &Effect{
				Name:      "Madness",
				Duration:  3,
				Magnitude: 5,
				Type:      EffectDebuffAttribute,
			},
		}
		monster.AbilitySet = append(monster.AbilitySet, cosmicHorrorAbility)
	}

	if monster.Level >= 6 {
		if monster.Name == "Mind Flayer" || monster.Name == "Elder Thing" {
			psionicScreamAbility := Ability{
				Name:        "Psionic Scream",
				Description: "Powerful psychic attack that affects all nearby enemies",
				ManaCost:    15,
				CoolDown:    8,
				Type:        AbilityTypeAttack,
				Power:       9 + monster.Level,
			}
			monster.AbilitySet = append(monster.AbilitySet, psionicScreamAbility)
		} else if monster.Name == "Beholder" {
			disintegrateAbility := Ability{
				Name:        "Disintegrate",
				Description: "Powerful ray that nearly disintegrates the target",
				ManaCost:    14,
				CoolDown:    8,
				Type:        AbilityTypeAttack,
				Power:       12 + monster.Level,
			}
			monster.AbilitySet = append(monster.AbilitySet, disintegrateAbility)
		}
	}
}

func GenerateAberration(level int) *Monster {
	templates := GetAberrationTemplates()

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
	baseHealth := 35 + int(levelScaling*12) // Aberrations are very resilient
	baseMana := 45 + int(levelScaling*10)   // And have lots of mana for abilities

	monster := &Monster{
		Name:   template.Name,
		Symbol: template.Symbol,
		Attributes: Attributes{
			Strength:     template.BaseAttributes.Strength + int(levelScaling),
			Agility:      template.BaseAttributes.Agility + int(levelScaling),
			Intelligence: template.BaseAttributes.Intelligence + int(levelScaling*1.5), // Aberrations are very intelligent
			Charisma:     template.BaseAttributes.Charisma - int(levelScaling/2),       // But often repulsive
		},
		Health:     baseHealth,
		MaxHealth:  baseHealth,
		Mana:       baseMana,
		MaxMana:    baseMana,
		Level:      level,
		Gold:       rand.Intn(20*level) + level*3, // Aberrations often have treasure
		Type:       MonsterTypeAberation,
		Behavior:   MonsterBehavior(rand.Intn(2) + 2), // Aberrations are always smart or aggressive
		DropRate:   0.6 + float64(level)*0.05,         // Best drop rate
		AbilitySet: []Ability{},
		ExpValue:   template.ExpValue * level,
	}

	CreateAberrationAbilities(monster)

	return monster
}
