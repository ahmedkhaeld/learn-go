# Composite Types
## 1. Arrays
arrays are typed by size, which is fixed at compile time
```go
// how to initialize an array
var a [3]int
var b [3]int{0,0,0}
var c [...]{0,0,0} // sized by initializer

var d [3]int 
d = b   // elements are copied, d is independent from b

var m [...]int{1,2,3,4}
c = m // type mismatch
```
* Arrays are passed by value, thus elements are copied
*  with an array  there's no descriptor an array is just a chunk of memory
.we just
copied the bytes physically
which is fine as long as there aren't
very many, if
these arrays get large then it's going
to be very inconvenient

---
## 2. Slice
* slice is like an array, but slice has a descriptor,
and it points at some other memory, and in fact what a slice
has, it always has an array behind it.

```go
package main

import "fmt"

var w = [...]int{1, 2, 3}
var x = []int{0, 0, 0}

func do(a [3]int, b []int) []int {
a[0] = 4

	b[0] = 3 // []int{3,0,0}

	c := make([]int, 5) // []int{0,0,0,0,0}

	c[4] = 42 // []int{0,0,0,0,42}

	copy(c, b) // copies only 3 elements to c  > {3,0,0,0,42}
	return c
}

func main() {
y := do(w, x)
fmt.Println("y", y)
fmt.Println("w", w)
fmt.Println("x", x)
}

// output:
//y [3 0 0 0 42]
//w [1 2 3]
//x [3 0 0]

```
* when we create a slice in its own, Go will magically put an array behind the slice
to hold the actual value 























