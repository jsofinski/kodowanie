package main

import (
	"fmt"
	"math"
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
	var dictCounter = 256 + 1
	var binaryLength = int(math.Ceil(math.Log2(float64(dictCounter))))
	for i := 0; i < 256; i++ {
		temp := make([]int, 0)
		temp = append(temp, i)
		dictionary = append(dictionary, temp)
	}
	fmt.Println(binaryLength)
	fmt.Println(dictionary)
	var pChar = make([]int, 0)
	var tempArray = make([]int, 0)
	var cChar = 0

	var outputBuffer = make([]int, 0)
	for i := 0; i < int(size); i++ {
		// fmt.Println(pChar)
		cChar = int(byteFile[i])
		tempArray = make([]int, len(pChar))
		copy(tempArray, pChar)
		tempArray = append(tempArray, cChar)
		// fmt.Println(tempArray)
		if arrayContains(dictionary, tempArray) {
			pChar = append(pChar, cChar)

		} else {
			// output P from dictionary
			outputBuffer = append(outputBuffer, getIndexFromDictionary(dictionary, pChar))
			fmt.Println(getIndexFromDictionary(dictionary, pChar))
			pChar = append(pChar, cChar)
			// fmt.Println(pChar)
			dictionary = append(dictionary, pChar)
			dictCounter++
			binaryLength = int(math.Ceil(math.Log2(float64(dictCounter))))
			pChar = make([]int, 0)
			pChar = append(pChar, cChar)
		}
	}

	fmt.Println(dictionary)

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
