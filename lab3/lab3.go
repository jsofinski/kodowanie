package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hej")

	argsWithProg := os.Args
	var counter [256][2]int

	var f *os.File
	var size int64

	var encode = argsWithProg[1] == "encode"
	var inputFileName = argsWithProg[2]
	// var dictionaryFileName = argsWithProg[3]
	var outputFileName = argsWithProg[3]

	if encode { //encode
		fmt.Println("encode")
		f, size = getFile(inputFileName)
		byteFile := make([]byte, size)
		NObytes, err := f.Read(byteFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		encodeFile(inputFileName, outputFileName)
		//input data
		counter = getCounter(byteFile, NObytes)
		entropy := getEntropy(counter, byteFile, NObytes)
		fmt.Println("Entropia: \t\t", entropy)

	} else { //decode
		fmt.Println("decode")
		decodeFile(inputFileName, outputFileName)
	}
}
