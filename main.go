package main

import (
	"log"
	"time"

	"github.com/simonbredeche/simonbredeche/manager"
	"github.com/simonbredeche/simonbredeche/shared"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/simonbredeche/simonbredeche/game"
)

var gameManager *manager.GameManager

var gameStarted = false
var currentGeneration = 0
var tileSize = 16
var timeTick = int64(20)

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
	game.DetectInput(&timeTick, &gameStarted, &tileSize, gameManager)
	if currentTick >= timeTick {
		game.UpdateGrid(&gameStarted, &currentGeneration, gameManager)
		currentTick = 0
	}
	return nil
}

// Dessine la grille
func (g *Game) Draw(screen *ebiten.Image) {
	game.DrawGrid(shared.GridSize, &tileSize, gameManager, screen)
	game.DrawGUI(timeTick, &currentGeneration, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return shared.SCREEN_WIDTH, shared.SCREEN_HEIGHT
}

func main() {
	gameManager = &manager.GameManager{}
	game.InitGrids(&shared.GridSize, gameManager)
	ebiten.SetWindowSize(shared.SCREEN_WIDTH, shared.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game of life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
