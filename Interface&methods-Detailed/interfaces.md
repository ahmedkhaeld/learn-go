# Interfaces in details
#### PRAGMATICS OF HOW WE USE INTERFACES
>if you create a variable of interface type what value does it have by default<br>
> for interfaces it's nil

#### Nil interfaces
An interface variable is nil until initialized.

* it really has two parts:<br>
    * it has some way of referring to the actual concrete type
    * a value 
  
```go
var r io.Reader  // nil until initialized 
var b *bytes.Buffer  // b is a pointer to bytes buffer

r = b            // r is no longer nil, its type is b, but it has no value in it
```

---
#### Error is really an interface
We called error a special type, but it's really an interface
```
type error interface {
func Error() string
}
```
We can compare it to nil unless we make a mistake <br>
The mistake is to store a nil pointer to a concrete type in the error variable

##### explain what error is
treat it like a primitive type it's either nil or not<br>
but actually error is an interface

it's an interface with one method error() that returns some sort of string describing 
what the error is
```go
var err error
if err != nil{
}
```
a var err of type error
if i don't initialize it with anything it's nil 

we've been saying
things like if error is not equal to nil `if err != nil{ }`
that means an error happened because
this
error interface is no longer
uninitialized it has
an error in it of some sort now

>the reason it's an interface is that sometimes we customize errors
>>certain kinds of
errors where you want to report more
information
so you have some sort of type it
encapsulates a bit more knowledge
or in some cases people want errors with
a traceback
and there are ways of creating errors
with a traceback not using the standard
library but using some extension
packages

* how do we figure out if an error happened?
  we compare the error interface to nil
if it's nil nothing's ever been put into
it there's no error

```go
package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}
//XYZ returns nil pointer to concrete type errFoo 
func XYZ(a int) *errFoo {
	return nil
}
func main() {
	// err of type error points to a nil concrete type, but it is not a nil interface now
	// when interface points to a nil type, it is not nil 
	var err error = XYZ(1) // BAD: interface gets a nil concrete ptr
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok!")
	}
	// oops
}

```
we can fix this by make `XYZ func` return a nil to an interface, means the err  is nil
```go
package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}
//XYZ returns nil pointer to concrete type errFoo 
func XYZ(a int) error  {
	return nil
}
func main() {
	// err of type error points to a nil concrete type, but it is not a nil interface now
	// when interface points to a nil type, it is not nil 
	var err error = XYZ(1) // BAD: interface gets a nil concrete ptr
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok!")
	}
	// ok!
	// because XYZ return a nil error
}

```


---
#### Currying functions
Currying
takes a function and reduces its argument
one
argument gets bound, and a new function is returned
```go
package main

import "fmt"

func Add(a, b int) int {
  return a + b
}
func AddToA(a int) func(int) int {
  return func(b int) int {
    return Add(a, b)
  }
}
func main() {

  addTo1 := AddToA(1)
  fmt.Println(Add(1, 2) == addTo1(2)) // true
}
```
currying is a process in
which we take a function that has say
three parameters
we bind one of them and turn it into a
function that has two parameters
