package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumLoss function below.
func minimumLoss(price []int64) int64 {

	var minloss int64 = math.MaxInt64

	indices := make(map[string]int)
	for i := 0; i < len(price); i++ {
		indices[strconv.FormatInt(price[i], 10)] = i
	}

	MergeSort(price)

	for i := len(price) - 1; i > 0; i-- {
		if price[i]-price[i-1] < minloss && indices[strconv.FormatInt(price[i], 10)] < indices[strconv.FormatInt(price[i-1], 10)] {
			minloss = price[i] - price[i-1]
		}
	}

	return minloss

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 2048*2048)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	priceTemp := strings.Split(readLine(reader), " ")

	var price []int64

	for i := 0; i < int(n); i++ {
		priceItem, err := strconv.ParseInt(priceTemp[i], 10, 64)
		checkError(err)
		price = append(price, priceItem)
	}

	result := minimumLoss(price)

	fmt.Fprintf(writer, "%d\n", result)

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

//MergeSort function"
func MergeSort(A []int64) {

	mergeSortRec(A, 0, len(A)-1, make([]int64, len(A)))
}

func mergeSortRec(A []int64, i, j int, B []int64) {
	if i < j {
		m := (i + j) / 2

		mergeSortRec(A, i, m, B)
		mergeSortRec(A, m+1, j, B)
	}

	merge(A, i, j, B)
}

func merge(A []int64, i, j int, B []int64) {
	m := (i + j) / 2

	k := i
	p := i
	q := m + 1

	for p <= m && q <= j {
		if A[p] < A[q] {
			B[k] = A[p]
			k++
			p++
		} else {
			B[k] = A[q]
			k++
			q++
		}
	}

	for p <= m {
		B[k] = A[p]
		k++
		p++
	}

	for q <= j {
		B[k] = A[q]
		k++
		q++
	}

	for k = i; k <= j; k++ {
		A[k] = B[k]
	}
}
