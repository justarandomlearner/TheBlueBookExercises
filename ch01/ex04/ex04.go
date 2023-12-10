package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
we can test the else condition (reading from os.Stdin) doing:
cat batata.txt | go run ex04.go
*/
type lineInfo struct {
	ocurrences int
	atFiles    map[string]struct{}
}

func main() {
	counts := make(map[string]lineInfo)

	if len(os.Args[1:]) >= 1 {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "error openning file: %v", err)

				continue
			}

			countLines(f, counts)

			f.Close()
		}
	} else {
		countLines(os.Stdin, counts)
	}

	for line, count := range counts {
		filesSlice := make([]string, 0)
		for file := range count.atFiles {
			filesSlice = append(filesSlice, file)
		}

		fmt.Printf("'%s'--> %d occurences at files: %s\n", line, count.ocurrences, strings.Join(filesSlice, ", "))
	}
}

// maps are always passed as a reference to its subjacent data structure
func countLines(f *os.File, counts map[string]lineInfo) {
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error during file scanning: %v", err)
		}
		if _, b := counts[sc.Text()]; !b {
			tempAtFiles := make(map[string]struct{})
			tempAtFiles[f.Name()] = struct{}{}

			counts[sc.Text()] = lineInfo{
				ocurrences: 1,
				atFiles:    tempAtFiles,
			}

			continue
		}

		st := counts[sc.Text()]
		st.ocurrences += 1
		st.atFiles[f.Name()] = struct{}{}
		counts[sc.Text()] = st
	}
}
