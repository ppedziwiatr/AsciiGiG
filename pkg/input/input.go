package input

import (
	"bufio"
	"os"
	"syscall"

	"golang.org/x/term"
)

const (
	KeyArrowUp    = 1000
	KeyArrowDown  = 1001
	KeyArrowLeft  = 1002
	KeyArrowRight = 1003
)

func ReadKey() (rune, int, error) {

	fd := int(syscall.Stdin)

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return 0, 0, err
	}

	defer term.Restore(fd, oldState)

	reader := bufio.NewReader(os.Stdin)
	r, _, err := reader.ReadRune()
	if err != nil {
		return 0, 0, err
	}

	if r == '\x1b' {

		escape := make([]rune, 3)
		escape[0] = r

		for i := 1; i < 3; i++ {
			r, _, err = reader.ReadRune()
			if err != nil {
				return 0, 0, err
			}
			escape[i] = r
		}

		if escape[1] == '[' {
			switch escape[2] {
			case 'A':
				return 0, KeyArrowUp, nil
			case 'B':
				return 0, KeyArrowDown, nil
			case 'C':
				return 0, KeyArrowRight, nil
			case 'D':
				return 0, KeyArrowLeft, nil
			}
		}
	}

	return r, 0, nil
}

func HandlePlayerInput(gameMode string) (string, int) {

	char, key, err := ReadKey()
	if err != nil {
		return "none", -1
	}

	input := ""

	if key != 0 {
		switch key {
		case KeyArrowUp:
			input = "up"
		case KeyArrowDown:
			input = "down"
		case KeyArrowLeft:
			input = "left"
		case KeyArrowRight:
			input = "right"
		}
	} else if char > 0 {

		input = string(char)
	}

	if gameMode == "exploring" {
		switch input {
		case "w", "up": // Up
			return "move", 0 // North
		case "s", "down": // Down
			return "move", 1 // South
		case "a", "left": // Left
			return "move", 2 // West
		case "d", "right": // Right
			return "move", 3 // East
		case "g": // Get item
			return "pickup", 0
		case "i": // Inventory
			return "inventory", 0
		case "c": // Character sheet
			return "character", 0
		case "q": // Quit
			return "quit", 0
		}
	} else if gameMode == "combat" {
		abilityIndex := -1
		if input >= "1" && input <= "9" {
			abilityIndex = int(input[0] - '1')
			return "ability", abilityIndex
		} else if input == "r" {
			return "run", 0
		}
	} else if gameMode == "inventory" {

		if input == "i" || input == "\r" || input == "\n" || input == " " {
			return "close_inventory", 0
		}

	}

	return "none", -1
}
