package dungeon

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/common"
	"github.com/ppedziwiatr/ascii-gig/pkg/ui"
)

const (
	TileFloor      ui.Tile = '.'
	TileWall       ui.Tile = '#'
	TileStairs     ui.Tile = '>'
	TileDoor       ui.Tile = '+'
	TilePlayer     ui.Tile = '@'
	TileMonster    ui.Tile = 'M'
	TileItem       ui.Tile = 'i'
	TileEmpty      ui.Tile = ' '
	TileUnexplored ui.Tile = '?'
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Position struct {
	X, Y int
}

type Room struct {
	X, Y, Width, Height int
}

type Dungeon struct {
	Level     int
	Width     int
	Height    int
	Tiles     [][]ui.Tile
	Rooms     []Room
	Visited   [][]bool
	Visible   [][]bool
	StairsPos Position
	Theme     Theme
}

func (d *Dungeon) GetVisibleTiles(playerX, playerY int) [][]bool {

	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			d.Visible[y][x] = false
		}
	}

	visionRadius := 10

	for y := playerY - visionRadius; y <= playerY+visionRadius; y++ {
		for x := playerX - visionRadius; x <= playerX+visionRadius; x++ {

			if y >= 0 && y < d.Height && x >= 0 && x < d.Width {

				dx := playerX - x
				dy := playerY - y
				distance := dx*dx + dy*dy

				if distance <= visionRadius*visionRadius {

					if hasLineOfSight(d, playerX, playerY, x, y) {
						d.Visible[y][x] = true
						d.Visited[y][x] = true
					}
				}
			}
		}
	}

	return d.Visible
}

func (d *Dungeon) IsWalkable(x, y int) bool {

	if x < 0 || x >= d.Width || y < 0 || y >= d.Height {
		return false
	}

	if d.Tiles[y][x] == TileWall {
		return false
	}

	return true
}

func (d *Dungeon) IsPositionEmpty(x, y int, entities []common.Entity) bool {

	if !d.IsWalkable(x, y) {
		return false
	}

	for _, entity := range entities {
		// TODO: Fix
		if entity == nil {
			continue
		}

		entX, entY := entity.GetPosition()
		if entX == x && entY == y {
			return false
		}
	}

	return true
}

func hasLineOfSight(dungeon *Dungeon, x1, y1, x2, y2 int) bool {

	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}
	if y1 >= y2 {
		sy = -1
	}
	err := dx - dy

	for {

		if x1 != x2 || y1 != y2 { // Don't check the destination
			if dungeon.Tiles[y1][x1] == TileWall {
				return false
			}
		}

		if x1 == x2 && y1 == y2 {
			return true
		}

		e2 := 2 * err
		if e2 > -dy {
			err = err - dy
			x1 = x1 + sx
		}
		if e2 < dx {
			err = err + dx
			y1 = y1 + sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
