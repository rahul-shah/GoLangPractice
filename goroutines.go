package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/pkg/errors"
)

func returnType(url string) {
	content, err := http.Get(url)
	if err != nil {
		errors.Wrap(err, "Can't call url")
		return
	}

	defer content.Body.Close()
	contentType := content.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, contentType)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		//Get get blocked on every call so creating a go routine here
		//returnType(url)
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
