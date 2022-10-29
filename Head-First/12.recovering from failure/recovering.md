#back on your feet: Recovering from failure
>Whoo! I really panicked there, when I thought the data was
> corrupted! Give me a moment to recover, and then I'll close the file

##### Every Program encounters errors. You should plan for them
Sometimes handling an error can be as simple as reporting it and exiting the
program. But other errors may require additional action. You may need to
close opened files or network connections, or otherwise clean up, so your
program doesn’t leave a mess behind.

* how to defer cleanup actions, so they happen even when there's an error
* how to make your program panic in those(rare) situations where it's appropriate, and how to recover afterward.

### Reading numbers from a file
```go
package readfile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//OpenFile opens the file and return a pointer to it,
//along with any error encountered
func OpenFile(fn string) (*os.File, error) {
	fmt.Println("Opening", fn)
	return os.Open(fn)
}

//CloseFile closes a file
func CloseFile(f *os.File) {
	fmt.Println("Closing file")
	f.Close()
}

func GetFloats(fn string) ([]float64, error) {
	var nums []float64
	file, err := OpenFile(fn)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, number)
	}
	CloseFile(file)
	if scanner.Err() != nil {
		return nil, err
	}
	return nums, nil
}
```

```go
package main

import (
	"fmt"
	"log"
	"os"
	"recover/readfile"
)

func main() {

	// get the name of the file from cmd
	//os.Args is a slice of strings, [0] is the name of the program
	//get the name by accessing os.Args[1]

	//pass the fn to GetFloats, if any errors,
	//store them in  err var, then simply log it and exit

	numbers, err := readfile.GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f \n", sum)
}
```
`go run main.go data.txt`
```  
Opening data.txt
Closing file
Sum: 50.75 
```
We can see when the OpenFile and CloseFile functions get called, since
they both include calls to fmt.Println. And at the end of the output, we
can see the total of all the numbers in data.txt. Looks like everything’s
working!

### Any errors will prevent the file from being closed!
`go run main.go bad-data.txt`

```
Opening ddata.txt
2022/10/26 13:14:31 open ddata.txt: no such file or directory
exit status 1

```

Now, that in itself is fine; every program receives invalid data occasionally.
But the GetFloats function is supposed to call the CloseFile function
when it’s done. We don’t see “Closing file” in the program output, which
would suggest that CloseFile isn’t getting called!

The problem is that when we call strconv.ParseFloat with a string that
can’t be converted to a float64, it returns an error. Our code is set up to
return from the GetFloats function at that point.

But that return happens before the call to CloseFile, which means the file
never gets closed!

### Deferring functions calls
each file
that’s left open continues to consume operating system resources. Over
time, multiple files left open can build up and cause a program to fail, or
even hamper performance of the entire system. It’s really important to get in
the habit of ensuring that files are closed when your program is done with
them.

But how can we accomplish this?

If you have a function call that you want to ensure is run, no matter what,
you can use a defer statement. You can place the _**defer**_ keyword before
any ordinary function or method call, and GO will defer(delay)
making the function call until after current function exits

Normally, function calls are executed as soon as they’re encountered

>using defer for function calls that need to happen no matter what<br>
Normally, function calls are executed as soon as they’re encountered

### Ensuring files get closed using deferred function calls
```go
func GetFloats(fn string) ([]float64, error) {
	var nums []float64
	file, err := OpenFile(fn)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, number)
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return nums, nil
}

```
Now, CloseFile called no matter what, whether there is an error,
or GetFloats finishes normally well
>Note:moving CloseFile call immediately after the call to OpenFile

bad data within data.txt, strconv return an error 
``` 
20.25
5.0
hello
15.0
```
when running, file opens, while reading found error, still closed the file
``` 
Opening data.txt
Closing file
2022/10/26 13:31:05 strconv.ParseFloat: parsing "hello": invalid syntax
exit status 1
```

>Note:<br>
> Q: So I can defer function and method calls... Can I defer other
statements too, like for loops or variable assignments?
> 
> A: No, only function and method calls. You can write a function or method
to do whatever you want and then defer a call to that function or method,
but the defer keyword itself can only be used with a function or method
call.
---
### Listing the files in a directory
read the content of directory:<br>
So our program calls ReadDir, passing it the name of my_directory as an
argument. It then loops over each value in the slice it gets back. If IsDir
returns true for the value, it prints "Directory:" and the file’s name.
Otherwise, it prints "File:" and the file’s name.

The `io/ioutil` pkg includes a `ReadDir` func that will let us
read the directory contents. You pass ReadDir the name of a directory, and it will
return a slice of values, one for each file or subdirectory the directory
contains (along with any error it encounters).

Each of the slice’s values satisfies the FileInfo interface, which includes a
Name method that returns the file’s name, and an IsDir method that returns
true if it’s a directory.
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	files, err := ioutil.ReadDir("my_directory")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
```

```  
File: a.txt
Directory: subdir
File: z.txt

```

#### Listing the files in subdirectories(will be trickier)
showcase:<br>
* go >
    * src >
        * geo >
            * coordinates.go
            * landmark.go
        * locked >
            * secret.go
    * vehicle >
        * car.go

like a Go workspace directory. That would contain an entire
tree of subdirectories nested within subdirectories, some containing files,
some not.

``` 
A.Get the next file.
B.Is the file a dir?
    1. if yes: get a list of files in the dir.
        a.Get the next file
        b.Is the file a dir?
            01. if yes: Get a list of the files in the dir...
            
    2. if no:just print the filename.
```
simplified logic
``` 
1.Get a list of file in the directory
    a.Get the next file
    b.Is the file a directory?
        i. If yes: start over at step 1 with this directory
        ii. if no: just print the filename.
```

#### Recursive function calls
Go supports recursion, which allows a function to call itself
>make sure that the recursion loop stops itself eventually

```go
package main

import (
	"fmt"
)

func count(start, end int) {
	fmt.Printf("count(%d,%d) called \n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning from count(%d,%d) call \n", start, end)

}

func main(){
	count(1, 3)

}
```
``` 
count(1,3) called 
1
count(2,3) called 
2
count(3,3) called 
3
Returning from count(3,3) call 
Returning from count(2,3) call 
Returning from count(1,3) call 
```
---
### Recursively listing directory contents
```go
package rdir

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//ScanDir takes the path of the directory it should scan,
//first it prints the current dir, so we know what dir we are working
//then it calls ReadDir method on that path to get the dir contents
// loops over the slice of FileInfo values ReadDir returns,
//processing each one.
//it calls filepath.Join to join the current dir path and current filename together with slashes
//(so "go" and "src" are joined to become "go/src")
//if the current file isn't directory, ScanDir just prints its full path
//and moves on to the next file(if there are anymore in the current directory)
//if the current file is dir, the recursion kicks in: ScanDir calls itself
//with the subdirectories' path
//if that subdirectory has any subdirectories, ScanDir will call itself
func ScanDir(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	//join dir path and file name with a slash
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := ScanDir(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

```
```go
package main

import (
	"log"
	"recover/rdir"
)

func main() {

	err := rdir.ScanDir("go")
	if err != nil {
		log.Fatal(err)
	}
}

```
```
go
go/src
go/src/geo
go/src/geo/coordinates.go
go/src/geo/landmark.go
go/src/locked
go/src/locked/secret.go
go/src/vehicle
go/src/vehicle/car.go

```
1. main calls scanDirectory with a path of "go"
2. scanDirectory prints the path it’s passed, "go", indicating the
   directory it’s working in
3. It calls ioutil.ReadDir with the "go" path
4. There’s only one entry in the returned slice: "src"
5. Calling filepath.Join with the current directory path of "go"
   and a filename of "src" gives a new path of "go/src"
6. src is a subdirectory, so scanDirectory is called again, this time
   with a path of "go/src"<br>
7. scanDirectory prints the new path: "go/src"
8. It calls ioutil.ReadDir with the "go/src" path
9. The first entry in the returned slice is "geo"
10. Calling filepath.Join with the current directory path of
    "go/src" and a filename of "geo" gives a new path of
    "go/src/geo"
11. geo is a subdirectory, so scanDirectory is called again, this time
    with a path of "go/src/geo"
12. 12. scanDirectory prints the new path: "go/src/geo"
13. It calls ioutil.ReadDir with the "go/src/geo" path
14. The first entry in the returned slice is "coordinates.go"
15. coordinates.go is not a directory, so its name is simply printed
16. And so on...

## Error handling in a recursive func
If scanDirectory encounters an error while scanning any subdirectory (for
example, if a user doesn’t have permission to access that directory), it will
return an error. This is expected behavior; the program doesn’t have any
control over the filesystem, and it’s important to report errors when they
inevitably occur

### Starting a panic
when a program panics, the current func stop running, and
the program prints a log msg and crashes.

* Stack traces
each func that's called need to return to the function that called it.
<br>Go keeps a _**call stack**_, a list of the function calls
that are active at any given point.

```go
package main

func main() {
	one() //this func call get added to the stack
}

func one() {
	two() //this func call get added to the stack
}
func two() {
	three() //this func call get added to the stack
}
func three() {
	panic("This call stack's too deep") 
	//panic! the stack trace will include the above calls
}

```

``` 
panic: This call stack's too deep

goroutine 1 [running]:
main.three(...)
        /home/ahmed/go-advanced/Head-First/12.recovering from failure/main.go:20
main.two(...)
        /home/ahmed/go-advanced/Head-First/12.recovering from failure/main.go:17
main.one(...)
        /home/ahmed/go-advanced/Head-First/12.recovering from failure/main.go:14
main.main()
        /home/ahmed/go-advanced/Head-First/12.recovering from failure/main.go:10 +0x29
exit status 2

```

* Deferred calls completed before crash
when a program panics, all the deferred calls will still be made
<br> if there's more than one deferred call, they'll be
made in the reverse order they were deferred in.

### Using panic with ScanDir
The scanDirectory function at the right has been updated to call panic
instead of returning an error value. This greatly simplifies the error
handling
````go
package rdir

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)


func ScanDir(path string)  {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	//join dir path and file name with a slash
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDir(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}
````
```go
package main

import (
	"recover/rdir"
)

func main() {
	 rdir.ScanDir("go")
}
```
Now, when scanDirectory encounters an error reading a directory, it
simply panics. All the recursive calls to scanDirectory exit.

## When to panic
calling panic may simplify the code, but it also crashes the
program! That does not seem like much of an improvement...

>calling panic should be reserved for “impossible”
situations: errors that indicate a bug in the program, not a mistake on the
user’s part.

>Things like inaccessible files, network failures, and bad user input should
usually be considered “normal,” and should be handled gracefully though
error values.

Here’s a program that uses panic to indicate a bug. It awards a prize hidden
behind one of three virtual doors. The doorNumber variable is populated not
with user input, but with a random number chosen by the rand.Intn
function. If doorNumber contains any number other than 1, 2, or 3, **it’s not**
user error, it’s a bug in the program.

So it makes sense to call panic if doorNumber contains an invalid value. It
should never happen, and if it does, we want to stop the program before it
behaves in unexpected ways.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}

func awardPrize() {
	//number generated internally, must 1 or 2 or 3
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("There is a bug! invalid door number")
	}
}
```

## The recover function
Changing our scanDirectory function to use panic instead of returning an
error greatly simplified the error handling code. But panicking is also
causing our program to crash with an ugly stack trace. We’d rather just
show users the error message

Go offers a built-in recover function that can stop a program from
panicking. We’ll need to use it to exit the program gracefully.


If you call recover when a program is panicking, it will stop the panic. But
when you call panic in a function, that function stops executing. So there’s
no point calling recover in the same function as panic, because the panic
will continue anyway:
```go
func feakOut(){
	panic("oh no")
	recover()
}
```
>Important Practice:<br>
But there is a way to call recover when a program is panicking... During a
panic, any deferred function calls are completed. So you can place a call to
recover in a separate function, and use defer to call that function before
the code that panics.

```go
func calmDown(){
	recover()
}
func feakOut(){
	defer calmDown()
	panic("oh no")
	fmt.Println("I won't be run!")
}

func main(){
	freakOut()
	fmt.Println("Exiting normally")
}
```
`Exiting normally` program exits normally<br>
Calling recover will not cause execution to resume at the point of the
panic, at least not exactly. The function that panicked will return
immediately, and none of the code in that function’s block following the
panic will be executed. After the function that panicked returns, however,
normal execution resumes.
The panic value is return

### The panic value is returned from recover
when there is no panic, a call to recover return nil

But when there is panic, recover returns whatever value was passed to panic
this can be used to gather information about the panic,
to aid recovering or to report errors to the user
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}

func awardPrize() {
	//number generated internally, must 1 or 2 or 3
	doorNumber := rand.Intn(3)
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		defer calmDown()
		panic("There is a bug! invalid door number")
	}
}

func calmDown() {
	fmt.Println(recover())
}

```
if 0 if randomly provided
`"There is a bug! invalid door number"` printed out 

* panic parameter type and recover return type<br>
    * panic takes an empty interface
    * recover returns an empty interface

You can pass recover’s return value to fmt functions like
Println (which accept interface{} values), but you won’t be able to call
methods on it directly.


Here’s some code that passes an error value to panic. But in doing so, the
error is converted to an interface{} value. When the deferred function
calls recover later, that interface{} value is what’s returned. So even
though the underlying error value has an Error method, attempting to call
Error on the interface{} value results in a compile error
```go
func calmDown() {
    p :=recover()  //return an empty interface has no methods
	fmt.Println(p.Error()) // so p does not include the Error() method on it
}
```
need to do type assertion, convert back it to the underlying type
```go
func calmDown() {
    p :=recover()  
	err, ok:=p.(error)
	if ok{
		fmt.Println(err.Error())
    }
}

func main (){
	defer calmDown()
	//define an error message
	err :=fmt.Errorf("there's an error") 
	//assign err message to the panic
	panic(err)
	//when panicking, recover can take that error type and print it
	//because  type asserted 
}
```


## Recovering from panic in ScanDir
