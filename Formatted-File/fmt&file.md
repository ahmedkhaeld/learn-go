# Input/Output
### Standard I/O
unix has the notion of three standard I/O streams
* Standard input
* Standard output
* Standard error (output) 

## File I/O
pkg os has function to open or create files, list directories, <br>
pkg io has utilities to read and write; bufio provide the buffered I/O scanners
<br>pkg io/ioutil has extra utilities such as reading an entire file to memory, or writing it out all at once.

```go
package main

import (
	"fmt"
	"io"
	"os"
)

// program works like unix cat(concatenate) command
// cat is a standard Unix utility that reads files sequentially, writing them to standard output.
// concatenate files, it will read each file and dump them together to the stdout
func main() {

	// loop through the cmd arguments from the second arg to the end 
	//as the first arg is for program name
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer file.Close()
	}
}


```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// calculate file size,

	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Println("The file has ", len(data), "bytes")

		file.Close()
	}
}

```


---
## fmt pkg
```go
package main

import "fmt"

func main() {
	a, b := 12, 325

	fmt.Printf("%d %d\n", a, b)   // decimal rep  12 325
	fmt.Printf("%X %X\n", a, b)   // hexadecimal rep  C 145
	fmt.Printf("%#X %#X\n", a, b) // hexadecimal rep 0XC 0X145

	c, d := 1.3, 5.49
	fmt.Printf("%f %f\n", c, d)     // float rep  1.300000 5.490000
	fmt.Printf("%.2f %.2f\n", c, d) // 1.30 5.49   %.2f limit the float point digits

	// the concept of using columns
	fmt.Printf("|%d|%d|\n", a, b)     //   |12|325|
	fmt.Printf("|%6d|%6d|\n", a, b)   // print with width of 6 char  |    12|   325|
	fmt.Printf("|%06d|%06d|\n", a, b) // |000012|000325|
	fmt.Printf("|%-6d|%-6d|\n", a, b) // left justify  |12    |325   |
}

```
```go
package main

import "fmt"

func main() {

	fmt.Println("#####################: Slice")
	s := []int{1, 2, 3}
	fmt.Printf("%T\n", s)  // type of s                      []int
	fmt.Printf("%v\n", s)  // value of s                     [1 2 3]
	fmt.Printf("%#v\n", s) // combination of type and value  []int{1, 2, 3}
	fmt.Println("#####################: ARRAY")

	a := [3]rune{'a', 'b', 'c'}
	fmt.Printf("%T\n", a)  // type of s                      [3]int32
	fmt.Printf("%q\n", a)  //                                ['a' 'b' 'c']
	fmt.Printf("%v\n", a)  // value of s to their byte rep   [97 98 99]
	fmt.Printf("%#v\n", a) // combination of type and value  [3]int32{97, 98, 99}

	fmt.Println("#####################: MAP")
	m := map[string]int{"and": 1, "or": 2}
	fmt.Printf("%T\n", m)  //map[string]int
	fmt.Printf("%v\n", m)  // map[and:1 or:2]
	fmt.Printf("%#v\n", m) // map[string]int{"and":1, "or":2}
}

```