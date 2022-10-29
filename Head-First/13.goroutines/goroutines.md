# sharing work: Goroutine and Channels
Working on one thing at a time isn’t always the fastest way to finish a
task. Some big problems can be broken into smaller tasks. Goroutines let
your program work on several different tasks at once. Your goroutines can
coordinate their work using channels, which let them send data to each
other and synchronize so that one goroutine doesn’t get ahead of another.

### Retrieving web pages
finishing work faster by doing several tasks simultaneously.
But first, we need a big task that we can break into little parts.

The smaller a web page is, the faster it loads in visitors’ browsers. We need
a tool that can measure the sizes of pages, in bytes.

This shouldn’t be too difficult, thanks to Go’s standard library. The program
below uses the net/http package to connect to a site and retrieve a web
page with just a few function calls.

We pass the URL of the site we want to the http.Get function. It will
return an http.Response object, plus any error it encountered.

The http.Response object is a struct with a Body field that represents the
content of the page. Body satisfies the io package’s ReadCloser interface,
meaning it has a Read method (which lets us read the page data), and a
Close method that releases the network connection when we’re done.

We defer a call to Close, so the connection gets released after we’re done
reading from it. Then we pass the response body to the ioutil package’s
ReadAll function, which will read its entire contents and return it as a slice
of byte values.

>Go type byte used for holding raw data, such as you might read
> from a file or network connection

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	responseSize("https://golang.org/")
	responseSize("https://golang.org/dev")
	responseSize("https://golang.org/doc")
}

func responseSize(url string) {
	//print the URL we’re retrieving just for debugging purposes
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//get the size of the page
	fmt.Println(len(body))
}
```
##Multitasking
finding a way to speed programs up by performing multiple tasks at the same time

Our program makes several calls to responseSize, one at a time.
each call establishes a network connection to the website, waits for
the site to respond, prints the size

only when one call returns can the next begin.
if one big long func where the all code was repeated three times
, it would take the same amount of time to run 
###Concurrency using goroutines
```go
func main() {
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/dev")
	go responseSize("https://golang.org/doc")
}

```

### we don't directly control when goroutines run
Under normal circumstances, Go makes no guarantees about when it will
switch between goroutines, or for how long. This allows goroutines to run
in whatever way is most efficient. But if the order your goroutines run in is
important to you, you’ll need to synchronize them using channels

### Go statements can't be used with return values
We’ll get compile errors. The compiler stops you from attempting to get a
return value from a function called with a **go** statement.

You can’t rely on the return values being
ready in time, and so the Go compiler blocks any attempt to use them

Go won’t let you use the return value from a function called with a go
statement, because there’s no guarantee the return value will be ready
before we attempt to use it:

### Channels
there is a way to communicate between goroutines: Channels.
<br>Not only do channel allow you to send values from one
goroutine to another, they ensure the sending goroutines has
sent the value before the receiving goroutine attempts to use it.

Each channel only carries values of a particular type, so you might have one
channel for int values, and another channel for values with a struct type.
<br>`var myChan chan float64` to declare a variable to holds a channel
<br>`myChan = make(chan float64)` to actually create a channel<br>

it is easier to use short declaration<br>
`myChan :=make(chan float64)` <br>

The only practical way to use a channel is to communicate from one
goroutine to another goroutine.

``` 
- create a channel
- write a function that receives a chan as a param. we'll run
this function in a separate goroutines, 
and use it to send values over the channel 
- Receive the sent values in our orininal goroutine
```

###Sending and Receiving values with channels
`myChan <- 3.14 ` store|send|write  3.14 to channel
`<-myChan` receive | read  a value from

```go
package main

import "fmt"

func main() {
	myChan := make(chan string)
	//pass a var myChan, will hold a value string
	go greeting(myChan)
	receivedValue:=<-myChan
	fmt.Println(receivedValue)
	// read the value stored within the channel
}

//greeting takes a channel, store a string value in the channel
func greeting(myChan chan string) {
	myChan <- "hi"
}
```
```go
package main

import "fmt"

func main() {
	myChan := make(chan string)
	//pass a var myChan, will hold a string
	go greeting(myChan, "take me to the channel")
	rsvMsg := <-myChan
	fmt.Println(rsvMsg)
	// read the msg stored within the channel
}

//greeting takes a channel, store a string value in the channel
func greeting(myChan chan string, msg string) {
	myChan <- msg
}

```
### Synchronizing goroutines with channels
channels also ensure the sending goroutine has sent the
value before the receiving channel attempts to use it.

Channels do this by
blocking—by pausing all further operations in the current goroutine.

A send operation blocks the sending goroutine until another goroutine
executes a receive operation on the same channel. And vice versa: a receive
operation blocks the receiving goroutine until another goroutine executes a
send operation on the same channel.

This behavior allows goroutines to
synchronize their actions—that is, to coordinate their timing.
<img src="https://user-images.githubusercontent.com/40498170/198146641-aeccf82a-6fa6-49c4-b028-c12f17203600.png">
```go
package main

import "fmt"

func main() {

	//create two channels
	chan1 := make(chan string)
	chan2 := make(chan string)

	//main routine handle two go routines
	//the main routine becomes the orchestrator of abc and def
	// allowing them to send only when it's ready to read what they are sending
	go abc(chan1)
	go def(chan2)
	fmt.Print(<-chan1)
	fmt.Print(<-chan2)
	fmt.Print(<-chan1)
	fmt.Print(<-chan2)
	fmt.Print(<-chan1)
	fmt.Print(<-chan2)
	//adbecf

}

func abc(c chan string) {
	c <- "a"
	c <- "b"
	c <- "c"
}

func def(c chan string) {
	c <- "d"
	c <- "e"
	c <- "f"
}

```

##Fixing our web page size program with channels
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	size:=make(chan int)
	urls :=[]string{
		"https://golang.org/",
		"https://golang.org/dev",
		"https://golang.org/doc",
	}
	for _, url :=range urls{
		go responseSize(url, size)
	}
	for i :=0; i<len(urls); i++{
		fmt.Println(<-size)
	}
	

}

//responseSize takes an url to process, 
//a chan to hold the size of the page size we can send it through that chan
func responseSize(url string, c chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//store the length in the channel
	c<-len(body)
}
```
**Which responseSize belongs to which URL?** 

There’s still one issue we need to fix with the responseSize function. We
have no idea which order the websites will respond in. And because we’re
not keeping the page URL together with the response size, we have no idea
which size belongs to which page!

* Channels can carry composite type like slices, maps, and structs.

we can create a struct type that will store page URL with its size.
so we can send both over the channel together.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type Page struct {
	URL string 
	Size int 
}

//responseSize takes a url to process,
//a chan to hold the size of the page size we can send it through that chan
func responseSize(url string, c chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//store the length in the channel
	c <- Page{URL: url, Size: len(body)}
}

func main() {
	page := make(chan Page)
	urls := []string{
		"https://golang.org/",
		"https://golang.org/dev",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		go responseSize(url, page)
	}
	for i := 0; i < len(urls); i++ {
		p:=<-page
		fmt.Printf("%s: with %d bytes \n",p.URL,p.Size )
	}

}

```
Now the output will pair the page sizes with their URLs. It’ll finally be
clear again which size belongs to which page.