package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	result := ""

	if strings.HasSuffix(s, "AM") {
		if strings.HasPrefix(s, "12") {
			result = "00" + strings.TrimLeft(strings.TrimRight(s, "AMP"), "0123456789")
		} else {
			result = strings.TrimRight(s, "AMP")
		}
	} else if strings.HasSuffix(s, "PM") {
		if strings.HasPrefix(s, "12") {
			result = strings.TrimRight(s, "AMP")
		} else {
			a := strings.Split(s, ":")
			b, _ := strconv.Atoi(a[0])
			b = b + 12
			result = strconv.Itoa(b) + strings.TrimLeft(strings.TrimRight(s, "AMP"), "0123456789")
		}
	}

	return result

}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	//outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer outputFile.Close()

	//writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

	//s := readLine(reader)
	s := ""

	for i := 1; i < 10; i++ {
		s = "0" + strconv.Itoa(i) + ":05:45PM"
		fmt.Print(timeConversion(s))
		fmt.Print("|||")
	}

	for i := 0; i < 3; i++ {
		s = "1" + strconv.Itoa(i) + ":05:45PM"
		fmt.Print(timeConversion(s))
		fmt.Print("|||")
	}

	for i := 1; i < 10; i++ {
		s = "0" + strconv.Itoa(i) + ":05:45AM"
		fmt.Print(timeConversion(s))
		fmt.Print("|||")
	}

	for i := 0; i < 3; i++ {
		s = "1" + strconv.Itoa(i) + ":05:45AM"
		fmt.Print(timeConversion(s))
		fmt.Print("|||")
	}

	//result := timeConversion(s)

	//fmt.Print(result)

	//fmt.Fprintf(writer, "%s\n", result)

	//writer.Flush()
}
