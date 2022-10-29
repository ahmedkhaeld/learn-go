package main

import (
	"autotest/prose"
	"fmt"
)

func main() {
	phrases := []string{"my parents!", "a rodeo"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))

	phrases = []string{"my parents!", "a rodeo", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))

}
