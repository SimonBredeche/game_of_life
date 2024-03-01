package export

import (
	"fmt"
	"github.com/simonbredeche/simonbredeche/shared"
	"io"
	"os"
)

func ExportGrid(gridArray *[][]bool) {
	file, err := os.Create("export.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for x := 0; x < shared.GridSize; x++ {

		data := ""
		for y := 0; y < shared.GridSize; y++ {
			if (*gridArray)[x][y] {
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

func LoadFromFile(gridArray *[][]bool) {
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
	for i := 0; i < shared.GridSize; i++ {
		for j := 0; j < shared.GridSize; j++ {
			if fileContent[fileIndex] == 'A' {
				(*gridArray)[i][j] = true
			} else {
				(*gridArray)[i][j] = false
			}
			fileIndex++
		}
	}

}
