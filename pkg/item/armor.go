package item

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/common"
	"math/rand"
	"strings"
)

type ArmorType string

const (
	ArmorTypeLight  ArmorType = "light"
	ArmorTypeMedium ArmorType = "medium"
	ArmorTypeHeavy  ArmorType = "heavy"
)

type ArmorTemplate struct {
	Name        string
	Symbol      rune
	Slot        EquipmentSlot
	BaseValue   int
	BaseDefense int
	ReqStr      int
	ArmorType   ArmorType
}

func GetArmorTemplates() []ArmorTemplate {
	return []ArmorTemplate{
		{"Cloth Robe", '[', SlotBody, 8, 2, 3, ArmorTypeLight},
		{"Leather Armor", '[', SlotBody, 15, 3, 6, ArmorTypeLight},
		{"Chain Mail", '[', SlotBody, 25, 5, 10, ArmorTypeMedium},
		{"Plate Armor", '[', SlotBody, 40, 8, 15, ArmorTypeHeavy},
		{"Cap", '^', SlotHead, 5, 1, 2, ArmorTypeLight},
		{"Leather Cap", '^', SlotHead, 8, 2, 3, ArmorTypeLight},
		{"Helmet", '^', SlotHead, 15, 3, 6, ArmorTypeMedium},
		{"Full Helm", '^', SlotHead, 25, 5, 10, ArmorTypeHeavy},
		{"Gloves", '}', SlotHands, 5, 1, 2, ArmorTypeLight},
		{"Bracers", '}', SlotHands, 12, 2, 5, ArmorTypeMedium},
		{"Gauntlets", '}', SlotHands, 20, 3, 8, ArmorTypeHeavy},
		{"Shoes", '{', SlotFeet, 6, 1, 2, ArmorTypeLight},
		{"Boots", '{', SlotFeet, 12, 2, 4, ArmorTypeMedium},
		{"Greaves", '{', SlotFeet, 18, 3, 7, ArmorTypeHeavy},
		{"Amulet", '"', SlotAccessory, 20, 0, 0, ArmorTypeLight},
		{"Ring", '=', SlotAccessory, 25, 0, 0, ArmorTypeLight},
	}
}

func GetArmorPrefixes() []string {
	return []string{"", "", "", "Sturdy", "Reinforced", "Ornate", "Masterwork", "Enchanted", "Ancient", "Legendary"}
}

func GetArmorSuffixes() []string {
	return []string{"", "", "", "of Protection", "of Warding", "of the Knight", "of the Wizard", "of Deflection", "of Resistance", "of Invulnerability"}
}

func GenerateArmor(level int) *Item {
	armorTypes := GetArmorTemplates()
	armorType := armorTypes[rand.Intn(len(armorTypes))]

	prefixes := GetArmorPrefixes()
	suffixes := GetArmorSuffixes()

	prefix := prefixes[rand.Intn(len(prefixes))]
	suffix := suffixes[rand.Intn(len(suffixes))]

	qualityMod := 1.0
	if prefix != "" {
		qualityMod += 0.2
	}
	if suffix != "" {
		qualityMod += 0.3
	}

	levelMod := float64(level) * 0.2

	name := armorType.Name
	if prefix != "" {
		name = prefix + " " + name
	}
	if suffix != "" {
		name = name + " " + suffix
	}

	defense := int(float64(armorType.BaseDefense) * (1.0 + levelMod) * qualityMod)
	value := int(float64(armorType.BaseValue) * (1.0 + levelMod) * qualityMod)

	strReq := int(float64(armorType.ReqStr) * (1.0 + levelMod*0.5))

	strBonus := 0
	agiBonus := 0
	chaBonus := 0
	intBonus := 0

	if strings.Contains(name, "Protection") || strings.Contains(name, "Knight") || strings.Contains(name, "Sturdy") {
		strBonus = int(1.0 + levelMod)
	}
	if strings.Contains(name, "Deflection") || armorType.ArmorType == ArmorTypeLight {
		agiBonus = int(1.0 + levelMod)
	}
	if strings.Contains(name, "Ornate") {
		chaBonus = int(1.0 + levelMod)
	}
	if strings.Contains(name, "Wizard") || strings.Contains(name, "Enchanted") {
		intBonus = int(1.0 + levelMod)
	}

	return &Item{
		Name:   name,
		Symbol: armorType.Symbol,
		Type:   ItemTypeArmor,
		Slot:   armorType.Slot,
		Value:  value,
		BonusAttributes: common.Attributes{
			Strength:     strBonus,
			Agility:      agiBonus,
			Charisma:     chaBonus,
			Intelligence: intBonus,
		},
		RequiredAttributes: common.Attributes{
			Strength: strReq,
		},
		Effects: []Effect{
			{
				Name:      "Defense",
				Magnitude: defense,
				Type:      common.EffectBuffAttribute,
			},
		},
	}
}
