package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// demo shows concurrent tasks each has a different workload
//  GOMAXPROCS=1 GOGC=off GODEBUG=schedtrace=200,scheddetail=1 ./exec
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task5(&wg)
	wg.Wait()
}

// task1 workload is make a request to server
func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("could not make http request: %v", err)
	}
	fmt.Println("task 1: done")
}

// task2 workload do some calculations simulate io on the cpu level
func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	var count int
	for i := 0; i < 1_000_000_000; i++ {
		count += i
	}
	fmt.Println("task 2: done")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task 3: done")
}

// task4 workload sleep simulate heavy work
func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("task 4: done")
}

// task5 workload sleep simulate heavy work
func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(600 * time.Millisecond)
	fmt.Println("task 5: done")
}
