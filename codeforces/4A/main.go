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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	wTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	if isSplitInEvenPosible(wTemp) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func isSplitInEvenPosible(w int64) bool {
	if w < 4 {
		return false
	}
	if w%2 == 0 {
		return true
	}
	return false
}
