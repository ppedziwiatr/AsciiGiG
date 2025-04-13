package item

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/common"
	"github.com/ppedziwiatr/ascii-gig/pkg/dungeon"
)

type Item struct {
	Name               string
	Symbol             rune
	Position           dungeon.Position
	Type               ItemType
	Slot               EquipmentSlot
	Value              int               // Gold value
	BonusAttributes    common.Attributes // Bonus attributes
	Effects            []Effect
	RequiredAttributes common.Attributes // Required attributes to use
}

type ItemType string

const (
	ItemTypeWeapon     ItemType = "weapon"
	ItemTypeArmor      ItemType = "armor"
	ItemTypeConsumable ItemType = "consumable"
	ItemTypeResource   ItemType = "resource"
	ItemTypeSpecial    ItemType = "special"
)

type EquipmentSlot string

const (
	SlotHead      EquipmentSlot = "head"
	SlotBody      EquipmentSlot = "body"
	SlotHands     EquipmentSlot = "hands"
	SlotWeapon    EquipmentSlot = "weapon"
	SlotOffhand   EquipmentSlot = "offhand"
	SlotFeet      EquipmentSlot = "feet"
	SlotAccessory EquipmentSlot = "accessory"
)

type Effect struct {
	Name      string
	Duration  int
	Magnitude int
	Type      common.EffectType
}

type ItemTemplate struct {
	Name           string
	Type           ItemType
	BaseValue      int
	Rarity         float64
	Prefixes       []string
	Suffixes       []string
	BaseAttributes common.Attributes
}

func (i *Item) GetPosition() (int, int) {
	return i.Position.X, i.Position.Y
}

func (i *Item) SetPosition(x, y int) {
	i.Position.X = x
	i.Position.Y = y
}

func (i *Item) GetSymbol() rune {
	return i.Symbol
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) IsEquippable() bool {
	return i.Type == ItemTypeWeapon || i.Type == ItemTypeArmor
}

func (i *Item) IsConsumable() bool {
	return i.Type == ItemTypeConsumable
}
