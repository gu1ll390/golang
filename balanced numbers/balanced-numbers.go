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

func balancedNumbers(B, L int32, input []int64) []int64 {
	return []int64{math.MaxInt64, math.MaxInt64}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	BL := strings.Split(readLine(reader), " ")

	Btemp, err := strconv.ParseInt(BL[0], 10, 64)
	checkError(err)
	B := int32(Btemp)

	Ltemp, err := strconv.ParseInt(BL[1], 10, 64)
	checkError(err)
	L := int32(Ltemp)

	inputTemp := strings.Split(readLine(reader), " ")

	var input []int64

	for i := 0; i < int(L); i++ {
		inputItem, err := strconv.ParseInt(inputTemp[i], 10, 64)
		checkError(err)
		input = append(input, inputItem)
	}

	result := balancedNumbers(B, L, input)

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
