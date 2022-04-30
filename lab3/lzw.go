package main

import (
	"fmt"
	"os"
)

func test() {
	fmt.Println("test")
}

func encodeFile(inputFileName string, outputFileName string) {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	intputFile, size := getFile(inputFileName)
	byteFile := make([]byte, size)
	_, err = intputFile.Read(byteFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Input size:  ", size, len(byteFile))
	// code

	var dictionary = make([][]int, 0)

	for i := 0; i < 256; i++ {
		temp := make([]int, 0)
		temp = append(temp, i)
		dictionary = append(dictionary, temp)
	}

	// fmt.Println(dictionary)
	// for i := 0; i < int(size); i++ {
	// 	fmt.Println(byteFile[i])
	// }

	// print info
	fileStat, err := outputFile.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		panic(err)
	}
	outputSize := fileStat.Size()
	fmt.Println("Output size: ", outputSize)
	fmt.Println("Compression rate: ", float64(size)/float64(outputSize))

	defer outputFile.Close()
	defer intputFile.Close()

	// fmt.Println(stringValue)
}

func decodeFile(inputFileName string, outputFileName string) {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}

	intputFile, size := getFile(inputFileName)
	bytes := make([]byte, size)
	NObytes, err := intputFile.Read(bytes)

	var dictionary = make([][]int, 0)

	for i := 0; i < 256; i++ {
		temp := make([]int, 0)
		temp = append(temp, i)
		dictionary = append(dictionary, temp)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(NObytes)

	defer outputFile.Close()
}
