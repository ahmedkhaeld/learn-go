# Arrays
>I have a huge list of things to do today! well, I'll just
handle them one at a time. I'll get done eventually!

A lot of programs deal with lists of things. List of addresses
List of phone numbers. List of products. Go has built-in  way
of storing lists
* create arrays
* fill arrays with data
* get data from arrays
* process all the elements in array

#### Arrays hold collections of values

```go
package main

import "fmt"

func main() {
	notes := [7]string{"do", "ri", "mi", "fa", "so", "la", "ti"}
	for i, v := range notes {
		fmt.Println(i, v)
	}

}

```

```go
package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	count := float64(len(numbers))
	fmt.Println(sum)
	fmt.Printf("average: %0.2f ", sum/count)

}

```
reading a text file<br>
parse file into float64 <br>
calculate the average and sum
```go
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

```
>[71.8 56.2 89.5]<br>
average: 72.50 
