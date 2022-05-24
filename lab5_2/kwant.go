package main

import (
	"fmt"
	"math"
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
	// fmt.Println("cyk ", rbits, gbits, bbits)
	return byteOutput
}

func printMSE(bitmap []byte, byteFile []byte) {
	length := len(bitmap)
	sum := 0.0
	// for i := 10000; i < 10010; i++ {
	// 	fmt.Println(math.Pow(float64(int(bitmap[i])-int(byteFile[i])), 2))
	// }
	for i := 0; i < length; i++ {
		sum += math.Pow(float64(int(bitmap[i])-int(byteFile[i])), 2)
	}
	fmt.Println("mse:\t=", sum/float64(length))

	sum = 0
	for i := 0; i < length/3; i++ {
		// sum += math.Pow(float64(bitmap[i*3+2]-byteFile[i*3+2]), 2)
		sum += math.Pow(float64(int(bitmap[i*3+2])-int(byteFile[i*3+2])), 2)
	}
	fmt.Println("mse:(r)\t=", sum/float64(length/3))
	sum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(bitmap[i*3+1])-int(byteFile[i*3+1])), 2)
	}
	fmt.Println("mse:(g)\t=", sum/float64(length/3))
	sum = 0
	for i := 0; i < length/3; i++ {
		// sum += math.Pow(float64(bitmap[i*3]-byteFile[i*3]), 2)
		sum += math.Pow(float64(int(bitmap[i*3])-int(byteFile[i*3])), 2)
	}
	fmt.Println("mse:(b)\t=", sum/float64(length/3))
}
func printSNR(bitmap []byte, byteFile []byte) {
	length := len(bitmap)
	sum := 0.0
	mseSum := 0.0

	for i := 0; i < length; i++ {
		sum += math.Pow(float64(int(byteFile[i])), 2)
		mseSum += math.Pow(float64(int(bitmap[i])-int(byteFile[i])), 2)
	}
	dec := 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:\t=", sum/mseSum, "\t(", dec, "dB)")

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3+2])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3+2])-int(byteFile[i*3+2])), 2)
	}
	dec = 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:(r)\t=", sum/mseSum, "\t(", dec, "dB)")

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3+1])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3+1])-int(byteFile[i*3+1])), 2)
	}
	dec = 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:(g)\t=", sum/mseSum, "\t(", dec, "dB)")

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3])-int(byteFile[i*3])), 2)
	}
	dec = 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:(b)\t=", sum/mseSum, "\t(", dec, "dB)")
}

func getMaxMSE(bitmap []byte, byteFile []byte) float64 {
	length := len(bitmap)
	sum := 0.0

	sum = 0
	for i := 0; i < length/3; i++ {
		// sum += math.Pow(float64(bitmap[i*3+2]-byteFile[i*3+2]), 2)
		sum += math.Pow(float64(int(bitmap[i*3+2])-int(byteFile[i*3+2])), 2)
	}
	maxMSE := sum / float64(length/3)

	sum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(bitmap[i*3+1])-int(byteFile[i*3+1])), 2)
	}
	if sum/float64(length/3) > maxMSE {
		maxMSE = sum / float64(length/3)
	}

	sum = 0
	for i := 0; i < length/3; i++ {
		// sum += math.Pow(float64(bitmap[i*3]-byteFile[i*3]), 2)
		sum += math.Pow(float64(int(bitmap[i*3])-int(byteFile[i*3])), 2)
	}
	if sum/float64(length/3) > maxMSE {
		maxMSE = sum / float64(length/3)
	}
	return maxMSE
}

func getMinSNR(bitmap []byte, byteFile []byte) float64 {
	length := len(bitmap)
	sum := 0.0
	mseSum := 0.0

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3+2])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3+2])-int(byteFile[i*3+2])), 2)
	}
	minSNR := sum / mseSum

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3+1])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3+1])-int(byteFile[i*3+1])), 2)
	}
	if sum/mseSum < minSNR {
		minSNR = sum / mseSum
	}
	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3])-int(byteFile[i*3])), 2)
	}
	if sum/mseSum < minSNR {
		minSNR = sum / mseSum
	}
	return minSNR
}
