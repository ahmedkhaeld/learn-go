# GOROTINES
* Creating goroutines
* Synchronization
    * waitGroups
    * Mutexes
* Paralleism
* Best Practices
- - -  
<br/>

most programming
languages working with os threads or use operating
system threads. And what that means is that they've got an individual function call stack
dedicated to the execution of whatever code is handed to that thread<br/>
traditionally,
these tend to be very, very large. They have, for example, about one megabyte of RAM, they
take quite a bit of time for the application to set up. And so you want to be very conservative
about how you use your threads. And that's where you get into concepts of thread pooling, because the creation and destruction of threads is very expensive.
<br/>
<br/>

in Go, it follows a little bit of a different model. using what's called
Green threads. So instead of creating these very massive heavy overhead threads, we're
going to create an abstraction of a thread that we're going to call a go routine.
<br/>
>inside of the go runtime, we've got a scheduler that's going to map these go routines onto
these operating system threads for periods of time, and the scheduler will then take
turns with every CPU thread that's available and assign the different go routines, a certain
amount of processing time on those threads. But we don't have to interact with those low
level threads directly. we're interacting with these high level go routines.<br/> - the
advantage of that is since we have this abstraction go routines can start with very, very small
stack spaces, because they can be reallocated very, very quickly. And so they're very cheap
to create and to destroy. So it's not uncommon in a go application to see 1000s or 10s of
1000s of go routines running at the same time. And the application is no problem with that
at all.

if you compare that to other languages that rely on operating system threads that
have one megabyte of overhead, there's no way you're going to run 10,000 threads in
an environment like that. So by using go routines, we get this nice lightweight abstraction over
a thread, and we no longer have to be afraid of creating and destroying them
* To trun function invokation to go routine use keyword  `go` : what is is going to do is tell go to spin of a green thread, and run that function in the greenthread
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)

}

func sayHello() {
	fmt.Println("Hello")
}
```
> main function is actually executing in a go routine itself and it exits as soon as it is done.
for now  delay main function, so the sayHello() has some time to execute to print 

```go

func main() {
	var message = "hello"
	go func() {
		fmt.Println(message)

	}()
	time.Sleep(100 * time.Millisecond)
}
```
> I'm using this anonymous function this anonymously declared function, and I'm invoking it immediately. And I'm launching it with go routine.
what's interesting about it is I'm printing out the message variable that I've defined up.the reason that it works is go has the concept of closures, which means that this anonymous function actually does have access to the variables in the outer scope. So it can take
advantage of this message variable that we declared up. and use it inside of the go routine. Even though the go routine is running with a completely different execution
stack. The go runtime understands where to get that MSG variable, and it takes care of
it for us.

 the problem with this is that we've actually created a dependency between the variable in the main function and the variable in the go routine. will create a race condition  

 to illustrate how that can be a problem.
 ```go
 func main() {
	var message = "hello"
	go func() {
		fmt.Println(message)

	}()
	message = "goodbye"
	time.Sleep(100 * time.Millisecond)
}
```
>  I'm declaring the variable message and setting it equal to Hello, and then printing it out in the
go routine. And then right after I launched the go routine, I'm
reassigning the variable to goodbye. <br />  
> *  we will get goodbye as output.

>the reason for that. And it's not always going
to be guaranteed to execute this way. But most of the time, the go scheduler is not
going to interrupt the main thread until it hits this sleep call, Which means
even though it launches another go routine, it doesn't actually give it any love yet, it's still executing the main function. 
And so it actually gets to and reassigns, the value of the message variable before the go routine has a chance
to print it out. And this is actually creating what's called a race condition  

* what are the other options? 
```go
func main() {
	var message = "hello"
	go func(message string) {
		fmt.Println(message)

	}(message)
	message = "goodbye"
	time.Sleep(100 * time.Millisecond)
}
```
> if we add a message argument here, and then down in the prints, where we're actually invoking
the function? What if we pass in the message parameter? Well, since we're passing this
in by value, so we're actually going to copy the string hello into the function, then we've
actually decoupled the message variable in the main function from the go routine. 


* it's not best practice is because we're using this sleep call<br/>
 So
we're actually binding the applications performance and the applications clock cycles to the real
world clock. And that's very unreliable.
- - -  
## wait-group
what
a wait-group does is it's designed to synchronize multiple go routines together.

in this
application, we've got two go routines that we care about, we've got the go routine that's
executing the main function. And we've got this anonymous
go routine
```go
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // create a WaitGroup and intialize it

func main() {

	var message = "hello"
	wg.Add(1) // tell wait-group we have another go-routine to sync with
	go func(message string) {
		fmt.Println(message)
		wg.Done() // anonymous routine anounce its done, so main routine can exit 
	}(message)
	message = "goodbye"
	wg.Wait() // main routine waits for the added routine
}
```

- -  - 
we can have multiple go routines that are working on the same data. And we might need to synchronize those together

in this example here, I'm creating wait group, And then I'm initializing a counter variable. inside of my main function,
I'm actually spawning 20 go routines, because inside of this for loop, each time I run through,
I add two to the wait group to let it know there are two more go routines that are running.
And then I spawn a say hello, and then I spawn an increment
```go
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}

	wg.Wait()
}

func sayHello() {
	fmt.Printf(" Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
```
>output:<br /> Hello #5<br/>
 Hello #7<br/>
 Hello #8<br/>
 Hello #9<br/>
 Hello #2<br/>
 Hello #1<br/>
 Hello #2<br/>
 Hello #2<br/>
 Hello #5<br/>
 Hello #2
>>we see that we get a mess, we in fact, don't have any kind of reliable behavior going
on here

>what's happening here is our go routines are actually racing against each other. So we have no synchronization between the go routines. they going as fast as they can to accomplish the work that we've asked them to do, regardless of what else is going on in the application.

### in order to correct this we're going to introduce this concept of a <b>mutex</b>.
a mutex is basically a lock that the application is going to honor.<br/>
a simple mutex is simply locked or unlocked. So if the mutex
is locked, and something tries to manipulate that value, it has to wait until the mutex
is unlocked. And they can obtain the mutex lock itself. <br/>So what we can do with that is
we can actually protect parts of our code so that only one entity can be manipulating
that code at a time. And typically what we're going to use that for is to protect data to
ensure that only one thing can access the data at a single time.
<br/>

With an<b> RW mutex</b>. as many time as want to can read this
data, but only one can write it at a time. And if anything is reading, then we can't
write to it at all.

So we can have an infinite number of readers, but only one writer. And
so when something comes in and makes a write request, it's going to wait till all the readers
are done. And then the writer is going to lock the mutex. So nothing can read it or
write it until the writer is done.

* two case for where to put the locks

I'm attempting to use a mutex to synchronize things together.<br> So the modification is down
here in my `sayHello()`, I'm just reading the value of the <b>counter</b> variable. And that's
what I'm trying to protect. I'm trying to protect the counter variable from concurrent
reading and writing because that's what was getting us into trouble
```go
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}

	wg.Wait()
}

func sayHello() {
	m.RLock()
	fmt.Printf(" Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
```
 > I
obtained a read Rlock on the mutex and then I print out my message and And then I released
that lock using the RUnlock method.

> in the increment, that's where I'm actually mutating
the data. So I need to write lock. And so I'm going to call the lock method on the mutex,
increment the value. And then I'm going to call unlock.


> we actually haven't gotten quite where I want to be. So I don't get the weird random behavior
that I was seeing before. But you notice that something seems to be out of sync still, because
I get Hello, one, hello, two, and then it stays at two. And if I keep running, this
actually can get different behaviors. But notice that I'm always going in the proper
order.

> output:<br/>
 Hello #0<br/>
 Hello #0<br/>
 Hello #1<br/>
 Hello #1<br/>
 Hello #1<br/>
 Hello #1<br/>
 Hello #1<br/>
 Hello #1<br/>
 Hello #1<br/>
 Hello #1<br/>

- - -
* <b>second use case</b>: we actually have to lock the mutex outside of the context of the go routine.
So we have a reliable execution model.

 moved the locks out. So the locks are now executing
before each go routine executes. And then I unlock them when the go routine is done.
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100) //
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}

	wg.Wait()
}

func sayHello() {
	fmt.Printf(" Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
```
> I basically have completely destroyed concurrency
and parallelism in this application. Because all of these mutexes are forcing the data
to be synchronized and run in a single threaded way. So any potential benefits that I would
get from the go routines are actually gone. 
---
<b>runtime.GOMAXPROCS(-1)</b> <br/>
from the runtime package, it's going to tell me the number of threads that are available.

So by default, what go is going
to do is it's going to give you the number of operating system threads equal to the number
of cores that are available on the machine. So in this virtual machine, I've exposed four
cores to the VM.

can change that value
to anything I want.
```go
runtime.GOMAXPROCS(20) // this change to 20 threads
fmt.Printf("threads %v\n ", runtime.GOMAXPROCS(-1))
// this negative -1 returns the numbers of threads thar are set
```
> if you get up too high Like, for example, 100, then
you can run into other problems, because now you've got additional memory overhead. Because
you're maintaining 100 operating system threads, your scheduler has to work harder because
it's got all these different threads to manage. And so eventually, the performance peaks and
it starts to fall back off, because your application is constantly rescheduling go routines on
different threads. And so you're losing time every time that occurs.
---
## Best practice
* Don't create goroutines in libraries
    * Let consumer control concurrency

>it's better to let the consumer control the concurrency of the library,
not the library itself. If you force your library to work concurrently, then that can
actually cause your consumers to have more problems synchronizing data together. So in
general, keep things simple, keep things single threaded, and let the consumer of your library
decide when to use a go routine and when not 
* When creating a goroutine, know how it will end
    * Avoid subtle leaks      
> if you don't have a way to stop that
go routine, that go routine is going to continue on forever. And so it's constantly going
to be a drain on the resources of your application. And eventually, as the go routine ages, it
could even cause other issues and cause your application to crash
* Check for race conditions at compile time<br/>
detect that without running the application 
` go run -race main.go`