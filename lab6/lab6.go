package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if args[1] == "compare" {
		compare(args[2], args[3])
		return
	}
	inputFileName := args[2]
	NObits, errr := strconv.Atoi(args[4])
	if errr != nil {
		panic(errr)
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

	var height = int64(header[15])*256 + int64(header[14])
	var width = int64(header[13])*256 + int64(header[12])
	// fmt.Println("Width: ", width)
	// fmt.Println("Height: ", height)

	output, err := os.Create(args[3])
	check(err)
	output.Write(header)

	bitmap := make([]byte, 0)
	if args[1] == "encode" {
		bitmap = encodePixelMatrix(byteFile, int(width), int(height), NObits)
		// printMSE(bitmap, byteFile)
		// printSNR(bitmap, byteFile)
	} else if args[1] == "decode" {
		bitmap = decodePixelMatrix(byteFile, int(width), int(height), NObits)
	}

	output.Write(bitmap)
	output.Write(footer)

	defer output.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
