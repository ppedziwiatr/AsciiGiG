package ui

import (
	"fmt"
	"strings"
)

type Tile rune

const (
	TileFloor      Tile = '.'
	TileWall       Tile = '#'
	TileStairs     Tile = '>'
	TileDoor       Tile = '+'
	TilePlayer     Tile = '@'
	TileMonster    Tile = 'M'
	TileItem       Tile = 'i'
	TileEmpty      Tile = ' '
	TileUnexplored Tile = '?'
)

func RenderGameScreen(
	dungeonWidth, dungeonHeight int,
	dungeonLevel, maxLevels int,
	tiles [][]Tile,
	visible [][]bool,
	visited [][]bool,
	playerX, playerY int,
	playerHealth, playerMaxHealth int,
	playerMana, playerMaxMana int,
	playerExp, playerLevelUpExp int,
	playerGold int,
	playerStrength, playerAgility, playerCharisma, playerIntelligence int,
	monsters []MonsterInfo,
	items []ItemInfo,
	messages []string,
	mode string,
	currentTarget *MonsterInfo,
	playerAbilities []AbilityDisplayInfo,
) {

	fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen

	themeColors := GetDungeonThemeColors(dungeonLevel)

	for y := 0; y < dungeonHeight; y++ {
		for x := 0; x < dungeonWidth; x++ {
			if visible[y][x] {

				if playerX == x && playerY == y {
					fmt.Print(ColorBrightYellow + TextBold + "@" + ColorReset)
					continue
				}

				monsterFound := false
				for _, monster := range monsters {
					if monster.X == x && monster.Y == y {

						monsterColor := GetMonsterColor(monster.Type)
						fmt.Print(monsterColor + TextBold + string(monster.Symbol) + ColorReset)
						monsterFound = true
						break
					}
				}
				if monsterFound {
					continue
				}

				itemFound := false
				for _, item := range items {
					if item.X == x && item.Y == y {

						itemColor := GetItemColor(item.Type)
						fmt.Print(itemColor + string(item.Symbol) + ColorReset)
						itemFound = true
						break
					}
				}
				if itemFound {
					continue
				}

				switch tiles[y][x] {
				case TileWall:
					fmt.Print(themeColors.Wall + string(tiles[y][x]) + ColorReset)
				case TileStairs:
					fmt.Print(ColorBrightCyan + string(tiles[y][x]) + ColorReset)
				case TileDoor:
					fmt.Print(ColorYellow + string(tiles[y][x]) + ColorReset)
				default:
					fmt.Print(themeColors.Floor + string(tiles[y][x]) + ColorReset)
				}
			} else if visited[y][x] {

				fmt.Print(ColorBlack + TextBold + string(TileUnexplored) + ColorReset)
			} else {

				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println(ColorCyan + strings.Repeat("-", dungeonWidth) + ColorReset)
	fmt.Printf("%sLevel: %d/%d | HP: %s%d/%d%s | Mana: %s%d/%d%s | XP: %d/%d | Gold: %s%d%s\n",
		ColorWhite, dungeonLevel+1, maxLevels,
		ColorRed, playerHealth, playerMaxHealth, ColorWhite,
		ColorBlue, playerMana, playerMaxMana, ColorWhite,
		playerExp, playerLevelUpExp,
		ColorYellow, playerGold, ColorReset)

	fmt.Printf("%sSTR: %s%d%s | AGI: %s%d%s | CHA: %s%d%s | INT: %s%d%s\n",
		ColorWhite,
		ColorBrightRed, playerStrength, ColorWhite,
		ColorBrightGreen, playerAgility, ColorWhite,
		ColorBrightYellow, playerCharisma, ColorWhite,
		ColorBrightBlue, playerIntelligence, ColorReset)

	fmt.Println(ColorCyan + strings.Repeat("-", dungeonWidth) + ColorReset)
	msgCount := len(messages)
	for i := max(0, msgCount-3); i < msgCount; i++ {
		fmt.Println(ColorBrightWhite + messages[i] + ColorReset)
	}

	fmt.Println(ColorCyan + strings.Repeat("-", dungeonWidth) + ColorReset)
	fmt.Println(ColorGreen + "Move: [↑][↓][←][→] | [g]et | [i]nventory | [c]haracter | [q]uit" + ColorReset)

	if mode == "combat" && currentTarget != nil {
		fmt.Println(ColorRed + strings.Repeat("=", dungeonWidth) + ColorReset)
		fmt.Printf("%sCombat with %s%s%s (HP: %s%d/%d%s)\n",
			ColorWhite,
			ColorBrightRed+TextBold, currentTarget.Name, ColorWhite,
			ColorRed, currentTarget.Health, currentTarget.MaxHealth, ColorReset)

		fmt.Println(ColorYellow + "Abilities:" + ColorReset)
		for i, ability := range playerAbilities {
			ready := ""
			if ability.Cooldown > 0 {
				ready = fmt.Sprintf(" %s(CD: %d)%s", ColorRed, ability.Cooldown, ColorReset)
			}

			abilityColor := ColorWhite
			switch ability.Type {
			case "attack":
				abilityColor = ColorRed
			case "heal":
				abilityColor = ColorGreen
			case "buff":
				abilityColor = ColorCyan
			case "debuff":
				abilityColor = ColorMagenta
			}

			fmt.Printf("%s[%d]%s %s%s%s - %s (Power: %d, Mana: %s%d%s)%s\n",
				ColorYellow, i+1, ColorReset,
				abilityColor+TextBold, ability.Name, ColorReset,
				ability.Description, ability.Power,
				ColorBlue, ability.ManaCost, ColorReset,
				ready)
		}

		fmt.Printf("%s[r]%s Run away\n", ColorYellow, ColorReset)
	}
}

type MonsterInfo struct {
	X, Y      int
	Symbol    rune
	Type      string
	Name      string
	Health    int
	MaxHealth int
}

type ItemInfo struct {
	X, Y   int
	Symbol rune
	Type   string
	Name   string
}

type AbilityDisplayInfo struct {
	AbilityInfo
	Cooldown int
}
