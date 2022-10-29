# You are my type: Defined Types
>almost done with the definition for my `Name` type! its underlying
type is string, and you'll be able to call `Capitalize` method
on any `Name` value. so convenient!

#### Type errors in real life
-when you have a number, it's best to certain what that number
is measuring. You want to know if it's liters or gallon,
kilograms or pounds, dollars or yen.

##### Defined types with underlying basic types
if you have the following variable:<br>
`var fuel float64 =10`<br>
does that represent 10 gallons or 10 liters?<br>
the person who wrote that declaration knows, but no
one else does, not for sure.

-You can use Go’s defined types to make it clear what a value is to be used
for. Although defined types most commonly use structs as their underlying
types, they can be based on int, float64, string, bool, or any other type.

```go
package main

import "fmt"

func main() {
	var carFuel Gallons //def a var with a type of Gallons
	var busFuel Liters  //def a var with a type of Liters

	carFuel = Gallons(10.0) //convert a float64 to Gallons
	busFuel = Liters(240.0) //convert a float64 to Liters
	fmt.Print(carFuel, busFuel)
}

//define two new types, each with an underlying type of float64

type Liters float64
type Gallons float64
```
Once you’ve defined a type, you can do a conversion to that type from any
value of the underlying type.

If we had wanted, we could have written short variable declarations in the
code above using type conversions:
``` 
//use short var declarations together with type conversions
carFuel:=Gallons(10.0)
busFuel:=Liters(240.0)
```

you cannot assign a value of
a different defined type to it,
````go
package main

import "fmt"

func main() {
	var carFuel Gallons
	var busFuel Liters

	//you cannot assign a value of
	//a different defined type to it,
	//even if the other type has the same underlying type.
	//This helps protect developers from confusing the two types.
	carFuel = Liters(10.0)
	busFuel = Gallons(240.0)
	fmt.Print(carFuel, busFuel)
}

type Liters float64
type Gallons float64

````
```
Errors
./main.go:13:10: cannot use Liters(10) (type Liters) as type Gallons in assignment
./main.go:14:10: cannot use Gallons(240) (type Gallons) as type Liters in assignment

```

But you can convert between types that have the same underlying type.
Liters can be converted to Gallons and vice versa, because both have an
underlying type of float64.
>A quick web search shows that one liter equals roughly 0.264 gallons, and
that one gallon equals roughly 3.785 liters. We can multiply by these
conversion rates to convert from Gallons to Liters, and vice versa.
```go
package main

import "fmt"

func main() {
	var carFuel Gallons
	var busFuel Liters

	//a car vehicle owner knows fuel with Gallons but at the station gives by Liters
	// so convert the liters to gallons
	carFuel = Gallons(Liters(40.0) * 0.264) 
	//40 liters gives 10.6 gallons
	
	busFuel = Liters(Gallons(63.0) * 3.785) 
	// 63 gallons gives 238.5 liters
	
	fmt.Print(carFuel, busFuel)
}

type Liters float64
type Gallons float64
```
#### Defined types and operators
* A defined type supports all the same operations as its underlying type.
Types based on float64, for example, support arithmetic operators like +,
-, *, and /, as well as comparison operators like ==, >, and <.

* A defined type can be used in operations together with literal values:
* But defined types cannot be used in operations together with values of a
  different type, even if the other type has the same underlying type. Again,
  this is to protect developers from accidentally mixing the two types.

####Converting between types using functions
Suppose we wanted to take a car whose fuel level is measured in Gallons
and refill it at a gas pump that measures in Liters. Or take a bus whose
fuel is measured in Liters and refill it at a gas pump that measures in
Gallons. To protect us from inaccurate measurements, Go will give us a
compile error if we try to combine values of different types:
```go
package main

import "fmt"

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)

	carFuel += toGallons(Liters(40.0))
	busFuel += toLiters(Gallons(30.0))

	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
	//Car fuel: 11.8 gallons
	//Bus fuel: 118.1 liters
}

type Liters float64

func toGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

type Gallons float64

func toLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

```
The metric system has other units of
measure as well, but the milliliter (1/1000 of a liter) is the most commonly
used. `type Milliliters floa64`<br>
we are also going to want a way to convert from Milliliters to 
Gallons, we run into a problem: we can not have two `ToGallons f`unctions in the same
pkg
>Wouldn't it be dreamy if you could write a ToGallons function
> that worked with Liters values, and another ToGallons
> function that worked with Milliliters values?

We can define methods of our own to help with our type conversion
problem.
#### Defining methods
```go
package main
import "fmt"

type MyType string 

func (m MyType)sayHi(){
	fmt.Println("HI!")
}

func main(){
	value:=MyType("a MyType value")
	value.sayHi()
	anotherValue:=MyType("another value")
	anotherValue.sayHi()
}
```
`HI<br> HI`
Once a method is defined on a type, it can be called on any value of that type

The receiver parameter is (pretty much) just
another parameter
```go
package main
import "fmt"

type MyType string 

func (m MyType)sayHi(){
	fmt.Println("Hi from", m) 
}

func main(){
	value:=MyType("a MyType value")
	value.sayHi()
	anotherValue:=MyType("another value")
	anotherValue.sayHi()
}
```
`Hi from a MyType value <br> Hi from another value`
>Go uses receiver parameters instead of the “self” or “this” values seen
in other languages.

>important!<br>
> Q: Can i define new methods on any type?<br>
> A: Only type that are defined in the same pkg where you define
> the method. That means no defining methods for types from
> someone else's security pkg from your hacking pkg, and no defining
> methods on universal types like int or string
> ---
> Q: But I need to be able to use methods of my own with someone else's type!<br>
> A: First you should consider whether a function would work well enough;
> a function can take any type you want as a parameter. But if you really need
> a value that has some methods of your own, plus some methods from a type
> in another pkg, you can make a struct type that embeds the other
> pkg's type as anonymous field. 

#### Pointer receiver parameters
```go
package main

import "fmt"

type Number int 

func(n Number)Double(){
	n*=2
}

func main(){
	num :=Number(4) //type cast the 4 from int to Number type
	fmt.Println("Original value of number:", num)
	num.Double()
	fmt.Println("number after calling Double:", num)
	
}
```
`Original value of number: 4 `<br>` number after calling Double: 4`

* We learned that func parameters receive a copy of the values the
function is called with, not the original values, and that
any updates to the copy would be lost when the function exited.
* To make the Double fun work, we had to pass a pointer to the original value we wanted to update
```go
package main

import "fmt"

type Number int 

//Double has  a pointer receiver to Number
func(n *Number)Double(){
	*n*=2 // update the value at the pointer 
	//*n value in the address 
}

func main(){
	num :=Number(4) //type cast the 4 from int to Number type
	fmt.Println("Original value of number:", num) //4
	num.Double()
	fmt.Println("number after calling Double:", num) //8
	
}
```
>you should avoid mixing type receivers, pointer or value


#### Converting Liters and Milliliters to Gallons using methods
```go
package main

import "fmt"

func main() {
	//soda is instance Liters that contains 2 liters
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.toGallons())

	//water is instance of Milliliters that contains 500ml
	water := Milliliters(500)
	fmt.Printf("%0.0f milli-liters equals %0.3f gallons\n", water, water.toGallons())
	
	//milk is instance of Gallons, that contains 2 gallons of milk
	milk:=Gallons(2)
	fmt.Printf("%0.3f gallons equal %0.3f liters", milk, milk.toLiters())
	fmt.Printf("%0.3f gallons equal %0.3f milli-liters", milk, milk.toMilliliters())

}

type Milliliters float64

func (ml Milliliters) toGallons() Gallons {
	return Gallons(ml * 0.000264)
}

type Liters float64

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

type Gallons float64

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) toMilliliters() Liters {
	return Liters(g * 3785.41)
}

```
