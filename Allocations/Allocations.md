# Understanding Allocations: the Stack and the Heap
in Program's Memory, there's two kinds of memory, stack and heap

Go allocates memory in two places: a global heap for dynamic allocations and a local stack for
each goroutine. Go prefer allocation on the stack-- most of the allocations withing a given go 
program will be on the stack, it's cheap because it only requires two cpu instructions:
one ot push onto the stack for allocation, and another to release from the stack.

Unfortunately not all data can use memory allocated on the stack. Stack requires that the 
lifetime and memory footprint of a variable can be determined at compile time. otherwise a 
dynamic allocation onto the heap occurs at runtime. malloc must search for a chunk of free memory large enough to hold the new value. Later down the line, the garbage collector scans the heap for objects which are no longer referenced. It probably goes without saying that it is significantly more expensive than the two instructions used by stack allocation.


## Escape Analysis
The compiler uses a technique called escape analysis to choose between these two options. The basic idea is to do the work of garbage collection at compile time. The compiler tracks the scope of variables across regions of code. It uses this data to determine which variables hold to a set of checks that prove their lifetime is entirely knowable at runtime. If the variable passes these checks, the value can be allocated on the stack. If not, it is said to escape, and must be heap allocated.

Go manages memory allocation automatically. This prevents a whole class of potential bugs, but it doesn’t completely free the programmer from reasoning about the mechanics of allocation. Since Go doesn’t provide a direct way to manipulate allocation, developers must understand the rules of this system so that it can be maximized for our own benefit.

****The rules for escape analysis aren’t part of the Go language specification.****
For Go programmers, the most straightforward way to learn about these rules is experimentation. The compiler will output the results of the escape analysis by building with 
`go build -gcflags '-m'`

Let’s look at an example:
```go
package main

import "fmt"

func main() {
        x := 42
        fmt.Println(x)
}
```
`$ go build -gcflags '-m' ./main.go`
> command-line-arguments <br>
>./main.go:7: x escapes to heap <br>
>./main.go:7: main ... argument does not escape

See here that the variable x “escapes to the heap,” which means it will be dynamically allocated on the heap at runtime
x escapes because it is passed to a function argument which escapes itself 

## Patterns which typically cause variable to escape to the heap

* **Sending pointers or values containing pointers to channels**.
  At compile time there’s no way to know which goroutine will receive the data on a channel. Therefore the compiler cannot determine when this data will no longer be referenced.
* **Storing pointers or values containing pointers in a slice**
  An example of this is a type like []*string. This always causes the contents of the slice to escape. Even though the backing array of the slice may still be on the stack, the referenced data escapes to the heap.
* **Calling methods on an interface type.**
  Method calls on interface types are a dynamic dispatch — the actual concrete implementation to use is only determinable at runtime. Consider a variable r with an interface type of io.Reader. A call to r.Read(b) will cause both the value of r and the backing array of the byte slice b to escape and therefore be allocated on the heap.

## Some Pointers
The rule of thumb is: pointers point to data allocated on the heap. Ergo, reducing the number of pointers in a program reduces the number of heap allocations. This is not an axiom, but we’ve found it to be the common case in real-world Go programs.

in many cases copying a value is much less expensive than the overhead of using a pointer. “Why” you might ask?
* **The compiler generates checks when dereferencing a pointer.**<br>
  The purpose is to avoid memory corruption by running panic() if the pointer is nil. This is extra code that must be executed at runtime. When data is passed by value, it cannot be nil.
* **Pointers often have poor locality of reference**<br>
All of the values used within a function are collocated in memory on the stack. Locality of reference is an important aspect of efficient code. It dramatically increases the chance that a value is warm in CPU caches and reduces the risk of a miss penalty during prefetching.
* **Copying objects within a cache line is the roughly equivalent to copying a single pointer.** <br>
  CPUs move memory between caching layers and main memory on cache lines of constant size. On x86 this is 64 bytes. Further, Go uses a technique called Duff’s device to make common memory operations like copies very efficient.

Pointers should primarily be used to reflect ownership semantics and mutability. In practice, the use of pointers to avoid copies should be infrequent. Don’t fall into the trap of premature optimization. It’s good to develop a habit of passing data by value, only falling back to passing pointers when necessary. An extra bonus is the increased safety of eliminating nil.

Reducing the number of pointers in a program can yield another helpful result as **the garbage collector will skip regions of memory that it can prove will contain no pointers.** For example, regions of the heap which back slices of type []byte aren’t scanned at all. This also holds true for arrays of struct types that don’t contain any fields with pointer types.

Not only does reducing pointers result in less work for the garbage collector, it produces more cache-friendly code. Reading memory moves data from main memory into the CPU caches. Caches are finite, so some other piece of data must be evicted to make room. Evicted data may still be relevant to other portions of the program. The resulting cache thrashing can cause unexpected and sudden shifts the behavior of production services.
