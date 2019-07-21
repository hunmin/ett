package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/hunmin/ett/timetagger"
)

func main() {
	tt := timetagger.NewTimeTagger()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if isPrefix {
			fmt.Printf("%s", line)
		} else {
			fmt.Printf("%-8s | %s \n", tt.Tag(), line)
		}
	}
}
