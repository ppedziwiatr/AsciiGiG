package item

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/common"
	"math/rand"
	"strings"
)

type WeaponType string

const (
	WeaponTypeSword  WeaponType = "sword"
	WeaponTypeAxe    WeaponType = "axe"
	WeaponTypeBow    WeaponType = "bow"
	WeaponTypeStaff  WeaponType = "staff"
	WeaponTypeDagger WeaponType = "dagger"
	WeaponTypeHammer WeaponType = "hammer"
)

type WeaponTemplate struct {
	Name       string
	Symbol     rune
	Slot       EquipmentSlot
	BaseValue  int
	BaseDamage int
	ReqStr     int
	ReqAgi     int
	ReqInt     int
}

func GetWeaponTemplates() []WeaponTemplate {
	return []WeaponTemplate{
		{"Dagger", '/', SlotWeapon, 10, 3, 5, 10, 5},
		{"Sword", '/', SlotWeapon, 20, 5, 10, 8, 5},
		{"Axe", '\\', SlotWeapon, 25, 6, 12, 6, 5},
		{"Mace", '!', SlotWeapon, 22, 5, 11, 5, 5},
		{"Staff", '|', SlotWeapon, 15, 4, 6, 6, 10},
		{"Bow", ')', SlotWeapon, 18, 4, 6, 12, 5},
		{"Wand", '~', SlotWeapon, 16, 3, 4, 5, 12},
	}
}

func GetWeaponPrefixes() []string {
	return []string{"", "", "", "Fine", "Sharp", "Sturdy", "Masterwork", "Enchanted", "Ancient", "Legendary"}
}

func GetWeaponSuffixes() []string {
	return []string{"", "", "", "of Power", "of Quickness", "of the Warrior", "of the Mage", "of Destruction", "of Slaying", "of Doom"}
}

func GenerateWeapon(level int) *Item {
	weaponTypes := GetWeaponTemplates()
	weaponType := weaponTypes[rand.Intn(len(weaponTypes))]

	prefixes := GetWeaponPrefixes()
	suffixes := GetWeaponSuffixes()

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

	name := weaponType.Name
	if prefix != "" {
		name = prefix + " " + name
	}
	if suffix != "" {
		name = name + " " + suffix
	}

	damage := int(float64(weaponType.BaseDamage) * (1.0 + levelMod) * qualityMod)
	value := int(float64(weaponType.BaseValue) * (1.0 + levelMod) * qualityMod)

	strReq := int(float64(weaponType.ReqStr) * (1.0 + levelMod*0.5))
	agiReq := int(float64(weaponType.ReqAgi) * (1.0 + levelMod*0.5))
	intReq := int(float64(weaponType.ReqInt) * (1.0 + levelMod*0.5))

	strBonus := 0
	agiBonus := 0
	chaBonus := 0
	intBonus := 0

	if strings.Contains(name, "Power") || strings.Contains(name, "Warrior") || strings.Contains(name, "Sturdy") {
		strBonus = int(1.0 + levelMod)
	}
	if strings.Contains(name, "Quickness") || strings.Contains(name, "Sharp") {
		agiBonus = int(1.0 + levelMod)
	}
	if strings.Contains(name, "Mage") || strings.Contains(name, "Enchanted") {
		intBonus = int(1.0 + levelMod)
	}

	return &Item{
		Name:   name,
		Symbol: weaponType.Symbol,
		Type:   ItemTypeWeapon,
		Slot:   weaponType.Slot,
		Value:  value,
		BonusAttributes: common.Attributes{
			Strength:     strBonus,
			Agility:      agiBonus,
			Charisma:     chaBonus,
			Intelligence: intBonus,
		},
		RequiredAttributes: common.Attributes{
			Strength:     strReq,
			Agility:      agiReq,
			Intelligence: intReq,
		},
		Effects: []Effect{
			{
				Name:      "Damage",
				Magnitude: damage,
				Type:      common.EffectDamageOverTime,
			},
		},
	}
}
