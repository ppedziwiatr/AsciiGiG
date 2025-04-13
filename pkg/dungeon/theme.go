package dungeon

type Theme struct {
	Name            string
	Description     string
	WallSymbol      rune
	FloorSymbol     rune
	MonsterTypes    []string
	ColorScheme     string
	SpecialFeatures []string
}

func GetThemes() []Theme {
	return []Theme{
		{
			Name:         "Church Catacombs",
			Description:  "Dark and damp stone passages beneath the church.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"undead", "animal"},
			ColorScheme:  "blue",
		},
		{
			Name:         "Underground Passages",
			Description:  "Natural caves with occasional constructed walls.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"animal", "undead"},
			ColorScheme:  "blue",
		},
		{
			Name:         "Forgotten Tombs",
			Description:  "Ancient burial chambers with dusty sarcophagi.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"undead", "demon"},
			ColorScheme:  "purple",
		},
		{
			Name:         "Torture Chambers",
			Description:  "Blood-stained rooms with instruments of pain.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"demon", "undead"},
			ColorScheme:  "red",
		},
		{
			Name:         "Hellish Caves",
			Description:  "Caverns heated by infernal fires below.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"demon", "elemental"},
			ColorScheme:  "red",
		},
		{
			Name:         "Burning Hell",
			Description:  "Lakes of fire and brimstone surround you.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"demon", "elemental"},
			ColorScheme:  "red",
		},
		{
			Name:         "Realm of Hatred",
			Description:  "A twisted landscape of malice and spite.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"demon", "aberration"},
			ColorScheme:  "purple",
		},
		{
			Name:         "Diablo's Lair",
			Description:  "The final resting place of the Lord of Terror.",
			WallSymbol:   '#',
			FloorSymbol:  '.',
			MonsterTypes: []string{"demon", "aberration"},
			ColorScheme:  "red",
		},
	}
}

func GetThemeForLevel(level int) Theme {
	themes := GetThemes()
	if level > 0 && level <= len(themes) {
		return themes[level-1]
	}
	return themes[0] // Default to first theme
}
