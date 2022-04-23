# Understanding Allocations: the Stack and Heap

* **How do i know if my variable lives on the stack or the heap?**
the answer is you don't.

* **Does it matter to know where the variable is ?** <br>
  it does not matter that you
know where your variable is as far as
the correctness of your program is
concerned go is going to make sure that
your variable lives in the correct place
either the stack or the heap 
* **Does affect the performance of your program**<br>
  you may need to know and the
reason it affects the performance is
that anything that's on the heap gets
managed by the garbage collector the
garbage collector is very good but it
does cause some latency and and it can
cause latency for your whole program not
just the part creating garbage you may be concerned about how much garbage you're putting on the heap 

* **Do You need to know?** <br>
    * if your program is NOT FAST ENOUGH according to the benchmarks 
    * excessive heap allocations
> Optimize for correctness first, not performance

* **The Stack**

what's happening in memory as we advance through this
```go
package main

import "fmt"

func main() {
	n := 4
	n2 := square(n)

	fmt.Println(n2)
}

func square(x int) int {
	return x * x
}
```
what's gonna happen is go it's
gonna take a certain amount of memory
from this stack that's large enough to
hold all  the function local variables
and this function main we have n and n2
we see n has a value of 4 n2
as just gonna start here as value
zero it's got these memory addresses
this section here this section of the
stack is called a stack frame and this
is this stack frame is for a particular
function in this case func main when we
it move forward we call this function
Square and we pass in the value of n
which is four goes going to take some
more stack memory it frames off another
section here to create this stack frame
for this function square we have a new
memory address here and has this value 4
and this variable called X when this
function returns it's going to multiply
four by four so we have 16 which gets
put into this variable n2


* **The Stack with pointers** <br>
  program here that has a value of 4 and
we're gonna take the address of n and
we're gonna share it down into this
function Inc, so this function Inc takes
this pointer variable X it's going to
dereference the pointer and increment
the value of the pointer points to
```go
package main

import "fmt"

func main() {
	n := 4
	inc(&n)
	fmt.Println(n)
}

func inc(x *int) {
	*x++
}

```
> sharing down typically stays on the stack, meaning passing pointers,
> passing references to things typically stays on the stack.

---
### Returning Pointers?
what if we go the other way what if we start returning pointers
```go
package main

import "fmt"

func main() {
	n := answer()
	fmt.Println(*n/2)
}

func answer() *int{
	x := 42
	return &x 
}

```
in this program here in func main the  first thing we do is we call another
function answer<br>
answer has a variable
called X with the value of 42 and we
return the address of X the address of X
is stored in this variable in then main
is going to dereference that pointer get
the value of points

walking
through that one step at a time
when we first call main it has enough
room for this variable in just we'll say
it's nil to begin with we call answer
and we had we have a stack frame here
for answer we have this variable X with
the value of 42 and then we return the
address of X so now in has this address
here with the address of X which is 4 4
7 7 0 but right now we've already got a
problem notice we have a pointer that is
pointing down into the invalid section
of memory so what happens when we call
print line when we call print line we
know it's going to reclaim this space so
we dereference the pointer we get that
value we divide it by 2 we get 21 we're
gonna call print line with the value of
21

<img src="https://user-images.githubusercontent.com/40498170/164790096-158f7c0d-e1dd-42de-9006-0fcdf7ee540c.png" height="300" width="300"> <img src="https://user-images.githubusercontent.com/40498170/164790176-d55328d8-5a4f-402a-83d6-0b935701829b.png" height="300" width="300">

what
actually would happen is this same
program now in this program when it
first starts we got our stack frame for
main that's great we call answer but the
compiler knows it was not safe to leave
that variable on the stack it cannot be
in the stack frame for that function
answer so instead X is over here X gets
declared somewhere on the heap when we
return the address of X while we're
returning is this address here which
points out here out of the heap
so when we call print line there's no
problem we don't accidentally clobber
that variable that we had just asked for
so we say this escapes to the heap

to be clear I it says escapes to
the heap but it doesn't get moved it's
not like gets moved at runtime this
happens at compile time this very
variable X is going to be constructed on
the heap initially because the compiler
knows so here's my second major point
sharing up typically escapes to the heap
and I say sharing up I mean like
returning pointers returning references
returning things that have pointers in
them will typically escape to the heap
> sharing up typically escapes to the heap, like returning pointers, returning references 

when possible, the Go compilers will allocate variables that are local to a function in
that function's stack frame.
However, if the compiler cannot prove that the variable is not referenced after the 
function returns, then the compiler must allocate the variable on the garbage-collected heap
to avoid dangling pointers errors.


### Escape Analysis
what the compiler does is it's
going to look at our code and it's gonna
see just do any of these variables need
to be put on the heap?

In the current compilers, if a variable has its address taken, that variable is a candidate
for allocation on the heap. However, a basic escape analysis recognizes some cases when such
variables will not live past the return from the function and can reside on the stack.

* Let's Ask the Compiler!
`go build -gcflags "-m"`
 1.  first example where I use
pointers, but they're staying on the
stack when I build with - M we can see
right here that the address of n does
not escape on line 5 when you take the
address of this variable it could
potentially go to the heap escape
analysis has proven that that is not
going to be referenced after func main
returns, so it's fine

  <img src="https://user-images.githubusercontent.com/40498170/164793259-3834b6a6-c366-41d1-84d7-f4de410c9bbc.png" width="300" height="300">

2. in the other example where we do escape
to the heap right here when we line 10
says when you take the address of X
that's gonna cause it to escape to the
heat because this value X is going to be
referenced after func answer returns and
like I mentioned before it doesn't get
moved at runtime right at compile time
it says right here move to the heap X we
know this variable has to live on the
heap

  <img src="https://user-images.githubusercontent.com/40498170/164793638-2e2024a1-4f89-4bd1-b884-21f472997c0c.png" height="300" width="300">

`go build -gcflags "-m=2"` for more verbose

---
#### when are values are constructed on the heap?
1. when a value could possibly be referenced after the function that constructed the value returns
2. when the compiler determines a value is too large to fit on the stack
3. when the compiler doesn't know the size of a value at compile time

#### some commonly allocated values on the heap
* values shared with pointers
* variables stored in interface variable
* Func literal variables
  * variables captured by a closure
* backing data for Maps, Channels, Slices, and Strings


#### which stays on the stack?
've got two programs they both do basically the same thing

<img src="https://user-images.githubusercontent.com/40498170/164795446-12146745-6089-4489-b94c-36c466e89f1d.png" width="300" height="300">

1. program on the left here the first
thing it's going to do is going to call
this function read this function read is
going to return a slice of bytes
now this slice has a constant size it's
only 32 bytes it's so it's small it
could fit on the stack and it has a
constant size so the compiler knows what
to do with it and so this function is
gonna create the slice assume that we
like put some data in here and then we
return it 
2. function on the right
does it differently here on the right we
create the slice of bytes up in Maine
and we pass that slice down into read
the read function takes the slice and
it's going to write into it and then
return now in this scenario like when
you run this what you find is that here
this one on the right is gonna stay on
the stack specifically this this slice
variable and the 32 bytes that are
behind it whereas on the left every time
you call the read function this slice of
bytes is going to be referenced after
func read returns and so this will
frequently be going to the heap instead


* this actually explains a question  why is the i/o reader interface the way it is
<br>i/o reader is an interface that's something
  you can read its frequently implemented
  by files network connection bytes
  buffers things like that
```go
type Reader interface{
	Read(p []byte) (n int, err erro)
}
//make
//the slice and you pass the slice into
//the read method and then it returns a
//number to tell you how much of your
//slice it filled
```
instead of 
```go
type Reader interface{
	Read( n int) (b []byte, err error)
}

// here  it
//return me the slice I'm saying I want to
//read some bytes it would make sense that
//I would say hey please read some bytes
// and it says here's your bytes
```
this second example we would have so much garbage on the heap
