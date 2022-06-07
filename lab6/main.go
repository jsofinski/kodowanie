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
				// diffPixel := pixelDiff(avgPixel, lastPixel)
				avgPixelArray[(index-1)/2] = avgPixel
				r := quantizationEncode(avgPixel.r-lastPixel.r, NObits)
				g := quantizationEncode(avgPixel.g-lastPixel.g, NObits)
				b := quantizationEncode(avgPixel.b-lastPixel.b, NObits)
				diffPixelArray[(index-1)/2] = pixel{r: r, g: g, b: b}
			}
			index++
		}
	}

	byteOutput := make([]byte, len(byteFile))
	index = 0
	lastAvgPixel := pixel{0, 0, 0}
	for i := 0; i < len(avgPixelArray); i++ {
		//dolnoprzepustowy, srednia roznicowo
		// fmt.Println(lastAvgPixel)
		// fmt.Println(avgPixelArray[i])
		byteOutput[index+2] = (byte(avgPixelArray[i].r - lastAvgPixel.r))
		byteOutput[index+1] = (byte(avgPixelArray[i].g - lastAvgPixel.g))
		byteOutput[index] = (byte(avgPixelArray[i].b - lastAvgPixel.b))

		// gornoprzepustowy, odchylenie za pomocoa kwantyzacji rownomiernej
		byteOutput[index+5] = (byte(diffPixelArray[i].r))
		byteOutput[index+4] = (byte(diffPixelArray[i].g))
		byteOutput[index+3] = (byte(diffPixelArray[i].b))
		// fmt.Println((diffPixelArray[i]))
		index += 6
		lastAvgPixel.r = avgPixelArray[i].r
		lastAvgPixel.g = avgPixelArray[i].g
		lastAvgPixel.b = avgPixelArray[i].b

		// fmt.Println(diffPixelArray[i])
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
	arrayIndex := 0
	index = 0
	tempAvgPixel := pixel{0, 0, 0}
	lastAvgPixel := pixel{0, 0, 0}
	tempDiffPixel := pixel{0, 0, 0}
	byteOutput := make([]byte, len(byteFile))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if arrayIndex%2 == 0 {
				r := int(byteFile[index+2])
				g := int(byteFile[index+1])
				b := int(byteFile[index])
				if lastAvgPixel.r+r >= 256 {
					r -= 256
				}
				if lastAvgPixel.g+g >= 256 {
					g -= 256
				}
				if lastAvgPixel.b+b >= 256 {
					b -= 256
				}
				tempAvgPixel = pixel{r: lastAvgPixel.r + r, g: lastAvgPixel.g + g, b: lastAvgPixel.b + b}
				// fmt.Println(tempAvgPixel)
			} else {
				tempR := int(byteFile[index+2])
				tempG := int(byteFile[index+1])
				tempB := int(byteFile[index])
				// fmt.Println(tempR)
				if tempR > 128 {
					tempR -= 256
				}
				if tempG > 128 {
					tempG -= 256
				}
				if tempB > 128 {
					tempB -= 256
				}
				r := quantizationDecode(tempR, NObits)
				g := quantizationDecode(tempG, NObits)
				b := quantizationDecode(tempB, NObits)
				tempDiffPixel = pixel{r: r, g: g, b: b}
				// fmt.Println(tempDiffPixel)

				// fmt.Println(int(byteFile[index+2]))
				// tempSecondPixel
				tempFirstPixel := pixel{tempAvgPixel.r - tempDiffPixel.r, tempAvgPixel.g - tempDiffPixel.g, tempAvgPixel.b - tempDiffPixel.b}
				tempSecondPixel := pixel{tempAvgPixel.r + tempDiffPixel.r, tempAvgPixel.g + tempDiffPixel.g, tempAvgPixel.b + tempDiffPixel.b}
				// tempFirstPixel = pixel{r: tempAvgPixel.r, g: tempAvgPixel.g, b: tempAvgPixel.b}
				// tempSecondPixel = pixel{r: tempAvgPixel.r, g: tempAvgPixel.g, b: tempAvgPixel.b}
				tempFirstPixel = normalizePixel(tempFirstPixel)
				tempSecondPixel = normalizePixel(tempSecondPixel)

				// byteOutput[index+2] = (byte(tempFirstPixel.r))
				// byteOutput[index+1] = (byte(tempFirstPixel.g))
				// byteOutput[index] = (byte(tempFirstPixel.b))

				// byteOutput[index+5] = (byte(tempSecondPixel.r))
				// byteOutput[index+4] = (byte(tempSecondPixel.g))
				// byteOutput[index+3] = (byte(tempSecondPixel.b))

				byteOutput[index-1] = (byte(tempFirstPixel.r))
				byteOutput[index-2] = (byte(tempFirstPixel.g))
				byteOutput[index-3] = (byte(tempFirstPixel.b))

				byteOutput[index+2] = (byte(tempSecondPixel.r))
				byteOutput[index+1] = (byte(tempSecondPixel.g))
				byteOutput[index] = (byte(tempSecondPixel.b))

				lastAvgPixel = tempAvgPixel
			}
			index += 3
			arrayIndex += 1
		}
	}

	return byteOutput
}
