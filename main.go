package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const tileSize = 16
const gridSize = 100

var gameStarted = false

var gridArray [gridSize][gridSize]bool

var tempGrid [gridSize][gridSize]bool

type Game struct{}

func (g *Game) Update() error {

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			tempGrid[x][y] = false
		}
	}

	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Printf("Mouse is pressed at (%d, %d)\n", mx, my)
		gridX := convertGlobalToGrid(mx)
		gridY := convertGlobalToGrid(my)
		gridArray[gridX][gridY] = true
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		fmt.Printf("Game started")
		gameStarted = true
	}

	if gameStarted {
		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				checkCell(x, y)
			}
		}
		gridNextState()
	}

	return nil
}

func gridNextState() {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			gridArray[x][y] = tempGrid[x][y]
		}
	}
}

func checkCell(posX int, posY int) {
	alive := 0
	for i := posX - 1; i <= posX+1; i++ {
		for j := posY - 1; j <= posY+1; j++ {
			if isInBoundCoordinates(i, j) && !(i == posX && j == posY) {
				if gridArray[i][j] {
					alive++
				}
			}
		}
	}
	if !gridArray[posX][posY] {
		if alive == 3 {
			tempGrid[posX][posY] = true
		} else {
			tempGrid[posX][posY] = false
		}
	} else if gridArray[posX][posY] {
		if alive == 2 || alive == 3 {
			tempGrid[posX][posY] = true
		} else {
			tempGrid[posX][posY] = false
		}
	}

}

func isInBoundCoordinates(x int, y int) bool {
	return x >= 0 && y >= 0 && x < gridSize && y < gridSize
}

func convertGlobalToGrid(coord int) int {
	return coord / tileSize
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			vector.StrokeRect(screen, float32(x*tileSize), float32(y*tileSize), tileSize, tileSize, 1, color.RGBA{255, 255, 255, 255}, false)
			if gridArray[x][y] {
				vector.DrawFilledRect(screen, float32(x*tileSize), float32(y*tileSize), tileSize, tileSize, color.RGBA{255, 255, 255, 255}, false)
			}
		}
	}
	//vector.DrawFilledRect(screen, 0, 0, 32, 32, color.RGBA{255, 255, 255, 255}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 728
}

func main() {
	initGame()
	ebiten.SetWindowSize(1024, 728)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func initGame() {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			gridArray[x][y] = false
		}
	}
}
