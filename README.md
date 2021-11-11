# learn-go
* Strong and statically typed:
> means that the type of a variable cannot change over time. So when
you declare a variable a to hold an integer, it's always going to hold an integer, you
can't put a Boolean in it, you can't put a string in it. And static typing means that
all of those variables have to be defined at compile time.
* Key features:
  * simlicity
  * Fast compile times
  * Garbage collected
  > which means that you're not going to have to manage your own memory. Now you can manage
your own memory. But by and large, the go runtime is going to manage that for you.
  * Built-in concurrency
  * Compile to standalone binaries
  ---
  # Variables
  * Variable declaration 
  * Redeclaration and shadowing
  * Visiblity
  * Naming conventions
  * Type conventions
  

````go
package main

import (
	"fmt"
)

func main() {
	var i int               //var declration, case is not ready to assign a value yet 
	i=45                   // var assignment
	var j int =66         //declration and assignment same line, case we need to force a type not let it to the value 
  k :=10               // var assingment the declration is figured out by go compiler
	fmt.Println(i, j, k)
}

````

* at the package level you can not use k:= synatx you have to declare the variable
````go
package main

import (
	"fmt"
)

var j int =66         //package level

func main() {
 
	fmt.Println( j)
}
````
* declare a block of variables
````go
var(
	actorName string="ahmed"
	companion string="hamo"
	doctorNumber int=7
	season       int=9
)

func main() {
 
	fmt.Println()
}
````

* variable shadwing: function level takes the precedence on the package level for that case, but the package level still avialable for others
````go

var i int=27

func main() {
 	var i int=9
	fmt.Println(i)
}
````
* declared variable must be used in go
### rules to declare vars
* how naming convetion can affect scope visiblity
1. lower case only visible to the package 
````go
package main

import (
	"fmt"
)

var i int=27

func main() {
	
	fmt.Println(i)
}
````
2. upper case visible to the outsider packages
````go 
package main

import (
	"fmt"
)

var I int=27

func main() {
	
	fmt.Println(i)
}
````
* the scope of the variable defines how long is the name to a var
* acronyms should be upper case like HTTP, URL

## how to convert var type
* show case to convert an integer to string
1. this will output the string with 27 in the unicode not a 27 conveted to string type

```go
package main

import (
	"fmt"
)

var i int=27

func main() {

	fmt.Printf("%v, %T\n", i, i)
	
	var j string
	j=string(i)
	fmt.Printf("%v, %T\n", j, j)
}
```
2. you have to import strconv to actually change the type of i for j  

```go
package main

import (
	"fmt"
	"strconv"
)

var i int=27

func main() {

	fmt.Printf("%v, %T\n", i, i)
	
	var j string
	j=strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
```

arithmatic ops 
```go
package main

import (
	"fmt"
)


func main() {
	a :=10
	b :=3
	fmt.Printf("%v\t, %T, %T\t\n", a + b, a ,b)      // 13
	fmt.Printf("%v\t, %T, %T\t\n", a - b, a ,b)     // 7
	fmt.Printf("%v\t, %T, %T\t\n", a * b, a ,b)    // 30
	fmt.Printf("%v\t, %T, %T\t\n", a / b, a ,b)   // 3 dops the float because it the two operands are int
	fmt.Printf("%v\t, %T, %T\t\n", a % b, a ,b)  //1  pick up the remainder
}
```

if two var one is int and the second is int8 you have to type convert
```go
package main

import (
	"fmt"
)


func main() {
	var a int =10
	var b int8 =3
	fmt.Println(a + int(b)) // or 	fmt.Println(int8(a) + b)
	
}
```
wisbit operations 
```go

package main

import (
	"fmt"
)


func main() {
	a :=10                 // 0b1010
	b :=3                 // 0b0011
	
	fmt.Println(a & b)     // AND   =0010    =2
	fmt.Println(a | b)    // OR    =1011    =11
	fmt.Println(a ^ b)   //XOR    =1001    =9
	fmt.Println(a &^ b) //ANDNOT =0100    =8
	
}
```
shifting means 2^x : base 2 to the power x<br/>
shiftleft means multiply, shiftright means divid
```go
package main

import (
	"fmt"
)


func main() {
	a :=8  //2^3=8
	fmt.Println(a<<2)  //shiftleft means multiply 2^3 * 2^2=2^5 =32
	fmt.Println(a>>3)  //shiftright mean divid   2^3 / 2^3 =2^1 =1
	
	
}
```
> shifting and remainder operators are only with intgers

## complex numbers
there are two types of complex numbers. There's complex 64, and complex 128. go undersand the equations of the complex numbers

```go
package main

import (
	"fmt"
)


func main() {
	var n complex64 =1 + 2i
	var m complex64 = 2i
	fmt.Printf("%v, %T\n", n, n)  // (1+2i), complex64
	fmt.Printf("%v, %T", m,m)    // (0+2i), complex64
	
}
```
operations with complex nubmers
```go 
package main

import (
	"fmt"
)


func main() {
	var n complex64 =1 + 2i
	var m complex64 = 2i
	fmt.Println( n+m)  // (1+4i)
	fmt.Println( n-m)  //(1+0i)
	fmt.Println( n*m)  //(-4+2i)
	fmt.Println( n/m)  //(1-0.5i)
}
```
destructing the complex number to get real and imagine number
```go
package main

import (
	"fmt"
)


func main() {
	var n complex64 =1 + 2i
  	fmt.Printf("%v , %T\n", real(n), real(n)) // 1 , float32
	fmt.Printf("%v , %T\n", imag(n), imag(n))  // 2 , float32

}
```
crate a complex number by Complex function
```go 
package main

import (
	"fmt"
)


func main() {
	var n complex128 =complex(5, 15)
	fmt.Printf("%v, %T\n", n, n)  //(5+15i), complex128
}
```
## Texting types
* String type
represent UTf-8 charactars
```go

func main() {
	s := "this a string"
	fmt.Printf("%v, %T\n", s, s)
}
```
one of the interesting aspects of a string is we can actually treat it sort of like an array. treat the string of text as a collection of letters.

```go

func main() {
	s := "this a string"
	fmt.Printf("%v, %T\n", s[2], s[2])  // 105, uint8
}
```
>  * what the heck happened there? Well, what's happening is that strings in go are actually aliases for bytes.<br/>
>  * strings are generally immutable.<br/>
>  * there is one arithmetic or pseudo arithmetic operation that we can do with strings, and that is string concatenation. Or in simpler terms, we can add strings together.
```go
func main() {
	s := "this a string"
	s2 :="iam the second string "
	fmt.Println( s+s2)
}
```
convert them to collections of bytes
```go
func main() {
	s := "this a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b,b)
}
```
> [116 104 105 115 32 97 32 115 116 114 105 110 103], []uint8 <br/>
> we actually get this as a string comes out as the ASCII values, or the UTF valuesfor each character in that string.

>why would you use this one? It's a very
good question. A lot of the functions that we're going to use in go actually work with
byte slices. And that makes them much more generic and much more flexible than if we
work with hard coded strings. So for example, if you want to send as a response to a web
service call, if you want to send a string back, you can easily convert it to a collection
of bytes. But if you want to send a file back, well, a file on your hard disk is just a collection
of bytes, too. So you can work with those transparently and not have to worry about
line endings and things like that. So while in your go programs, you're going to work
with strings a lot as strings. When you're going to start sending them around to other
applications or to other services, you're very often going to take advantage of this
ability to just convert it to a byte slice.


* Rune 
represent UTF-32 charactars int32

---

# Constants
* Naming convention
* Typed constants
* Untyped constants
* Enumerated constants
* Enumeration expressions
```go

func main() {
	const myConst   // internal constant
	const MyConst   // global constant to be exported
  const i int =40  // named type constant
}
```
characteristic of a constant is that it has to be assignable at compile time. you can not assign a calculation of something to const
```go 

func main() {
	
	const myConst float64=math.Sin(1.57)
	fmt.Printf("%v, %T\n", myConst, myConst) // get error
}
```
## Enums
what is Iota? Well, Iota is a counter that we can use when we're creating what are called enumerated constants.
```go
package main

import "fmt"

const(
	a=iota
	b=iota
	c=iota
)

func main() {
	fmt.Printf("%v\n",a) //0
	fmt.Printf("%v\n",b) //1
	fmt.Printf("%v\n",c) //2

}
```
if we don't assign the value of a constant after the first one, then the compiler is going to try and figure the pattern of assignments.
<br/>that value of Iota is scoped to that constant block.

use iota as flag checking, also 
we can use iota to check a variable is assigned a value yet, or equal to zero value of the constant

```go
package main

import "fmt"

const(
	errorSpecialist=iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int
	fmt.Printf("%v\n",specialistType==catSpecialist)

}
```
we can use this underscore symbol if we don't care about zero, then we don't have any reason
to assign the memory to it.<br/>
And basically, what that tells the compiler
is yes, I know you're going to generate a value here, but I don't care what it is go
ahead and throw that away.


this can be valuable if you need some kind of a fixed offset. 
```go
package main

import "fmt"

const(
	_ =iota +5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int
	fmt.Printf("%v\n",specialistType==catSpecialist)
	fmt.Printf("%v\n",catSpecialist)
	fmt.Printf("%v\n",dogSpecialist)
	fmt.Printf("%v\n",snakeSpecialist)

}
```

use case 
shifleft is essentially multiply by 2 to the power of x

```go
package main

import "fmt"

const(
	_ =iota  // igonre first value by assigning to blank identifier
	KB=1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize :=4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
```
> - =iota is zero<br/>
> KB=1<<0 MEANS NO SHIFTING<BR/>
> MB=1<<10   MEANS 1*2^10                   <BR/>
> GB=1<<100  MEANS 1*2^100                      <BR/>
> TB=1<<1000 MEANS 1*2^1000

* another use case<br/>
let's just say that we've got an application and that application has
users and those users have certain roles. So inside of this constant block, here, I'm
defining various roles that we can have. So for example, you might be an admin, you might
be at the headquarters or out in the field somewhere, you might be able to see the financials
or see the monetary values. And then there may be some regional roles. So can you see
properties in Africa, can you see properties in Asia, Europe, North America, or South America.
So in order to define these constants, what I'm doing is I'm setting the value to one
bit shifted iota.

So the first constant is admin is one bit shifted zero places, so it's
a literal one, the second one is one bit shifted one place, that's two, and then four, and
then eight, and then 16, and so on.

in the main program, I'm defining the roles in a single byte.

```go
package main

import "fmt"

const(
	isAdmin=1<<iota            // 1
	isHeadquarters            //2
	canSeeFinancials         //4
	canSeeAfrica            //8
	canSeeAsia             //16
	canSeeEurope          //32
	canSeeNorthAmerica   //64
	canSeeSouthAmerica  //128
	
)

func main() {
	var roles byte= isAdmin | canSeeFinancials | canSeeEurope
  
	fmt.Printf("%b\n", roles)   // binary representation of ORing the three active roles and storing them in one variable
  
	// check if a user is admin or any other role to check against using bitwise bitmask mathimatics
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
  
	fmt.Printf("Can see Africa? %v", canSeeAfrica&roles == canSeeAfrica)

}
```



















 



















































































































