package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if nil != err {
				fmt.Fprintf(os.Stderr, "dupfile: %v\n", err)
				continue
			}
			countLine(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
	}
}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++
	}
}