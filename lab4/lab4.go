package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hej")
	args := os.Args
	inputFileName := args[1]
	// outputFileName := args[1]

	f, size := getFile(inputFileName)
	byteFile := make([]byte, size)
	NObytes, err := f.Read(byteFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(NObytes)

	counter := getCounter(byteFile, NObytes)
	entropy := getEntropy(counter, byteFile, NObytes)
	fmt.Println("Entropia: \t", entropy)

	var header = byteFile[:18]
	byteFile = byteFile[18:]
	var footer = byteFile[len(byteFile)-26:]
	// byteFile = byteFile[:len(byteFile)-26]

	fmt.Println("header len: ", len(header))
	fmt.Println("footer len: ", len(footer))

	var height = int64(header[15])*256 + int64(header[14])
	var width = int64(header[13])*256 + int64(header[12])
	fmt.Println("Width: ", width)
	fmt.Println("Height: ", height)

	output, err := os.Create(args[2])
	check(err)

	output.Write(header)

	// printData(byteFile, output)
	outputByteFile := makePixelMatrix(byteFile, int(width), int(height))
	output.Write(outputByteFile)
	output.Write(footer)

	counter = getCounter(outputByteFile, len(outputByteFile))
	entropy = getEntropy(counter, outputByteFile, len(outputByteFile))
	fmt.Println("Entropia output: \t", entropy)

	defer output.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
