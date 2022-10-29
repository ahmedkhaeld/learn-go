#  JustForFunc convering
* what is context value
* how to create values to pass them to functions that require them
* how to define functions that receive those context values
* how can context can make the HTTP clients and servers mor efficient and less wasteful
  
### How to use context pkg and why?
main important usage is cancellation, and cancellation propagation <br>

**cancellation** : is when you are requesting some service you can say actually never mind, example, imagine you asked me to make you a sandwich, and for me it will take some steps to prepare and hand the sandwich, if at some point you be like i don't want the sandwich, i should stop everything i was doing, even if im in a middle of doing something, i should be able to stop it, that is called cancellation<br>

**cancellation propagation** : means if i asked someone else to help me prepare the sandwich, I should be able to
 say to the other person, hey, i don't want to your help anymore, because i don't have to make the sandwich.


## what is a context useful for?
* is not only useful for defining a few different timeouts, when
working with external APIs, but also, at the same time to propagate
information between different services.

* context pkg was added as a way of providing common
method of cancellation for work in progress.

* context is not specifically about concurrency.

* is a way to tie a bunch of related work,
imagine, I got a microservice, I get a request, and 
that request spwans some other requests. i need to talk to two or three microservices, and maybe do a 
database transaction, and i want to something to tie them, for example if i want to cancel, i want to cancel everything
or maybe i want to timeout to apply all of these sub pieces at the same time

---

<div style="page-break-after: always;"></div>

* how context with background()
```go
package main

import (
 "context"
 "fmt"
 "time"
)

// this demo shows when we pass a context that has no cancellation, it will eventually be executed 
// even after w install it for sleeping
func main() {
 // define a root(parent) context that has no cancellation
 ctx := context.Background()
 // pass the ctx to the function
 sleepAndTalk(ctx, 5*time.Second, "hello")
}
func sleepAndTalk(ctx context.Context, duration time.Duration, s string) {
 time.Sleep(duration)
 fmt.Println(s)
}

```
* context withCancel()
````go
package main

import (
	"bufio"
	"context"
	"os"
	"time"
)
// this demo shows when we pass a context that has  cancellation
// for example, we want to cancel as soon as we are able to read something from the standard input
// after 5 Sec it prints hello, but for any reason we typed anything to the terminal withing 5 seconds it cancel
func main() {
	// define a root(parent) context that has no cancellation
	ctx := context.Background()

	// new context that is a child from the root context that add cancellation
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		// or use to sleep one-second then call cancel
		//time.Sleep(time.Second)
		cancel()
	}()

	// pass the ctx to the function
	sleepAndTalk(ctx, 5*time.Second, "hello")
}
````
* context WithTimeout()
````go
func main() {
	ctx := context.Background()

	// new context that is a child from the root context that add cancellation after timeout duration
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// pass the ctx to the function
	mySleepAndTalk(ctx, 5*time.Second, "hello")
}
````

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}

// mySleepAndTalk either it has enough time to fire a message or does not have time so context call
// Done() that close, because the timeout is elapsed and log the cancellation err
func mySleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}

}
```
so far we showed demos that runs on a single process, that works great!

### what happens if we do communication over network? over http?
- server to handle requests from clients at some address<br/>
the scenario, <br/>
- server is started
- client hit the request
- client waits 5 sec as the time-server takes to respond
 **server>main.go**
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// the request has a hidden context
	ctx := r.Context()

	log.Printf("handler started")
	defer log.Printf("handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello ")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
```
**client to hit the server** 
```go
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("15.http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalln(res.Status)
	}
	// copy the body of the resp to the standard output
	io.Copy(os.Stdout, res.Body)
}

```
the cool part if server started and listen to the client, if for any reason, the 
client decided to cancel, the server will also receive a context cancellation, that's awesome
because all the HTTP server able to stop processing the task they were asked to do it, if they 
knew the client will never receive it.

sleep here is like a simulation for something expensive like going through some bunch of data
files or through a database, we do not want to do this if the client will not receive it

---
add cancellation withing the client
```go
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// where to put the context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// separate the request into two steps
	// 1. create the request
	req, err := http.NewRequest(http.MethodGet, "15.http://localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// put the ctx within the request before sending
	req = req.WithContext(ctx)

	// 2. send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalln(res.Status)
	}
	// copy the body of the resp to the standard output
	io.Copy(os.Stdout, res.Body)
}

```
> 2022/04/05 07:06:36 Get "http://localhost:8080": context deadline exceeded
