package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/simonbredeche/simonbredeche/export"
	"github.com/simonbredeche/simonbredeche/shared"
	"image/color"
	"strconv"
)

const (
	statDisplayWidth  = 150
	statDisplayHeight = 80
	statDisplayX      = rectMinusX + rectMinusWidth + 10
	statDisplayY      = 10
)

const (
	rectMinusWidth  = 80
	rectMinusHeight = 80
	rectMinusX      = rectPlusX + rectPlusWidth + 10
	rectMinusY      = 10
)

const (
	rectPlusWidth  = 80
	rectPlusHeight = 80
	rectPlusX      = exportRectX + exportRectWidth + 10
	rectPlusY      = 10
)

const (
	exportRectX      = 10
	exportRectY      = 10
	exportRectWidth  = 100
	exportRectHeight = 80
)

const (
	loadRectWidth  = 100
	loadRectHeight = 80
	loadRectX      = shared.SCREEN_WIDTH - loadRectWidth - 10
	loadRectY      = 10
)

const (
	zoomPlusWidth  = 80
	zoomPlusHeight = 80
	zoomPlusX      = rectStartX + rectStartWidth + 10
	zoomPlusY      = 10
)

const (
	zoomMinusWidth  = 80
	zoomMinusHeight = 80
	zoomMinusX      = zoomPlusX + zoomPlusWidth + 10
	zoomMinusY      = 10
)

const (
	rectStartX      = shared.SCREEN_WIDTH/2 - buttonStartSize/2
	rectStartY      = 10
	rectStartWidth  = buttonStartSize
	rectStartHeight = 80
)

const buttonStartSize = 200

func DrawGUI(timeTick int64, currentGeneration *int, screen *ebiten.Image) {

	buttonColor := color.RGBA{34, 139, 34, 255}

	vector.DrawFilledRect(screen, float32(rectStartX), rectStartY, float32(buttonStartSize), rectStartHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "START", rectStartX, 10)

	vector.DrawFilledRect(screen, float32(exportRectX), exportRectY, float32(exportRectWidth), exportRectHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "EXPORT", exportRectX, 10)

	vector.DrawFilledRect(screen, float32(loadRectX), loadRectY, float32(loadRectWidth), loadRectHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "LOAD", loadRectX, 10)

	vector.DrawFilledRect(screen, float32(rectPlusX), rectPlusY, float32(rectPlusWidth), rectPlusHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "PLUS TICK", rectPlusX, 10)

	vector.DrawFilledRect(screen, float32(rectMinusX), rectMinusY, float32(rectMinusWidth), rectMinusHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "MINUS TICK", rectMinusX, 10)

	vector.DrawFilledRect(screen, float32(statDisplayX), statDisplayY, float32(statDisplayWidth), statDisplayHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "TICK SPEED : "+strconv.Itoa(int(timeTick)), statDisplayX, 10)
	ebitenutil.DebugPrintAt(screen, "GENERATION : "+strconv.Itoa(*currentGeneration), statDisplayX, 30)

	vector.DrawFilledRect(screen, float32(zoomPlusX), zoomPlusY, float32(zoomPlusWidth), zoomPlusHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "ZOOM PLUS", zoomPlusX, 10)

	vector.DrawFilledRect(screen, float32(zoomMinusX), zoomMinusY, float32(zoomMinusWidth), zoomMinusHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "ZOOM MINUS", zoomMinusX, 10)
}

func updateGui(timeTick *int64, gameStarted *bool, gridSize *int, tileSize *int, tempGrid *[][]bool, gridArray *[][]bool) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		//start button
		if mx >= rectStartX && mx <= rectStartX+rectStartWidth && my >= rectStartY && my <= rectStartY+rectStartHeight {
			*gameStarted = !*gameStarted
		}
		//Export button
		if mx >= exportRectX && mx <= exportRectX+exportRectWidth && my >= exportRectY && my <= exportRectY+exportRectHeight {
			export.ExportGrid(gridArray)
		}
		//Load button
		if mx >= loadRectX && mx <= loadRectX+loadRectWidth && my >= loadRectY && my <= loadRectY+loadRectHeight {
			export.LoadFromFile(gridArray)
		}
		//Plus tick
		if mx >= rectPlusX && mx <= rectPlusX+rectPlusWidth && my >= rectPlusY && my <= rectPlusY+rectPlusHeight {
			*timeTick += 1
		}
		//Minus tick
		if mx >= rectMinusX && mx <= rectMinusX+rectMinusWidth && my >= rectMinusY && my <= rectMinusY+rectMinusHeight {
			*timeTick -= 1
		}
		//Zoom plus
		if mx >= zoomPlusX && mx <= zoomPlusX+zoomPlusWidth && my >= zoomPlusY && my <= zoomPlusY+zoomPlusHeight {
			*gridSize = *gridSize / 2
			*tileSize = *tileSize * 2
			InitGrids(gridSize, tempGrid, gridArray)
		}
		//Zoom minus
		if mx >= zoomMinusX && mx <= zoomMinusX+zoomMinusWidth && my >= zoomMinusY && my <= zoomMinusY+zoomMinusHeight {
			*gridSize = *gridSize * 2
			*tileSize = *tileSize / 2
			InitGrids(gridSize, tempGrid, gridArray)
		}
	}
}

// Action qui s'execute quand on clique sur une cellule
func DetectInput(timeTick *int64, gameStarted *bool, tileSize *int, tempGrid *[][]bool, gridArray *[][]bool) {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		gridX := ConvertGlobalToGrid(mx, tileSize)
		gridY := ConvertGlobalToGrid(my-shared.OFFSET_Y, tileSize)
		if IsInBoundCoordinates(gridX, gridY) {
			(*gridArray)[gridX][gridY] = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		gridX := ConvertGlobalToGrid(mx, tileSize)
		gridY := ConvertGlobalToGrid(my-shared.OFFSET_Y, tileSize)
		if IsInBoundCoordinates(gridX, gridY) {
			(*gridArray)[gridX][gridY] = false
		}
	}
	updateGui(timeTick, gameStarted, &shared.GridSize, tileSize, tempGrid, gridArray)
}
