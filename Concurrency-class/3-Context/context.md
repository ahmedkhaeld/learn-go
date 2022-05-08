# Context
context a way of providing a common method of cancellation for work in progress

* the context package is a way to tie together a bunch of related work 
> imagine for a second i've
got a microservice
i get a request and that request spawns
some other requests i need to talk to
two or three other microservices and
maybe do a database transaction
and i want something to tie them
together for example
if i decide to cancel i want everything
to cancel or maybe i want a timeout and
i want the timeout to apply to all of
these sub pieces
at the same time 


>context was
originally invented primarily as a means
of providing that context cancellation
in other words tying together a bunch of
work so it all gets cancelled together

>there's another utility about passing values 

* the context pkg offers a common method to cancel requests
    * **explicit cancellation**: <br> i get some sort of cancel function that
i can just call
and say hey this context is cancelled
drop all the work related to it
    * **implicit cancellation based on a timeout or deadline**: <br> maybe
the database has a timeout, or got
a request and i'm going to handle it
within three seconds or i'm going to
drop it
regardless of how many other
microservices i talk to
  
* the context offer two controls
    * a channel that closes when the cancellation occurs: <br>  is we get a channel
there's a function called` done()` if we
call it, it gives us a channel
that we can listen to and if that
channel becomes ready to read
then we've been signaled, we should stop working
    * an error that's readable once the channel closes: <br> if we get that signal then there's
another function error
and if we call that we'll get a reason
why we were told to stop working maybe
the error is because we timed out
or maybe the error is because there was
an explicit cancellation

---
##### the context is not one thing
the context is actually a tree, usually it has some root context that has
no information
and as we add things to it if we add a
cancellation property
or we add a timeout or we cast the
context to carry a value along with it
we keep making more nodes in this tree
>the individual nodes once they're
created
are immutable okay you never modify an
existing context<br>
> what you do is you actually create a new
context that points to the
parents above it 

>if i have a context that somebody passes
into me and i want to add a timeout
again
i'm going to create a new context that
points at the old one
and the timeout applies downward it
applies to the context i just created
and it applies to any context that is
derived from the one that i have now

>i've got an incoming
http request
and i add a timeout to it, and then i
start calling
other outgoing requests to other
microservices,
okay my the timeout applies to me and
those other calls below me,
but it doesn't apply to the http request
before i got it and started processing
it,
because the context above me doesn't
have my timeout.
 the parent context may have its own
thing but it doesn't have mine

--- 
* **Context as a tree structure**

1. `ctx := context.Background()` root context which gives you basically a context that you can point
to that doesn't have anything in it
2. `ctx = context.withValue(ctx, "v", 7)` to add a value that we can carry along 
like a trace id, so i call
the function context with value
i pass in a context and it gives me a
context back
which typically i just assign over top
of the old one
but in reality it's a new object with a
pointer
going upwards to the root
3. `ctx, cancel :=context.WithTimeout(ctx, t)` now i want to put a
timeout
so i call context with timeout i pass in
the parent context
and the timeout value and again it gives
me a new context
```
req, _ := http.NewRequest(method, url, nil) 
req = req.WithContext(ctx)
resp, err := http.DefaultClient.Do(req)
```
show you how we use a context then
i'm going to make an http request and
one of my options
is to take my request and actually add a
context
into it 

```go
package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()
	// create a req that contains the ctx and url,
	//this will inject a timeout into the http Get request
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {

		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}

func main() {

	results := make(chan result)

	list := []string{
		"https://amazon.com",
		"https://wsj.com",
		"https://google.com",
		"https://nytimes.com",
		"https://youtube.com",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for _, url := range list {
		go get(ctx, url, results)
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s \n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}

}

//2022/05/08 17:51:48 https://wsj.com      1.673s
//2022/05/08 17:51:48 https://youtube.com  1.829s
//2022/05/08 17:51:48 https://google.com   1.848s
//2022/05/08 17:51:48 https://amazon.com   2.12s
//2022/05/08 17:51:49 https://nytimes.com  Get "https://www.nytimes.com/": context deadline exceeded 
```


---

```go
package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	var r result
	start := time.Now()
	ticker := time.NewTicker(1 * time.Second).C

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {

		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

	for {
		select {
		case ch <- r:
			return
		case <-ticker:
			log.Println("tick", r)
		}
	}

}

func first(ctx context.Context, urls []string) (*result, error) {
	results := make(chan result)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, url := range urls {
		go get(ctx, url, results)
	}

	select {
	case r := <-results:
		return &r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {

	list := []string{
		"https://amazon.com",
		"https://wsj.com",
		"https://google.com",
		"https://nytimes.com",
		"https://youtube.com",
	}

	r, _ := first(context.Background(), list)

	if r.err != nil {
		log.Printf("%-20s %s \n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Second)
	log.Println("quit anyway...", runtime.NumGoroutine(), "still running")

}

```