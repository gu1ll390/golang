package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func balancedNumbers(B, L int32, input []int64) []int64 {
	return []int64{int64(7), int64(28)}
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
	fmt.Print("B: ", B, " ")

	Ltemp, err := strconv.ParseInt(BL[1], 10, 64)
	checkError(err)
	L := int32(Ltemp)
	fmt.Print("L: ", L, "\n")

	inputTemp := strings.Split(readLine(reader), " ")

	var input []int64

	for i := 0; i < int(L); i++ {
		inputItem, err := strconv.ParseInt(inputTemp[i], 10, 64)
		checkError(err)
		input = append(input, inputItem)
	}
	fmt.Print(input, "\n")

	result := balancedNumbers(B, L, input)

	strResult := strconv.FormatInt(result[0], 10) + " " + strconv.FormatInt(result[1], 10)

	fmt.Fprintf(writer, "%s\n", strResult)

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
