package character

import "github.com/ppedziwiatr/ascii-gig/pkg/common"

type AttributeBonus struct {
	StrengthBonus     int
	AgilityBonus      int
	CharismaBonus     int
	IntelligenceBonus int
}

func GetAttributeBonus(player *Player, attributeName string) int {
	bonus := 0

	for _, item := range player.Equipment {
		switch attributeName {
		case "strength":
			bonus += item.BonusAttributes.Strength
		case "agility":
			bonus += item.BonusAttributes.Agility
		case "charisma":
			bonus += item.BonusAttributes.Charisma
		case "intelligence":
			bonus += item.BonusAttributes.Intelligence
		}
	}

	return bonus
}

func AddAttributes(a, b common.Attributes) common.Attributes {
	return common.Attributes{
		Strength:     a.Strength + b.Strength,
		Agility:      a.Agility + b.Agility,
		Charisma:     a.Charisma + b.Charisma,
		Intelligence: a.Intelligence + b.Intelligence,
	}
}

func CanUseItem(attrs common.Attributes, itemRequirements common.Attributes) bool {
	return attrs.Strength >= itemRequirements.Strength &&
		attrs.Agility >= itemRequirements.Agility &&
		attrs.Intelligence >= itemRequirements.Intelligence &&
		attrs.Charisma >= itemRequirements.Charisma
}

func (p *Player) GetTotalAttributes() common.Attributes {
	total := p.Attributes

	for _, item := range p.Equipment {
		total = AddAttributes(total, item.BonusAttributes)
	}

	return total
}
