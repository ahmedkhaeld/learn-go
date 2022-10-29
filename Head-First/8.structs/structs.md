#building storage: Structs
**Sometimes you need to store more than one type of data**
>Then we decided to just connect the string, the int, and the
> bool together!

#### Gopher Fancy magazine
to start, we need to store the subscriber's name, the monthly
rate, we're charging them, and whether their subscription
is active, But the name is a string, the rate float64, the
active status is bool

You can create a struct that holds string, float64, bool values
all in one convenient group
``` 
struct {
    field1 type
    field2 type
}
```

``` 
var myStruct struct {
    number float64
    word   string
    toggle bool
}
```

#### Defined types and structs
``` 
var myStruct1 struct {
    number float64
    word   string
    toggle bool
}
myStruct1.number=1

var myStruct2 struct {
    number float64
    word   string
    toggle bool
}
myStruct2.number=3
```
**Type Definitions** allows you to create type of your own.
They let you create a new defined type that's based on an
underlying type
```go
package main

import "fmt"

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println(bolts.description)
	fmt.Println(bolts.count)

}

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}
```
#### Using defined types with functions
```go
package main

import "fmt"

func main() {
	P:=minOrder("Hex bolts")
	showInfo(P)
	fmt.Println(P.count,P.description)
}

type part struct {
	description string
	count       int
}

func showInfo(p part){
	fmt.Println("Description:", p.description)
	fmt.Println("Description:", p.count)
}

func minOrder(description string)part{
	var p part
	p.description=description
	p.count=100
	return p
}
```

subscriber example
```go
package main

import "fmt"

func main() {
	sub := defaultSubscriber("ahmed")
	applyDiscount(&sub)
	printInfo(sub)
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s

}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

//applyDiscount update the rate fo a subscriber
//it takes a pointer to subscriber
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
```
#### Pass large structs using pointers
>so function parameters receive a copy of the arguments from
> the function call, even for structs... <br>
> if you pass a big struct with a lot fo fields, won't
> that take up a lot of the computer's memory?<br>
> Yes, it will. it has to make room for the original struct and the copy.

Functions receive a copy of the arguments they’re called with, even if
they’re a big value like a struct.

That’s why, unless your struct has only a couple small fields, it’s often a
good idea to pass functions a pointer to a struct, rather than the struct itself.
(This is true even if the function doesn’t need to modify the struct.) When
you pass a struct pointer, only one copy of the original struct exists in
memory. The function just receives the memory address of that single
struct, and can read the struct, modify it, or whatever else it needs to do, all
without making an extra copy.

Here’s our defaultSubscriber function, updated to return a pointer, and
our printInfo function, updated to receive a pointer. Neither of these
functions needs to change an existing struct like applyDiscount does. But
using pointers ensures that only one copy of each struct needs to be kept in
memory, while still allowing the program to work as normal.
```go
package main

import "fmt"

func main() {
	//sub is a struct pointer
	sub := defaultSubscriber("ahmed")
	applyDiscount(sub)
	printInfo(sub)
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

//defaultSubscriber return a pointer to a subscriber
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s

}

//printInfo takes a pointer to subscriber
func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

//applyDiscount update the rate fo a subscriber
//it takes a pointer to subscriber
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

```

#### Struct literals
Go offers struct literals to let you create a struct and set
its fields at the same time, `myCar :=car {name:"Corvette", topSpeed: 337}`

* make use of packages
* Export struct and its fields
* anonymous struct within struct
```go
//Package magazine
package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

//Employee track the Names and salaries of our employees
type Employee struct {
	Name   string
	Salary float64
	Address
}

//Address store the mailing addresses for both sub and emp
type Address struct {
	Street, City, State, PostCode string
}

//DefaultSubscriber return a pointer to a Subscriber
func DefaultSubscriber(Name string) *Subscriber {
	var s Subscriber
	s.Name = Name
	s.Rate = 5.99
	s.Active = true
	return &s

}

//PrintInfo takes a pointer to Subscriber
func PrintInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly Rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

//ApplyDiscount update the Rate fo a Subscriber
//it takes a pointer to Subscriber
func ApplyDiscount(s *Subscriber) {
	s.Rate = 4.99
}
```
### Embedding structs
An inner struct that is stored within an outer struct using an anonymous
field is said to be embedded within the outer struct. Fields for an embedded
struct are promoted to the outer struct, meaning you can access them as if
they belong to the outer struct.

So now that the Address struct type is embedded within the Subscriber
and Employee struct types, you don’t have to write out
subscriber.Address.City to get at the City field; you can just write
subscriber.City. You don’t need to write employee.Address.State;
you can just write employee.State.