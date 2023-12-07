package util

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

// streamLines returns a scanner to stream a file
func StreamLines(path string) bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return *bufio.NewScanner(file)
}

func GetLines(path string) []string {
	input := make([]string, 0)
	scanner := StreamLines(path)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func LineCounter(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// Merge copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the two sets of values will be merged.
func Merge[M1 ~map[K]V, M2 ~map[K]V, K comparable, V []int](dst M1, src M2) {
	for k, v := range src {
		_, ok := dst[k]
		if ok {
			dst[k] = append(dst[k], v...)
		} else {
			dst[k] = v
		}
	}
}

// Returns a ptr to the value passed in
func ToPtr[T any](value T) *T {
	return &value
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
