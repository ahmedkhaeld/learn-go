# code quality assurance: Automated Testing
>I test all the equipments before every shift.
> That way, if there's a problem, we can fix it before we
> send out defective products!

* Are you sure your software is working right now? Really sure?
<br>Before you sent that new version to your users, you presumably
tried out the new features to ensure they all worked. But did
you try the old features to ensure you didn't beak any of them?

###Automated testing
ensure your program's components work correctly. even after you
change your code. Go's testing pkg and go test tool make it easy
to write automated tests

An automated test is a separate
program that executes components of your main program, and verifies they
behave as expected.

>I run my programs every time I add a new feature, to test it out.
> Isn't that enough?
> 
> Not unless you’re going to test all the old features as well, to make sure
your changes haven’t broken anything. Automated tests save time over
manual testing, and they’re usually more thorough, too

###A function we should have had automated tests for
prose pkg has JoinWithComma func

``` 
1. accept slice of strings []phrases
2. comma should place after each phrase
    a. but not the last phrase
    b. also we neet to place " and" before the last phrase

- slice the phrases and exclude the last phrase
- place the comma after them
- concatenate them with " and"
- include back the last phrase
```
```go
package prose

import "strings"

//JoinWithCommas accepts a slice of strings to join, return the resulted phrase
func JoinWithCommas(phrases []string) string {
	res := strings.Join(phrases[:len(phrases)-1], ", ")
	res += ", and "
	res += phrases[len(phrases)-1]
	return res
}
```
```go
package main

import (
	"autotest/prose"
	"fmt"
)

func main() {
	phrases := []string{"my parents!", "a rodeo"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))

	phrases = []string{"my parents!", "a rodeo", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))

}

```
``` 
output:
A photo of my parents!, and a rodeo
A photo of my parents!, a rodeo, and a prize bull
```
> comma placed with more than two phrases<br>
the comma doesn't belong here<br>
`A photo of my parents, and a rodeo clown` 

* We've introduced a bug
    * the code is working correctly with three phrases
    * but not with two phrases

An automated test runs your code with a particular set of inputs and looks
for a particular result. As long as your code’s output matches the expected
value, the test will “pass.”

### Writing tests
Go includes a testing package that you can use to write automated tests
for your code, and a go test command that you can use to run those tests.

Let’s start by writing a simple test.

`join_test.go`
```go
package prose

import "testing"

func TestThreePhrases(t *testing.T){
	t.Error("no test here either")
}

```
* this test code will be part of the same pkg as the code we're testing
* import the std lib testing
* func name should start with Test, then what we are testing
* Test func will be passed a pointer to a testing.T value
* t.Error to fail the test

```go
package prose

import "testing"

func TestTwoPhrases(t *testing.T) {
	list := []string{"apple", "orange"}
	expected := "apple and orange"
	if JoinWithCommas(list) != expected {
		t.Error("did not match expected value") //failing
	}
}

func TestThreePhrases(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	exp := "apple, orange, and pear"
	if JoinWithCommas(list) != exp {
		t.Error("did not match expected value")
	}
}

```
``` 
--- FAIL: TestTwoPhrases (0.00s)
    join_test.go:9: did not match expected value
FAIL
exit status 1
FAIL    autotest/prose  0.002s
```
This is a good thing; it matches what we expected to see based on the
output of our join program. It means that we’ll be able to rely on our tests
as an indicator of whether JoinWithCommas is working as it should be!

### More detailed test failure messages with the “Errorf” method
Our test failure message isn’t very helpful in diagnosing the problem right
now. We know there was some value that was expected, and we know the
return value from JoinWithCommas was different than that, but we don’t
know what those values were.
````go
package prose

import "testing"

func TestTwoPhrases(t *testing.T) {
	list := []string{"apple", "orange"}
	exp := "apple and orange"
	got := JoinWithCommas(list)

	if got != exp {
		t.Errorf("expected: %s ,but got: %s \n", exp, got)
	}
}

func TestThreePhrases(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	exp := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != exp {
		t.Errorf("expected: %s, but got: %s \n", exp, got)
	}
}

````

### Test “helper” functions
use helper func to format errors and decrease code duplication
```go
package prose

import (
	"fmt"
	"testing"
)

func TestTwoPhrases(t *testing.T) {
	list := []string{"apple", "orange"}
	got := JoinWithCommas(list)

	exp := "apple and orange"
	if got != exp {
		t.Errorf(errorString(list, got, exp))
	}
}

func TestThreePhrases(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	got := JoinWithCommas(list)

	exp := "apple, orange, and pear"
	if got != exp {
		t.Errorf(errorString(list, got, exp))
	}
}

func errorString(list []string, got, exp string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) gives: \"%s\", Expected: \"%s \"", list, got, exp)
}

```

 ###Getting the tests to pass
Let’s modify JoinWithCommas to fix this. If there are just two elements in
the slice of strings, we’ll simply join them together with " and ", then
return the resulting string. Otherwise, we’ll follow the same logic we
always have.
```go
package prose

import "strings"

//JoinWithCommas accepts a slice of strings to join, return the resulted phrase
func JoinWithCommas(phrases []string) string {
	if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		res := strings.Join(phrases[:len(phrases)-1], ", ")
		res += ", and "
		res += phrases[len(phrases)-1]
		return res
	}
}

```
---
## Test-driven development
1. Write the test: <br>
   You write a test for the feature you want, even
   though it doesn’t exist yet. Then you run the test to ensure that it
   fails.
2. Make it pass: <br>
   You implement the feature in your main code. Don’t
   worry about whether the code you’re writing is sloppy or
   inefficient; your only goal is to get it working. Then you run the
   test to ensure that it passes.
3. Refactor your code:<br>
   Now, you’re free to refactor the code, to
   change and improve it, however you please. You’ve watched the
   test fail, so you know it will fail again if your app code breaks.
   You’ve watched the test pass, so you know it will continue passing
   as long as your code is working correctly.

### Another bug to fix
It’s possible that JoinWithCommas could be called with a slice containing
only a single phrase. But it doesn’t behave very well in that case, treating
that one item as if it appeared at the end of a list:
`phrases=[]string{"my parents"}`
>treating that one item as if it appeared at the end of a list:<br>
> `A photo of , and my parents`

What should JoinWithCommas return in this case? If we have a list of one
item, we don’t really need commas, the word and, or anything at all. We
could simply return a string with that one item.

```go
package prose

import "strings"

//JoinWithCommas accepts a slice of strings to join, return the resulted phrase
func JoinWithCommas(phrases []string) string {
	if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		res := strings.Join(phrases[:len(phrases)-1], ", ")
		res += ", and "
		res += phrases[len(phrases)-1]
		return res
	}
}

```

### Table-driven test
* test all the cases we want using Table tests 
* reduce the code duplication withing the tests

````go
package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	exp  string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{list: []string{"apple"}, exp: "apple"},
		{list: []string{"apple", "orange"}, exp: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, exp: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.exp {
			t.Errorf(errorString(test.list, got, test.exp))
		}
	}
}

func errorString(list []string, got, exp string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) gives: \"%s\", Expected: \"%s \"", list, got, exp)
}

````


### fixing a panicking
The best thing about table-driven tests, though, is that it’s easy to add new
tests when you need them. Suppose we weren’t sure how JoinWithCommas
would behave when it’s passed an empty slice. To find out, we simply add a
new testData struct in the tests slice. We’ll specify that if an empty slice
is passed to JoinWithCommas, an empty string should be returned:

```go
func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{list: []string{}, exp: ""}, //panic
		{list: []string{"apple"}, exp: "apple"},
		{list: []string{"apple", "orange"}, exp: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, exp: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.exp {
			t.Errorf(errorString(test.list, got, test.exp))
		}
	}
}
```
``` 
$ go test -v
=== RUN   TestJoinWithCommas
--- FAIL: TestJoinWithCommas (0.00s)
panic: runtime error: slice bounds out of range [:-1] [recovered]
        panic: runtime error: slice bounds out of range [:-1]
```
fix the empty slice
````go
package prose

import "strings"

//JoinWithCommas accepts a slice of strings to join, return the resulted phrase
func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		res := strings.Join(phrases[:len(phrases)-1], ", ")
		res += ", and "
		res += phrases[len(phrases)-1]
		return res
	}
}

````
