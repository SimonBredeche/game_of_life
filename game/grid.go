package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/simonbredeche/simonbredeche/shared"
	"image/color"
)

func DrawGrid(gridSize int, tileSize *int, gridArray *[][]bool, screen *ebiten.Image) {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			vector.StrokeRect(screen, float32(x**tileSize), shared.OFFSET_Y+float32(y**tileSize), float32(*tileSize), float32(*tileSize), 1, color.RGBA{255, 255, 255, 255}, false)
			if (*gridArray)[x][y] {
				vector.DrawFilledRect(screen, float32(x**tileSize), shared.OFFSET_Y+float32(y**tileSize), float32(*tileSize), float32(*tileSize), color.RGBA{255, 255, 255, 255}, false)
			}
		}
	}
}

func InitGrids(gridSize *int, tempGrid *[][]bool, gridArray *[][]bool) {

	var newGrid = make([][]bool, *gridSize)
	var newTempGrid = make([][]bool, *gridSize)

	for x := 0; x < *gridSize; x++ {
		newGrid[x] = make([]bool, *gridSize)
		newTempGrid[x] = make([]bool, *gridSize)
		for y := 0; y < *gridSize; y++ {
			newGrid[x][y] = false
			newTempGrid[x][y] = false
		}
	}

	*gridArray = newGrid
	*tempGrid = newTempGrid
}

func UpdateGrid(gameStarted *bool, currentGeneration *int, tempGrid *[][]bool, gridArray *[][]bool) {
	for x := 0; x < shared.GridSize; x++ {
		for y := 0; y < shared.GridSize; y++ {
			(*tempGrid)[x][y] = false
		}
	}

	if *gameStarted {
		*currentGeneration++
		for x := 0; x < shared.GridSize; x++ {
			for y := 0; y < shared.GridSize; y++ {
				CheckCell(x, y, tempGrid, gridArray)
			}
		}
		GridNextState(tempGrid, gridArray)
	}
}

// Copie la grille temportaite dans la grille principale
func GridNextState(tempGrid *[][]bool, gridArray *[][]bool) {
	for x := 0; x < shared.GridSize; x++ {
		for y := 0; y < shared.GridSize; y++ {
			(*gridArray)[x][y] = (*tempGrid)[x][y]
		}
	}
}

// Vérifie si une cellule doit être vivante ou morte
func CheckCell(posX int, posY int, tempGrid *[][]bool, gridArray *[][]bool) {
	alive := 0
	for i := posX - 1; i <= posX+1; i++ {
		for j := posY - 1; j <= posY+1; j++ {
			if IsInBoundCoordinates(i, j) && !(i == posX && j == posY) {
				if (*gridArray)[i][j] {
					alive++
				}
			}
		}
	}
	if !(*gridArray)[posX][posY] {
		if alive == 3 {
			(*tempGrid)[posX][posY] = true
		} else {
			(*tempGrid)[posX][posY] = false
		}
	} else if (*gridArray)[posX][posY] {
		if alive == 2 || alive == 3 {
			(*tempGrid)[posX][posY] = true
		} else {
			(*tempGrid)[posX][posY] = false
		}
	}
}

func IsInBoundCoordinates(x int, y int) bool {
	return x >= 0 && y >= 0 && x < shared.GridSize && y < shared.GridSize
}

func ConvertGlobalToGrid(coord int, tileSize *int) int {
	return coord / *tileSize
}
