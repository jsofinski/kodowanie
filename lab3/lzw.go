package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var MaximumBits = 12

type buffer struct {
	value  int
	length int
}

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
	// fmt.Println(binaryLength)
	// fmt.Println(dictionary)
	var pChar = make([]int, 0)
	var tempArray = make([]int, 0)
	var cChar = 0

	var outputBufferInt = make([]buffer, 0)

	for i := 0; i < int(size); i++ {
		if i%10000 == 0 {
			fmt.Println(i)
			fmt.Println(binaryLength)
		}
		cChar = int(byteFile[i])
		// fmt.Println("save next byte: ", cChar)
		tempArray = make([]int, len(pChar))
		copy(tempArray, pChar)
		tempArray = append(tempArray, cChar)
		// fmt.Println(tempArray)
		if arrayContains(dictionary, tempArray) {
			pChar = append(pChar, cChar)
		} else {
			// output P from dictionary
			tempIndex := getIndexFromDictionary(dictionary, pChar)
			outputBufferInt = append(outputBufferInt, buffer{value: tempIndex, length: binaryLength})

			for sumOfLengths(outputBufferInt)%8 == 0 && sumOfLengths(outputBufferInt) > 0 {
				saveIntBufferToOutput(outputBufferInt, outputFile)
				outputBufferInt = make([]buffer, 0)
			}

			pChar = append(pChar, cChar)
			// fmt.Println(pChar)
			if binaryLength <= MaximumBits {
				dictionary = append(dictionary, pChar)
				dictCounter++
			}
			binaryLength = int(math.Ceil(math.Log2(float64(dictCounter))))
			pChar = make([]int, 0)
			pChar = append(pChar, cChar)
		}
	}
	// fmt.Println(outputBuffer)
	// Save the rest of outputBuffer to outputFile
	for sumOfLengths(outputBufferInt) > 0 {
		saveIntBufferToOutput(outputBufferInt, outputFile)
		outputBufferInt = make([]buffer, 0)
	}

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
}

func decodeFile(inputFileName string, outputFileName string) {
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

	var dictionary = make([][]int, 0)
	var dictCounter = 256 + 2
	var binaryLength = int(math.Ceil(math.Log2(float64(dictCounter))))
	// dictionary init
	for i := 0; i < 256; i++ {
		temp := make([]int, 0)
		temp = append(temp, i)
		dictionary = append(dictionary, temp)
	}

	var old = -1
	var new = -1
	// variable names s, c from LZW pseudocode
	var s = make([]int, 0)
	var c = make([]int, 0)
	var lastInputIndex = 0
	var inputBufferMinLength = 9
	var inputStringBuffer = ""
	for len(inputStringBuffer) < inputBufferMinLength {
		inputStringBuffer += getNextByteString(byteFile, lastInputIndex)
		lastInputIndex++
	}
	var tempByte int64
	tempByte, err = strconv.ParseInt(inputStringBuffer[:binaryLength], 2, 9)
	inputStringBuffer = inputStringBuffer[binaryLength:]
	if err != nil {
		panic(err)
	}
	old = int(tempByte)
	// ouput translation of Old; or dont?
	if _, err := outputFile.Write([]byte{byte(old)}); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	endOfDictionary := false

	for int64(lastInputIndex) < size {
		// make sure there at least binaryLength buffer in inputStringBuffer
		for len(inputStringBuffer) < binaryLength {
			inputStringBuffer += getNextByteString(byteFile, lastInputIndex)
			lastInputIndex += 1
		}
		// fmt.Println(inputStringBuffer, " ", lastInputIndex)
		tempByte, err = strconv.ParseInt(inputStringBuffer[:binaryLength], 2, 32)
		// fmt.Println(tempByte)

		inputStringBuffer = inputStringBuffer[binaryLength:]
		if err != nil {
			panic(err)
		}
		new = int(tempByte)

		// if !arrayContains(dictionary, new) {
		if new > dictCounter {
			// s = translation of old
			s = []int{}
			// fmt.Println("looking in dict: ", old)
			s = dictionary[old]
			// ... allows to pass multiple values to append(), nice
			s = append(s, c...)
		} else {
			// s = translation of new
			s = []int{}
			// fmt.Println("dict len: ", len(dictionary))
			if new == len(dictionary) {
				s = append(s, dictionary[old]...)
				s = append(s, dictionary[old][0])
			} else {
				s = append(s, dictionary[new]...)
			}
		}

		// output S
		// fmt.Println("Output: ", s)

		if _, err := outputFile.Write(intArrayToByteArray(s)); err != nil {
			panic(err)
		}
		// c = first char of S
		c = s[:1]
		// add old + c to stringTable
		tempOldC := make([]int, 0)
		tempOldC = append(tempOldC, dictionary[old]...)
		tempOldC = append(tempOldC, c...)

		if !endOfDictionary {
			dictionary = append(dictionary, tempOldC)
			dictCounter++
		}
		if binaryLength > MaximumBits {
			endOfDictionary = true
		}
		binaryLength = int(math.Ceil(math.Log2(float64(dictCounter))))
		// fmt.Println(binaryLength)

		old = new
	}

	// fmt.Println(dictionary)

	defer outputFile.Close()
}

func getNextByteString(byteFile []byte, lastInputIndex int) string {
	tempByte := byteFile[lastInputIndex]
	tempStringByte := strconv.FormatInt(int64(tempByte), 2)
	tempStringByte = rightjust(tempStringByte, 8, "0")
	return tempStringByte
}

func saveIntBufferToOutput(buffer []buffer, outputFile *os.File) {
	totalLength := sumOfLengths(buffer)
	bytes := make([]byte, totalLength/8)
	if totalLength%8 != 0 {
		// fmt.Println("DUPA")
		bytes = append(bytes, 0)
	}

	currentByte := 0
	currentByteIndex := 0
	for i := 0; i < len(buffer); i++ {
		for j := buffer[i].length - 1; j >= 0; j-- {
			if hasBit(buffer[i].value, uint(j)) {
				bytes[currentByte] = byte(setBit(int(bytes[currentByte]), 7-uint(currentByteIndex)))
			}
			currentByteIndex += 1
			if currentByteIndex%8 == 0 {
				currentByteIndex = 0
				currentByte += 1
			}
		}
	}

	fmt.Println("bytes:")
	fmt.Println(bytes)
	if _, err := outputFile.Write(bytes); err != nil {
		panic(err)
	}
}

func getIntBufferToOutput(buffer []buffer, inputFile *os.File) {
	totalLength := sumOfLengths(buffer)
	bytes := make([]byte, totalLength/8)
	if totalLength%8 != 0 {
		// fmt.Println("DUPA")
		bytes = append(bytes, 0)
	}

	currentByte := 0
	currentByteIndex := 0
	for i := 0; i < len(buffer); i++ {
		for j := buffer[i].length - 1; j >= 0; j-- {
			if hasBit(buffer[i].value, uint(j)) {
				bytes[currentByte] = byte(setBit(int(bytes[currentByte]), 7-uint(currentByteIndex)))
			}
			currentByteIndex += 1
			if currentByteIndex%8 == 0 {
				currentByteIndex = 0
				currentByte += 1
			}
		}
	}

	fmt.Println("bytes:")
	fmt.Println(bytes)
	if _, err := outputFile.Write(bytes); err != nil {
		panic(err)
	}
}
