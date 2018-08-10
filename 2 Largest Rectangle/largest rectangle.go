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

func largestRectangle(h []int32) int64 {
	var maxArea int64
	var area int64
	var base int64
	var altura int64

	maxArea = math.MinInt64

	for i := 0; i < len(h); i++ {
		area = math.MinInt64
		altura = math.MaxInt64

		for j := i; j < len(h); j++ {
			base = int64(j - i + 1)
			if int64(h[j]) < altura {
				if area > base*int64(h[j]) {
					base--
					break
				}
				altura = int64(h[j])
			}
			area = base * altura
			if maxArea < area {
				maxArea = area
			}
		}

		baseDer := base

		for j := i; j >= 0; j-- {
			base = int64(i-j) + baseDer
			if int64(h[j]) < altura {
				if area > base*int64(h[j]) {
					base--
					break
				}
				altura = int64(h[j])
			}
			area = base * altura
			if maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	hTemp := strings.Split(readLine(reader), " ")

	var h []int32

	for i := 0; i < int(n); i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	result := largestRectangle(h)

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
