# Code and project organization
---
* Organizing our code idiomatically
* Dealing efficiently with abstractions: interfaces and generics
* Best practices regarding how to structure a project

---
## 2.1 Unintended variable shadowing
The scope of a variable refers to the places a variable can be referenced: in other words, the part of an application where a name binding is valid. In Go, a variable name declared in a block can be redeclared in an inner block. This principle, called variable shadowing, is prone to common mistakes.

How can we ensure that a value is assigned to the original client variable?<br>
There are two different options.

The first option uses temporary variable in the inner blocks

The second option uses the assignment operator in the inner blocks to
directly assign the function results to the outer variable


---
## 2.2 Unnecessary nested code
Code is qualified as readable based on multiple criteria such as naming, consistency, formatting, and so forth. Readable code requires less cognitive effort to maintain a mental model; hence, it is easier to read and maintain.

>A critical aspect of readability is the number of nested levels.

Calls a concatenate function to perform some specific concatenation but may return errors
```go
func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}
```
This join function concatenates two strings and returns a substring if the length is greater than max. Meanwhile, it handles checks on s1 and s2 and whether the call to concatenate returns an error

From an implementation perspective, this function is correct. However, building a mental model encompassing all the different cases is probably not a straightforward task. Why? Because of the number of nested levels


Now, let’s try this exercise again with the same function but implemented differently:
```go
func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}
```

Let’s see some different applications of this rule to optimize our code for readability:

* When an if block returns, we should omit the else block in all cases. For example, we shouldn’t write
```go
if foo(){
	//...
	return true
}else{
	//..
}
```
Instead, we omit the else block like this:
```go
if foo(){
	//...
	return true
}
//...
```
With this new version, the code living previously in the else block is moved to the top level, making it easier to read.

* we can aslo follow this logic with a non-happy path:
```go
if s != ""{
	//...
}else{
	return errors.New("empty string")
}
```
Here, an empty s represents the non-happy path. Hence, we should flip the condition like so:
```go
if s==""{
	return errors.New("empty string")
}
//...
```
---
## 2.3 Misusing init functions
#### Concepts :
an init func is a function used to initialize the state
of an application. It takes no args and returns no results.
when a package is initialized, <br>
all the constant and variable declarations
in the package are evaluated.<br>
Then, the init functions are executed
```go
package main

import "fmt"

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
//var 
//init
//main
```

* define two packages, main and redis, where main depends
on redis. 
* because main depends on redis, the redis pkg's init func is executed
first, followed by the init of the main pkg, and then the main func itself

>We can define multiple init functions per package. When we do, the execution order of the init function inside the package is based on the source files’ alphabetical order.
> <br>if a package contains an a.go file and a b.go file and both have an init function, the a.go init function is executed first.

>We can also define multiple init functions within the same source file.
> <br>The first init function executed is the first one in the source order.


>We can also use init functions for side effects. In the next example, we define a main package that doesn’t have a strong dependency on foo (for example, there’s no direct use of a public function). However, the example requires the foo package to be initialized. We can do that by using the _ operator this way:
```go
package main

import (
	"fmt"
	_ "foo"
)

func main(){
	//...
}
```
import foo for side effects.<br>
In this case, the foo pkg is initialized before main. Hence
the init functions of foo are executed.

>an init func can't be invoked directly

### When to use init functions
an example where using an init function can be considered inappropriate: holding a database connection pool. In the init function in the example, we open a database using sql.Open. We make this database a global variable that other functions can later use:
```go
var db *sql.DB

func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = d
}
```
What should we think about this implementation?
<br>Let’s describe three main downsides.

First, error management in an init function is limited. Indeed, as an init function doesn’t return an error, one of the only ways to signal an error is to panic, leading the application to be stopped. In our example, it might be OK to stop the
application anyway if opening the database fails. However, it shouldn’t necessarily be up to the package itself to decide whether to stop the application. Perhaps a caller might have preferred implementing a retry or using a fallback mechanism. In this case, opening the database within an init function prevents client packages from implementing their error-handling logic.

Another important downside is related to testing. If we add tests to this file, the init function will be executed before running the test cases, which isn’t necessarily what we want (for example, if we add unit tests on a utility function that doesn’t require this connection to be created). Therefore, the init function in this example complicates writing unit tests.

The last downside is that the example requires assigning the database connection pool to a global variable.
<br>Global  variable have some severe drawback; for example:<br>
* any function cal alter global variables withing the pkg
* Unit test can be more complicated because a function depends on a
global variable won't be isolated anymore.

In most cases, we should favor encapsulating a variable rather than keeping it global.

For these reasons, the previous initialization should probably be handled as part of a plain old function like so
```go
func createClient(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
```
Using this function, we tackled the main downsides discussed previously. Here’s how:
* The responsibility of error handling is left up to the caller.
* It’s possible to create an integration test to check that this function works.
* The connection pool is encapsulated within the function.


Is it necessary to avoid init functions at all? Not really

uses an init function to set up the static HTTP configuration:
````go
// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command blog is a web server for the Go blog that can run on App Engine or
// as a stand-alone HTTP server.
package main

import (
	"net/http"
	"strings"
	"time"

	"golang.org/x/tools/blog"
	"golang.org/x/website/content/static"

	_ "golang.org/x/tools/playground"
)

const hostname = "blog.golang.org" // default hostname for blog server

var config = blog.Config{
	Hostname:     hostname,
	BaseURL:      "https://" + hostname,
	GodocURL:     "https://golang.org",
	HomeArticles: 5,  // articles to display on the home page
	FeedArticles: 10, // articles to include in Atom and JSON feeds
	PlayEnabled:  true,
	FeedTitle:    "The Go Programming Language Blog",
}

func init() {
	// Redirect "/blog/" to "/", because the menu bar link is to "/blog/"
	// but we're serving from the root.
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	// Keep these static file handlers in sync with app.yaml.
	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)

	http.Handle("/lib/godoc/", http.StripPrefix("/lib/godoc/", http.HandlerFunc(staticHandler)))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	b, ok := static.Files[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.ServeContent(w, r, name, time.Time{}, strings.NewReader(b))
}
````
In this example, the init function cannot fail (http.HandleFunc can panic, but only if the handler is nil, which isn’t the case here).

---
## 2.4 Overusing getters and setters
In programming, data encapsulation refers to hiding the values or state of an object. Getters and setters are means to enable encapsulation by providing exported methods on top of unexported object fields

the standard Go library doesn’t enforce using getters and/or setters even when we shouldn’t modify a field.


On the other hand, using getters and setters presents some advantages, including these:
* They encapsulate a behavior associated with getting or setting a field, allowing new functionality to be added later (for example, validating a field, returning a computed value, or wrapping the access to a field around a mutex).
* They hide the internal representation, giving us more flexibility in what we expose.
* They provide a debugging interception point for when the property changes at run time, making debugging easier.

> if we use them with a field called balance, we should follow these naming conventions:
> <br>The getter method should be named Balance (not GetBalance).
> <br>The setter method should be named SetBalance.

---
## 2.5 Interface pollution
Interface pollution is about overwhelming our code with unnecessary abstractions, making it harder to understand.

### Concepts
an interface provides a way to specify the behavior of an object.

We use interfaces to create common abstractions that multiple objects can implement.

* understand interfaces
The `io` pkg provides abstractions for I/O primitives.
<br>io.Reader relates to reading data from a data source<br>
and io.Writer to writing to a target

```go
type Reader interface{
	Read(p []byte) (n int, err error)
}
```
Custom implementations of the io.Reader interface should accept
a slice of bytes, filling it with its data, and returning either
the number of bytes read or an error

```go
type Writer interface{
	Write(p []byte) (n int, err error)
}
```
Custom implementations of io.Writer should write the data coming from a slice to a target and return either the number of bytes written or an error.

* io.Reader reads data from a source.
* io.Writer writes data to a target.

**What is the rationale for having these two interfaces in the language? What is the point of creating these abstractions?**<br>
Let’s assume we need to implement a function that should copy the content of one file to another. We could create a specific function that would take as input two *os.Files.

```go
func copySourceToDest(source io.Reader, dest io.Writer) error {
    // ...
}
```
This function would work with *os.File parameters (as *os.File implements both io.Reader and io.Writer)
<br>and any other type that would implement these interfaces.

For example, we could create our own io.Writer that writes to a database, and the code would remain the same. It increases the genericity of the function; hence, its reusability.

Furthermore, writing a unit test for this function is easier because, instead of having to handle files, we can use the strings and bytes packages that provide helpful implementations:
```go
package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCopySourceToDest(t *testing.T) {
	const input = "foo"
	source := strings.NewReader(input) //1
	dest := bytes.NewBuffer(make([]byte, 0)) //2

	err := copySourceToDest(source, dest)  //3
	if err != nil {
		t.FailNow()
	}

	got := dest.String()
	if got != input {
		t.Errorf("expected: %s, got: %s", input, got)
	}
}
```
1: creates an io.Reader<br>
2: creates an io.Writer<br>
3: Call copySourceToDest from *strings.Reader and a *bytes.Buffer

In the example, source is a *strings.Reader, whereas dest is a *bytes.Buffer. Here, we test the behavior of copySourceToDest without creating any files.


>While designing interfaces, the granularity (how many methods the interface contains) is also something to keep in mind.
> <br>The bigger the interface, the weaker the abstraction.


Furthermore, we can also combine fine-grained interfaces to create higher-level abstractions. This is the case with io.ReadWriter, which combines the reader and writer behaviors:
```go
type ReadWrite interface{
	Reader
	Writer
}
```

### When to use interfaces
* Common behavior
* Decoupling
* Restricting behavior

1. **Common behavior**
   use interfaces when multiple types implement a common behavior.

In such a case, we can factor out the behavior inside an interface.

If we look at the standard library, we can find many examples of such a use case. For example, sorting a collection can be factored out via three methods:

``` 
- Retrieving the number of elements in the collection
- Reporting whether one element must be sorted before another
- Swapping two elements
 ```
```go
type Interface interface{
	Len() int
	Less(i,j int)bool
	Swap(i,j int)
}
```
This interface has a strong potential for reusability because it encompasses the common behavior to sort any collection that is index-based.

Throughout the sort package, we can find dozens of implementations. If at some point we compute a collection of integers, for example, and we want to sort it, are we necessarily interested in the implementation type? Is it important whether the sorting algorithm is a merge sort or a quicksort? In many cases, we don’t care.
Hence, the sorting behavior can be abstracted, and we can depend on the sort.Interface.


Finding the right abstraction to factor out a behavior can also bring many benefits. For example, the sort package provides utility functions that also rely on sort.Interface, such as checking whether a collection is already sorted. For instance,

```go
func IsSorted(data Interface)bool{
	n :=data.Len() 
	for i:=n-1; i>0 ; i--{
		if data.Less(i,i-1)
		return false
        }
    }
	return true 
}

```


2. **Decoupling** <br>
   Another important use case is about decoupling our code from an implementation. If we rely on an abstraction instead of a concrete implementation, the implementation itself can be replaced with another without even having to change our code.

One benefit of decoupling can be related to unit testing. Let’s assume we want to implement a CreateNewCustomer method that creates a new customer and stores it. We decide to rely on the concrete implementation directly (let’s say a mysql.Store struct):
```go
type CustomerService struct{
	store mysql.Store
}

func(cs CustomerService) CreateNewCustomer(id sttring) error{
	customer:=Customer{id:id}
	return cs.store.StoreCustomer(customer)
}
```
Now, what if we want to test this method? Because customerService relies on the actual implementation to store a Customer, we are obliged to test it through integration tests, which requires spinning up a MySQL instance (unless we use an alternative technique such as go-sqlmock, but this isn’t the scope of this section). Although integration tests are helpful, that’s not always what we want to do. To give us more flexibility,

we should decouple CustomerService from the actual implementation, which can be done via an interface like so:
```go
type customerStorer interface{
	StoreCustomer(Customer) error
}

type CustomerService struct{
	storer customerStorer
}

func(cs CustomerService)CreateNewCustomer(id string)error{
	customer:=Customer{id:id}
	return cs.storer.StorerCustomer(customer)
}
```
Because storing a customer is now done via an interface, this gives us more flexibility in how we want to test the method.

>Use the concrete implementation via integration tests<br>
>Use a mock (or any kind of test double) via unit tests<br>
>Or both

3. Restricting behavior



---
## 2.6 Interface on the producer side
## 2.7 Returning interfaces

## 2.8 any says nothing

## 2.9 Being confused about when to use generics

## 2.10 Not being aware of the possible problems with type embedding

## 2.11 Not using functional options pattern

## 2.12 Project mis-organization

## 2.13 Creating utility packages

## 2.14 Ignoring package name collisions

## 2.16 Not using linters


