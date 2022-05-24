package main

import (
	"fmt"
	"os"
)

type pixel struct {
	r int
	g int
	b int
}

func printData(byteFile []byte, output *os.File) {
	index := 0
	for i := 0; i < 512; i++ {
		fmt.Println("R: ", byteFile[index], "; G: ", byteFile[index+1], "; B: ", byteFile[index+2])
		index += 3
	}
}

func makePixelMatrix(byteFile []byte, width int, height int, rbits int, gbits int, bbits int) []byte {
	index := 0
	inputMatrix := make([][]pixel, height)
	outputMatrix := make([][]pixel, height)
	for i := range inputMatrix {
		inputMatrix[i] = make([]pixel, width)
		outputMatrix[i] = make([]pixel, width)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			inputMatrix[i][j] = pixel{r: int(byteFile[index+2]), g: int(byteFile[index+1]), b: int(byteFile[index])}
			index += 3
		}
	}

	index = 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			temp := inputMatrix[i][j]
			// outputMatrix[i][j] = pixel{r: 128, g: 128, b: 128}
			r := quantization(temp.r, rbits)
			g := quantization(temp.g, gbits)
			b := quantization(temp.b, bbits)
			outputMatrix[i][j] = pixel{r: r, g: g, b: b}
			// outputMatrix[i][j] = pixel{r: temp.r, g: 128, b: 128}
			// outputMatrix[i][j] = pixel{r: 128, g: temp.g, b: 128}
			// outputMatrix[i][j] = pixel{r: 128, g: 128, b: temp.b}
			// outputMatrix[i][j] = pixel{r: 128, g: temp.g, b: temp.b}
			// outputMatrix[i][j] = pixel{r: temp.r, g: 128, b: temp.b}
			// outputMatrix[i][j] = pixel{r: temp.r, g: temp.g, b: 128}

		}
	}

	byteOutput := make([]byte, len(byteFile))
	index = 0
	for i := range outputMatrix {
		for j := range outputMatrix[0] {
			byteOutput[index+2] = (byte(outputMatrix[i][j].r))
			byteOutput[index+1] = (byte(outputMatrix[i][j].g))
			byteOutput[index] = (byte(outputMatrix[i][j].b))
			index += 3
		}
	}
	return byteOutput

	// test1 := pixel{r: 12, g: 105, b: 201}
	// test2 := pixel{r: 24, g: 115, b: 181}
	// fmt.Println(pixelAdd(test1, test2, 256))
	// fmt.Println(pixelSub(test1, test2, 256))
	// fmt.Println((10 - 12) % 256)
	// fmt.Println(myMod((10 - 12), 256))

}
