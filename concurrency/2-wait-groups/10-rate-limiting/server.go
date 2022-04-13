package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main() {
	// create a tcp server
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not create listener: %v", err)
	}

	// in case we have a valid listener,
	//start listen to connections
	// connection counter to our accepted connection
	var connections int32
	for {
		conn, err := li.Accept()
		if err != nil {
			// if there is an error continue listen to other connections
			continue
		}
		connections++

		// run the connection inside its own go routine
		// so, we don't block other connections
		go func() {
			defer func() {
				_ = conn.Close()
				// don't accept multiple connections at the same time to avoid 1-race conditions
				atomic.AddInt32(&connections, -1)
			}()

			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			// sleeping simulate heave work
			time.Sleep(time.Second)
			_, err := conn.Write([]byte("success"))
			if err != nil {
				log.Fatalf("could not write to connection: %v", err)
			}
		}()
	}
}
