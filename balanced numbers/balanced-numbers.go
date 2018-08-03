package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

func balancedNumbers(B, L int64, input []int64) []int64 {
	var N10 int64
	for i := int64(L - 1); i >= 0; i-- {
		N10 += input[i] * int64(math.Pow(float64(B), float64(L-1-i)))
	}
	fmt.Print(N10)

	var suma int64
	var count int64
	fmt.Print("\n")
	for i := int64(1); i <= N10; i++ {
		if isBalanced(i, B) {
			count++
			suma += i
		}

	}
	fmt.Print("\n")
	return []int64{count, suma}
}

func isBalanced(n int64, B int64) bool {
	if n < B {
		return true
	}
	var signTotal float64
	for i := float64(1); float64(n) > math.Pow(float64(B), i); i++ {
		signTotal = i + 1
	}
	mid := math.Ceil(signTotal / 2)
	//fmt.Print(signTotal)

	c := n
	var sign float64
	var balance int64

	for c > 0 {
		sign++
		r := int64(math.Mod(float64(c), float64(B)))
		c = int64(float64(c) / float64(B))

		if math.Mod(signTotal, float64(2)) != float64(0) {
			if sign < mid {
				balance += r
			} else if sign > mid {
				balance -= r
			}
		} else {
			if sign <= mid {
				balance += r
			} else if sign > mid {
				balance -= r
			}
		}
	}

	return balance == int64(0)
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// BL := strings.Split(readLine(reader), " ")

	// Btemp, err := strconv.ParseInt(BL[0], 10, 64)
	// checkError(err)
	// B := int32(Btemp)
	// fmt.Print("B: ", B, " ")

	// Ltemp, err := strconv.ParseInt(BL[1], 10, 64)
	// checkError(err)
	// L := int32(Ltemp)
	// fmt.Print("L: ", L, "\n")

	// inputTemp := strings.Split(readLine(reader), " ")

	// var input []int64

	// for i := 0; i < int(L); i++ {
	// 	inputItem, err := strconv.ParseInt(inputTemp[i], 10, 64)
	// 	checkError(err)
	// 	input = append(input, inputItem)
	// }
	// fmt.Print(input, "\n")

	// B := int64(10)
	// input := []int64{7}

	B := int64(10)
	input := []int64{4, 8, 5, 7}

	////-----------------------------
	// B := int64(11)
	// input := []int64{1, 10, 9}

	result := balancedNumbers(B, int64(len(input)), input)

	fmt.Print(result)
	////-----------------------------

	//fmt.Print(isBalanced(int64(122), int64(11)))
	//fmt.Print(isBalanced(int64(11), int64(11)))

	// strResult := strconv.FormatInt(result[0], 10) + " " + strconv.FormatInt(result[1], 10)

	// fmt.Fprintf(writer, "%s\n", strResult)

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
