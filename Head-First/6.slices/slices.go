package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3}
	severalInt(intSlice...)
	stringSlice := []string{"a", "b", "c", "d"}
	mix(1, true, stringSlice...)

}

func severalInt(numbers ...int) {
	fmt.Println(numbers)
}

func mix(n int, f bool, strings ...string) {
	fmt.Println(n, f, strings)
}
