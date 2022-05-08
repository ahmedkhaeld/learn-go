# Select
another control structure, but it's going to allow us to work with channels and go routines
<br>it allows us to multiplex channels



```go
package main

import (
	"log"
	"time"
)

func main() {
	// slice of channels
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	// range over the channels and write to them
	for i := range chans {
		//start two go routines to write to the channels numbers one two
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	// read the channels, select the ready channel and read from it
	// select allows to listen to both channels at the same time, and whichever one
	// is ready first, read from it, then go back to listening to both
	//to read when a channel produce a data

	for i := 0; i < 12; i++ {
		select {
		case m0 := <-chans[0]:
			log.Println("received", m0)
		case m1 := <-chans[1]:
			log.Println("received", m1)
		}
	}
}
```
```
//2022/05/08 13:31:10 received 1
//2022/05/08 13:31:11 received 1
//2022/05/08 13:31:11 received 2
//2022/05/08 13:31:12 received 1
//2022/05/08 13:31:13 received 1
//2022/05/08 13:31:13 received 2
//2022/05/08 13:31:14 received 1
//2022/05/08 13:31:15 received 1
//2022/05/08 13:31:15 received 2
//2022/05/08 13:31:16 received 1
//2022/05/08 13:31:17 received 2
//2022/05/08 13:31:17 received 1

```
>notice that there's a pattern
here i'm getting the ones twice
as often as the twos
>>because one go routine is
sending a message every second and the
other one is sending a message
every other second right um

> select save time, read which one is ready

* demo without using select
```go
package main

import (
	"log"
	"time"
)

func main() {
	// slice of channels
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	// range over the channels and write to them
	for i := range chans {
		//start two go routines to write to the channels numbers one two
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {

		m0 := <-chans[0]
		log.Println("received", m0)
		m1 := <-chans[1]
		log.Println("received", m1)

	}
}

```
``` 

//2022/05/08 13:37:15 received 1
//2022/05/08 13:37:16 received 2
//2022/05/08 13:37:16 received 1
//2022/05/08 13:37:18 received 2
//2022/05/08 13:37:18 received 1
//2022/05/08 13:37:20 received 2
//2022/05/08 13:37:20 received 1
//2022/05/08 13:37:22 received 2
//2022/05/08 13:37:22 received 1
//2022/05/08 13:37:24 received 2
//2022/05/08 13:37:24 received 1
//2022/05/08 13:37:26 received 2
//2022/05/08 13:37:26 received 1
//2022/05/08 13:37:28 received 2
//2022/05/08 13:37:28 received 1
//2022/05/08 13:37:30 received 2
//2022/05/08 13:37:30 received 1
//2022/05/08 13:37:32 received 2
//2022/05/08 13:37:32 received 1
//2022/05/08 13:37:34 received 2
//2022/05/08 13:37:34 received 1
//2022/05/08 13:37:36 received 2
//2022/05/08 13:37:36 received 1
//2022/05/08 13:37:38 received 2
```
> i'm going to get 1
and then i'm going to wait 2 seconds get
a 2 then i'm going to get a 1 then i'm
going to get a 2 then i'm going to get a
1
because in this case this is the not
select case. <br>ping-ponging back and
forth between these channels
one of these channels would like to send
something every second and it can't
because the other channel is only
sending stuff every other second 


---
* a program stops when reading a responses take more than 3 seconds


```go
package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	// measure the start and end to calc the latency
	start := time.Now()

	if resp, err := http.Get(url); err != nil {

		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}

func main() {
	// stopper is a signal in channel, send a msg when time is up(waiting for 3 sec)
	stopper := time.After(3 * time.Second)

	results := make(chan result)

	// list has some actual urls
	list := []string{
		"https://amazon.com",
		"https://wsj.com",
		"https://google.com",
		"https://nytimes.com",
		"https://youtube.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		select {
		case r := <-results:
			if r.err != nil {
				log.Printf("%-20s %s \n", r.url, r.err)
			} else {
				log.Printf("%-20s %s\n", r.url, r.latency)
			}

		case t := <-stopper:
			log.Fatalf("timeout %s", t)
		}

	}

}

```
``` 
2022/05/08 15:11:29 https://wsj.com      672ms
2022/05/08 15:11:29 https://youtube.com  979ms
2022/05/08 15:11:29 https://google.com   981ms
2022/05/08 15:11:30 https://amazon.com   1.734s
2022/05/08 15:11:31 timeout 2022-05-08 15:11:31.379448328 +0200 EET m=+3.001472077

```
>select run two channels simultaneously:<br>
> * normal chan that read the result
> * stopper chan to exit the program


---
* program that ticks every 2 seconds, but there is a timeout after 5 seconds

```go
package main

import (
	"log"
	"time"
)

// create two things:
//1. stopper, which is worth of five ticks(time out), gives a signal when time is up
//2. ticker, which is a periodic timer
func main() {
	log.Println("start")

	const tickRate = 2 * time.Second

	stopper := time.After(5 * tickRate)  //stopper is a writer channel
	ticker := time.NewTicker(tickRate).C // ticker is a writer channel

loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}

	log.Println("finished")

}

```

``` 
2022/05/08 15:28:22 start
2022/05/08 15:28:24 tick
2022/05/08 15:28:26 tick
2022/05/08 15:28:28 tick
2022/05/08 15:28:30 tick
2022/05/08 15:28:32 finished

```

>this is very common
right it's very common to have a select
statement either with a timeout or a
ticker
that causes me to do something every so
often for example i have a server
whose purpose is every 10 minutes to go
and update a redis cache to pre-warm the
cache with some data
and so basically once the program starts
up it just sits in a for loop with a
ticker
the ticker in this case is like every 10
minutes so it wakes up once every 10
minutes
and goes and says hey is there something
i need to put in the cache to keep it
warm
