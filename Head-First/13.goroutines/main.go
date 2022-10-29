package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

//responseSize takes a url to process,
//a chan to hold the size of the page size we can send it through that chan
func responseSize(url string, c chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//store the length in the channel
	c <- Page{URL: url, Size: len(body)}
}

func main() {
	page := make(chan Page)
	urls := []string{
		"https://golang.org/",
		"https://golang.org/dev",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		go responseSize(url, page)
	}
	for i := 0; i < len(urls); i++ {
		p := <-page
		fmt.Printf("%s: with %d bytes \n", p.URL, p.Size)
	}

}
