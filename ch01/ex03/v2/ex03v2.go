package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

/*
time go run v2/ex03v2.go oi vc eh muito legal mesmo parceiro
/tmp/go-build1862554161/b001/exe/ex03v2 oi vc eh muito legal mesmo parceiro

real    0m0,103s
user    0m0,131s
sys     0m0,102s
*/

//Embora seja mais lenta neste caso, strings.Join geralmente é mais rápida para n grande
