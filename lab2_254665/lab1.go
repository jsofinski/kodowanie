// Przygotować jako bibliotekę
package main

import (
    "fmt"
    "os"
    "math"
)


func calc() {
	argsWithProg := os.Args
	var counter[256][2]int;
	var condCounter[256][256]int;
	for i := 0; i < 256; i++ {
		counter[i][0] = i
		counter[i][1] = 0
		// fmt.Printf("%08b\n", i)
	}
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			condCounter[i][j] = 0
		}
	}
	var f *os.File
	var size int64

	if len(argsWithProg) > 1 {
		f, size = getFile(argsWithProg[1])
	}

	b1 := make([]byte, size)
    n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}

    fmt.Printf("%d bytes:\n", n1)

	counter = getCounter(b1, n1)
	condCounter = getCondCounter(b1, n1)

	defer f.Close()
	entropy := getEntropy(counter, b1, n1)
	fmt.Println("Entropia: \t\t", entropy)
	condEntropy := getCondEntropy(counter, condCounter, b1, n1)
	fmt.Println("Entropia warunkowa:\t", condEntropy)
	fmt.Println("Różnica:\t\t", entropy - condEntropy)
}

func getFile(fileName string) (*os.File, int64) {
	var f *os.File
	var err error
	f, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
	  // Could not obtain stat, handle error
	  panic(err)
	}	
	var size = fi.Size()
	return f, size
}

func getCounter(b1 []byte, n1 int)([256][2]int) {
	var inCounter[256][2]int


	for i := 0; i < 256; i++ {
		inCounter[i][0] = i
		inCounter[i][1] = 0
	}

	for i := 0; i < n1; i++ {
		inCounter[b1[i]][1] += 1
	}
	return inCounter
}

func getCondCounter(b1 []byte, n1 int)([256][256]int) {
	var inCondCounter[256][256]int
	inCondCounter[0][b1[0]]++
	for i := 0; i < n1-1; i++ {
		inCondCounter[b1[i]][b1[i+1]]++
	}
	return inCondCounter
}

func printCounter(counter [256][2]int) {
	for i := 0; i < 256; i++ {
		fmt.Printf("%08b		%d\n", counter[i][0], counter[i][1])
	}
}

func printCond(condCounter [256][2]int) {
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			fmt.Printf("%d ", condCounter[i][j])
		}			
		fmt.Printf("\n")

	}
}

func getEntropy(counter [256][2]int, b1 []byte, n1 int) float64 {
	entropy := 0.0
	for i:= 0; i < 256; i++ {
		if (counter[i][1] != 0) {
			entropy += float64(counter[i][1]) * (math.Log2(float64(n1)) - math.Log2(float64(counter[i][1])))
		}
	}
	entropy = entropy/float64(n1)
	return entropy
}

func getCondEntropy(counter [256][2]int, condCounter [256][256]int, b1 []byte, n1 int) float64 {
	entropy := 0.0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			if (counter[i][1] != 0 && condCounter[i][j] != 0) {
				entropy += float64(condCounter[i][j]) * (math.Log2(float64(counter[i][1])) - math.Log2(float64(condCounter[i][j])))
			}		
		}		

	}
	for i:= 0; i < 256; i++ {
		
	}
	entropy = entropy/float64(n1)
	return entropy
}