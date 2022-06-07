// Przygotować jako bibliotekę
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func sortCounter(counter [256][2]int) [256][2]int {
	sort.Slice(counter[:], func(i, j int) bool {
		return counter[i][1] > counter[j][1]
	})

	return counter
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

func getCounter(b1 []byte, n1 int) [256][2]int {
	var inCounter [256][2]int

	for i := 0; i < 256; i++ {
		inCounter[i][0] = i
		inCounter[i][1] = 0
	}

	for i := 0; i < n1; i++ {
		inCounter[b1[i]][1] += 1
	}
	return inCounter
}

func getCondCounter(b1 []byte, n1 int) [256][256]int {
	var inCondCounter [256][256]int
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
	for i := 0; i < 256; i++ {
		if counter[i][1] != 0 {
			entropy += float64(counter[i][1]) * (math.Log2(float64(n1)) - math.Log2(float64(counter[i][1])))
		}
	}
	entropy = entropy / float64(n1)
	return entropy
}

func getCondEntropy(counter [256][2]int, condCounter [256][256]int, b1 []byte, n1 int) float64 {
	entropy := 0.0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			if counter[i][1] != 0 && condCounter[i][j] != 0 {
				entropy += float64(condCounter[i][j]) * (math.Log2(float64(counter[i][1])) - math.Log2(float64(condCounter[i][j])))
			}
		}

	}
	for i := 0; i < 256; i++ {

	}
	entropy = entropy / float64(n1)
	return entropy
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func arrayContains(matrix [][]int, array []int) bool {
	// fmt.Println(matrix)
	// fmt.Println(array)
	for _, a := range matrix {
		if compareArrays(a, array) {
			// fmt.Println("true")
			return true
		}
	}
	// fmt.Println("false")
	return false
}

func compareArrays(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func intArrayToByteArray(intArray []int) []byte {
	tempByteArray := make([]byte, 0)
	for i := 0; i < len(intArray); i++ {
		tempByteArray = append(tempByteArray, byte(intArray[i]))
	}
	return tempByteArray
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func pixelAdd(a pixel, b pixel, mod int) pixel {
	return pixel{r: (a.r + b.r) % mod, g: (a.g + b.g) % mod, b: (a.b + b.b) % mod}
}

func pixelSub(a pixel, b pixel, mod int) pixel {
	return pixel{r: myMod((a.r - b.r), mod), g: myMod((a.g - b.g), mod), b: myMod((a.b - b.b), mod)}
}

func myMod(x int, y int) int {
	var res int = x % y
	if (res < 0 && y > 0) || (res > 0 && y < 0) {
		return res + y
	}
	return res
}

func pixelDiv2(a pixel) pixel {
	return pixel{r: a.r / 2, g: a.g / 2, b: a.b / 2}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxPixel(x, y pixel) pixel {
	sumX := pixelSum(x)
	sumY := pixelSum(y)
	if sumX > sumY {
		return x
	} else {
		return y
	}
}

func minPixel(x, y pixel) pixel {
	sumX := pixelSum(x)
	sumY := pixelSum(y)
	if sumX > sumY {
		return y
	} else {
		return x
	}

}

func pixelSum(x pixel) int {
	return (x.r + x.g + x.b)
}

func pixelAvg(x pixel, y pixel) pixel {
	return pixel{(x.r + y.r) / 2, (x.g + y.g) / 2, (x.b + y.b) / 2}
}

func pixelDiff(x pixel, y pixel) pixel {
	return pixel{x.r - y.r, x.g - y.g, x.b - y.b}
}

func quantization(input int, bits int) int {
	if bits == 0 {
		return 128
	}
	temp := float64(input) / (255 / (math.Pow(2, float64(bits)) - 1))
	temp = math.Round(temp)

	// fmt.Println(temp)

	// fmt.Println(255 / (math.Pow(2, float64(bits)) - 1))
	temp = temp * (255 / (math.Pow(2, float64(bits)) - 1))
	temp = math.Round(temp)
	return int(temp)
}

func quantizationEncode(input int, bits int) int {
	if bits == 0 {
		return 128
	}
	temp := float64(input) / (255 / (math.Pow(2, float64(bits)) - 1))
	temp = math.Round(temp)
	return int(temp)
}

func quantizationDecode(input int, bits int) int {
	if bits == 0 {
		return 128
	}
	temp := float64(input) * (255 / (math.Pow(2, float64(bits)) - 1))
	temp = math.Round(temp)
	return int(temp)
}

func normalizePixel(x pixel) pixel {
	if x.r < 0 {
		// fmt.Println(x.r)
		x.r = 0
	}
	if x.g < 0 {
		// fmt.Println(x.g)
		x.g = 0
	}
	if x.b < 0 {
		// fmt.Println(x.b)
		x.b = 0
	}

	if x.r >= 256 {
		// fmt.Println(x.r)
		x.r = 255
	}
	if x.g >= 256 {
		// fmt.Println(x.g)
		x.g = 255
	}
	if x.b >= 256 {
		// fmt.Println(x.b)
		x.b = 255
	}
	return x
}

func printMSE(bitmap []byte, byteFile []byte) {
	length := len(bitmap)
	sum := 0.0
	// for i := 10000; i < 10010; i++ {
	// 	fmt.Println(math.Pow(float64(int(bitmap[i])-int(byteFile[i])), 2))
	// }
	for i := 0; i < length; i++ {
		sum += math.Pow(float64(int(bitmap[i])-int(byteFile[i])), 2)
	}
	fmt.Println("mse:\t=", sum/float64(length))

	sum = 0
	for i := 0; i < length/3; i++ {
		// sum += math.Pow(float64(bitmap[i*3+2]-byteFile[i*3+2]), 2)
		sum += math.Pow(float64(int(bitmap[i*3+2])-int(byteFile[i*3+2])), 2)
	}
	fmt.Println("mse:(r)\t=", sum/float64(length/3))
	sum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(bitmap[i*3+1])-int(byteFile[i*3+1])), 2)
	}
	fmt.Println("mse:(g)\t=", sum/float64(length/3))
	sum = 0
	for i := 0; i < length/3; i++ {
		// sum += math.Pow(float64(bitmap[i*3]-byteFile[i*3]), 2)
		sum += math.Pow(float64(int(bitmap[i*3])-int(byteFile[i*3])), 2)
	}
	fmt.Println("mse:(b)\t=", sum/float64(length/3))
}

func printSNR(bitmap []byte, byteFile []byte) {
	length := len(bitmap)
	sum := 0.0
	mseSum := 0.0

	for i := 0; i < length; i++ {
		sum += math.Pow(float64(int(byteFile[i])), 2)
		mseSum += math.Pow(float64(int(bitmap[i])-int(byteFile[i])), 2)
	}
	dec := 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:\t=", sum/mseSum, "\t(", dec, "dB)")

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3+2])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3+2])-int(byteFile[i*3+2])), 2)
	}
	dec = 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:(r)\t=", sum/mseSum, "\t(", dec, "dB)")

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3+1])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3+1])-int(byteFile[i*3+1])), 2)
	}
	dec = 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:(g)\t=", sum/mseSum, "\t(", dec, "dB)")

	sum = 0
	mseSum = 0
	for i := 0; i < length/3; i++ {
		sum += math.Pow(float64(int(byteFile[i*3])), 2)
		mseSum += math.Pow(float64(int(bitmap[i*3])-int(byteFile[i*3])), 2)
	}
	dec = 10 * math.Log10(sum/mseSum)
	fmt.Println("SNR:(b)\t=", sum/mseSum, "\t(", dec, "dB)")
}
