# Slices
* talk about the difference between a nil slice and an empty slice
* illustrate some things about the difference between capacity and the length of the slice

```go
package main

import "fmt"

func main() {
	var s []int  // nil slice
	t := []int{} // empty slice

	u := make([]int, 5)    // make slice of length 5
	v := make([]int, 0, 5) // make slice of length 0 with capacity 5

	fmt.Printf("%d, %d, %T, %5t, %#[3]v \n", len(s), cap(s), s, s == nil)
	// 0, 0, []int,  true, []int(nil)

	fmt.Printf("%d, %d, %T, %5t, %#[3]v \n", len(t), cap(t), t, t == nil)
	// 0, 0, []int, false, []int{}

	fmt.Printf("%d, %d, %T, %5t, %#[3]v \n", len(u), cap(u), u, u == nil)
	//5, 5, []int, false, []int{0, 0, 0, 0, 0}

	fmt.Printf("%d, %d, %T, %5t, %#[3]v \n", len(v), cap(v), v, v == nil)
	//0, 5, []int, false, []int{}

}

```

print out the `length of a slice`,
`the capacity of the slice`, `the type,
whether the slice is nil or no`t, and then
finally i want it to give me it's it's
debugging view of what the slice looks
like in other words the
`pound v view of the variable`

---
page #2
##### showing what a slice descriptor looks like for some slice
1. `a:=[]int{0,1,2,3}` <br>
   is a slice descriptor that has length four capacity four
and a pointer which points to the first
element in the slice which is just a
pointer to a sequence in memory because we know the values in the slice are all going to get lined up in memory
one after another so it's enough to have a pointer to the first one
2. `s` is a nil slice; slice descriptor just has zeros<br>
   so it has zero length zero capacity and
a null pointer
which is very convenient because when we
declared you know var
s int and it wants to provide a default
value for the slice variable
it just writes a bunch of zeros into
memory
3. `t` empty slice<br>
   the slice descriptor zero length zero capacity but it has a
pointer to something. what go does under the hood is it has a pointer off to a special location
in memory
4. `u` a regular slice <br>
   it has five zeros so it has length five
capacity five
and it has a pointer to the actual
storage 
5. `v` is a slice that has zero length capacity five and it has a pointer

---
### why do we care  where it matters whether it's nil or empty?

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a []int

	j1, _ := json.Marshal(a)
	fmt.Println(string(j1)) // null

	b := []int{}

	j2, _ := json.Marshal(b)
	fmt.Println(string(j2)) // []
}

```
* if i take the nil slice
and encode it into json it's going to come out as null because there's nothing there
* if i take the empty slice and encode it
as json
i'm going to get a pair of square
brackets with nothing in them in other
words
an array or list of length 0.
> the same thing with a map

>a mistake to ask is `a == nil`? that is the wrong question!<br>
> to check is some slice if full or not<br>
> that's the wrong question because there's two cases where a slice can be empty<br>
> **one:** it can be a nil slice.  **two:** it can be an empty slice


> the right question is: ask the length of a slice is equal to zero `if len(a) == 0{}` <br>
> because that's what we really want to know is the slice empty or not
and the length is 0 if it's nil and the
length of 0  it's empty

#### one more thing on u & v
 `u :=make([]int,5)` `v:=make([]int, 0, 5)`
* if i've created `v` this way it has space and i can append to it 
* if i create a slice`t` of length 5 it starts out with five zeros. if i do an append where does the
next thing get appended well it gets appended after the five zeros
> if you want to reserve space for a slice make sure you put the zero `v`

> **it is perfectly okay to append to a nil slice** `var s []int`<br>
> `s=append(s, something)`

---
### length vs capacity
```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println("a=", a) //a= [1 2 3]

	s := a[:1]
	fmt.Println("s=", s) // s= [1]
	fmt.Println(len(s))  // 1
	fmt.Println(cap(s))  // 3

	c := s[0:2]
	fmt.Println("c=", c) // c= [1 2]
	fmt.Println(len(c))  // 2
	fmt.Println(cap(c))  // 3

}

```
s&c has the capacity 3 of the underlying array which is a

* use the three index operator
```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println("a=", a) //a= [1 2 3]

	//three index slicing operator controls both the length and capacity of the slice
	d := a[0:1:1]       // [i:j:k] len=j-i  cap=k-i
	fmt.Println(len(d)) // 1
	fmt.Println(cap(d)) // 1
}
```

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[0:1]
	c := b[0:2]

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	//a[0xc00001a228] = [1 2 3]
	//b[0xc00001a228] = [1]
	//c[0xc00001a228] = [1 2]
}

```
not surprisingly look at these addresses
they're all the same address so `a` took the address of the first element
`b`took the address of the first
element that `b` refers to
which is the same first element of `a` and
the same thing with `c`


* when we do **append()**

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[0:1]
	c := b[0:2]

	fmt.Printf("a[%p] = %v\n", &a, a) //a[0xc00001a228] = [1 2 3]
	fmt.Printf("b[%p] = %[1]v\n", b)  //b[0xc00001a228] = [1]
	fmt.Printf("c[%p] = %[1]v\n", c)  // b[0xc00001a228] = [1 2]

	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
	//a[0xc00001a228] = [1 2 5]
	//c[0xc00001a228] = [1 2 5]

}
```
explain:<br>
do an append
to c and i look at my values for **a**
and **c** and they're the same okay
well c was a slice of a's first two
elements it has
the one and the two ,but it has capacity
3
which means there's one more spot in c
to put something
but that spot is part of **a** and so when i
append to **c**
what i'm actually doing is i'm mutating
the **a** that's underneath it i'm
overwriting
**a's** third value when i append to **c**
and this is again this is part of what i
think is unintuitive


* append after using three index operator
```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[0:1]
	//c is a slice of length two capacity two
	//it represents the first two elements of a,
	// and it's got no extra space to put anything
	c := b[0:2:2]

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	//a[0xc0000b8000] = [1 2 3]
	//b[0xc0000b8000] = [1]
	//c[0xc0000b8000] = [1 2]

	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
	//a[0xc0000b8000] = [1 2 3]
	//c[0xc0000be020] = [1 2 5]
}

```
what
happens is by limiting
how much capacity **c** had if i append to
it. I force a reallocation there's not enough
space
in the **c** that i've just created here
with the three index slicing operator
to put another element so the append
forces a new piece of memory to be
allocated
we copy what was in **a** into it and then
we put this new element at the end in a
different location
so **a** is not changed and **c** now has the
values of [1 2 5]






