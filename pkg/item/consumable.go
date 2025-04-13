package item

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/common"
	"math/rand"
)

type ConsumableTemplate struct {
	Name      string
	Symbol    rune
	BaseValue int
	Effect    Effect
}

func GetConsumableTemplates() []ConsumableTemplate {
	return []ConsumableTemplate{
		{"Health Potion", '!', 15, Effect{Name: "Healing", Duration: 1, Magnitude: 20, Type: common.EffectHealOverTime}},
		{"Mana Potion", '!', 15, Effect{Name: "Mana Restore", Duration: 1, Magnitude: 15, Type: common.EffectBuffAttribute}},
		{"Strength Elixir", '!', 25, Effect{Name: "Strength Boost", Duration: 10, Magnitude: 3, Type: common.EffectBuffAttribute}},
		{"Agility Elixir", '!', 25, Effect{Name: "Agility Boost", Duration: 10, Magnitude: 3, Type: common.EffectBuffAttribute}},
		{"Intelligence Elixir", '!', 25, Effect{Name: "Intelligence Boost", Duration: 10, Magnitude: 3, Type: common.EffectBuffAttribute}},
		{"Charisma Elixir", '!', 25, Effect{Name: "Charisma Boost", Duration: 10, Magnitude: 3, Type: common.EffectBuffAttribute}},
		{"Scroll of Fireball", '?', 30, Effect{Name: "Fireball", Duration: 1, Magnitude: 25, Type: common.EffectDamageOverTime}},
		{"Scroll of Protection", '?', 35, Effect{Name: "Protection", Duration: 20, Magnitude: 5, Type: common.EffectBuffAttribute}},
	}
}

func GetConsumablePrefixes() []string {
	return []string{"", "", "Minor", "Standard", "Greater", "Superior", "Exceptional", "Perfect"}
}

func GenerateConsumable(level int) *Item {
	itemTypes := GetConsumableTemplates()
	itemType := itemTypes[rand.Intn(len(itemTypes))]

	prefixes := GetConsumablePrefixes()

	prefix := prefixes[rand.Intn(len(prefixes))]

	qualityMod := 1.0
	switch prefix {
	case "Minor":
		qualityMod = 0.8
	case "Standard":
		qualityMod = 1.0
	case "Greater":
		qualityMod = 1.2
	case "Superior":
		qualityMod = 1.5
	case "Exceptional":
		qualityMod = 2.0
	case "Perfect":
		qualityMod = 2.5
	}

	levelMod := float64(level) * 0.2

	name := itemType.Name
	if prefix != "" && prefix != "Standard" {
		name = prefix + " " + name
	}

	value := int(float64(itemType.BaseValue) * (1.0 + levelMod) * qualityMod)
	magnitude := int(float64(itemType.Effect.Magnitude) * (1.0 + levelMod) * qualityMod)

	effect := itemType.Effect
	effect.Magnitude = magnitude

	return &Item{
		Name:    name,
		Symbol:  itemType.Symbol,
		Type:    ItemTypeConsumable,
		Value:   value,
		Effects: []Effect{effect},
	}
}

func GenerateItem(level int) *Item {
	itemTypes := []ItemType{
		ItemTypeWeapon,
		ItemTypeArmor,
		ItemTypeConsumable,
	}

	itemType := itemTypes[rand.Intn(len(itemTypes))]

	var item *Item

	switch itemType {
	case ItemTypeWeapon:
		item = GenerateWeapon(level)
	case ItemTypeArmor:
		item = GenerateArmor(level)
	case ItemTypeConsumable:
		item = GenerateConsumable(level)
	}

	return item
}
