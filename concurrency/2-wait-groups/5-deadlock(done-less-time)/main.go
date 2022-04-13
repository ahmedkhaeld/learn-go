package main

import "sync"

// not call Done enough time as many Added to the Add()
// call Dane() less time than indicated in the Add() will result a deadlock
// which means we are waiting on something that will never happen
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// wait forever for operation that never calls done()
	wg.Wait()
}
