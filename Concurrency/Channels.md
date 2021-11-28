## Channels

* <b>Channel basics:</b>  how to create them, how we can use them how we can pass data through them
* <b>Restricting data flow</b>
* <b>Buffered channels</b> and how we can actually design channels to have an internal data store so that they can store several messages at once just in case the sender and the receiver aren't processing data at the same rate.
* <b>Closing Channels</b>
* <b>For... range loops with channels</b>
*<b> Select statments</b>

 how those can be used to pass data between different go routines in a way
that is safe, and prevents issues such as race conditions, and memory sharing problems
that can cause issues in your application that are very difficult to debug.

we're almost always going to
be working with them in the context of go routines. And the reason is because channels
are really designed to synchronize data transmission between multiple go routines

```go
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // create a channel for intergers

	wg.Add(2)

	go func() {
		i := <-ch //  this goroutine recieves data from the channel and stores it in i 
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // this goroutine sends values to the channel
		wg.Done()
	}()

	wg.Wait()

}
```
* the first thing that I need to do is I need to create a wait group. Because as you remember from the last
video, we use wait groups in order to synchronize go routines together.  to make sure that our main routine waits for all of our other go routines to finish.
*  then we're going to use channels in order to synchronize the data flow between them. 
---------------------------------------------------------------
 create five sets of go routines.we're going to spawn 10 go routines here, five senders and five receivers
 ```go
 var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for x := 0; x < 5; x++ {
		wg.Add(2)

		// go routine receives from the channel
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		// go routine sends data to the channel
		go func() {
			ch <- 42
			wg.Done()
		}()

	}

	wg.Wait()

}
```
> 42 42 42 42 42   

move the receiver outside of the for loop. So you're gonna have one receiver and multiple senders. make them <b>asymmetrical</b>

```go
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	// go routine receives from the channel
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	for x := 0; x < 5; x++ {
		wg.Add(2)
		// go routine sends data to the channel
		go func() {
			ch <- 42
			wg.Done()
		}()

	}

	wg.Wait()

}
```
> fatal error: all goroutines are asleep - deadlock!

if you think about how this go routine is going
to process, it's going to receive the message coming in from the channel, it's going to
print and then it's going to exit, but then in the loop, we're actually going
to spawn five messages into that channel. So we can only receive one, but we're sending
five
*  Now an important thing to keep in mind here
is the reason that this is a deadlock. And the reason for that is this line of code `ch <- 42`
is actually going to pause the execution of this go routine  until there's a space available in the channel.
* by default, we're working with unbuffered channels, which means only one message can be in the channel at one time


---    
## bi-directional communication between go routines
each routine send and receive 
```go
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

    // this go routine receives from the chanel and print data out 44, then send to the other routine channel
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 20
		wg.Done()
	}()
    // this go routine sends the data to the channel, and print out the data from the other routine 20
	go func() {
		ch <- 44
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()

}
```

### Restric each routine to receive or send 
```go
func main() {
	ch := make(chan int)
	wg.Add(2)
	// receive only
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	// send only
	go func(ch chan<- int) {
		ch <- 44
		wg.Done()
	}(ch)

	wg.Wait()

}
```
---
 ## solve the problem of asymmetrical communication
by adding a buffered channels
 ` ch := make(chan int, 50) ` <br> and using for... range loop to ouput all the data stored in the buffer
 add a second parameter to the
make function up here, and provide an integer, that's actually going to tell go to create
a channel that's got an internal data store that can store in this case, 50 integers.

 what a buffered channel is really designed to do<br> is if the sender or the receiver
operate at a different frequency than the other side
<br> is when your sender or your receiver needs a
little bit more time to process. And so you don't want to block the other side, because
you have a little bit of a delay 

```go
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	// go routine receives from the channel
	go func(ch <-chan int) {

		for i := range ch {
			fmt.Println(i)

		}
		wg.Done()
	}(ch)

	// go routine sends multiple data to the channel
	go func(ch chan<- int) {
		ch <- 30
		ch <- 40
		ch <- 50
		close(ch) // close the channel after send all the data. so for...range detect the length of the channel buffer
		wg.Done()
	}(ch)

	wg.Wait()

}
```
---------------------------------------------------------------

### SELECT 
there can be situations where you create go routines that don't have an obvious way to close.
> remember, an application is shut down as soon as the last statement of the main function finishes execution. all resources are reclaimed as the go runtime returns all of the resources that it was using back to the operating system.


Now the problem I want you to consider is when does the logger go routine closed down.
```go
package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting..."}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down!"}

	time.Sleep(100 * time.Second)
}

func logger() {
	for entry := range logCh {
		fmt.Printf(" %v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05Z"), entry.severity, entry.message)
	}
}
```
our logger go routine is being
torn down forcibly. There's no graceful shutdown for this go routine. It's just being ripped out because the main function has done.

* But there are many situations where you want to have much more control over a go routine. 
*  you should always have a strategy for how your go routine is going to shut down when you create your go
routine. Otherwise, it can be a subtle resource leak, and eventually, it can leak enough resources
that it can bring your application down.
1. use a defer.
```go 
func main() {
	go logger()
    defer func(){
        close(logCh)
    }
	logCh <- logEntry{time.Now(), logInfo, "App is starting..."}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down!"}

	time.Sleep(100 * time.Second)
}
```
2. use SELECT

```go
package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // this called a signal only channel

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting..."}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down!"}

	time.Sleep(100 * time.Second)
	doneCh <- struct{}{}
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf(" %v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05Z"), entry.severity, entry.message)

		case <-doneCh:
			break
		}

	}
}
```
>  `var doneCh = make(chan struct{}` <br> 
struct with no fields
in the go language is unique in that it requires zero memory allocations. So a lot of times you will see a channel set up like this, So this is what's called a signal
only channel. There's zero memory allocations required in sending the message. But we do
have the ability to just let the receiving side know that a message was sent. 

inside of our logger function, we've got an infinite loop now,
and we're using this select block. 

 what this SELECT statement does is the entire statement
is going to block until a message is received on one of the channels that it's listening
for. So in this case, we've got a case listening for messages from the log channel, and the
case listening for messages from the done channel

So if we get a message from the log
channel, then we're going to print out our log entry. If we get a message from the done
channel, then we're going to go ahead and break out of this for loop. So what this allows
us to do is at the end of our application

--- 

