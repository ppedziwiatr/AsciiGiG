package character

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/common"
	"github.com/ppedziwiatr/ascii-gig/pkg/dungeon"
	"github.com/ppedziwiatr/ascii-gig/pkg/item"
)

type Player struct {
	Name       string
	Symbol     rune
	Position   dungeon.Position
	Attributes common.Attributes
	Health     int
	MaxHealth  int
	Mana       int
	MaxMana    int
	Level      int
	Experience int
	Inventory  []*item.Item
	Gold       int
	Equipment  map[item.EquipmentSlot]*item.Item
	Class      PlayerClass
	Abilities  []Ability
	LevelUpExp int
}

const (
	MaxInventorySize = 20
)

func NewPlayer(name string, class PlayerClass) *Player {
	classInfo := GetClassInfo(class)

	player := &Player{
		Name:       name,
		Symbol:     '@',
		Position:   dungeon.Position{X: 0, Y: 0},
		Attributes: classInfo.BaseStats,
		Health:     classInfo.StartingHP,
		MaxHealth:  classInfo.StartingHP,
		Mana:       classInfo.StartingMP,
		MaxMana:    classInfo.StartingMP,
		Level:      1,
		Experience: 0,
		Inventory:  make([]*item.Item, 0, MaxInventorySize),
		Gold:       10,
		Equipment:  make(map[item.EquipmentSlot]*item.Item),
		Class:      class,
		Abilities:  []Ability{GetBasicAttack(class)},
		LevelUpExp: 100,
	}

	return player
}

func (p *Player) GetPosition() (int, int) {
	return p.Position.X, p.Position.Y
}

func (p *Player) SetPosition(x, y int) {
	p.Position.X = x
	p.Position.Y = y
}

func (p *Player) GetSymbol() rune {
	return p.Symbol
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GainExperience(amount int) bool {
	p.Experience += amount
	if p.Experience >= p.LevelUpExp {
		p.LevelUp()
		return true
	}
	return false
}

func (p *Player) LevelUp() {
	p.Level++
	p.Experience -= p.LevelUpExp
	p.LevelUpExp = p.Level * 100 // Increase exp needed for next level

	levelUpStats := GetLevelUpStats(p.Class)

	p.Attributes.Strength += levelUpStats.Strength
	p.Attributes.Agility += levelUpStats.Agility
	p.Attributes.Intelligence += levelUpStats.Intelligence
	p.Attributes.Charisma += levelUpStats.Charisma

	healthInc, manaInc := GetHealthAndManaIncrease(p.Class, p.Attributes)
	p.MaxHealth += healthInc
	p.Health = p.MaxHealth // Fully heal on level up
	p.MaxMana += manaInc
	p.Mana = p.MaxMana // Fully restore mana on level up

	if p.Level == 3 || p.Level == 6 || p.Level == 9 {
		newAbility := GetNewAbilityForLevel(p.Class, p.Level)
		p.Abilities = append(p.Abilities, newAbility)
	}
}

func (p *Player) PickUpItem(i *item.Item) bool {
	if len(p.Inventory) >= MaxInventorySize {
		return false
	}

	p.Inventory = append(p.Inventory, i)
	return true
}

func (p *Player) EquipItem(index int) bool {
	if index < 0 || index >= len(p.Inventory) {
		return false
	}

	itemToEquip := p.Inventory[index]

	totalAttrs := p.GetTotalAttributes()
	if !CanUseItem(totalAttrs, common.Attributes{
		Strength:     itemToEquip.RequiredAttributes.Strength,
		Agility:      itemToEquip.RequiredAttributes.Agility,
		Intelligence: itemToEquip.RequiredAttributes.Intelligence,
		Charisma:     itemToEquip.RequiredAttributes.Charisma,
	}) {
		return false
	}

	if existingItem, ok := p.Equipment[itemToEquip.Slot]; ok {
		p.Inventory = append(p.Inventory, existingItem)
	}

	p.Equipment[itemToEquip.Slot] = itemToEquip

	p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)

	return true
}

func (p *Player) UnequipItem(slot item.EquipmentSlot) bool {
	if len(p.Inventory) >= MaxInventorySize {
		return false // Inventory full
	}

	if playerItem, ok := p.Equipment[slot]; ok {
		p.Inventory = append(p.Inventory, playerItem)
		delete(p.Equipment, slot)
		return true
	}

	return false
}

func (p *Player) UseItem(index int) bool {
	if index < 0 || index >= len(p.Inventory) {
		return false
	}

	itemToUse := p.Inventory[index]

	if itemToUse.Type != item.ItemTypeConsumable {
		return false
	}

	for _, effect := range itemToUse.Effects {
		switch effect.Type {
		case common.EffectHealOverTime:
			p.Health += effect.Magnitude
			if p.Health > p.MaxHealth {
				p.Health = p.MaxHealth
			}

		}
	}

	p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)

	return true
}

func GetEquipmentSlotName(slot item.EquipmentSlot) string {
	switch slot {
	case item.SlotHead:
		return "Head"
	case item.SlotBody:
		return "Body"
	case item.SlotHands:
		return "Hands"
	case item.SlotWeapon:
		return "Weapon"
	case item.SlotOffhand:
		return "Offhand"
	case item.SlotFeet:
		return "Feet"
	case item.SlotAccessory:
		return "Accessory"
	default:
		return string(slot)
	}
}
