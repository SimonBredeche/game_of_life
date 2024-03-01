package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const offSetY = 100
const SCREEN_WIDTH = 1280
const SCREEN_HEIGHT = 720
const buttonStartSize = 200

var tileSize = 16

var gridSize = 80

var gameStarted = false

var gridArray [][]bool

var tempGrid [][]bool

var timeTick = int64(20)

var currentGeneration = 0

var (
	lastUpdateTime time.Time
	currentTick    int64
)

const (
	rectStartX      = SCREEN_WIDTH/2 - buttonStartSize/2
	rectStartY      = 10
	rectStartWidth  = buttonStartSize
	rectStartHeight = 80
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
	loadRectX      = SCREEN_WIDTH - loadRectWidth - 10
	loadRectY      = 10
)

const (
	rectPlusWidth  = 80
	rectPlusHeight = 80
	rectPlusX      = exportRectX + exportRectWidth + 10
	rectPlusY      = 10
)

const (
	rectMinusWidth  = 80
	rectMinusHeight = 80
	rectMinusX      = rectPlusX + rectPlusWidth + 10
	rectMinusY      = 10
)

const (
	statDisplayWidth  = 150
	statDisplayHeight = 80
	statDisplayX      = rectMinusX + rectMinusWidth + 10
	statDisplayY      = 10
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

type Game struct{}

func (g *Game) Update() error {

	currentTime := time.Now()
	deltaTime := currentTime.Sub(lastUpdateTime).Milliseconds()
	lastUpdateTime = currentTime
	currentTick += deltaTime
	detectInput()
	if currentTick >= timeTick {
		updateGrid()
		currentTick = 0
	}
	return nil
}

func detectInput() {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		gridX := convertGlobalToGrid(mx)
		gridY := convertGlobalToGrid(my - offSetY)
		if isInBoundCoordinates(gridX, gridY) {
			gridArray[gridX][gridY] = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		gridX := convertGlobalToGrid(mx)
		gridY := convertGlobalToGrid(my - offSetY)
		if isInBoundCoordinates(gridX, gridY) {
			gridArray[gridX][gridY] = false
		}
	}
	updateGui()
}

func updateGrid() {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			tempGrid[x][y] = false
		}
	}

	if gameStarted {
		currentGeneration++
		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				checkCell(x, y)
			}
		}
		gridNextState()
	}
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
			vector.StrokeRect(screen, float32(x*tileSize), offSetY+float32(y*tileSize), float32(tileSize), float32(tileSize), 1, color.RGBA{255, 255, 255, 255}, false)
			if gridArray[x][y] {
				vector.DrawFilledRect(screen, float32(x*tileSize), offSetY+float32(y*tileSize), float32(tileSize), float32(tileSize), color.RGBA{255, 255, 255, 255}, false)
			}
		}
	}
	drawGUI(screen)
}

func drawGUI(screen *ebiten.Image) {

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
	ebitenutil.DebugPrintAt(screen, "GENERATION : "+strconv.Itoa(currentGeneration), statDisplayX, 30)

	vector.DrawFilledRect(screen, float32(zoomPlusX), zoomPlusY, float32(zoomPlusWidth), zoomPlusHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "ZOOM PLUS", zoomPlusX, 10)

	vector.DrawFilledRect(screen, float32(zoomMinusX), zoomMinusY, float32(zoomMinusWidth), zoomMinusHeight, buttonColor, false)

	ebitenutil.DebugPrintAt(screen, "ZOOM MINUS", zoomMinusX, 10)

}

func updateGui() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		//start button
		if mx >= rectStartX && mx <= rectStartX+rectStartWidth && my >= rectStartY && my <= rectStartY+rectStartHeight {
			gameStarted = !gameStarted
		}
		//Export button
		if mx >= exportRectX && mx <= exportRectX+exportRectWidth && my >= exportRectY && my <= exportRectY+exportRectHeight {
			exportGrid()
		}
		//Load button
		if mx >= loadRectX && mx <= loadRectX+loadRectWidth && my >= loadRectY && my <= loadRectY+loadRectHeight {
			loadFromFile()
		}
		//Plus tick
		if mx >= rectPlusX && mx <= rectPlusX+rectPlusWidth && my >= rectPlusY && my <= rectPlusY+rectPlusHeight {
			timeTick += 1
		}
		//Minus tick
		if mx >= rectMinusX && mx <= rectMinusX+rectMinusWidth && my >= rectMinusY && my <= rectMinusY+rectMinusHeight {
			timeTick -= 1
		}
		//Zoom plus
		if mx >= zoomPlusX && mx <= zoomPlusX+zoomPlusWidth && my >= zoomPlusY && my <= zoomPlusY+zoomPlusHeight {
			gridSize = gridSize / 2
			tileSize = tileSize * 2
			initGrids()
		}
		//Zoom minus
		if mx >= zoomMinusX && mx <= zoomMinusX+zoomMinusWidth && my >= zoomMinusY && my <= zoomMinusY+zoomMinusHeight {
			gridSize = gridSize * 2
			tileSize = tileSize / 2
			initGrids()
		}
	}

}

func loadFromFile() {
	filePath := "export.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var content []byte
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		content = append(content, buffer[:n]...)
	}

	fileContent := string(content)
	fileIndex := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if fileContent[fileIndex] == 'A' {
				gridArray[i][j] = true
			} else {
				gridArray[i][j] = false
			}
			fileIndex++
		}
	}

}

func exportGrid() {
	file, err := os.Create("export.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for x := 0; x < gridSize; x++ {

		data := ""
		for y := 0; y < gridSize; y++ {
			if gridArray[x][y] {
				data += "A"
			} else {
				data += "D"
			}
		}
		byteSlice := []byte(data)
		_, err = file.Write(byteSlice)
	}

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data has been written to the file.")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	initGrids()
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Game of life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func initGrids() {

	var newGrid = make([][]bool, gridSize)
	var newTempGrid = make([][]bool, gridSize)

	for x := 0; x < gridSize; x++ {
		newGrid[x] = make([]bool, gridSize)
		newTempGrid[x] = make([]bool, gridSize)
		for y := 0; y < gridSize; y++ {
			newGrid[x][y] = false
			newTempGrid[x][y] = false
		}
	}

	gridArray = newGrid
	tempGrid = newTempGrid
}
