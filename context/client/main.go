package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// where to put the context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// separate the request into two steps
	// 1. create the request
	req, err := http.NewRequest(http.MethodGet, "15.http://localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// put the ctx within the request before sending
	req = req.WithContext(ctx)

	// 2. send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalln(res.Status)
	}
	// copy the body of the resp to the standard output
	io.Copy(os.Stdout, res.Body)
}
