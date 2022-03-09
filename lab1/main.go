package main

import (
    "fmt"
    "os"
    "math"
)

var counter[256][2]int;
var condCounter[256][256]int;

func main() {
	fmt.Println("hello")
	argsWithProg := os.Args

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

	f, err := os.Open("test.txt")
	if len(argsWithProg) > 1 {
		f, err = os.Open(argsWithProg[1])
	}
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
	  // Could not obtain stat, handle error
	}	
	var size = fi.Size()

	b1 := make([]byte, size)
    n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}

    fmt.Printf("%d bytes:\n", n1)

	fillCounter(b1, n1)
	fillCondCounter(b1, n1)

	defer f.Close()
	entropy := getEntropy(b1, n1)
	fmt.Println(entropy)
	condEntropy := getCondEntropy(b1, n1)
	fmt.Println(condEntropy)
}

func fillCounter(b1 []byte, n1 int) {
	for i := 0; i < n1; i++ {
		addByte(int(b1[i]))
	}
}

func fillCondCounter(b1 []byte, n1 int) {
	condCounter[0][b1[0]]++
	for i := 0; i < n1-1; i++ {
		condCounter[b1[i]][b1[i+1]]++
	}
}

func addByte(byteVal int) {
	counter[byteVal][1] += 1
}

func printCounter() {
	for i := 0; i < 256; i++ {
		fmt.Printf("%08b		%d\n", counter[i][0], counter[i][1])
	}
}
func printCond() {
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			fmt.Printf("%d ", condCounter[i][j])
		}			
		fmt.Printf("\n")

	}
}

func getEntropy(b1 []byte, n1 int) float64 {
	entropy := 0.0
	for i:= 0; i < 256; i++ {
		if (counter[i][1] != 0) {
			entropy += float64(counter[i][1]) * (math.Log2(float64(n1)) - math.Log2(float64(counter[i][1])))
		}
	}
	entropy = entropy/float64(n1)
	return entropy
}

func getCondEntropy(b1 []byte, n1 int) float64 {
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