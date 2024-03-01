package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/simonbredeche/simonbredeche/manager"
	"github.com/simonbredeche/simonbredeche/shared"
)

func DrawGrid(gridSize int, tileSize *int, gameManager *manager.GameManager, screen *ebiten.Image) {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			vector.StrokeRect(screen, float32(x**tileSize), shared.OFFSET_Y+float32(y**tileSize), float32(*tileSize), float32(*tileSize), 1, color.RGBA{255, 255, 255, 255}, false)
			if gameManager.GridArray[x][y] {
				vector.DrawFilledRect(screen, float32(x**tileSize), shared.OFFSET_Y+float32(y**tileSize), float32(*tileSize), float32(*tileSize), color.RGBA{255, 255, 255, 255}, false)
			}
		}
	}
}

func InitGrids(gridSize *int, gameManager *manager.GameManager) {

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

	gameManager.GridArray = newGrid
	gameManager.TempGrid = newTempGrid
}

func UpdateGrid(gameStarted *bool, currentGeneration *int, gameManager *manager.GameManager) {
	for x := 0; x < shared.GridSize; x++ {
		for y := 0; y < shared.GridSize; y++ {
			gameManager.TempGrid[x][y] = false
		}
	}

	if *gameStarted {
		*currentGeneration++
		for x := 0; x < shared.GridSize; x++ {
			for y := 0; y < shared.GridSize; y++ {
				CheckCell(x, y, gameManager)
			}
		}
		GridNextState(gameManager)
	}
}

// Copie la grille temportaite dans la grille principale
func GridNextState(gameManager *manager.GameManager) {
	for x := 0; x < shared.GridSize; x++ {
		for y := 0; y < shared.GridSize; y++ {
			gameManager.GridArray[x][y] = gameManager.TempGrid[x][y]
		}
	}
}

// Vérifie si une cellule doit être vivante ou morte
func CheckCell(posX int, posY int, gameManager *manager.GameManager) {
	alive := 0
	for i := posX - 1; i <= posX+1; i++ {
		for j := posY - 1; j <= posY+1; j++ {
			if IsInBoundCoordinates(i, j) && !(i == posX && j == posY) {
				if gameManager.GridArray[i][j] {
					alive++
				}
			}
		}
	}
	if !gameManager.GridArray[posX][posY] {
		if alive == 3 {
			gameManager.TempGrid[posX][posY] = true
		} else {
			gameManager.TempGrid[posX][posY] = false
		}
	} else if (gameManager.GridArray)[posX][posY] {
		if alive == 2 || alive == 3 {
			gameManager.TempGrid[posX][posY] = true
		} else {
			gameManager.TempGrid[posX][posY] = false
		}
	}
}

func IsInBoundCoordinates(x int, y int) bool {
	return x >= 0 && y >= 0 && x < shared.GridSize && y < shared.GridSize
}

func ConvertGlobalToGrid(coord int, tileSize *int) int {
	return coord / *tileSize
}
