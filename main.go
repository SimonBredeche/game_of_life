package main

import (
	"log"
	"time"

	"github.com/simonbredeche/simonbredeche/manager"
	"github.com/simonbredeche/simonbredeche/shared"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/simonbredeche/simonbredeche/game"
)

var gridManager *manager.GridManager

var gameState *manager.GameState

var (
	lastUpdateTime time.Time
	currentTick    int64
)

type Game struct{}

func (g *Game) Update() error {

	currentTime := time.Now()
	deltaTime := currentTime.Sub(lastUpdateTime).Milliseconds()
	lastUpdateTime = currentTime
	currentTick += deltaTime
	game.DetectInput(gameState, gridManager)
	if currentTick >= gameState.TimeTick {
		game.UpdateGrid(gameState, gridManager)
		currentTick = 0
	}
	return nil
}

// Dessine la grille
func (g *Game) Draw(screen *ebiten.Image) {
	game.DrawGrid(gameState, gridManager, screen)
	game.DrawGUI(gameState, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return shared.SCREEN_WIDTH, shared.SCREEN_HEIGHT
}

func main() {
	gridManager = &manager.GridManager{}
	gameState = &manager.GameState{}
	gameState.CurrentGeneration = 0
	gameState.GameStarted = false
	gameState.TileSize = 16
	gameState.TimeTick = int64(20)
	gameState.GridSize = 80
	game.InitGrids(gameState, gridManager)
	ebiten.SetWindowSize(shared.SCREEN_WIDTH, shared.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game of life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
