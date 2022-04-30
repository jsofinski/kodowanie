package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

	var outputBuffer = ""
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
			tempIndex := getIndexFromDictionary(dictionary, pChar)
			tempString := strconv.FormatInt(int64(tempIndex), 2)
			tempString = rightjust(tempString, binaryLength, "0")
			outputBuffer = outputBuffer + tempString
			for len(outputBuffer)%8 == 0 && len(outputBuffer) > 0 {
				var b, err = strconv.ParseInt(outputBuffer[:8], 2, 9)
				if err != nil {
					panic(err)
				}
				// fmt.Println(outputBuffer)
				// fmt.Println(byte(b))
				// fmt.Print(binaryLength)
				// byte(b) save to outputFile
				if _, err := outputFile.Write([]byte{byte(b)}); err != nil {
					panic(err)
				}
				outputBuffer = outputBuffer[8:]
			}

			pChar = append(pChar, cChar)
			// fmt.Println(pChar)
			dictionary = append(dictionary, pChar)
			dictCounter++
			binaryLength = int(math.Ceil(math.Log2(float64(dictCounter))))
			pChar = make([]int, 0)
			pChar = append(pChar, cChar)
		}
	}

	// output buffer to outputFile
	// int -> string -> bytes

	// for i := 0; i < len(outputBuffer); i++ {
	// 	tempString := strconv.FormatInt(int64(outputBuffer[i]), 2)
	// 	tempString = rightjust(tempString, 9, "0")
	// 	fmt.Println(tempString)
	// }

	// fmt.Println(outputBuffer)
	// fmt.Println(dictionary)

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
