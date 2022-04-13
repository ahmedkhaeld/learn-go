# Concurrency

> **benchmarks** : 
> is the number of cores the cpu has. <br/>\
> **thread** : list of task which allocated by the go app. <br>\
> **Concurrency**: dealing with a lot of tasks at the same time. <br>\
> **parallelism**: doing a lot of task at the same time <br>\
> **context** **switching**: when the cpu will go for eac task and try to execute each one
>for a certain amount of time, it will to try to move on to a different task,
> which means when a task is taking too much time, probably there is another 
> task is waiting which might take less time and can be executed


> ******Where does GO fit in?******<br>
we have `cpu ` -> `cores`-> `threads` -> `tasks`
> <br> simply GO application wrapped around go runtime, which has everything that
> a binary needs in order to execute e.g. Garbage Collector, Scheduler, which go 
> will use in the end when it is going to execute that binary

---
## Fork/Join Model
in **fork/join** model has a main function, the main function has a couple
of task at certain point in time, you have go keyword which
means is going to tell the scheduler to go ahead and schedule
that go routine.

<img src="https://user-images.githubusercontent.com/40498170/162205219-8992e9a3-a374-47eb-aa8f-18901742b8e9.png" height="300"> 

from this point in time, this specific processes going in its separate way, 
now that child process has its own list of task, but at certain time or 
after the child processes finishes executing, it has to do the join,
it has to go back to the main function which called the join point.
if it does not go back into the main process or if the main process does not expect a join 
back from that fork which was done initially, basically the main function is going to say
oh! we are done, we don't have anything to wait for.

#### How we create a join point?
* wait-groups
* channels

### concurrency can be quite messy
1. Race conditions
2. deadlocks
3. live locks
4. starvation
**race condition** when multiple pieces of concurrent code try to access some kind of
shared data, then you **deadlocks** when processes wait for one another
forever.<br>
**live locks** when processes pretend they are modifying a common/shared
state then you have **starvation** when a process kind violate cpu for
longer than other process

---
## wait-groups pt1
when it comes to concurrency there is two options
* channels
* primitives
    - wait-groups
    - cond
    - mutex
    - RWMutex
    - locker
    - Map
    - pool

##### what is a wait-group? why it is called that?
the approach is very simple, you will have some kind of wait point at some point in the program
and will have to tell each concurrent process go routine that this action is done

<img src="https://user-images.githubusercontent.com/40498170/162230795-0b3ca479-983f-4d70-8d6a-426c1d62853f.png" height="300">

there is a couple of problems when it comes to concurrent code, like wait for a condition.
You have a concurrent actions/functions. you have to execute those functions first before you
execute anything else, whether it's asynchronous or synchronous.<br>
You have a starting point where all those actions have to be executed, and then we have a wait point

#### Wait-group  issues
* when don't call the method add() before call wait()
wait will return immediately
* not calling done() enough times, let's say you call done() fewer
times than you indicated in the add(), will result a deadlock
* call done() more times than call() result a panic which is deadly for the application
* pass wait-group by value, will panic, we always have to pass wg by ref
* trying to reuse wait-group, try to reuse a wg while it is still waiting will result a panic

---
## Wait-group pt2
another practical use case of wg is 
- wait limiting <br>
let's imagine we have application(server) has a capacity of 5 requests to database, we pull all 
5 requests data at a time, because we want to utilize all the resources.
<br> at the another side there an external API able to receive only 3 requests max at a time
this will result the fourth and fifth request will be denied

<img src="https://user-images.githubusercontent.com/40498170/162264601-1d0247c3-159f-41ee-a0b8-b701d96e0488.png" height="300">

server made 5 req to db, got 5 records, and want to throw all at a time<br>
so in this case it required to implement a batching to the requests, so we can comply with 
the rate limiting <br>
Apply some Batching on the application with rate limiting, this batching easily done
with wait-group. limit the # the number of requests| limit the # of go routines
group those requests in a very small chunk and throw them as batches
and once the batch is done, the next batch

---
## Wait-groups pt3
<img src="https://user-images.githubusercontent.com/40498170/162499188-ff4da807-c314-4b6c-8e36-bf8d93ef401b.png" height="300">

go scheduler and how it does
schedule things how it does the heavy
lifting when it comes to concurrency for
you so when we use the go keyword in
front of a function what exactly will
happen
is the go scheduler will get triggered
and it will try and schedule that go
routine
it will try and schedule that function
somewhere on a specific thread which is
going to get run by a specific processor
or cpu

the question is not about
the go scheduler or how it does things
the question is when we write co
code and when we use this go keyword in
front of a function
does the code execute sequentially which
means does the code execute in the exact
same order as we specify the go keyword?

the answer is very simple
first of all the go scheduler doesn't
know or doesn't even
have any kind of idea how much work
there is to be done
inside each function which means it's
going to try and schedule each function
on a specific thread

let's say the
go scheduler has done its job
it's time for these tasks to actually
execute and when these tasks execute
it's also not guaranteed which tasks is
going to finish first because again
it depends on what kind of work they're
doing so the rule of thumb
is every time you write go code it's not
guaranteed that the order in which you
write your code in which you use the go
keyword is actually going to be
sequential or
execute one after the other it has to go
for the go scheduler and even if the go
scheduler does its best job to
actually schedule that and make sure it
executes as efficiently as possible it's
still not guaranteed that
it's going to actually execute in the
exact same order the go scheduler
actually scheduled them which means
everything is going to be executed
pretty much
randomly which means you don't have to
really have any kind of expectations
as to what is going to be the order in
which the go routines execute

* another problem order dependency(preserve-order)

<img src="https://user-images.githubusercontent.com/40498170/162515546-6eb20b3b-f36d-434c-b274-c923a3b25621.png">

another problem which arises when
executing concurrent code is let's say
you want to preserve the order let's say
some of the tasks can execute
concurrently can execute in parallel but
somehow you want to preserve order you
want to execute a certain amount of
tasks first
then you want to execute a different
amount of tasks second and so on and so
forth

---
## Atomics
* **what are atomics how to use atomics?**<br>
  primarily called
atomics because of their nature because
they are pretty much like
atoms in an electron they are
indivisible they are the smallest thing
you can get
and this is a very important component
for concurrency because every operation
that happens concurrently has to be
atomic it has to be deterministic it has
to
not be divisible or interruptable but
some other process which also executes
concurrently

within the defined context
something is atomic
if it happens in its entirety without
anything else
happening simultaneously or interrupting
that specific process

<img src="https://user-images.githubusercontent.com/40498170/162537112-75fdd4c1-e03a-44c3-8d71-1947e4ba47a7.png" height="300">

let's suppose you have
go routine one and go routine two and
you have some kind of data which is
shared let's say both
is trying to access that shared data
first goroutine is trying to read the
data and the second goroutine is trying to
write to the same data now the way that
we would reason about that
is they have to have some kind of
mechanism they would communicate you
would think
we might have some kind of locking
mechanism in place which is the way to
go and this is the basic principle of
atomicity we are trying to
accomplish something and we don't want
to be interrupted by some other process

>

<img src="https://user-images.githubusercontent.com/40498170/162537696-3b6f0eed-4920-4fb1-ac8d-446a766a95d5.png" height="400">

let's imagine this time you have three
routines you have g1 g2 and g3 let's say
the first go routine is trying to
increment the i variable by one the
second is trying to decrement it by one
and the third is trying to increment it
with two and we have an initial
declaration of the i variable which is
going to be equals to zero
and this way our code will still execute
deterministically let's say g2 will
execute first and g1 will get to execute
or work with that register and then
in the end g3 will do its job

#### race condition problem
to solve race condition is to use atomics <br>
so when it comes to the atomic
package it will primarily have support
for functions which will give you the
possibility to atomically and
concurrently modify
numbers

```
Load_            int32         
Add_             int64
Store_           uint32
Swap_            uint64
CompareAndSwap_  uintptr
                 pointer

```

--- 
#### atomic with other types rather than numbers
atomic package it's not limited to just
numbers you can also manipulate a little
bit more complex types
atomically and concurrently types like
structs and all that stuff
and that specific type in the atomic
package is called atomic.value and it
will primarily have two methods it will
have a method called
load and another one called store
````
atomic.Value

Store()
Load()
````























