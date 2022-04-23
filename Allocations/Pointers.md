# Pointers

#### Variable
if you
think of a variable in programming
it would look something like this it's
going to have a name type and
value and it's going to be stored
somewhere in the memory address is going to represent the
location of the box
inside the warehouse so if we need to go
and get that box

<img src="https://user-images.githubusercontent.com/40498170/164616727-6875936b-f284-4120-aa5c-8706ac5f8aa9.png" height="100">

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(i,j)
}
```
first we declare two variables with the
shorthand declaration
and assign these integer values 42 and 2701 . i and j is going to get some space in
the memory and they'll also get an
address
if you print out i and j you'll get the
value like what's in the box

* if you want to check the address for
i and j
you can do that by adding an ampersand
in front of the variable name
```go
package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(&i,&j)
}
// 0xc0000b8000 0xc0000b8008
```
&i address of i

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(&i, &j)

	p := &i
	// p is a pointer to the same address that i points to
	fmt.Println(p)
}
// address of i 0xc0000b8000   address of j  0xc0000b8008
// address of p 0xc0000b8000 
```


what if we do print *p (the value in p) this is called dereferencing
```go
package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(&i, &j)

	p := &i
	// print the value at the address of p 
	fmt.Println(*p)
}

```
changing the value in *p will change the value in i


---
