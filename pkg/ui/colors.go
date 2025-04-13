package ui

const (
	ColorReset   = "\033[0m"
	ColorBlack   = "\033[30m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"

	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	TextBold      = "\033[1m"
	TextUnderline = "\033[4m"
)

type ThemeColors struct {
	Wall  string
	Floor string
}

func GetDungeonThemeColors(level int) ThemeColors {

	switch level {
	case 1: // Church Catacombs
		return ThemeColors{
			Wall:  ColorCyan,
			Floor: ColorBlue,
		}
	case 2: // Underground Passages
		return ThemeColors{
			Wall:  ColorWhite,
			Floor: ColorBlue,
		}
	case 3: // Forgotten Tombs
		return ThemeColors{
			Wall:  ColorWhite,
			Floor: ColorMagenta,
		}
	case 4: // Torture Chambers
		return ThemeColors{
			Wall:  ColorYellow,
			Floor: ColorRed,
		}
	case 5: // Hellish Caves
		return ThemeColors{
			Wall:  ColorYellow,
			Floor: ColorRed,
		}
	case 6: // Burning Hell
		return ThemeColors{
			Wall:  ColorBrightRed,
			Floor: ColorRed,
		}
	case 7: // Realm of Hatred
		return ThemeColors{
			Wall:  ColorBrightMagenta,
			Floor: ColorMagenta,
		}
	case 8: // Diablo's Lair
		return ThemeColors{
			Wall:  ColorBrightRed,
			Floor: ColorRed,
		}
	default:
		return ThemeColors{
			Wall:  ColorWhite,
			Floor: ColorBlue,
		}
	}
}

func GetMonsterColor(monsterType string) string {
	switch monsterType {
	case "undead":
		return ColorCyan
	case "demon":
		return ColorRed
	case "animal":
		return ColorYellow
	case "elemental":
		return ColorBlue
	case "aberration":
		return ColorMagenta
	default:
		return ColorWhite
	}
}

func GetItemColor(itemType string) string {
	switch itemType {
	case "weapon":
		return ColorBrightRed
	case "armor":
		return ColorBrightCyan
	case "consumable":
		return ColorBrightGreen
	case "resource":
		return ColorBrightYellow
	case "special":
		return ColorBrightMagenta
	default:
		return ColorBrightWhite
	}
}
