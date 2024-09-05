package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBufio(t *testing.T) {
	input := strings.NewReader("this is long\ntext paragraf\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("hello world\n")
	writer.Flush()
}
