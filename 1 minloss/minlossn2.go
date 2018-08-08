package main

// import (
// 	"math"
// )

// // Complete the minimumLoss function below.
// func minimumLoss(price []int64) int64 {

// 	var minloss int64
// 	minloss = math.MaxInt64

// 	for i := 0; i < len(price); i++ {
// 		var dif int64
// 		for j := i; j < len(price); j++ {
// 			dif = price[i] - price[j]
// 			if dif > 0 && dif < minloss {
// 				minloss = dif
// 			}
// 		}
// 	}

// 	return minloss

// }

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 1024*1024)

// 	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	priceTemp := strings.Split(readLine(reader), " ")

// 	var price []int64

// 	for i := 0; i < int(n); i++ {
// 		priceItem, err := strconv.ParseInt(priceTemp[i], 10, 64)
// 		checkError(err)
// 		price = append(price, priceItem)
// 	}

// 	result := minimumLoss(price)

// 	fmt.Fprintf(writer, "%d\n", result)

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
