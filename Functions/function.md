## Functions, Parameter, Defer
* Functions
* How parameters passed to functions
* Defer stmt which is a very useful way for making sure when function is Done  that certain things gets done with it

functions in go are first class objects you can do anything with them

* pass by value to func
```go
package main

import "fmt"

// do take in array of integers and returns array after changing the second element value
func do(b [3]int) [3]int {
	b[0] = 0
	return b
}

func main() {
	// declare an array (copied)
	a := [3]int{1, 2, 3}
	// calling the function do, it takes a copy from array 'a' the do whatever it wants to 
	// without affecting the array 'a' values
	v := do(a)

	fmt.Println("a => ", a, "v =>", v) // [1 2 3] [0 2 3]

}

```
> the array 'a' does not change even what the function do() as it change the first element of the passed array<br>
> when we pass an array, the array is being copied, array 'b' is different array to 'a' 


* pass by ref
a&b are slices and the pointer withing the slice descriptor points to the same block of values
<br>so when we change through b it also changes a
```go
package main

import "fmt"

// do take in slice of integers and slice after changing the second element value
func do(b []int) []int {
	b[0] = 0
	// print the address of b
	fmt.Printf("a address %p\n", b) //b address 0xc00001a228
	return b
}

func main() {
	// declare an slice (reference)
	a := []int{1, 2, 3}
	fmt.Printf("b address %p\n", a) //a address 0xc00001a228

	// calling the function do, it takes a copy from slice 'a' the do whatever it wants to
	// and affecting the slice 'a' values
	// because a&b both referring to the same underlying array of values
	v := do(a)

	fmt.Println("a => ", a, "v =>", v) // [0 2 3] [0 2 3]

}

```
* pass by ref
using map: 
```go
package main

import "fmt"

func do(m1 map[int]int) {
	// add a new key to the map
	m1[4] = 400
	
	// make map creates a new different map with only one key 
	// m1 is local map that has no relation to the passed map
	// changes in m1 now does not affect the passed map
	m1 = make(map[int]int)
	m1[4] = 4

	fmt.Println("m1", m1)

}

func main() {
	// declare an map (reference)
	m := map[int]int{1: 100, 2: 200, 3: 300}
	fmt.Println("m", m)
	//m map[1:100 2:200 3:300]

	do(m)
	fmt.Println("m", m)
	//m   map[1:100 2:200 3:300 4:400]  key 4 is added to with value 400

}

```



