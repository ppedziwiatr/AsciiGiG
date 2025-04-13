# Diablo-like ASCII Roguelike

A Diablo-inspired roguelike game with ASCII graphics, procedurally generated dungeons, and turn-based combat.

## Features

- 8 procedurally generated dungeon levels inspired by Diablo's progression into Hell
- Turn-based combat system similar
- 4 character attributes that affect gameplay (ASCI ...):
    - **Agility**: Affects turn order, dodge chance, and ability to use ranged weapons
    - **Strength**: Affects maximum health, attack power, and ability to use heavy weapons/armor
    - **Charisma**: Affects persuasion attempts and item prices
    - **Intelligence**: Affects magic power, mana pool, and spell variety
- Multiple character classes with unique abilities
- Procedurally generated monsters with different behaviors and abilities
- Procedurally generated items (weapons, armor, and consumables)
- Colorful terminal interface

## Requirements

- Go 1.24 or higher
- Terminal that supports ANSI colors

## Installation

1. Install Go from https://golang.org/
2. Clone this repository:
```
git clone https://github.com/ppedziwiatr/ascii-gig.git
cd diablo-roguelike
```
3. Install dependencies:
```
go get golang.org/x/term
```
4. Build and run the game:
```
go build -o roguelike ./cmd/game
./roguelike
```

## Controls

- Arrow keys or WASD: Move
- G: Pick up item
- I: Inventory
- C: Character sheet
- Q: Quit game
- 1-9: Use ability (in combat)
- R: Run away (in combat)

## Dungeon Levels

1. **Church Catacombs**: Dark and damp stone passages beneath the church
2. **Underground Passages**: Natural caves with occasional constructed walls
3. **Forgotten Tombs**: Ancient burial chambers with dusty sarcophagi
4. **Torture Chambers**: Blood-stained rooms with instruments of pain
5. **Hellish Caves**: Caverns heated by infernal fires below
6. **Burning Hell**: Lakes of fire and brimstone
7. **Realm of Hatred**: A twisted landscape of malice and spite
8. **Diablo's Lair**: The final resting place of the Lord of Terror

## Game Structure

The codebase is organized into the following packages:

- `cmd/game`: Main application entry point
- `pkg/ui`: User interface rendering and menus
- `pkg/input`: Keyboard input handling
- `pkg/dungeon`: Dungeon generation and management
- `pkg/character`: Player character implementation
- `pkg/monster`: Monster generation and behavior
- `pkg/item`: Item generation and effects
- `pkg/game`: Core game logic and state management

## License

This project is licensed under the MIT License - see the LICENSE file for details.