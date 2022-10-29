# Slices
we can't add more elements to an array. that's a real problem
for our program, because we don't know in advance how many pieces of
data our file contains

**Slices** are a collection type that can grow to hold
additional items `var myslice []string` this is just like the syntax
for declaring an array, except that you don't specify the size.


Unlike with array variable, declaring a slice variable doesn't automatically
create a slice. for that you can call the built-in `make` function
You pass
make the type of the slice you want to create (which should be the same as
the type of the variable you’re going to assign it to), and the length of slice
it should create.

```go
package main

import "fmt"

func main() {
	var notes []string        //declare a slice variable
	notes = make([]string, 7) //create a slice with seven string elements
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "me"
	notes[3] = "fa"
	notes[4] = "so"
	notes[5] = "la"
	notes[6] = "ti"
	fmt.Println(notes)

	//create a slice with 5 int and set up a var to hold it
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes)

	//using slice literals
	//if you know in advance what values a slice will start with
	//you can initialize the slice
	letters := []string{"a", "b", "c"}
	fmt.Println(letters)
}

```
>Hold up! It looks like can do everything arrays can do, and
>you say we can add additional values to them! why didn't you
> just show us slices, and skip that array nonsense?

Because slices are built on top of arrays. You can't
understand how slices work without understanding arrays.

#### The slice operator
every slice is built-in on top of an underlying array. it's the underlying
array that actually holds the slice's data; the slice is merely
a view into some or all of the array's elements

when you use make func or a slice literal to create a slice, 
the underlying array is created for you automatically(you can't access it, except through the slice)
. But you can also create the array yourself, and then create
a slice based on it with the **slice operator**
`mySlice:=myArray[1:3]`

````go
package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3] //elements 0 to 2
	slice2 := underlyingArray[2:5] //elements 2 to 4
	slice3 := underlyingArray[:]   //elements from 0 to 4 start to end
	fmt.Println(slice1)            //[a b c]
	fmt.Println(slice2)            //[c d e]
	fmt.Println(slice3)            //[a b c d e]
}
````
#### append
Go offers a built-in append function that takes a slice, and one or more
values you want to append to the end of that slice.
```go
package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "f")
	fmt.Println(slice, len(slice))

}
```
>Notice that we’re making sure to assign the return value of append back to
the same slice variable we passed to append. This is to avoid some
potentially inconsistent behavior in the slices returned from append.

A slice’s underlying array can’t grow in size. If there isn’t room in the array
to add elements, all its elements will be copied to a new, larger array, and
the slice will be updated to refer to this new array.
#### Slices and zero values

```go
package main

import "fmt"

func main() {
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v \n", intSlice, stringSlice)
	fmt.Println(len(intSlice), len(stringSlice)) // 0 0

	// asking the len of slice eq 0 means ask slice empty or nil

	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%#v \n", slice)

}
```
#### Reading additional file lines using slice and append
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

func getFloat(fn string) ([]float64, error) {
	var numbers []float64
	data, err := readFile(fn)
	if err != nil {
		return nil, err // return a nil slice if err
	}
	for i := range data {
		number, err := strconv.ParseFloat(data[i], 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func readFile(fn string) ([]string, error) {
	var lines []string

	// open the data file for reading
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	//create a new scanner for the file
	scanner := bufio.NewScanner(file)
	//Scan read a single line of text from the file
	//loops until the end of the file is reached
	//and scanner.scan return false
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		i++
		lines = append(lines, line)
	}
	//once the loop exits, we're done with the file
	err = file.Close() //close the file to free resources
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return lines, nil

}

```
> Note: we are returning a nil slice in case of an error, because
> if we return numbers this might return slice with invalid data

#### using cmd to provide new elements to the slice from the user
* Getting command-line args from `os.Args` slice

The os pkg has a pkg variable os.Args, that gets set to
a slice of strings representing the cmd args the currently running program
was executed with. 
`first element is the program name`
```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var sum float64 = 0
	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	count := float64(len(args))
	fmt.Printf("Average: %0.2f \n", sum/count)
}
```
#### Variadic functions
Have you noticed that some function calls can takes as few
or as many, arguments as needed?
look at `fmt.Println` or append
```go
package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(1, 2, 3, 4, 5)
	letters := []string{"a"}
	letters = append(letters, "b")
	letters = append(letters, "c", "d", "e", "f", "g")
}
```
so how? They are declared as variadic functions.
A variadic function is one that can be called with a varying
numbers of arguments. `func myFunc(param1 int, param2 ...string){}`
```go
package main

import "fmt"

func main() {
	severalInt(1)
	severalInt(1,3,4,5)
}

func severalInt(numbers ...int) {
	fmt.Println(numbers)
}

```

```go
package main

import "fmt"

func main() {
	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d")
	severalStrings()
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}
```

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
}

//maximum takes any number of float64 arguments
func maximum (numbers ...float64) float64{
	max:=math.Inf(-1) //starts with a very low value
	for _, number :=range numbers{
		if number > max {
			max = number
		}
	}
	return max 
}
```

```go
package main

import "fmt"

func main() {
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5, -5.2))
}

//inRange takes a min & max and any additional float64 args
// it will discard any argument that is below the given minimum or above the given maximum
// returning a slice containing only the arguments
//that were in the specified range
func inRange(min float64, max float64, nums ...float64) []float64 {
	var res []float64
	for _, num := range nums {
		if num >= min && num <= max {
			res = append(res, num)
		}
	}
	return res
}

```

```go
package main

import "fmt"

func main() {

	fmt.Println(avg(10, 20, 30, 50))
}

func avg(nums ...float64) float64 {
	var sum float64 = 0 //set up a var to hold the sum of args
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}
```

#### Passing slices to variadic functions
```go
package main

import "fmt"

func main() {
	intSlice :=[]int{1,2,3}
	severalInt(intSlice...)
	stringSlice :=[]string{"a","b","c","d"}
	mix(1,true,stringSlice...)

}


func severalInt(numbers ...int){
	fmt.Println(numbers)
}

func mix(n int, f bool, strings ...string)  {
	fmt.Println(n,f,strings)
}

```





