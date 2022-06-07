package maindupa

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

func makePixelMatrix(byteFile []byte, width int, height int, NObits int) []byte {
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
	lastPixel := inputMatrix[0][0]
	avgPixelArray := make([]pixel, width*height/2)
	diffPixelArray := make([]pixel, width*height/2)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			temp := inputMatrix[i][j]

			if index%2 == 0 {
				lastPixel = temp
			} else {
				avgPixel := pixelAvg(lastPixel, temp)
				diffPixel := pixelDiff(avgPixel, lastPixel)
				avgPixelArray[index] = avgPixel
				diffPixelArray[index] = diffPixel
				index++
			}
			// outputMatrix[i][j] = pixel{r: 128, g: 128, b: 128}
			tempVal := quantizationEncode(temp.r, NObits)
			r := quantizationDecode(tempVal, NObits)
			tempVal = quantizationEncode(temp.g, NObits)
			g := quantizationDecode(tempVal, NObits)
			tempVal = quantizationEncode(temp.b, NObits)
			b := quantizationDecode(tempVal, NObits)
			outputMatrix[i][j] = pixel{r: r, g: g, b: b}
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
}

func encodePixelMatrix(byteFile []byte, width int, height int, NObits int) []byte {
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
	lastR := 0
	lastG := 0
	lastB := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			temp := inputMatrix[i][j]
			// outputMatrix[i][j] = pixel{r: 128, g: 128, b: 128}
			r := quantizationEncode(temp.r, NObits)
			g := quantizationEncode(temp.g, NObits)
			b := quantizationEncode(temp.b, NObits)
			outputMatrix[i][j] = pixel{r: r - lastR, g: g - lastG, b: b - lastB}
			lastR = r
			lastG = g
			lastB = b

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
}
func decodePixelMatrix(byteFile []byte, width int, height int, NObits int) []byte {
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
	lastR := 0
	lastG := 0
	lastB := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			temp := inputMatrix[i][j]
			// outputMatrix[i][j] = pixel{r: 128, g: 128, b: 128}
			r := quantizationDecode(temp.r+lastR, NObits)
			g := quantizationDecode(temp.g+lastG, NObits)
			b := quantizationDecode(temp.b+lastB, NObits)
			outputMatrix[i][j] = pixel{r: r, g: g, b: b}
			lastR = temp.r + lastR
			lastG = temp.g + lastG
			lastB = temp.b + lastB
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
}
