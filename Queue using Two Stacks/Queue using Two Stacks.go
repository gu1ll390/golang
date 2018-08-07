package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queue []int32

	for qItr := 0; qItr < int(q); qItr++ {
		query := strings.Split(readLine(reader), " ")

		switch query[0] {
		case "1":
			x, err := strconv.ParseInt(query[1], 10, 32)
			checkError(err)
			queue = append(queue, int32(x))
		case "2":
			if len(queue) > 0 {
				queue = queue[1:]
			}
		default:
			fmt.Fprintf(writer, "%d\n", queue[0])
		}
	}

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
