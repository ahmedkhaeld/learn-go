# Pointer/Value Semantics
* pointer and value semantics and how they work in go
* when to use pointers and when to use values? and what's the trade-off?

### fundamental of pointers and values
* POINTERS:<br>
when we use a pointer to something in many cases
what we mean by that is we want to share
it. we don't want to copy it we want to
share it so that different
functions or parts of the program are
modifying the same data
* VALUES: <br>
  we're copying it it's
because we don't want to share it
okay now in a lot of cases
not sharing things is safer and that's
particularly true when we talk about
concurrency. _the best approach to concurrency is don't share anything_


#### common uses of pointers
* **some objects can't be copied safely (mutex)<br>**
  if we can't copy it we must share it
and we must therefore use reference
semantics
* **some objects are too large to copy efficiently (consider pointers when size>64bytes)** <br>
  once an object gets large enough it's
more expensive to copy it
than it is to use a pointer to get to it
and there's a cost to pointers
right the short version is if i use a
pointer to get to something
well first i have to read the pointer
variable out of memory
and then i have to do a second memory
read to get the actual data
* **some methods need to change(mutate) the receiver**<br>
  we may want to have a method that
changes the object it's called
on and that's going to have to be
reference semantics obviously if the
method is called with value semantics we
get a copy of the object we calculate
something
but we can't change the object
permanently from a method uses a value
semantic
* **when decoding protocol data into an object**<br>
   * (JSON, etc.; often in variable argument list)
  
```go
var r Response
err:=json.Unmarshal(j, &r)
```
example decode protocol data right
so when we did our json unmarshall
we had to pass a reference so that the
json code would have a place to copy the
data back into
when it decodes the json
* **when using a pointer to signal a "null" object**<br>
  if i build a tree structure and we've
done that a couple times already
typically we have pointers to nodes and
we want the whole concept of being able
to have a nil pointer to say
hey this node doesn't have children

#### Must not copy
* any struct with a mutex and wait-group must be passed by reference:
```go
type Employee struct{
	mu sync.Mutex
	Name string
	
}

func do(emp *Employee){
	emp.mu.Lock()
	
	defer emp.mu.Unlock()
}
```

#### May copy
any small struct under 64 bytes probably should be copied:
```go
type Widget struct{
	ID int
	Count int
}

// Expend take a Widget instance and modify it, then returns it back
func Expend(w Widget) Widget{
	w.Count
	
	return w
}
```
it's perfectly okay to write a function
that copies a struct and changes it and
passes it back. f i don't return it the
modification will never be visible


on a system where we have
 a data structure that represented a 4k
block of data all right and if you start
copying 4k blocks of data your program
is going to slow down
really badly and so you're not going to
want to do that you're going to want to
use a pointer
to that thing that's 4k and you're going
to want to deal with it
through the pointer whether you meant to
share it or not doesn't really matter
the cost of copying the 4k block of data
from one place in memory to another is
just going to take a huge whack on the
performance of your program
so that's a case where you're going to
use discretion and use reference
semantics for efficiency

---
#### Semantic consistency
> if a thing is to be shared, then always pass a pointer
```go
 type Employee struct{}
 f1(emp *Employee) 
 f2(emp *Employee)
 f3(emp  Employee) // passes a copy
 f4(emp *Employee) // changes are LOST!
```
example here is i'm gonna i've
got some process i'm gonna do employee
relocation
in my software that manages the
employees in the company
and to do that i'm gonna call a bunch of
functions so i start out with my
employee record
and i'm going to pass it around through
a pointer fine i had a reason for doing
that
which is great so long as all the
functions take pointers

and the problem here is well somebody
said hey for function 3
i'm going to just take a copy and
doing that causes a problem because
let's suppose i pass that to function 4
function 4 makes a change it changes my
employee record
ok but that change will never be seen
above function 3 because function 3 made
a copy
so in the middle of this call chain i
made a copy of the data
and below that i'm mutating the copy and
any change made really in f3 or below is
no longer visible
to the actual original employee record
that was passed into the top
so if we're going to go down the road of
using reference semantics particularly
if we're doing it for efficiency
we need to make sure we're consistent in
our function signatures that they're all
on board
using reference semantics

---
#### Stack Allocation
* stack allocation is more efficient.
* Accessing  variable directly is more efficient than following a pointer
* Accessing a dense sequence of data is more efficient than sparse data<br>
  it's better to have a bunch of
data in a slice than a bunch of data
spread around in memory

####### Go would prefer to allocate on the stack, but sometimes can't
#### Heap Allocation
* a function returns a pointer to a local object
* a local object is captured in a function closure
* a pointer to local object is sent via a channel
* any object is assigned into interface
* any object whose size is variable at runtime(slice)

---
#### For Loops
* the value returned by `range` is always a copy
```go
for i, thing :=range things{
	// thing is a copy
	thing.which=whatever // this mutation will disappear when the loop is done
}
```
when you use the range operator in a for loop
the value you get is a copy so thing
in this top case is a copy of an element
of a slice or an array <br>
if i change it in the copy
that change will no longer be visible
outside this loop iteration

* use the index if you need to mutate the element:
```go
for i := range things{
	things[i].which= whatever  
}
```
it's much more useful to do the
one variable range
get the index use the index to subscript
into the original slicer array
and change it this change down here will
be visible
after the loop is done


#### Slice safety
anytime a function mutates a slice that's passed in, we must return a copy

* when to use append you need to assign the result of append
back to the original slice variable, because append could cause the slice to get reallocated if it grows beyond a
certain limit
* if i have a function that takes a slice
and i want that function to be able to
modify the slice maybe it's going to put
some more stuff in it
then i have to return the slice
```go
func update(things []thing) []thing{
	things =append(things, x)
	return things
}
```

--- 
```go
package main

import "fmt"

func main() {
	// items is a slice of arrays of two bytes
	items := [][2]byte{{1, 2}, {2, 2}, {3, 4}}
	var a [][]byte // slice of slice

	for _, item := range items {
		a = append(a, item[:]) // item[:] take the whole thing
	}

	fmt.Println(items)
	fmt.Println(a)
	//[[1 2] [2 2] [3 4]]
	//[[3 4] [3 4] [3 4]]

}
```
* items: is a slice the type of the slice is an array of two bytes
* a slice of slice of byte which means i'm actually going to create
a slice from each of the array entries
in items and put that into a

* explain the result  what's going on here <br>
  * remember item is a copy of an entry in the items slice
  * so item is a two byte array it's a copy of a two byte array at a particular location in memory
  * every time i go through the loop that copy is in the same place
  * and my slice gets a reference to that location
  * so my slice here right is not picking up the values it's picking up a pointer to where the values are stored
  * when my for loop is done what does item end up having in it it has the last thing it had going through the for loop and the last thing in the for loop was the last two byte array
  * so all three of these slices
that are inside the big a slice because
it's a slice of slices
all three of those slices point to the
same array with the same values
  * we're keeping those references
to the loop iteration variable
item and using them afterwards when the
loop is done
and they all refer to the final value of
item
  
>one way to fix this is to make a copy
```go
package main

import "fmt"

func main() {
	// items is a slice of arrays of two bytes
	items := [][2]byte{{1, 2}, {2, 2}, {3, 4}}
	var a [][]byte // slice of slice

	for _, item := range items {
		i := make([]byte, len(item))
		copy(i, item[:])
		a = append(a, i)
	}

	fmt.Println(items)
	fmt.Println(a)

	//[[1 2] [2 2] [3 4]]
	//[[1 2] [2 2] [3 4]]

}

```
each time through here i'm going to make a new slice
which means it has this new invisible backing array
and i'm going to copy all the values from items 
