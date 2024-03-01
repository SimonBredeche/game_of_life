package main

import (
	"github.com/simonbredeche/simonbredeche/shared"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/simonbredeche/simonbredeche/game"
)

var gameStarted = false
var currentGeneration = 0
var tileSize = 16
var tempGrid [][]bool
var gridArray [][]bool
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
	game.DetectInput(&timeTick, &gameStarted, &tileSize, &tempGrid, &gridArray)
	if currentTick >= timeTick {
		game.UpdateGrid(&gameStarted, &currentGeneration, &tempGrid, &gridArray)
		currentTick = 0
	}
	return nil
}

// Dessine la grille
func (g *Game) Draw(screen *ebiten.Image) {
	game.DrawGrid(shared.GridSize, &tileSize, &gridArray, screen)
	game.DrawGUI(timeTick, &currentGeneration, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return shared.SCREEN_WIDTH, shared.SCREEN_HEIGHT
}

func main() {
	game.InitGrids(&shared.GridSize, &tempGrid, &gridArray)
	ebiten.SetWindowSize(shared.SCREEN_WIDTH, shared.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game of life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
