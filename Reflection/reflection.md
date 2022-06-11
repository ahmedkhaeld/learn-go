# Reflection


* example on reflection with a Hard JSOn

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// response the actual response we want to fill in with json
type response struct {
	Item   string `json:"item"`
	Album  string
	Title  string
	Artist string
}
//respWrapper which embeds a response, 
//to create a custom UnmarshalJSON method on Unmarshal(because json.Unmarshal can't be recursive) 
type respWrapper struct {
	response
}

var j1 = `{
"item": "album",
"album": {"title": "Dark side"}
}`

var j2 = `{
"item": "song",
"song": {"title": "Bella Donna ", "artist":"steve"}
}`

func main() {
	var resp1, resp2 respWrapper
	var err error

	if err = json.Unmarshal([]byte(j1), &resp1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v \n", resp1.response)

	if err = json.Unmarshal([]byte(j2), &resp2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v \n", resp2.response)
}

func (r *respWrapper) UnmarshalJSON(b []byte) (err error) {
	var raw map[string]interface{}

	err = json.Unmarshal(b, &r.response)
	err = json.Unmarshal(b, &raw)

	switch r.Item {
	case "album":
		inner, ok := raw["album"].(map[string]interface{})
		if ok {
			if album, ok := inner["title"].(string); ok {
				r.Album = album
			}
		}

	case "song":
		inner, ok := raw["song"].(map[string]interface{})
		if ok {
			if title, ok := inner["title"].(string); ok {
				r.Title = title
			}
			if artist, ok := inner["artist"].(string); ok {
				r.Artist = artist
			}
		}
	}
	return err
}
//main.response{Item:"album", Album:"Dark side", Title:"", Artist:""} 
//main.response{Item:"song", Album:"", Title:"Bella Donna ", Artist:"steve"}

```


* example Testing JSON
we want to know if a subset of data in other set of data
<br> `{"id":"Z"} in? {"id": "Z", "part": "fizgig", "qty": 2}`<br>

    * piece of unknown data`var unknown`, sub pieces `var known`
    * look inside the sub pieces and run `checkData()`
    * `checkData` unmarshal what we want, and we have got, then call `contains()`
    * 'contains()' it says is my expected data to be found within the data I'm passing in 
 
run at the playground terminal 
```go
package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

// matchNum look for a key that has an expected float64 value, and data which is a map of string to interface that I'm going to look in
// return true if it matches, what we're asking does certain data contains in it a certain key with a certain value of type float
func matchNum(key string, exp float64, data map[string]interface{}) bool {
	// 1. is the key there?
	// 2. type assert if key is type float64 true,
	//and the value i got by doing type assertion is equal to the exp value
	if v, ok := data[key]; ok {
		if val, ok := v.(float64); ok && val == exp {

			return true // the key is there in the map, and it has data of floating point number, also matches the passed floated number
		}
	}

	return false
}

// matchString look for a key that has an expected string value, and data which is a map of string to interface that I'm going to look in
// return true if it matches, what we're asking does certain data contains in it a certain key with a certain value of type string
func matchString(key string, exp string, data map[string]interface{}) bool {
	// 1. is the key there?
	// 2. type assert if key is type string true,
	//and the value i got by doing type assertion is equal to the exp value
	if v, ok := data[key]; ok {
		if val, ok := v.(string); ok && strings.EqualFold(val, exp) {

			return true // the key is there in the map, and it has data of string, also matches the passed string
		}
	}

	return false
}

//contains what contains says is, is my expected data to be found within the data I'm passing in
func contains(exp, data map[string]interface{}) error {
	// 1. loop over expectations. switch on type
	// if each exp case not match return error
	for k, v := range exp {
		switch x := v.(type) {
		case float64:
			if !matchNum(k, x, data) {
				return fmt.Errorf("%s unmatched (%d)", k, int(x))
			}
		case string:
			if !matchString(k, x, data) {
				return fmt.Errorf("%s unmatched (%s)", k, x)
			}
		case map[string]interface{}:
			if val, ok := data[k]; !ok {
				return fmt.Errorf("%s missing", k) //  when missing key(not in the map)
			} else if unk, ok := val.(map[string]interface{}); ok {
				// the key is there, and it is a map of string to interface
				if err := contains(x, unk); err != nil {
					return fmt.Errorf("%s unmatched in %#v : %s", k, x, err)
				}
			} else {
				// the key is there, but it is the wrong type
				return fmt.Errorf("%s wrong in %#v", k, val)
			}
		}
	}

	return nil
}

func CheckData(known string, unknown []byte) error {
	var k, u map[string]interface{}

	if err := json.Unmarshal([]byte(known), &k); err != nil {
		return err
	}

	if err := json.Unmarshal(unknown, &u); err != nil {
		return err
	}

	return contains(k, u)
}

var unknown = `{
		"id": 1,
		"name": "bob",
		"addr": {
			"street": "Lazy Lane",
			"city": "Exit",
			"zip": "99999"
		},
		"extra": 21.1
	}`

func TestContains(t *testing.T) {
	var known = []string{
		`{"id": 1}`,
		`{"extra": 21.1}`,
		`{"name": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Exit"}}`,
	}

	for _, k := range known {
		if err := CheckData(k, []byte(unknown)); err != nil {
			t.Errorf("invalid: %s (%s)\n", k, err)
		}
	}
}

func TestNotContains(t *testing.T) {
	var known = []string{
		`{"id": 2}`,
		`{"pid": 1}`,
		`{"name": "bobby"}`,
		`{"first": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Alpha"}}`,
		// dup the above with "funk" and "extra" to up coverage
	}

	for _, k := range known {
		if err := CheckData(k, []byte(unknown)); err == nil {
			t.Errorf("false positive: %s\n", k)
		} else {
			t.Log(err)
		}
	}
}

```

````
=== RUN   TestContains
--- PASS: TestContains (0.00s)
=== RUN   TestNotContains
prog.go:128: id unmatched (2)
prog.go:128: pid unmatched (1)
prog.go:128: name unmatched (bobby)
prog.go:128: first unmatched (bob)
prog.go:128: addr unmatched in map[string]interface {}{"city":"Alpha", "street":"Lazy Lane"} : city unmatched (Alpha)
--- PASS: TestNotContains (0.00s)
PASS

Program exited.
````