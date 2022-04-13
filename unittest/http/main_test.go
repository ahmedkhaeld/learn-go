package main

import (
	"bytes"
	"fmt"
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
