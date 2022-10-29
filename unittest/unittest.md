## unit test
* Demo for sum of integers
```go
package main

import "fmt"

// Ints returns the sum of a list of integers.
func Ints(vs ...int) int {
	return ints(vs)
}

func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]) + vs[0]
}

func main() {
	vs := []int{5, 5, 5, 5}
	fmt.Println(ints(vs))   //20

	fmt.Println(Ints(1, 2, 3, 4))  //10
}

```


```go
package sum

// Ints returns the sum of a list of integers.
func Ints(vs ...int) int {
	return ints(vs)
}

func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]) + vs[0]
}

```
* unit test Ints() func with different scenarios
`$ go test -v `
```go
package sum

import "testing"

func TestInts(t *testing.T) {

	s := Ints(1, 2, 3, 4, 5)
	if s != 15 {
		t.Errorf("sum of 1 to 5 should be 15; go %v", s)
	}

	s = Ints()
	if s != 0 {
		t.Errorf("sum of no numbers should be 0; go %v", s)
	}

	s = Ints(1, -1)
	if s != 0 {
		t.Errorf("sum of 1 and -1 should be 0; go %v", s)
	}

}

```
* unit test using test tables
```go
package sum

import "testing"

func TestInts(t *testing.T) {

	// define a table of tests
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one ", []int{1, -1}, 0},
	}

	// loop through the test cases
	for _, tc := range tt {

		s := Ints(tc.numbers...)
		if s != tc.sum {
			t.Errorf("sum of %v should be %v ; got %v", tc.name, tc.sum, s)
		}

	}

}


```
* run sub tests
`$ go test -v -run Ints/one`
```go
package sum

import "testing"

func TestInts(t *testing.T) {

	// define a table of tests
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one ", []int{1, -1}, 0},
	}

	// loop through the test cases
	for _, tc := range tt {
		// run a sub test
		t.Run(tc.name, func(t *testing.T) {

			s := Ints(tc.numbers...)
			if s != tc.sum {
				// Fatalf stop only the execution of the sub test
				t.Fatalf("sum of %v should be %v ; got %v", tc.name, tc.sum, s)
			}
		})

	}

}

```

#### test http

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/double", doubleHandler)
	return r
}

// doubleHandler extract query string from the request, and double it the number
// if no query passed throw a missing value err, if text pass throw not a number err
func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("v")
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, v*2)
}

```

```go
package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	// this the request we expect and build the tests around it
	url := "localhost:8080/double?v=2"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	// rec provide implementation of 15.http.ResponseWriter that
	//// records response and its mutations for later inspection in tests.
	rec := httptest.NewRecorder()

	doubleHandler(rec, req)

	// get the result of the response  by the handler
	res := rec.Result()
	defer res.Body.Close()

	// check the status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; go %v", res.Status)
	}


	// check the body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v ", err)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("'expected intger; got %s", b)
	}
	if d != 4 {
		t.Fatalf("expected double to be 4; got %d", d)
	}
}

```
* how much are we testing? test coverage
add more test cases to test the failure cases(mistakes)

```go
package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	// create the test cases
	tt := []struct {
		name   string
		value  string
		double int
		err    string
	}{
		{name: "double of two ", value: "2", double: 4},
		{name: "missing value", value: "", err: "missing value"},
		{name: "not a number", value: "x", err: "not a number: x"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "localhost:8080/double?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not created request: %v", err)
			}

			rec := httptest.NewRecorder()
			doubleHandler(rec, req)

			// get the result of the response  by the handler
			res := rec.Result()
			res.Body.Close()

			// check the body
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v ", err)
			}

			// if we have no error means the status code is ok, else we check the reason
			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message%q; got %q", tc.err, msg)
				}
				return
			}

			// check the status code
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("'expected intger; got %s", b)
			}
			if d != tc.double {
				t.Fatalf("expected double to be %v; got %d", tc.double, d)
			}

		})

	}
}

```
lets go over it <br>
* using test cases table
* make the new request which to test against
* make the response writer, therefore we using httptest pkg which
that gives us a new recorder that actually satisfies
response writer interface and allows us to call result to 
to get the response the client would get and then
do alot of checks in there
* check if we get an error we do checks to make sure
that they're actually make sense
---
#### How we test this part? 
#### testing http routing
* how to test that whenever we send a request to 
slash double actually gets to double handler correctly
* test the request is routed to the right handler

```go
package main

func main() {
    err := http.ListenAndServe(":8080", handler())
    if err != nil {
    log.Fatal(err)
    }
}

func handler() http.Handler {
    r := http.NewServeMux()
    http.HandleFunc("/double", doubleHandler)
    return r
    }

```

```go

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/double?v=2", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("expected an integer; got %s", b)
	}
	if d != 4 {
		t.Fatalf("expected double to be 4; got %v", d)
	}
}

```