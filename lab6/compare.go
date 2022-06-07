package main

import "fmt"

func compare(inputFileName string, outputFileName string) {
	input, sizeI := getFile(inputFileName)
	inputByteFile := make([]byte, sizeI)
	NObytesI, err := input.Read(inputByteFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	counter := getCounter(inputByteFile, NObytesI)
	entropy := getEntropy(counter, inputByteFile, NObytesI)
	fmt.Println("Input entrophy: \t", entropy)

	output, sizeO := getFile(outputFileName)
	outputByteFile := make([]byte, sizeO)
	NObytesO, err := output.Read(outputByteFile)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	counter = getCounter(outputByteFile, NObytesO)
	entropy = getEntropy(counter, outputByteFile, NObytesO)
	fmt.Println("Output entrophy: \t", entropy)

	printMSE(inputByteFile, outputByteFile)
	printSNR(inputByteFile, outputByteFile)
}
