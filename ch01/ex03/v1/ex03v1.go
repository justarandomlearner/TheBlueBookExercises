package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

/*
time go run v1/ex03v1.go oi vc eh muito legal mesmo parceiro
oi vc eh muito legal mesmo parceiro

real    0m0,093s
user    0m0,110s
sys     0m0,106s
*/
