```go
func readLineString(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func readLineInt64(in *bufio.Reader) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(readLineString(in)), 10, 64)
}

func readLineInt(in *bufio.Reader) (int, error) {
	return strconv.Atoi(strings.TrimSpace(readLineString(in)))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
```
