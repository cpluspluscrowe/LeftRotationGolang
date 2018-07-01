package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func rotate(arr []int, numberOfRotations int) {
	r := ring.New(len(arr))

	// Initialize the ring with some integer values
	for _, item := range arr {
		r.Value = item
		r = r.Next()
	}
	r = r.Move(numberOfRotations)
	n := r.Len()
	for i := 0; i < n; i++ {
		fmt.Print(r.Value, " ")
		r = r.Next()
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	nd := strings.Split(readLine(reader), " ")
	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int(dTemp)

	aTemp := strings.Split(readLine(reader), " ")
	var a []int

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := aItemTemp
		a = append(a, int(aItem))
	}
	rotate(a, d)
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
