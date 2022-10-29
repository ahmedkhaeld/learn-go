package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	nums, err := getFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nums)
	var sum float64 = 0
	for _, number := range nums {
		sum += number
	}
	counts := float64(len(nums))
	fmt.Printf("average: %0.2f \n ", sum/counts)
}

func getFloat(fn string) ([3]float64, error) {
	var numbers [3]float64
	data, err := readFile(fn)
	if err != nil {
		return numbers, err
	}
	for i := range data {
		numbers[i], err = strconv.ParseFloat(data[i], 64)
		if err != nil {
			return numbers, err
		}
	}
	return numbers, nil
}

func readFile(fn string) ([3]string, error) {
	var lines [3]string

	// open the data file for reading
	file, err := os.Open(fn)
	if err != nil {
		return lines, err
	}

	//create a new scanner for the file
	scanner := bufio.NewScanner(file)
	//Scan read a single line of text from the file
	//loops until the end of the file is reached
	//and scanner.scan return false
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		//fmt.Println(scanner.Text()) // Text() returns a string with data that was read
		i++
	}
	//once the loop exits, we're done with the file
	err = file.Close() //close the file to free resources
	if err != nil {
		return lines, err
	}
	if scanner.Err() != nil {
		return lines, err
	}
	return lines, nil

}
