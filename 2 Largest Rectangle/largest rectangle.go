package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
				altura = int64(h[j])
			}
			if area > base*altura {
				break
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

	//h := []int32{1000000}
	//h := make([]int32, 100000)
	h := []int32{5, 1, 2, 3, 4, 5, 1}
	//h := []int32{3, 3, 3, 3, 3, 3} // 18
	//h := []int32{1, 2, 3, 4, 5, 1, 2, 2, 7, 7, 2, 2}
	//h := []int32{8979, 4570, 6436, 5083, 7780, 3269, 5400, 7579, 2324, 2116} // 1 - 26152
	// h := []int32{6873, 7005, 1372, 5438, 1323, 9238, 9184, 2396, 4605, 162, 7170, 9421, 4012, 5302, 6277, 2438, 4409, 3391, 4956, 4488, 622, 9365, 5088, 6926,
	// 	2691, 6909, 1050, 2824, 3538, 5801, 8468, 411, 9158, 9841, 2201, 481, 5431, 1385, 2877, 36, 1547, 48, 5809, 1911, 1702, 8439, 4349, 6111, 1830, 5657, 6951,
	// 	8804, 5022, 8391, 2083, 7713, 5300, 3133, 6890, 5190, 5286, 1710, 1953, 4445, 7903, 4154, 4926, 3335, 5539, 4156, 9723, 3438, 556, 1885, 5349, 2258, 324,
	// 	6050, 4722, 8506, 1707, 1673, 7310, 3081, 65, 9393, 7147, 1717, 8878, 389, 6908, 4165, 2099, 5213, 8610, 3, 9368, 3536, 9690, 1259} //3 - 51060
	fmt.Println(largestRectangle(h))

	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// hTemp := strings.Split(readLine(reader), " ")

	// var h []int32

	// for i := 0; i < int(n); i++ {
	// 	hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
	// 	checkError(err)
	// 	hItem := int32(hItemTemp)
	// 	h = append(h, hItem)
	// }

	// result := largestRectangle(h)

	// fmt.Fprintf(writer, "%d\n", result)

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
