package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	inputFileName := args[1]
	rbits, errr := strconv.Atoi(args[3])
	if errr != nil {
		panic(errr)
	}
	gbits, errg := strconv.Atoi(args[4])
	if errg != nil {
		panic(errg)
	}
	bbits, errb := strconv.Atoi(args[5])
	if errb != nil {
		panic(errb)
	}
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
	bitmap := makePixelMatrix(byteFile, int(width), int(height), rbits, gbits, bbits)

	printMSE(bitmap, byteFile)
	printSNR(bitmap, byteFile)

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
