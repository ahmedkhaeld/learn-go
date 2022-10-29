# func

####Function parameters receive copies of the arguments
when you call a function that has parameters declared,
you need to provide arguments to the call. The value in each argument is
copied to the corresponding parameter variable. (Programming languages
that do this are sometimes called “pass-by-value.”)

```go
package main
import "fmt"
func main() {
	num := 6
	double(num)
	fmt.Println(num) //6 print the original value
}
func double(n int) {
	n *= 2
}

```
This is fine in most cases. But if you want to pass a variable’s value to a
function and have it change the value in some way, you’ll run into trouble.
The function can only change the copy of the value in its parameter, not the
original. So any changes you make within the function won’t be visible
outside it!

We need a way to allow a function to alter the original value a variable
holds, rather than a copy.
####Pointers
You can get the address of a variable using & (an ampersand), which is Go’s
“address of” operator.
```go
package main

import "fmt"

func main() {

	myInt := 4 // address of myInt is [0xc0000ba000]
	var myIntPointer *int
	myIntPointer = &myInt // assign the address [0xc0000ba000] to myIntPointer

	fmt.Println(&myInt)        // prints the address  [0xc0000ba000]
	fmt.Println(myIntPointer)  // prints the address  [0xc0000ba000]
	fmt.Println(*myIntPointer) // prints the value in the address [4]

	*myIntPointer = 8          // change the value the address points to
	fmt.Println(*myIntPointer) //8
	fmt.Println(myInt)         //8

	//short  declaration for a pointer variable
	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)

}

```
#### Using pointers with functions
it's possible to return pointers from functions
```go
package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func main() {
	//declare and assign pointer to a variable
	var myFloatPointer = createPointer()
	fmt.Println(*myFloatPointer) //98.5
}

```

you can pass pointers to functions as arguments
```go
package main

import "fmt"

func printValueOfPointer(boolPointer *bool) {
	fmt.Print(*boolPointer)
}

func main() {
	var myBool = true
	printValueOfPointer(&myBool) //true
}
```

```go
package main

import "fmt"

//double takes an address to an integer, and double the value stored in this address
func double(n *int) {
	*n *= 2
}
func main() {
	num := 6
	double(&num)
	fmt.Println(num)//12
}
```
```go
package main

import "fmt"

func negate(myBool *bool) {
	*myBool = !*myBool
}
func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)//false
}

```