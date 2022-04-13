package main

import (
	"fmt"
	"sync"
)

/*
basic wait-group
this will output random  tasks order
*/
func main() {

	// 1-intro. create wait-group
	var wg sync.WaitGroup
	// 2-wait-groups. call Add method in order to indicate how many go routines to wait for
	// Add() call at the same layer as the Wait()
	wg.Add(3)

	// 3. run the go routines, inside each routine call Done method to decrement the waiting list
	go func() {
		defer wg.Done()
		fmt.Println("task 1-intro done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("task 2-wait-groups done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("task 3 done")

	}()

	// 4. join point, in order to wait for those go routines to execute
	// Wait() call at the same scope of Add()
	wg.Wait()
}
