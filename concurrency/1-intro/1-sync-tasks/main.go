package main

import (
	"fmt"
	"time"
)

/*
sync tasks
has 4 tasks, each task will take a certain amount of time after that they will exit
measure the time all 4 tasks take
capture time before start and time end
*/

func main() {
	now := time.Now()
	task1()
	task2()
	task3()
	task4()
	// measure how much time it take to execute tasks synchronously
	fmt.Println("elapsed:", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task2")
}

func task3() {
	fmt.Println("task3")
}

func task4() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4")
}
