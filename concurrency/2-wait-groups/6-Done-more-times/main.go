package main

import "sync"

// calling Done() more times than calling Add() will result a panic which is deadly for the application
// because Done() decrement the calling list, which will decrease zero to -ve counter
func main() {
	var wg sync.WaitGroup
	wg.Done()
}
