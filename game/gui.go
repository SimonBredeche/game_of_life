package game

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/simonbredeche/simonbredeche/export"
	"github.com/simonbredeche/simonbredeche/manager"
	"github.com/simonbredeche/simonbredeche/shared"
)

type Rect struct {
	x      int
	y      int
	width  int
	height int
}

// STAT BUTTON DIMENSION
var statRect = &Rect{
	x:      rectMinus.x + rectMinus.width + 10,
	y:      10,
	width:  150,
	height: 80,
}

// TICK MINUS BUTTON DIMENSION
var rectMinus = &Rect{
	x:      rectPlus.x + rectPlus.width + 10,
	y:      10,
	width:  80,
	height: 80,
}

// TICK PLUS BUTTON DIMENSION
var rectPlus = &Rect{
	x:      rectExport.x + rectExport.width + 10,
	y:      10,
	width:  80,
	height: 80,
}

// EXPORT BUTTON DIMENSION
var rectExport = &Rect{
	x:      10,
	y:      10,
	width:  100,
	height: 80,
}

// LOAD BUTTON DIMENSION
var loadRect = &Rect{
	x:      shared.SCREEN_WIDTH - 100 - 10,
	y:      10,
	width:  100,
	height: 80,
}

// ZOOM PLUS BUTTON DIMENSION
var zoomPlusRect = &Rect{
	x:      rectStart.x + rectStart.width + 10,
	y:      10,
	width:  80,
	height: 80,
}

// ZOOM MINUS BUTTON DIMENSION
var zoomMinusRect = &Rect{
	x:      zoomPlusRect.x + zoomPlusRect.width + 10,
	y:      10,
	width:  80,
	height: 80,
}

// CLEAR BUTTON DIMENSION
var clearRect = &Rect{
	x:      zoomMinusRect.x + zoomMinusRect.width + 10,
	y:      10,
	width:  80,
	height: 80,
}

// START BUTTON DIMENSION
var rectStart = &Rect{
	x:      shared.SCREEN_WIDTH/2 - buttonStartSize/2,
	y:      10,
	width:  buttonStartSize,
	height: 80,
}

const buttonStartSize = 200

func DrawGUI(gameState *manager.GameState, screen *ebiten.Image) {

	drawButton(rectStart, "START", screen)

	drawButton(rectExport, "EXPORT", screen)

	drawButton(loadRect, "LOAD", screen)

	drawButton(rectPlus, "PLUS TICK", screen)

	drawButton(rectMinus, "MINUS TICK", screen)

	drawButton(statRect, "TICK SPEED : "+strconv.Itoa(int(gameState.TimeTick))+"\n"+"GENERATION : "+strconv.Itoa(gameState.CurrentGeneration), screen)

	drawButton(zoomPlusRect, "ZOOM PLUS", screen)

	drawButton(zoomMinusRect, "ZOOM MINUS", screen)

	drawButton(clearRect, "CLEAR", screen)
}

func drawButton(rec *Rect, title string, screen *ebiten.Image) {
	buttonColor := color.RGBA{34, 139, 34, 255}
	vector.DrawFilledRect(screen, float32(rec.x), float32(rec.y), float32(rec.width), float32(rec.height), buttonColor, false)
	ebitenutil.DebugPrintAt(screen, title, rec.x, 10)
}

func checkCursorInRectangle(rec *Rect, mx int, my int) bool {
	return mx >= rec.x && mx <= rec.x+rec.width && my >= rec.y && my <= rec.y+rec.height
}

func updateGui(gameState *manager.GameState, gridManager *manager.GridManager) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		//start button
		if checkCursorInRectangle(rectStart, mx, my) {
			gameState.GameStarted = !gameState.GameStarted
		}
		//Export button
		if checkCursorInRectangle(rectExport, mx, my) {
			export.ExportGrid(gameState, gridManager)
		}
		//Load button
		if checkCursorInRectangle(loadRect, mx, my) {
			export.LoadFromFile(gameState, gridManager)
		}
		//Plus tick
		if checkCursorInRectangle(rectPlus, mx, my) {
			gameState.TimeTick += 1
		}
		//Minus tick
		if checkCursorInRectangle(rectMinus, mx, my) {
			gameState.TimeTick -= 1
		}
		//Zoom plus
		if checkCursorInRectangle(zoomPlusRect, mx, my) {
			gameState.GridSize = gameState.GridSize / 2
			gameState.TileSize = gameState.TileSize * 2
			InitGrids(gameState, gridManager)
		}
		//Zoom minus
		if checkCursorInRectangle(zoomMinusRect, mx, my) {
			gameState.GridSize = gameState.GridSize * 2
			gameState.TileSize = gameState.TileSize / 2
			InitGrids(gameState, gridManager)
		}
		//Clear rect
		if checkCursorInRectangle(clearRect, mx, my) {
			InitGrids(gameState, gridManager)
		}
	}
}

// Action qui s'execute quand on clique sur une cellule
func DetectInput(gameState *manager.GameState, gridManager *manager.GridManager) {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		gridX := ConvertGlobalToGrid(mx, &gameState.TileSize)
		gridY := ConvertGlobalToGrid(my-shared.OFFSET_Y, &gameState.TileSize)
		if IsInBoundCoordinates(gridX, gridY, gameState.GridSize) {
			gridManager.GridArray[gridX][gridY] = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		gridX := ConvertGlobalToGrid(mx, &gameState.TileSize)
		gridY := ConvertGlobalToGrid(my-shared.OFFSET_Y, &gameState.TileSize)
		if IsInBoundCoordinates(gridX, gridY, gameState.GridSize) {
			gridManager.GridArray[gridX][gridY] = false
		}
	}
	updateGui(gameState, gridManager)
}
