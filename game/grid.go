package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/simonbredeche/simonbredeche/manager"
	"github.com/simonbredeche/simonbredeche/shared"
)

func DrawGrid(gameState *manager.GameState, gridManager *manager.GridManager, screen *ebiten.Image) {
	for x := 0; x < gameState.GridSize; x++ {
		for y := 0; y < gameState.GridSize; y++ {
			vector.StrokeRect(screen, float32(x*gameState.TileSize), shared.OFFSET_Y+float32(y*gameState.TileSize), float32(gameState.TileSize), float32(gameState.TileSize), 1, color.RGBA{255, 255, 255, 255}, false)
			if gridManager.GridArray[x][y] {
				vector.DrawFilledRect(screen, float32(x*gameState.TileSize), shared.OFFSET_Y+float32(y*gameState.TileSize), float32(gameState.TileSize), float32(gameState.TileSize), color.RGBA{255, 255, 255, 255}, false)
			}
		}
	}
}

func InitGrids(gameState *manager.GameState, gridManager *manager.GridManager) {

	var newGrid = make([][]bool, gameState.GridSize)
	var newTempGrid = make([][]bool, gameState.GridSize)

	for x := 0; x < gameState.GridSize; x++ {
		newGrid[x] = make([]bool, gameState.GridSize)
		newTempGrid[x] = make([]bool, gameState.GridSize)
		for y := 0; y < gameState.GridSize; y++ {
			newGrid[x][y] = false
			newTempGrid[x][y] = false
		}
	}

	gridManager.GridArray = newGrid
	gridManager.TempGrid = newTempGrid
}

func UpdateGrid(gameState *manager.GameState, gridManager *manager.GridManager) {
	for x := 0; x < gameState.GridSize; x++ {
		for y := 0; y < gameState.GridSize; y++ {
			gridManager.TempGrid[x][y] = false
		}
	}

	if gameState.GameStarted {
		gameState.CurrentGeneration++
		for x := 0; x < gameState.GridSize; x++ {
			for y := 0; y < gameState.GridSize; y++ {
				CheckCell(x, y, gridManager, gameState)
			}
		}
		GridNextState(gameState, gridManager)
	}
}

// Copie la grille temportaite dans la grille principale
func GridNextState(gameState *manager.GameState, gridManager *manager.GridManager) {
	for x := 0; x < gameState.GridSize; x++ {
		for y := 0; y < gameState.GridSize; y++ {
			gridManager.GridArray[x][y] = gridManager.TempGrid[x][y]
		}
	}
}

// Vérifie si une cellule doit être vivante ou morte
func CheckCell(posX int, posY int, gridManager *manager.GridManager, gameState *manager.GameState) {
	alive := countAliveCell(posX, posY, gridManager, gameState)
	if !gridManager.GridArray[posX][posY] {
		if alive == 3 {
			gridManager.TempGrid[posX][posY] = true
		} else {
			gridManager.TempGrid[posX][posY] = false
		}
	} else {
		if alive == 2 || alive == 3 {
			gridManager.TempGrid[posX][posY] = true
		} else {
			gridManager.TempGrid[posX][posY] = false
		}
	}
}

func countAliveCell(posX int, posY int, gridManager *manager.GridManager, gameState *manager.GameState) int {
	alive := 0
	for i := posX - 1; i <= posX+1; i++ {
		for j := posY - 1; j <= posY+1; j++ {
			if IsInBoundCoordinates(i, j, gameState.GridSize) && !(i == posX && j == posY) {
				if gridManager.GridArray[i][j] {
					alive++
				}
			}
		}
	}
	return alive
}

func IsInBoundCoordinates(x int, y int, gridSize int) bool {
	return x >= 0 && y >= 0 && x < gridSize && y < gridSize
}

func ConvertGlobalToGrid(coord int, tileSize *int) int {
	return coord / *tileSize
}
