package dungeon

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/ui"
	"math/rand"
)

func GenerateDungeon(width, height, level int) *Dungeon {
	theme := GetThemeForLevel(level)

	dungeon := &Dungeon{
		Level:   level,
		Width:   width,
		Height:  height,
		Tiles:   make([][]ui.Tile, height),
		Visited: make([][]bool, height),
		Visible: make([][]bool, height),
		Rooms:   make([]Room, 0),
		Theme:   theme,
	}

	for y := 0; y < height; y++ {
		dungeon.Tiles[y] = make([]ui.Tile, width)
		dungeon.Visited[y] = make([]bool, width)
		dungeon.Visible[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			dungeon.Tiles[y][x] = TileWall
		}
	}

	numRooms := rand.Intn(5) + 5 // 5-10 rooms
	for i := 0; i < numRooms; i++ {
		roomWidth := rand.Intn(8) + 5  // 5-12 width
		roomHeight := rand.Intn(5) + 4 // 4-8 height
		roomX := rand.Intn(width-roomWidth-2) + 1
		roomY := rand.Intn(height-roomHeight-2) + 1

		overlaps := false
		for _, room := range dungeon.Rooms {
			if roomX <= room.X+room.Width+1 && roomX+roomWidth+1 >= room.X &&
				roomY <= room.Y+room.Height+1 && roomY+roomHeight+1 >= room.Y {
				overlaps = true
				break
			}
		}

		if !overlaps {

			room := Room{
				X:      roomX,
				Y:      roomY,
				Width:  roomWidth,
				Height: roomHeight,
			}
			dungeon.Rooms = append(dungeon.Rooms, room)

			for y := roomY; y < roomY+roomHeight; y++ {
				for x := roomX; x < roomX+roomWidth; x++ {
					dungeon.Tiles[y][x] = TileFloor
				}
			}

			if i > 0 && len(dungeon.Rooms) > 1 {

				prevRoomIndex := len(dungeon.Rooms) - 2 // Get the room just before the one we added
				prevRoom := dungeon.Rooms[prevRoomIndex]

				startX := roomX + roomWidth/2
				startY := roomY + roomHeight/2
				endX := prevRoom.X + prevRoom.Width/2
				endY := prevRoom.Y + prevRoom.Height/2

				if rand.Intn(2) == 0 {
					createHorizontalCorridor(dungeon, startX, endX, startY)
					createVerticalCorridor(dungeon, startY, endY, endX)
				} else {
					createVerticalCorridor(dungeon, startY, endY, startX)
					createHorizontalCorridor(dungeon, startX, endX, endY)
				}
			}
		}
	}

	if len(dungeon.Rooms) > 0 {
		lastRoom := dungeon.Rooms[len(dungeon.Rooms)-1]
		stairsX := lastRoom.X + lastRoom.Width/2
		stairsY := lastRoom.Y + lastRoom.Height/2
		dungeon.StairsPos = Position{X: stairsX, Y: stairsY}
		dungeon.Tiles[stairsY][stairsX] = TileStairs
	}

	return dungeon
}

func createHorizontalCorridor(dungeon *Dungeon, x1, x2, y int) {
	for x := min(x1, x2); x <= max(x1, x2); x++ {
		if y >= 0 && y < dungeon.Height && x >= 0 && x < dungeon.Width {
			dungeon.Tiles[y][x] = TileFloor
		}
	}
}

func createVerticalCorridor(dungeon *Dungeon, y1, y2, x int) {
	for y := min(y1, y2); y <= max(y1, y2); y++ {
		if y >= 0 && y < dungeon.Height && x >= 0 && x < dungeon.Width {
			dungeon.Tiles[y][x] = TileFloor
		}
	}
}

func GetStartingPosition(dungeon *Dungeon) Position {

	if len(dungeon.Rooms) > 0 {
		startRoom := dungeon.Rooms[0]
		return Position{
			X: startRoom.X + startRoom.Width/2,
			Y: startRoom.Y + startRoom.Height/2,
		}
	}

	return Position{
		X: dungeon.Width / 2,
		Y: dungeon.Height / 2,
	}
}
