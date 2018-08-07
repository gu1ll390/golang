package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the solve function below.
func solve(P []int32) float64 {
	return factorial(float64(len(P)))
}

func factorial(n float64) float64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func variacionesConRepeticion(P []int, m int) float64 {
	var total float64
	variacionesConRepeticionAux(P, make([]int, len(P)), 0, total)
	return total
}

func variacionesConRepeticionAux(P []int, variacion []int, pos int, total float64) {

	if pos == len(variacion) {
		total++
	} else {
		for i := 0; i < len(variacion); i++ {
			variacion[pos] = P[i]
			variacionesConRepeticionAux(P, variacion, pos+1, total)
		}
	}
}

func main() {

	P := []int{1, 2, 3, 4}
	t := variacionesConRepeticion(P, 0)
	fmt.Println(t)

	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// PCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)

	// PTemp := strings.Split(readLine(reader), " ")

	// var P []int32

	// for i := 0; i < int(PCount); i++ {
	// 	PItemTemp, err := strconv.ParseInt(PTemp[i], 10, 64)
	// 	checkError(err)
	// 	PItem := int32(PItemTemp)
	// 	P = append(P, PItem)
	// }

	// result := solve(P)

	// s := fmt.Sprintf("%.6f", result)

	// fmt.Fprintf(writer, "%v\n", s)

	// writer.Flush()

	// n := factorial(18)
	// s := fmt.Sprintf("%.6f", n)
	// fmt.Println(s)
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
