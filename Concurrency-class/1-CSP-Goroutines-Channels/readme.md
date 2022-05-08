# CSP communicating sequential processes
it's an important feature of
go because it provides us a powerful and
simple model
for solving a lot of basic concurrency
problems and
the features in go that do that are
channels and go routines

## Share memory by communicating
### channels
a channel is one-way communication pipe. <br>
pipe: it's a way of sending data from one program to another, and it redirects the output from one
program to the input of another.<br>
##### channels properties
* things go in one end, come out the other in the same order they went in 
until the channel is closed
* multiple readers and writers can share it safely

```
sequential processes talking through channels
the channel is the communicating part ofthe communicating sequential processes
and what it does is acts as either a buffer or a synchronization point
allowing these sequential processes individually to be concurrent as a group
```

* csp model allows us to write asynchronous code in a synchronous style
``` 
write something and it sends a message to something else
and how the messages get scheduled and how these somethings get scheduled across the cpus
```


---
* ****building out a program i want to use some channels and go routine**** 
   * **parallel GET through http**

* >start out with some
piece of data i'm going to pass around
and we're going to call it a result
it's going to have three things in it
okay we're going to have a url that we
looked at
we're going to have an error
and for now we'll just treat the url as
a string and we want one more thing
we're going to think about how long did
it take to get the web page


* > get something from a web page, have a func get() take a url and channel
  i'm going to run each get in a go
routine by itself
it's going to go out to the web get some
data and then it's going to put its
result into a channel
and send that back again it's going to
communicate its result
by putting it in a channel and it's
going to be a safe way to do that 
  >> `chan<- result` this says it is a channel that takes a results, means write only to channel

* >function well the first thing i need is a channel to actually send stuff in
  > some actual urls,
  
* >an important point whether we
have an error or not i'm going to put
something in the channel
so i know that for every url i started
i'm going to get a result of some kind
either it's an error result
or it's a normal result okay, routines based on how many urls i'm looking up



```go
package main

import (
	"log"
	"net/http"
	"time"
)

//result a piece of data to pass around, has url that we look at, error and duration of
// how long it takes to get the web page
type result struct {
	url     string
	err     error
	latency time.Duration
}

// get something from a web page
func get(url string, ch chan<- result) {
	// measure the start and end to calc the latency
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		//if there is an error, write to the result to the channel
		// the url, the error and no latency
		ch <- result{url, err, 0}
	} else {
		// if no error, write to the channel url, no error, the timer it took
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}

func main() {

	// create a channel of type result
	results := make(chan result)

	// list has some actual urls
	list := []string{
		"https://amazon.com",
		"https://wsj.com",
		"https://google.com",
		"https://nytimes.com",
		"https://youtube.com",
	}

	// loop through each url and make a get request to it
	// this will write a result to the channel for each request
	for _, url := range list {
		go get(url, results)
	}

	// read the results from the results channel, and put it to r variable
	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s \n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}

}

```
``` 
2022/05/08 03:32:11 https://wsj.com      302ms
2022/05/08 03:32:11 https://youtube.com  599ms
2022/05/08 03:32:11 https://google.com   615ms
2022/05/08 03:32:12 https://amazon.com   1.073s
2022/05/08 03:32:15 https://nytimes.com  3.762s
```

* **notice: why loop  this way?**
```
for range list {
r := <-results

		if r.err != nil {
			log.Printf("%-20s %s \n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
```
> the thing about
channels they block
if there's no data to read and i go to
read it i have to wait for data


```
if i wrote a program that read from the command line so i start a command
line program  it's going to read from standard input and i'm expected to type something so
what does the program do if i don't type anything 
well the answer is it waits
it waits for me either to type a line of text and hit return
or it waits for me to hit control d and say there's no more data
```
>so if i were to range over the
channel itself which i can do
right it would read until the channel
closes
but i'm not closing the channel 
i'm not closing the channel so i just
need to make sure i don't read more
times
than there could be data in the channel
to read or i'll get stuck
right here waiting for data that's not
going to arrive
and that's also why i made sure again
that every time i start a go routine
it's going to provide me a result
an error or a valid data but it's going
to give me a result so i'm going to get
four go routines and four results and
then stop <br>
> question would be well who closes the channel?<br>
which one of them and it's like this old
joke about well the last person out
turned the lights off
but there's no really way of knowing
between the go routines
which one is last so which one should
close the channel

---
##### show the power of go routines and channels to solve a data race a condition where we would otherwise have an unsafe operation in a concurrent program
* a main program that starts the server, a simple web server
* handler print out an incrementing number

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID int

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> You got %d <h1>", nextID)

	nextID++ //UNSAFE
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

this is unsafe `nextID++` the increment operation is a _read_ _modify_ _write_ <br>
web server built into the go
standard library is concurrent
if i were to hit this web page fast
enough i would start seeing some
problems
numbers would get skipped  it would cause a problem<br>

* how can I solve that?<br>
`nextID` to be a channel of integers<br>
  instead of actually trying to increment it I just read a value out of it

`func counter()` starts sending numbers. keep generating 
numbers and putting them into the channel

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	// <-nextID is a reader from the nextID channel
	fmt.Fprintf(w, "<h1> You got %d <h1>", <-nextID)

}

func counter() {
	for i := 0; ; i++ {
		// write i value to nextID channel
		nextID <- i
	}
}

func main() {
	go counter()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

>you may think well if if nobody's
using my web server that thing's just
going to spin and start just generating
all these numbers from 1 to a billion to
10 billion
to 100 billion billion and the answer is
well no
okay because in the normal use of
channels
i can't write to it unless somebody's
ready to read from it now we already
said
if the channel is not ready to be read
from
you block okay and that's again that's
no different than trying to read input
from standard input if i haven't typed
anything on standard input
the program has to wait for me to type
something
right so the writer here inside counter
where i write to the channel can't
actually do anything until there's
somebody ready to read from it
and the same thing up here the reader
can't read until there's somebody ready
to write

this is reasonably efficient. not super efficient.

* create a name type of chan int, and turn the handler into a method on that type
* the object-oriented version
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type nextCh chan int

//handler method on nextCha type, the handler can read from the nextCh through receiver
func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	// <-nextCh is a reader from the nextCh channel
	fmt.Fprintf(w, "<h1> You got %d <h1>", <-ch)

}

//counter takes a channel as a parameter, write to it the increments
func counter(ch chan<- int) {
	for i := 0; ; i++ {
		// write i value to ch channel
		ch <- i
	}
}

func main() {
	// create a channel of type nextCh
	var nextID nextCh = make(chan int)
	// when call the counter on the nextID channel, it increments the count
	go counter(nextID)

	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

---
* **Prime Sieve** with channels
    * calculate prime numbers
    * start a thing that generates numbers
    * and pipe it into a filter that drops all the numbers divisible by two
    * and i pipe that into some other filter that drops all the numbers divisible by three

#### flow of the demo
* the function of the generator is to start making numbers 
and as we get a new prime number from it.  we're going to start with 2,
we're going to ignore 1 (one is never considered prime)
* and as when we get that number we say okay so that's a number we need to filter out to and all the multiples of 2
* so we'll create another go routine called the 2 filter,
and we'll attach it on the channel in between generator and main
* which really means we'll take the channel that went from generator to main
and make that go into the 2 filter and
create a new channel coming out of the
2 filter going back to main
* then we get three well three is
also going to be prime so, we'll create
another filter
* when 4 gets generated it never gets there it gets removed by the 2 filter
* when 5 comes along we'll create a 5
filter, and again we'll add it to this little chain of channels and filters
* six will get gobbled up
* seven will make a filter
* eight will get gobbled up by number two, nine by number three and so on so
>every time it sees a new prime number
it's going to create another filtering
go routine, hook it into the channel that
it's getting numbers from,
and create a new channel back to itself


```go
package main

import "fmt"

//generate has a limit, runs until this limit hits, write to channel the numbers
// then read from this channel the numbers to be filtered
func generate(limit int, ch chan<- int) {
	// start of by 2, ignore 1, every time we make a number, put it into the channel
	for i := 2; i < limit; i++ {
		ch <- i
	}
	close(ch) // when get to the end, close the channel
}

//filter each filter needs to know a src chan(a channel to read from, where the numbers
//coming from), dst chan(where write numbers out), prime number to filter on
func filter(src <-chan int, dst chan<- int, prime int) {
	// loop over the src channel value until the channel closes
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst) // when the loop is finished close the destination
}

func sieve(limit int) {
	ch := make(chan int)

	go generate(limit, ch)
	for {
		// read from the channel the generated numbers, loop to get each number,
		// each number to be tested in the filter prime or not.
		prime, ok := <-ch
		if !ok {
			break
		}
		// make a new channel to represent the new filter
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		// pass the new filter to the main
		ch = ch1
		fmt.Print(prime, " ")
	}
}

func main() {
	sieve(100)
	//2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 
}

```
