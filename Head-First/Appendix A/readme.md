# understanding os.Openfile: Opening Files

---
some programs need to write data to files, not just read data

### Understanding os.OpenFile
```go
package main

import "os"

func main() {
	options:=os.O_WRONLY |os.O_APPEND | os.O_CREATE
	file, err:=os.OpenFile("signatures.txt", options,os.FileMode(0600))
}

```
**let's figure out what this `flag` means:**<br>
From the documentation, it looks like os.O_RDONLY is one of several int
constants intended for passing to the os.OpenFile function, which change
the function’s behavior.

`read file contents` pass os.O_RDONLY option to OpenFile func
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("signatures.txt", os.O_RDONLY, os.FileMode(0600))
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if scanner.Err() !=nil{
		log.Fatal(err)
	}
}
```
if you try to write to the same file, with `os.WRONLY` option
so, it opens the file for writing. it will produce no output, but it will update the
aardvark.txt file. But if we open aardvark.txt, we’ll see that instead of
appending the text to the end, the program overwrote part of the file!
```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("signatures.txt", os.O_WRONLY, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte("amazing!\n"))
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}

```
That’s not how we wanted the program to work. What can we do?<br>
`os.O_APPEND` flag should cause the programs to append data
to the file instead of overwriting it.
>But you can't just pass `os.O_APPEND` to `os.OpenFile` by itself
> you'll get an error if you try. <br>
> The docs says that os.O_APPEND and os.O_CREATE be OR'ED in.
> this referring to the binary OR operator.

---
## Binary notation
Bitwise operators: & Bitwise AND, | Bitwise OR

#### AND operator
if the bit at a given place in the first number is 1... and
the bit at the same place in the second number is 1... then
the bit at the same place in the result will be a 1

#### OR operator
he bitwise OR operator looks at the bits at a
given position in the two values it’s operating on to decide the value of the
bit at the same position in the result.
 
**the use of the bitwise OR operator to combine the
constant values together**

```go
	fmt.Printf("%016b read only\n", os.O_RDONLY)
fmt.Printf("%016b write only\n", os.O_WRONLY)
fmt.Printf("%016b read write\n", os.O_RDWR)
fmt.Printf("%016b create\n", os.O_CREATE)
fmt.Printf("%016b append\n", os.O_APPEND)
//0000000000000000 read only  0
//0000000000000001 write only 1
//0000000000000010 read write 2
//0000000001000000 create     64
//0000010000000000 append     1024
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf(
		"write only ORed create " +
		"\n%016b \n |\n%016b\n=\n%016b \n", 
		os.O_WRONLY, os.O_CREATE, os.O_WRONLY|os.O_CREATE)
	fmt.Println(os.O_WRONLY, "|", os.O_CREATE, "=", os.O_WRONLY|os.O_CREATE)
	//write only ORed create 
	//0000000000000001 
	// |
	//0000000001000000
	//=
	//0000000001000001 
	//1 | 64 = 65

}
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("write only ORed create ORed append \n%016b\n|\n%016b\n|\n%016b\n=\n%016b \n", os.O_WRONLY, os.O_CREATE, os.O_APPEND, os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	fmt.Println(os.O_WRONLY, "|", os.O_CREATE, "|", os.O_APPEND, "=", os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	//write only ORed create ORed append 
	//0000000000000001
	//|
	//0000000001000000
	//|
	//0000010000000000
	//=
	//0000010001000001 
	//1 | 64 | 1024 = 1089

}

```
The os.OpenFile function can check whether the first bit is a 1 to
determine whether the file should be write-only. If the seventh bit is a 1,
OpenFile will know to create the file if it doesn’t exist. And if the 11th bit
is a 1, OpenFile will append to the file.

---
### Using bitwise OR to fix our os.OpenFile options
Let’s see if we can
combine options so that it appends new data to the end of the file instead.
```go
package main

import (
	"log"
	"os"
)

func main() {
	options:=os.O_WRONLY|os.O_APPEND |os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte("amazing!\n"))
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}


```

### Unix-style file permissions
`os.OpenFile` permissions parameter<br>
which controls the file’s
permissions: which users will be permitted to read from and write to the file
after your program creates it.

>When developers talk about file permissions, they usually mean
permissions as they’re implemented on Unix-like systems like macOS and
Linux. Under Unix, there are three major permissions a user can have on a

| Abbreviation | Permission                                                                                |
|--------------|-------------------------------------------------------------------------------------------|
| r            | The user can Read the file's content                                                      |
| w            | The user can Write the file's content                                                     |
| x            | The user can execute the file.(this only appropriate for files that contain program code) |

If a user doesn’t have the permissions on a file, for example, any program
they run that tries to access the file’s contents will get an error from the operating system


* Notice <br
  Each file has three sets of permissions, affecting three different classes of
  users. The first set of permissions applies only to the user that owns the file.
  (By default, your user account is the owner of any files you create.) The
  second set of permissions is for the group of users that the file is assigned
  to. And the third set applies to other users on the system that are neither the
  file owner nor part of the file’s assigned group.
* `-rwx------` the file's owner will have full permissions 
* `fmt.Println(os.FileMode(0700)`
* `----rwx---` Users in the file's group will have full permissions
* `fmt.Println(os.FileMode(0070)`
* `-------rwx` All other users on the system will have full permissions
* `fmt.Println(os.FileMode(007)`
* 7 means full permission for the set, permission ranges from 0 to 7 (octal representation)

* `fmt.Println(os.FileMode(0000)` `----------`
* `fmt.Println(os.FileMode(0111)` `---x--x--x`
* `fmt.Println(os.FileMode(0222)` `--w--w--w-`
* `fmt.Println(os.FileMode(0333)` `--wx-wx-wx`
* `fmt.Println(os.FileMode(0444)` `-r--r--r--`
* `fmt.Println(os.FileMode(0555)` `-r-xr-xr-x`
* `fmt.Println(os.FileMode(0666)` `-rw-rw-rw-`
* `fmt.Println(os.FileMode(0777)` `-rwxrwxrwx`

```go
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
```
`os.FileMode(0644)` means rw for file's owner, r for group, r for any other user