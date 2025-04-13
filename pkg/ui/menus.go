package ui

import (
	"fmt"
	"strings"

	"github.com/ppedziwiatr/ascii-gig/pkg/input"
)

// DisplayTitleScreen displays the game's title screen
func DisplayTitleScreen() {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	// Small 2D ruby/gem logo that will appear inline with text
	rubyArt := []string{
		"    /\\    ",
		"   /  \\   ",
		"  / /\\ \\  ",
		" / /  \\ \\ ",
		" \\ \\  / / ",
		"  \\ \\/ /  ",
		"   \\  /   ",
		"    \\/    ",
	}

	// Title text in game over style
	titleTextLines := []string{
		"██████╗ ███████╗██████╗ ███████╗████████╗ ██████╗ ███╗   ██╗███████╗",
		"██╔══██╗██╔════╝██╔══██╗██╔════╝╚══██╔══╝██╔═══██╗████╗  ██║██╔════╝",
		"██████╔╝█████╗  ██║  ██║███████╗   ██║   ██║   ██║██╔██╗ ██║█████╗  ",
		"██╔══██╗██╔══╝  ██║  ██║╚════██║   ██║   ██║   ██║██║╚██╗██║██╔══╝  ",
		"██║  ██║███████╗██████╔╝███████║   ██║   ╚██████╔╝██║ ╚████║███████╗",
		"╚═╝  ╚═╝╚══════╝╚═════╝ ╚══════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═══╝╚══════╝",
		"",
		"          ██████╗  ██████╗  ██████╗ ██╗   ██╗███████╗                ",
		"          ██╔══██╗██╔═══██╗██╔════╝ ██║   ██║██╔════╝                ",
		"          ██████╔╝██║   ██║██║  ███╗██║   ██║█████╗                  ",
		"          ██╔══██╗██║   ██║██║   ██║██║   ██║██╔══╝                  ",
		"          ██║  ██║╚██████╔╝╚██████╔╝╚██████╔╝███████╗                ",
		"          ╚═╝  ╚═╝ ╚═════╝  ╚═════╝  ╚═════╝ ╚══════╝                ",
	}

	// Combined title with ruby
	fmt.Println()
	fmt.Println()

	// Calculate where to start the ruby to align it with the middle of REDSTONE
	rubyStartLine := 2 // Start ruby at line 2 of the title text (roughly aligns with middle of REDSTONE)

	for i := 0; i < len(titleTextLines); i++ {
		if i >= rubyStartLine && i < rubyStartLine+len(rubyArt) {
			// Print ruby and title on same line
			rubyIndex := i - rubyStartLine
			fmt.Print(ColorRed + rubyArt[rubyIndex] + ColorReset)
			fmt.Println(ColorRed + TextBold + titleTextLines[i] + ColorReset)
		} else {
			// Print just the title for lines without ruby
			fmt.Print("          ") // Space to align with ruby width
			fmt.Println(ColorRed + TextBold + titleTextLines[i] + ColorReset)
		}
	}

	// Display separator and tagline
	fmt.Println()
	fmt.Println(ColorYellow + strings.Repeat("=", 80) + ColorReset)
	fmt.Println(ColorBrightWhite + "\n            A roguelike adventure in the depths of hell\n" + ColorReset)
	fmt.Println(ColorYellow + strings.Repeat("=", 80) + ColorReset)
	fmt.Println(ColorBrightGreen + "\n\n      Press any key to begin your journey..." + ColorReset)

	// Wait for any key press
	input.ReadKey()
}

// DisplayGameOverScreen shows the game over or victory screen
func DisplayGameOverScreen(isVictory bool, playerLevel, currentLevel, maxLevels, gold int) {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	fmt.Println(ColorYellow + strings.Repeat("=", 50) + ColorReset)
	if isVictory {
		victoryText := `
             ██╗   ██╗██╗ ██████╗████████╗ ██████╗ ██████╗ ██╗   ██╗
             ██║   ██║██║██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗╚██╗ ██╔╝
             ██║   ██║██║██║        ██║   ██║   ██║██████╔╝ ╚████╔╝ 
             ╚██╗ ██╔╝██║██║        ██║   ██║   ██║██╔══██╗  ╚██╔╝  
              ╚████╔╝ ██║╚██████╗   ██║   ╚██████╔╝██║  ██║   ██║   
               ╚═══╝  ╚═╝ ╚═════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝   
`
		fmt.Println(ColorBrightGreen + TextBold + victoryText + ColorReset)
		fmt.Println(ColorBrightWhite + "       You have defeated Diablo" + ColorReset)
		fmt.Println(ColorBrightWhite + "       and saved the world from destruction!" + ColorReset)
	} else {
		gameOverText := `
              ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗ 
             ██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗
             ██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝
             ██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
             ╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║
              ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝
`
		fmt.Println(ColorBrightRed + TextBold + gameOverText + ColorReset)
	}
	fmt.Println(ColorBrightCyan + "\n\n       Statistics:" + ColorReset)
	fmt.Printf(ColorWhite+"       Level: "+ColorBrightYellow+"%d\n"+ColorReset, playerLevel)
	fmt.Printf(ColorWhite+"       Dungeon Depth: "+ColorBrightMagenta+"%d/%d\n"+ColorReset, currentLevel+1, maxLevels)
	fmt.Printf(ColorWhite+"       Gold Collected: "+ColorBrightYellow+"%d\n"+ColorReset, gold)

	fmt.Println(ColorBrightGreen + "\n\n       Press any key to exit..." + ColorReset)
	fmt.Println(ColorYellow + strings.Repeat("=", 50) + ColorReset)

	// Wait for any key press
	input.ReadKey()
}

// DisplayCharacterSheet shows the player's character information
func DisplayCharacterSheet(
	name string,
	level int,
	class string,
	health, maxHealth int,
	mana, maxMana int,
	experience, levelUpExp int,
	gold int,
	strength, baseStr, bonusStr int,
	agility, baseAgi, bonusAgi int,
	charisma, baseCha, bonusCha int,
	intelligence, baseInt, bonusInt int,
	equipment map[string]string,
	equipmentTypes map[string]string,
	abilities []AbilityInfo,
) {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	fmt.Println(ColorBrightCyan + TextBold + "CHARACTER SHEET" + ColorReset)
	fmt.Println(ColorYellow + strings.Repeat("=", 50) + ColorReset)
	fmt.Printf("%sName: %s%s%s (Level %s%d%s %s%s%s)\n",
		ColorWhite,
		ColorBrightYellow+TextBold, name, ColorWhite,
		ColorBrightGreen, level, ColorWhite,
		ColorBrightMagenta, class, ColorReset)

	fmt.Printf("%sHealth: %s%d/%d%s | Mana: %s%d/%d%s\n",
		ColorWhite,
		ColorBrightRed, health, maxHealth, ColorWhite,
		ColorBrightBlue, mana, maxMana, ColorReset)

	fmt.Printf("%sExperience: %s%d/%d%s | Gold: %s%d%s\n",
		ColorWhite,
		ColorBrightGreen, experience, levelUpExp, ColorWhite,
		ColorBrightYellow, gold, ColorReset)

	fmt.Println(ColorBrightWhite + "\nATTRIBUTES" + ColorReset)
	fmt.Println(ColorCyan + strings.Repeat("-", 50) + ColorReset)

	strColor := ColorBrightRed
	fmt.Printf("%sStrength:     %s%d%s (Base: %s%d%s + Bonus: %s%d%s)\n",
		ColorWhite, strColor, strength, ColorWhite,
		strColor, baseStr, ColorWhite,
		strColor, bonusStr, ColorReset)

	agiColor := ColorBrightGreen
	fmt.Printf("%sAgility:      %s%d%s (Base: %s%d%s + Bonus: %s%d%s)\n",
		ColorWhite, agiColor, agility, ColorWhite,
		agiColor, baseAgi, ColorWhite,
		agiColor, bonusAgi, ColorReset)

	chaColor := ColorBrightYellow
	fmt.Printf("%sCharisma:     %s%d%s (Base: %s%d%s + Bonus: %s%d%s)\n",
		ColorWhite, chaColor, charisma, ColorWhite,
		chaColor, baseCha, ColorWhite,
		chaColor, bonusCha, ColorReset)

	intColor := ColorBrightBlue
	fmt.Printf("%sIntelligence: %s%d%s (Base: %s%d%s + Bonus: %s%d%s)\n",
		ColorWhite, intColor, intelligence, ColorWhite,
		intColor, baseInt, ColorWhite,
		intColor, bonusInt, ColorReset)

	fmt.Println(ColorBrightWhite + "\nEQUIPMENT" + ColorReset)
	fmt.Println(ColorCyan + strings.Repeat("-", 50) + ColorReset)

	if len(equipment) == 0 {
		fmt.Println(ColorWhite + "No equipment" + ColorReset)
	} else {
		for slot, itemName := range equipment {
			// Color based on item type
			itemType, exists := equipmentTypes[slot]
			if !exists {
				itemType = "unknown"
			}
			itemColor := GetItemColor(itemType)
			fmt.Printf("%s%s: %s%s%s\n",
				ColorWhite, slot,
				itemColor+TextBold, itemName, ColorReset)
		}
	}

	fmt.Println(ColorBrightWhite + "\nABILITIES" + ColorReset)
	fmt.Println(ColorCyan + strings.Repeat("-", 50) + ColorReset)

	for _, ability := range abilities {
		// Color based on ability type
		abilityColor := ColorWhite
		switch ability.Type {
		case "attack":
			abilityColor = ColorBrightRed
		case "heal":
			abilityColor = ColorBrightGreen
		case "buff":
			abilityColor = ColorBrightCyan
		case "debuff":
			abilityColor = ColorBrightMagenta
		}

		fmt.Printf("%s%s: %s%s (Power: %s%d%s, Mana: %s%d%s)\n",
			abilityColor+TextBold, ability.Name, ColorWhite,
			ability.Description,
			ColorYellow, ability.Power, ColorWhite,
			ColorBlue, ability.ManaCost, ColorReset)
	}

	fmt.Println(ColorBrightGreen + "\nPress any key to return..." + ColorReset)
	// Wait for any key press
	input.ReadKey()
}

// AbilityInfo represents the data needed to display an ability
type AbilityInfo struct {
	Name        string
	Description string
	Type        string
	Power       int
	ManaCost    int
}
