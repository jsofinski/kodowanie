package main

import (
    "fmt"
    "os"
)

var counter[256][2]int;

func main() {
	fmt.Println("hello")
	argsWithProg := os.Args

	for i := 0; i < 256; i++ {
		counter[i][0] = i
		counter[i][1] = 0
		// fmt.Printf("%08b\n", i)
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

	for i := 0; i < n1; i++ {
		// fmt.Printf("%08b\n", b1[i])
		addByte(int(b1[i]))
	}


	defer f.Close()
	printCounter()
}

func addByte(byteVal int) {
	counter[byteVal][1] += 1
}

func printCounter() {
	for i := 0; i < 256; i++ {
		fmt.Printf("%08b		%d\n", counter[i][0], counter[i][1])
	}
}