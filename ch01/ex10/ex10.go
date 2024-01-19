package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for i, url := range os.Args[1:] {
		go fetch(i, url, ch)
	}

	f, err := os.OpenFile("./PERFORMANCE.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if errors.Is(err, os.ErrNotExist) {
		f, err = os.Create("PERFORMANCE.txt")
		fmt.Println(f.Name())
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error openning file: %v\n", err)
		if f != nil {
			f.Close()
		}
	}

	_, err = f.Write([]byte("---------------\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	for range os.Args[1:] {
		result := <-ch
		result += "\n"
		// _, err := fmt.Fprintf(f, "%s\n", result)

		_, err := f.Write([]byte(result))
		if err != nil {
			fmt.Printf("Printando o ERRO: %v\n", err)
		}
	}

	f.Close()
	f.Sync()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(positionInList int, url string, ch chan<- string) {
	start := time.Now()

	url, _ = strings.CutPrefix(url, "https://")
	url, _ = strings.CutPrefix(url, "http://")

	if !strings.HasPrefix(url, "www.") {
		url = "www." + url
	}

	file, err := os.Create(fmt.Sprintf("./output/%s-arg-position_%d.html", url, positionInList))
	if err != nil {
		fmt.Println(err) // send to channel ch
		file.Close()
		return
	}
	defer file.Close()

	url = "http://" + url

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading arg number %d: %s: %v", positionInList, url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Arg number %d: %.2f %7d %s", positionInList, secs, nbytes, url)
}
