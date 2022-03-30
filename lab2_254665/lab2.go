package main

import (
    "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"bufio"
)

type probType struct {
    id int
    value float64
}

type dictType struct {
    id int
    value int
}

type HuffmanNode struct {
    freq int
    symbol int

	left *HuffmanNode
	right *HuffmanNode

	huff int
}

func main() {
	argsWithProg := os.Args
	var counter[256][2]int
	var probability [256]probType

	var f *os.File
	var size int64

	var encode = argsWithProg[1] == "encode"
	var inputFileName = argsWithProg[2]
	var dictionaryFileName = argsWithProg[3]
	var outputFileName = argsWithProg[4]

	
	if (encode) {
		f, size = getFile(inputFileName)
		b1 := make([]byte, size)
		n1, err := f.Read(b1)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		counter = getCounter(b1, n1)
		entropy := getEntropy(counter, b1, n1)
		fmt.Println("Entropia: \t\t", entropy)

		probability = getProbability(counter, n1)
		counter = sortCounter(counter)
		probability = sortProbability(probability)

		var newTree = getTree(counter)
		var dictionary [256]string
		fillDictionary(&dictionary, newTree, "")
		fmt.Println("Average code length: ", getAvgCodeLength(dictionary))

		dictionaryToFile(dictionary, dictionaryFileName)

		encodeFile(b1, size, dictionary, outputFileName)
	} else { //decode
		dictFromFile := dictionaryFromFile(dictionaryFileName) 
		decodeFile(inputFileName, dictFromFile, outputFileName)
	}
}


func getAvgCodeLength(dictionary [256]string)float64 {
	numberOfCodes := 0
	totalLength := 0
	for i := 0; i < 256; i++ {
		if (len(dictionary[i]) > 0) {
			numberOfCodes++
			totalLength += len(dictionary[i])
		}
	}
	return float64(totalLength) / float64(numberOfCodes)
}

func getProbability(counter [256][2]int, size int)([256]probType) {
	var probability[256]probType;
	for i := 0; i < 256; i++ {
		probability[i].id = i
		probability[i].value = float64(float64(counter[i][1])/float64(size))
	}
	return probability
}

func printProbability(probability [256]probType) {
	for i := 0; i < 256; i++ {
		fmt.Printf("%08b		%f\n", int(probability[i].id), probability[i].value)
	}
}

func sortProbability(probability [256]probType)([256]probType) {
	// fmt.Println("dupa1")
	sort.Slice(probability[:], func(i, j int) bool {
		return probability[i].value > probability[j].value
	})	
	// fmt.Println("dupa2")

	return probability
}

func sortCounter(counter [256][2]int)([256][2]int) {
	// fmt.Println("dupa1")
	sort.Slice(counter[:], func(i, j int) bool {
		return counter[i][1] > counter[j][1]
	})	
	// fmt.Println("dupa2")

	return counter
}
func sortNodes(nodes []HuffmanNode)([]HuffmanNode) {
	// fmt.Println("dupa1")
	sort.Slice(nodes[:], func(i, j int) bool {
		return nodes[i].freq > nodes[j].freq
	})	
	// fmt.Println("dupa2")

	return nodes
}

func getTree(counter [256][2]int)(HuffmanNode) {
	var symbolsNumber = 0
	var nodes []HuffmanNode
	for i := 0; i < 256; i++ {
		if counter[i][1] != 0 {
			symbolsNumber++;
			var newNode HuffmanNode
			newNode.freq = counter[i][1]
			newNode.symbol = counter[i][0]
			newNode.left = nil
			newNode.right = nil
			newNode.huff = -1
			nodes = append(nodes, newNode)
			// fmt.Println(counter[i][0])
		}
	}
	// fmt.Println(symbolsNumber)

	
	for {
		if (len(nodes) <= 1) {
			break
		}
		var lastNode = nodes[len(nodes)-1]
		var secondLastNode = nodes[len(nodes)-2]
		nodes = nodes[:len(nodes)-2]
		// fmt.Println(secondLastNode.symbol)

		lastNode.huff = 0
		secondLastNode.huff = 1

		var newNode HuffmanNode
		newNode.left = &lastNode
		newNode.right = &secondLastNode

		newNode.freq = lastNode.freq + secondLastNode.freq

		nodes = append(nodes, newNode)
		nodes = sortNodes(nodes)
		// fmt.Println(nodes)
		// fmt.Println()
	}

	return nodes[0]
} 

func printTree(node HuffmanNode, value string) {
	// fmt.Println((node.huff))

	if (node.right != nil) {
		printTree(*node.right, value + strconv.Itoa(node.right.huff))
	}
	if (node.left != nil) {
		printTree(*node.left, value + strconv.Itoa(node.left.huff))
	}

	if (node.left == nil && node.right == nil) {
		fmt.Println(node.symbol, " : ", value)
	}
}

func fillDictionary(dictionary *[256]string, node HuffmanNode, value string) {
	if (node.right != nil) {
		fillDictionary(dictionary, *node.right, value + strconv.Itoa(node.right.huff))
	}
	if (node.left != nil) {
		fillDictionary(dictionary, *node.left, value + strconv.Itoa(node.left.huff))
	}

	if (node.left == nil && node.right == nil) {
		dictionary[node.symbol] = value
	}
}

func encodeFile(bytes []byte, size int64, dictionary [256]string, outputFileName string) {
	file, err := os.Create(outputFileName)
	if err != nil {
        panic(err)
    }
	
	var currentString = ""
	var prefixString = ""
	var i int64
	for i = 0; i < size; i++ {
		// if (i%50000 == int64(size%50000)) {
		// 	fmt.Println(i)
		// }
		// fmt.Println(string(bytes[i]) + " : " + dictionary[bytes[i]])
		currentString += dictionary[bytes[i]]
		if (len(currentString) >= 8) {
			prefixString = currentString[0:8]
			currentString = currentString[8:]

			if i, err := strconv.ParseInt(prefixString, 2, 64); err != nil {
				fmt.Println(err)
			} else {
				// save next byte to file
				b1 := []byte{byte(i)}
				if _, err := file.Write(b1); err != nil {
					panic(err)
				}			
			}

			// fmt.Printf("%s ", prefixString)
		}
	}
	if (len(currentString) != 0) {
		if (len(currentString) < 8) {
			currentString = currentString + strings.Repeat("0", 8 - len(currentString))
		}
		if i, err := strconv.ParseInt(currentString, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			// save next byte to file
			b1 := []byte{byte(i)}
			if _, err := file.Write(b1); err != nil {
				panic(err)
			}			
		}
		// fmt.Println(currentString)
	}
	fmt.Println("Input size:  ", size)
	fileStat, err := file.Stat()
	if err != nil {
	  // Could not obtain stat, handle error
	  panic(err)
	}	
	outputSize := fileStat.Size()
	fmt.Println("Output size: ", outputSize)
	fmt.Println("Compression rate: ", float64(size)/float64(outputSize))

	defer file.Close()

	// fmt.Println(stringValue)
}

func decodeFile(inputFileName string, dictionary [256]string, outputFileName string) {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
        panic(err)
    }


	intputFile, size := getFile(inputFileName)
	bytes := make([]byte, size)
    _, err = intputFile.Read(bytes)
	if err != nil {
		panic(err)
	}
    // fmt.Printf("%d bytes:\n", n1)

	var minLength = 128
	var maxLength = 0

	for i := 0; i < 256; i++ {
		currentLength := len(dictionary[i])
		if (currentLength == 0) {
			continue
		}
		if (currentLength < minLength) {
			minLength = currentLength
		} 
		if (currentLength > maxLength) {
			maxLength = currentLength
		}
	}
	// fmt.Println("min: ", minLength)
	// fmt.Println("max: ", maxLength)
	
	var currentString = ""
	// var prefixString = ""
	var i int64
	for i = 0; i < size; i++ {

		// fmt.Println("currentString still left: ", currentString)
		// if (i%50000 == int64(size%50000)) {
		// 	fmt.Println(i)
		// }
		

		newString := strconv.FormatInt(int64(bytes[i]), 2)
		if (len(newString) < 8) {
			// fmt.Println(len(newString))
			newString = strings.Repeat("0", 8 - len(newString)) + newString
		}
		currentString = currentString + newString
		// fmt.Println(currentString)

		j := 0
		currentCode := ""
		for {
			if (len(currentString) < minLength) {
				break
			}
			if (j >= len(currentString)) {
				break
			}
			currentCode = currentCode + string((currentString)[j])
			// fmt.Println(currentCode)
			j++
			
			if (len(currentCode) >= minLength) {
				for k := 0; k < 256; k++ {
					if (len(dictionary[k]) == 0) {
						continue
					}
					if (strings.Compare(currentCode, dictionary[k]) == 0) {
						// code found in dictionary
						// fmt.Println("found: ", currentCode, " : ", k)
						// save to file
						b1 := []byte{byte(k)}
						if _, err := outputFile.Write(b1); err != nil {
							panic(err)
						}	
						currentString = currentString[len(currentCode):]
						// fmt.Printf("%08b: %c\n", k, k)
						// fmt.Println("currentString left: ", currentString)
						currentCode = ""
						j = 0
						break
					}
				}
			}
		}

		// fmt.Println(currentString)
		// fmt.Println(len(currentString))
		
	}
	// fmt.Println(stringValue)
}


func dictionaryToFile(dictionary [256]string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
        panic(err)
    }
	for i := 0; i < 256; i++ {
		_, err := file.WriteString(dictionary[i] + "\n")

		if err != nil {
			panic(err)
		}
	}
	defer file.Close()
}

func dictionaryFromFile(fileName string)([256]string) {
	file, err := os.Open(fileName)
	if err != nil {
        panic(err)
    }
	var dictionary [256]string
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
        dictionary[i] = scanner.Text()
		i++
    }
	if (len(dictionary) != 256) {
		fmt.Println("File length error")
	}
	defer file.Close()
	return dictionary
}