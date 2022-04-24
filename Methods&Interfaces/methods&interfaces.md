# Methods and Interfaces
An Interface specifies abstract behavior in terms of methods.

Concrete types offer methods that satisfy the interface.

A method may be defined on any user-declared(named) type

```go
type Stringer interface{
	String() string 
}
```

### what is a method?
is a special type of function. in go it's just a little different
because we don't have classes and we
don't define the methods in the class
we define them separately.
**method is a function that also
has a receiver
and the receiver is specified before the
function name**

* example a method in user declared type 
```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)
}
//0: 1
//1: 2
//2: 3
//main.IntSlice [1;2;3]
```
**what is percent v going to print?** `[1;2;3]` **why we had this result?**<br>
when we are doing the Printf it printed out in this particular format
when printf went to print
my type with percent v it used my string function
>if we looked inside printf or print
line or any of these print functions
what we'd see is some interesting logic
the first thing it does is say
is the argument a string because if the
argument is a string
well just copy the string to the output
that's easy
and the second thing it does is to check
to see if the
item is a stringer because if it is a
stringer
then it can just call the string method
generate the string that way
and put it on the output 


* example assign type to interface after satisfying the interface
```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}
	var s fmt.Stringer = v // cause v satisfy Stringer interface, then I can assign v to s

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)
	fmt.Printf("%T %[1]v\n", s)
}

//0: 1
//1: 2
//2: 3
//main.IntSlice [1;2;3]

```
fmt.stringer is an interface
so it's an interface variable
and what is it what can i assign to it
well i can assign to it anything that
satisfies the interface
any actual type that has a string method

**if i print that interestingly enough i'm going to get the same output** <br>
s is a variable of interface
type
but at runtime when i do like a percent
v to printf
well what we're looking at is the type
of the thing that the interface is
holding on to

---
## why interfaces?
if we didn't have interfaces and we
didn't have a way to abstract behavior
then we would have a bunch of methods
that took different real types
concrete types like file or buffer or
socket or whatever
those met we'd have to create a
different method for each one and we'd
have to know that we're getting a file
or a socket and probably do
certain things that are special for that
and we don't want to we want to abstract
out the notion
of something we can write to

`we think about
interfaces in GO usually is not from
the provider side
but from the consumer side 
` <br>
i'm
some piece of function
and i need a parameter that provides
certain behavior
i therefore am the one to say i want
this particular interface with this
particular method set
to provide this behavior i the consumer
okay and it's up to various providers to
provide the right methods
i can therefore have existing objects in
a program
that have various methods and i can
create a new interface
somewhere else that says hey i want a
subset of those i just want this one
method like write
and that is the behavior that i want as
a consumer
and that may be behavior a specification
of behavior that was unknown
when those other types were written
maybe what i've done is actually
reduce so there's already some objects
out there some
types that have methods and there were
some interfaces with more than one
method
and but i say i'm going to create an
interface with just one method because i
only want to capture that one
piece of behavior

---
## Receivers
a method may take a pointer or value receiver, but not both
```go
type Point struct{
	X,Y float64
}

// with value receiver
func(p Point) Offset(x,y float64) Point{
	return Point{p.x+x, p.y+y}
}

```
Offset take a value receiver then the method gets a copy then it returns a new point<br>
the receiver p is not changed it
can't be changed
because it came into this function as a
copy


```go
type Point struct{
	X,Y float64
}

// with pointer receiver
func(p Point) Move(x,y float64) Point{
	p.x+=x
	p.y+=y
}

```
if i want to change p permanently the original object
which is passed into the method is
changed then i need to take a pointer
receiver. that's the case of Move i want to
Move this point by actually changing its
coordinates


---
* open a file "a.txt",
* create "out.txt"
* copy a to out
```go
package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	f1, _ := os.Open("a.txt")
	f2, _ := os.Create("out.txt")

	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
}
```
word count flag to a file `$ wc -c out.txt`

> os.Copy it reads from Reader interface, and writes to a Writer interface<br>
> so anything that has a read method can be the source, and anything with the write
method can be the destination

```go
package main

import (
	"fmt"
	"io"
	"os"
)

// ByteCounter is a Writer because it has the Write method
// ,so it can be putted in the destination for the io.Copy
// it just as good as any Writer
type ByteCounter int

//Write return the length of the bytes
func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l) // cast l to byte counter and increment the counter
	return l, nil
}
func main() {
	// declare a variable of type ByteCounter
	var c ByteCounter
	f1, _ := os.Open("a.txt")
	f2 := &c // f2 is a byte counter

	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
	fmt.Println(c)
}

//copied 15 bytes
//15

```

---
* demo for composite structs
```go
package main

import (
	"fmt"
	"math"
)

// Point has x,y coordinates
type Point struct {
	X, Y float64
}

// Line has two points, beginning and ending
type Line struct {
	Begin, End Point
}

// Distance method figure out what is the distance between begin point and end point
func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func main() {

	// create a line by passing two points literals(must put Point)
	side := Line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side.Distance())
	//5

}
```

* a more complex composition
```go
package main

import (
	"fmt"
	"math"
)

// Point has x,y coordinates
type Point struct {
	X, Y float64
}

// Line has two points, beginning and ending
type Line struct {
	Begin, End Point
}

// Path represent a non-straight line
type Path []Point

// Distance method figure out what is the distance between begin point and end point
func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

//Distance calculate some distances between some points
func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}

	return sum

}

func main() {

	// create a line by passing two points literals(must put Point)
	side := Line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side.Distance())

	// create an actual path
	// these points make a triangle, now cal the perimeter of it
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perimeter.Distance())

}

```
* create a Path that has a slice of Points
* create a Distance method on Path
    * calculate the distances between the points
    * between each two points there's a Line, and i already know how to cal Distance of 
      Line
    * sum is going to gather the points in the slice
    * init the i to 1 not zero
    * `[i-1]` previous point  `[i]` current point
    * `Line{p[i-1], p[i]}.Distance()` is a Line literal, is a value on the fly,
    does not have an address, calculate the distance and throw the value away


```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}

	return sum

}

type Distancer interface {
	Distance() float64
}

func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}

func main() {

	side := Line{Point{1, 2}, Point{4, 6}}
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

	PrintDistance(side)
	PrintDistance(perimeter)

}

```


```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

//ScaleBy scales the line, make it longer
func (l *Line) ScaleBy(f float64) {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
}

func main() {

	side := Line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side.Distance()) // prints 5

	// scale side to twice long
	side.ScaleBy(2)
	fmt.Println(side.Distance()) // prints 10

	side.ScaleBy(3)
	fmt.Println(side.Distance()) // prints 15
}

```


* l receiver by value
```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

//ScaleBy scales the line, make it longer
func (l Line) ScaleBy(f float64) Line {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)

	return Line{l.Begin, Point{l.End.X, l.End.Y}}
}

func main() {

	side := Line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side.Distance()) // prints 5

	s2 := side.ScaleBy(2)
	fmt.Println(s2.Distance()) //10

	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleBy(2.5).Distance())
	// 12.5
}

```