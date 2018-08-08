package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(P []int) float64 {
	if sort.IntsAreSorted(P) {
		return 0
	}

	sort.Ints(P)
	N := len(P)
	n := factorial(float64(N))
	div := float64(1)
	sameValues := 1
	for i := 0; i < N-1; i++ {
		if P[i] == P[i+1] {
			sameValues++
		} else {
			if sameValues != 1 {
				div = div * factorial(float64(sameValues))
				sameValues = 1
			}
		}
	}

	if sameValues != 1 {
		div = div * factorial(float64(sameValues))
		sameValues = 1
	}

	return n / div

}

func factorial(n float64) float64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	PCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	PTemp := strings.Split(readLine(reader), " ")

	var P []int

	for i := 0; i < int(PCount); i++ {
		PItemTemp, err := strconv.ParseInt(PTemp[i], 10, 64)
		checkError(err)
		PItem := int(PItemTemp)
		P = append(P, PItem)
	}

	result := solve(P)

	s := fmt.Sprintf("%.6f", result)

	fmt.Fprintf(writer, "%v\n", s)

	writer.Flush()
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
