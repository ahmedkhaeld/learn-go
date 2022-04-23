# Closures
* a closure is really about
functions that live inside functions
and refer to the enclosing functions
data

* Go functions are first-class citizens. Functions can be assigned to variables, stored in collections, created and deleted dynamically, or passed as arguments.
* Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
* A nested function, also called an inner function, is a function defined inside another function. An anonymous function is a function definition that is not bound to an identifier. Anonymous functions are often arguments being passed to higher-order function


### Variable Scope vs Lifetime
* scope is static, based on the code at compile time
* lifetime depends on program execution(runtime)

1. scope: <br>
for example
```go
package xyz
func doIt() *int{
	var b int
	// do something
    return &b
}
```
variable b can only be seen inside doIt, but its value will live on<br>
The value(object) will live so long as part of the program keeps a pointer to it.


2. lifetime: <br>
   in some older languages the only way
you ever got lifetime
greater than a function was you know you
use like a new operator or something
to deliberately allocate something in
the heap and you pass pointers around. <br>
**_****the code above****_** _in a traditional language like c you'd
look at it and say
you're declaring a local variable b and
you're referring a pointer to it
and that's going to blow up and the
reason it's going to blow up
gets back to this notion of stack frames
right because i have a stack frame for
some function that's calling do it
and when i start calling do it i create
a new stack frame
and it has things like the return
address and stack pointer
and then somewhere in here it has the
actual storage for the variable b
now if i return a pointer to that so
somebody else has a way to point to that
particular location
but i return from the function
and that space sort of goes away i mean
it doesn't physically go
away but the the next thing i do is
maybe i call some other function
you know db and the function db
reuses all this stack space but somebody
has a pointer into the middle of it
and now we have corruption this is
memory corruption_
#### why is this okay in GO?
the answer is

that in go as soon as the compiler sees
this it says hey b
is going to have to live longer than
this function
and what in go the term that's used is
escape analysis<br>
what's called
escape analysis to figure out
that the value you put in b we're going
to return a pointer to it
and it's going to live beyond the scope
of the function its lifetime is going to be as
long as whoever has that pointer
and so of course it immediately
allocates b on the heap not on the stack

allocating stuff in the heap one it's
inefficient in terms you have to go
through a pointer to get to it all the
time
and two it makes more work for the
garbage collector

go's approach is to allocate as much as
possible on the stack
but when it has to to put it in the heap
when the lifetime will exceed the scope
of
whatever context that variable is

### what is a closure?
A closure is when a function inside another function "closes over" one or more local 
variables of the outer function<br>
* here an example:<br>
  it's a function returning a function<br>
  the inner function here
is an anonymous function it doesn't have
a name it's just a function literal with
some code in it
okay it uses variables a and b it didn't
declare a and b
it just changes them the variables a and
b
<br>when i return this inner function
it's going to continue to refer to a and
b
even though the function fib is going to
go away right<br>
  fib runs returns a and b and theory are
gone
okay but they're not really they're
going to live on and they have to live
on
because this inner function that we're
returning closes over a
and b and so it keeps a and b alive
and that's why it's called a closure

```go
func fib() func int {
	a,b := 0,1
	
	return func() int{
		a,b = b, a+b
		return b
        }

}
```
The inner function get a ****reference**** to the outer function's vars<br>
Those variables may end up with a much longer _lifetime_ than expected as long
as there's a reference to the inner function

this piece of text right here is NOT the
closure, but it is what's being return
```return func() int{
a,b = b, a+b
return b
}
```
---
### closures: how they work

* left side of the slide _page3_<br>
`f := fib`
f is declared and assigned the value of
fib, now here i'm just using fib without
any parentheses i'm not calling fib i'm
just taking the name
fib which refers to the function and
assigning it to the variable f so what is f well f is essentially a
function pointer
it's something that points at the code
that is the function fib<br>
`f()` if i call f so now i put the
actual parentheses there to
indicate a function call, i'm calling the function fib f and fib
are just two names for the same thing

* right side of the slide _page3_ <br>
`f := fib()`
  call fib and assign the result from fib to f<br>
  f is no longer just a function
it's a closure<br>
  what happens in a closure there's really two pieces<br>
  that's part of what's returned from fib
is the anonymous function
that will be called when i call
f 
<br>but there's another piece that holds on to the information about
where to find
a and b because that function the
anonymous inner function here
 it changes a and b and it returns b
and it can't do that if it doesn't know
where they are


---
###### playground
```go
package main

import (
	"fmt"
)

func fib() func() int{
	a, b := 0, 1
	return func() int {
        a,b = b, a+b 
		return b 
	}
}

func main() {
	f :=fib()
	
	for x := f(); x < 100; x=f(){
		fmt.Println(x)
    }
}
// output:
//1  2  3  5  8  13  21  34  55  89

```
- call fib which return an anonymous func, but this anonymous here return a closure and put that into f
- in the for loop `for x := f(); x < 100; x=f()` started the initialization by calling f once
- and then change x each time by calling f again
- here if is a call, which is doing arithmetic on `a` and `b` and returning a value
- by looping each time it's reusing these variable `a` and `b`
- the variable `a` & `b` belong to the fib func they are closed over by the anonymous func
which holds onto them and as long as this `f` exists `a` & `b` they continue to have their
existence, so their values keep changing

---

```go
package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f, g := fib(), fib()
	
	fmt.Println(f(), f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g(), g())
}

// output:
// 1 2 3 5 8
// 1 2 3 5 8
```
get the same numbers twice over and the reason for that is that f and g get different a and b variables
<br> f closure has one copy of a and b, and it mutates them over time, and the
same with g they're distinct

so every time i call fib i get a new
fibonacci number generator that has its
own state
because it closed over unique copies of
a and b


---
#### issues with closures
called the closure
in the same loop iteration it was in the
loop body
```go
package main

import "fmt"

func do(d func()) {
	d()
}
func main() {

	//iterate 4 times
	for i := 0; i < 4; i++ {
		// v holds a value which of type func that print current i value and address
		v := func() {
			fmt.Printf("%d @ %p \n", i, &i)
		}

		// pass v to do() that immediately call itself
		do(v)
	}
}

// 0 @ 0xc00001e0c8
// 1 @ 0xc00001e0c8
// 2 @ 0xc00001e0c8
// 3 @ 0xc00001e0c8

```
the print out i have the values 0 1 2 3 those are the values of i
they all live at the same address

---
```go
package main

import "fmt"

func main() {
	s := make([]func(), 4)

	// loop 4 times, til slice has an index [4] then break out
	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	// when i get to the second loop i reached 4
	// repeat call the func in the slice four times
	// every time i call the closure is going to be the same i
	// and i has that final value 4
	for i := 0; i < 4; i++ {
		s[i]()
	}
}

//4 @ 0xc00001e0c8
//4 @ 0xc00001e0c8
//4 @ 0xc00001e0c8
//4 @ 0xc00001e0c8

```

>  if there's time between when the closure is created
and when it's called the variable i'm
pointing to could change, and it did
i created the first closure when 'i' was
zero
and then one when i was one and so on i
kept doing that but by the time i went
to
call those closures i was four and all
the closures
used the same 'i' and they got the same
value 

> this is a bug in my program
but it's a bug because i really didn't
understand the nature of a closure


* a simple way to fix this

```go
package main

import "fmt"

func main() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i // fix closure capture
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &2i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

//0 @ 0xc0000b8000
//1 @ 0xc0000b8008
//2 @ 0xc0000b8010
//3 @ 0xc0000b8018
```
not only does it produce 0 1 2
3 but it produces four distinct
addresses
#### how did that happen?
because essentially i'm saying
i'm declaring i2 to be i 

i'm creating this new
variable i2 to initialize it to the
variable i in the loop because it's now a distinct variable

this i2 is a new variable every time
through the loop
its location is different than i and its
location is different each time i go
through the loop
they're distinct and so now each closure
has its own value of i because it has
its own reference to a copy of i that
was placed in its own version of i2
okay every closure has a different
address for i2

---
##  Useful Ways to Use Closures in Go

### 1. Isolating data
Lets say you want to create a function that has access to data that persists even after the function exits. For example, you want to count how many times the function has been called, or you want to create a fibonacci number generator, but you don’t want anyone else to have access to that data (so they can’t accidentally change it). You can use closures to achieve this.

```go
package main

import "fmt"

func main() {
  gen := makeFibGen()
  for i := 0; i < 10; i++ {
    fmt.Println(gen())
  }
}

func makeFibGen() func() int {
  f1 := 0
  f2 := 1
  return func() int {
    f2, f1 = (f1 + f2), f2
    return f1
  }
}
```

### 2. Wrapping functions and creating middleware
Functions in Go are first-class citizens. What this means is that you can not only create anonymous functions dynamically, but you can also pass functions as parameters to a function. For example, when creating a web server it is common to provide a function that processes a web request to a specific route.
```go
package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}
```
While this code doesn’t require a closure, closures are incredibly helpful if we want to wrap our handlers with more logic. A perfect example of this is when we want to create middleware to do work before or after our handler executes.
> What is middleware?<br>
Middleware is basically a fancy term for reusable function that can run code both before and after your code designed to handle a web requst. In Go these are typically accomplished with closures, but in different programming languages they may be achieved in other ways.
Using middleware is common when writing web applications, and they can be useful for more than just timers (which you will see an example of below). For instance, middleware can be used to write code to verify if a user is logged in once, then apply it to all of your member-only pages.

```go
package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/hello", timed(hello))
  http.ListenAndServe(":3000", nil)
}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    f(w, r)
    end := time.Now()
    fmt.Println("The request took", end.Sub(start))
  }
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}
```
Notice that our timed() function takes in a function that could be used as a handler function, and returns a function of the same type, but the returned function is different that the one passed it. The closure being returned logs the current time, calls the original function, and finally logs the end time and prints out the duration of the request. All while being agnostic to what is actually happening inside of our handler function.

Now all we need to do to time our handlers is to wrap them in timed(handler) and pass the closure to the http.HandleFunc() function call.

###  3. Accessing data that typically isn’t available
A closure can also be used to wrap data inside of a function that otherwise wouldn’t typically have access to that data. For example, if you wanted to provide a handler access to a database without using a global variable you could write code like the following.
```go
package main

import (
  "fmt"
  "net/http"
)

type Database struct {
  Url string
}

func NewDatabase(url string) Database {
  return Database{url}
}

func main() {
  db := NewDatabase("localhost:5432")

  http.HandleFunc("/hello", hello(db))
  http.ListenAndServe(":3000", nil)
}

func hello(db Database) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, db.Url)
  }
}
```
Now we can write handler functions as if they had access to a Database object while still returning a function with the signature that http.HandleFunc() expects. This allows us to bypass the fact that http.HandleFunc() doesn’t permit us passing in custom variables without resorting to global variables or anything of that sort.


