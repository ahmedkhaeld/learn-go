package main

import (
	"fmt"
	"time"
)

/*
has 4 go routines
when executing these four go routines, will produce non-deterministic behavior
as the main() function did not wait for the go routines to finish

go handle concurrency by fork/join model

why main is not waiting for go routines to execute?
the main function in the main pkg when executed it's going to run in its own go routine which called main go routine

in fork/join model
has main function, the main function has a couple of tasks at certain point in time which will tell the go scheduler
go ahead and schedule that go routine, from this point(fork), this process go in separate way, now that child process
has its own list of tasks, but at certain time, or after the child process has finished, it has to join the main process
go back(join point)  to the main
if it does not go back, or the main process does not expect a join back from that fork which was done initially
basically, the main function is not waiting for anything, I'm done here!

*/
func main() {
	now := time.Now()
	go task5()
	go task6()
	go task7()
	go task8()
	fmt.Println("elapsed:", time.Since(now))
}

func task5() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 5")
}

func task6() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 6")
}

func task7() {
	fmt.Println("Task 7")
}

func task8() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 8")
}
