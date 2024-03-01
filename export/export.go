package export

import (
	"fmt"
	"io"
	"os"

	"github.com/simonbredeche/simonbredeche/manager"
)

func ExportGrid(gameState *manager.GameState, gridManager *manager.GridManager) {
	file, err := os.Create("export.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for x := 0; x < gameState.GridSize; x++ {

		data := ""
		for y := 0; y < gameState.GridSize; y++ {
			if gridManager.GridArray[x][y] {
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

func LoadFromFile(gameState *manager.GameState, gridManager *manager.GridManager) {
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
	for i := 0; i < gameState.GridSize; i++ {
		for j := 0; j < gameState.GridSize; j++ {
			if fileContent[fileIndex] == 'A' {
				gridManager.GridArray[i][j] = true
			} else {
				gridManager.GridArray[i][j] = false
			}
			fileIndex++
		}
	}

}
