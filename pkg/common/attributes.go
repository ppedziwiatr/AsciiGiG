package common

type Attributes struct {
	Agility      int // Affects turn order, dodge chance, and ranged weapons
	Strength     int // Affects HP, melee damage, and heavy weapons
	Charisma     int // Affects persuasion and item prices
	Intelligence int // Affects magic power, mana, and spell variety
}

func NewAttributes(str, agi, cha, intel int) *Attributes {
	return &Attributes{
		Strength:     str,
		Agility:      agi,
		Charisma:     cha,
		Intelligence: intel,
	}
}
