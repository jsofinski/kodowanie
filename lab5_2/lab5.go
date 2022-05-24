package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	inputFileName := args[1]
	bitsAmount, err := strconv.Atoi(args[3])
	if err != nil {
		panic(err)
	}
	selectionType := args[4]
	fmt.Println(selectionType)
	// outputFileName := args[1]

	f, size := getFile(inputFileName)
	byteFile := make([]byte, size)
	NObytes, err := f.Read(byteFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// fmt.Println(NObytes)

	counter := getCounter(byteFile, NObytes)
	entropy := getEntropy(counter, byteFile, NObytes)
	fmt.Println("Entropia: \t", entropy)

	var header = byteFile[:18]
	byteFile = byteFile[18:]
	var footer = byteFile[len(byteFile)-26:]
	// byteFile = byteFile[:len(byteFile)-26]

	// fmt.Println("header len: ", len(header))
	// fmt.Println("footer len: ", len(footer))

	var height = int64(header[15])*256 + int64(header[14])
	var width = int64(header[13])*256 + int64(header[12])
	// fmt.Println("Width: ", width)
	// fmt.Println("Height: ", height)

	output, err := os.Create(args[2])
	check(err)
	output.Write(header)

	bitmap := make([]byte, 0)

	optimalBits := [3]int{0, 0, 0}
	_ = optimalBits
	minmaxValue := -1.0

	for i := 0; i <= bitsAmount; i++ {
		for j := i; j <= bitsAmount; j++ {
			rbits := i
			gbits := j - i
			bbits := bitsAmount - j
			if rbits > 8 || gbits > 8 || bbits > 8 {
				continue
			}
			// fmt.Println("r: ", rbits, "; g: ", gbits, "; b: ", bbits)
			bitmap = makePixelMatrix(byteFile, int(width), int(height), rbits, gbits, bbits)

			if selectionType == "SNR" {
				temp := getMinSNR(bitmap, byteFile)
				if minmaxValue == -1 || minmaxValue < temp {
					minmaxValue = temp
					optimalBits = [3]int{rbits, gbits, bbits}
				}
			} else if selectionType == "MSE" {
				temp := getMaxMSE(bitmap, byteFile)
				if minmaxValue == -1 || minmaxValue > temp {
					minmaxValue = temp
					optimalBits = [3]int{rbits, gbits, bbits}
				}
			}
		}
	}
	fmt.Println("Optimal solution: ", optimalBits[0], " ", optimalBits[1], " ", optimalBits[2])
	bitmap = makePixelMatrix(byteFile, int(width), int(height), optimalBits[0], optimalBits[1], optimalBits[2])

	output.Write(bitmap)

	output.Write(footer)
	defer output.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 0
// 36.5
// 73
// 109.5
// 146
// 182.5
// 219
// 255.5
