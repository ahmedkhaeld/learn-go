# what can you do?: Interfaces
>No, it's not quite the same as a car...<br>
But as long as there's a method for steering it, I think i can handle it!

>Sometimes you don't care about the particular type of value.<br>
> you don't care about what it is. You just need to know that it will
> be able to do certain things.

Go interfaces let you define variables and func parameters that
will hold _any_ type, as long as that type defines certain methods

#### Two different types that have the same methods
```go
package gadget

import "fmt"

type TapPlayer struct {
	Batteries string
}

func (t TapPlayer) Play(song string) {
	fmt.Println("playing", song)
}

func (t TapPlayer) Stop() {
	fmt.Println("stopped!")
}

type TapRecorder struct {
	Microphones int
}

func (r TapRecorder) Record() {
	fmt.Println("Recording")
}
func (r TapRecorder) Play(song string) {
	fmt.Println("playing", song)
}

func (r TapRecorder) Stop() {
	fmt.Println("stopped!")
}

```
```go
package main

import "gadget/gadget"

func main() {
	player := gadget.TapPlayer{}
	mixTape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixTape)

}

//playList takes a TapePlayer value, and a slice of song titles
//to play on it, loop over the songs, when done playing, it stops the TapePlayer
func playList(device gadget.TapPlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

```
the playList function works great with a TapePlayer value.
You might hope it would work with a TapeRecorder as well.
<br> No! will not work<br>

That's too bad...All the playList need is a value whose type
defines Play() and Stop() methods. Both TapePlayer and TapeRecorder have those!<br>

In this case, it does seem like the Go language’s type safety is getting in our
way, rather than helping us. The TapeRecorder type defines all the
methods that the playList function needs, but we’re being blocked from
using it because playList only accepts TapePlayer values.

## Interfaces
* a set of actions you need a type to be able to perform
* a set of methods that certain values are expected to have

In Go, any type that has all the methods listed in an interface
definition is said to **satisfy** that interface

A type that satisfies an interface can be used anywhere that
interface is called for

A type can have methods in addition to those listed in the interface
,but it mustn't be missing any, or it doesn't satisfy that interface.

A variable declared with an interface type can hold any value whose type
satisfies that interface.

### Concrete type, interface types
A concrete type specifies not only what its values can do (what methods
you can call on them), but also what they are: they specify the underlying
type that holds the value’s data.

Interface types don’t describe what a value is: they don’t say what its
underlying type is, or how its data is stored. They only describe what a
value can do: what methods it has.

>Suppose you need to write down a quick note. In your desk drawer, you
have values of several concrete types: Pen, Pencil, and Marker. Each of
these concrete types defines a Write method, so you don’t really care which
type you grab. You just want a WritingInstrument: an interface type that
is satisfied by any concrete type with a Write method.
> ---
> interface type: I need something I can write with. <br>
> concrete types: Pen, Pencils, Marker


#### Assign any type that satisfies the interface
When you have a variable with an interface type, it can hold values of any
type that satisfies the interface.
```go
package main

import "fmt"

func main() {
	var toy NoiseMaker
	toy = Whistle("Toy Canary")
	toy.MakeSound()
}

type NoiseMaker interface {
	MakeSound()
}
type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

```

You can declare function parameters with interface types as well.
(After all, function parameters are really just variables too.)
```go
func main() {
	play(Whistle("Toy Canary"))
}
func play(n NoiseMaker){
	n.MakeSound()
}
```

#### You can only call methods defined as part of the interface
```go
package main

import "fmt"

func main() {
	play(Robot("Counting..."))
}
func play(n NoiseMaker) {
	n.MakeSound()
	//Walk() doesn't belong the NoiseMaker interface
	//n.Walk()
	//will not compile
}

type NoiseMaker interface {
	MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Tic...tic")
}
func (r Robot) Walk() {
	fmt.Println("walking...")
}
```
`n.Walk undefined (type NoiseMaker has no field or metoh Walk)`

####Fixing our playlist func using an interface
```go
package main

import "gadget/gadget"

func main() {
	//player var assigned a Taper type which implements the Player interface
	var player Player = gadget.Taper{}
	mixTape := []string{"Jessie's Girl", "Whip it", "9 to 5"}

	playList(player, mixTape)
	// re-assign/ modify player with Recorder
	player = gadget.Recorder{}
	playList(player, mixTape)

}

//playList takes any Player  value, and a slice of song titles
//to play on it, loop over the songs, when done playing, it stops the TapePlayer
func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

type Player interface {
	Play(string)
	Stop()
}

```

>IMPORTANT!<br>
> If a type declares methods with pointer receivers, then you’ll only
be able to use pointers to that type when assigning to interface
variables.

The toggle method on thw Switch type has to use a pointer
receiver, so it can modify the receiver
```go
package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	var t Toggleable = s
	t.toggle()
	t.toggle()
}
```
`error:Switch does not implment Toggleable
(toggle method has pointer receiver`<br>
When Go decides whether a value satisfies an interface, pointer
methods aren’t included for direct values. But they are included for
pointers. So the solution is to assign a pointer to a Switch to the
Toggleable variable, instead of a direct Switch value:
`var t Toggleable = &s`

---
### Type assertions
We’ve defined a new TryOut function that will let us test the various
methods of our Taper and Recorder types. TryOut has a single
parameter with the Player interface as its type, so that we can pass in either
a Taper or Recorder.

```go
package main

import "assets/gadget"

func main() {
	TryOut(gadget.Recorder{})
}

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player){
	player.Play("test track")
	player.Stop()
	player.Record()
}
```
`player.Record undefined (type Player has no field or method Record)`
<br>We need a way to get the concrete type value (which does have a Record
method) back.
* Your first instinct might be to try a type conversion to convert the Player
  value to a Recorder value. But type conversions aren’t meant for use
  with interface types, so that generates an error. The error message suggests
  trying something else
```go
func TryOut(player Player) {
player.Play("test track")
player.Stop()
recorder:=gadget.Recorder(player)  //bug
recorder.Record()
}

```
`./main.go:17:29: cannot convert player (type Player) to type gadget.Recorder: need type assertion
`
* a "type assertion"? What's that?
<br>When you have a value of a concrete type assigned to a variable with an
  interface type, a type assertion lets you get the concrete type back.
``` 
var noiseMaker NoiseMaker = Robot("Botco Ambler")
var robot Robot = noiseMaker.(Robot)
                  interface value.(asserted type)
```
In plain language, the type assertion above says something like “I know this
variable uses the interface type NoiseMaker, but I’m pretty sure this
NoiseMaker is actually a Robot.”

Once you’ve used a type assertion to get a value of a concrete type back,
you can call methods on it that are defined on that type, but aren’t part of
the interface.
```go
package main

import "fmt"

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep peep")
}
func (r Robot) Walk() {
	fmt.Println("walking...")
}


type NoiseMaker interface {
	MakeSound()
}

func main(){
	//define a var with an interface type 
	//with assigned value of a type that satisfies the interface
	var noiseMaker NoiseMaker=Robot("robotics")
	noiseMaker.MakeSound() //call the method that's part of the interface
	
	//convert back to the concrete type using a type assertion
	var robot =noiseMaker.(Robot) 
	//call a method that's defined on the concrete type(not the interface)
	robot.Walk()
}
```
### Type assertion failures
```go
func main() {
	TryOut(gadget.Recorder{})
	TryOut(gadget.Taper{})
}

func TryOut(player Player) {
	player.Play("test track")
	player.Stop()
	recorder := player.(gadget.Recorder)
	recorder.Record()
}
```
``` 
playing test track
stopped!
Recording
playing test track
stopped!
panic: interface conversion: main.Player is gadget.Taper, not gadget.Recorder

```
Everything compiles successfully, but when we try to run it, we get a
runtime panic! As you might expect, trying to assert that a Taper is
actually a Recorder did not go well. (It’s simply not true, after all.)

###Avoiding panics when type assertions fail
If a type assertion is used in a context that expects only one return value,
and the original type doesn’t match the type in the assertion, the program
will panic at runtime (not when compiling):
```go
var player Player=gadget.Taper{}
recorder :=player.(gadget.Recorder)
```
`panic: interface conversion: player is gadget.Taper, not gadget.Recorder`

if type assertions are used in a context where multiple return
values are expected, a return indicates whether the
assertion was successful or not
```go
var player Player=gadget.Taper{}
recorder, ok :=player.(gadget.Recorder)
if ok{
	recorder.Record()
}else{
	fmt.Println("player wan not a Recorder")
}
```

### Testing Taper and Recorder using type assertions
Let’s see if we can fix our TryOut function for
Taper and Recorder values. Instead of ignoring the second
return value from our type assertion.
```go
package main

import "assets/gadget"

func main() {
	TryOut(gadget.Recorder{})
	TryOut(gadget.Taper{})
}

func TryOut(player Player) {
	player.Play("test track")
	player.Stop()
	recorder, ok := player.(gadget.Recorder)
	if ok {
		recorder.Record()
	}
}

type Player interface {
	Play(string)
	Stop()
}

```

--- 
## The error interface 
```go
type error interface{
	Error () string
}
```
Declaring the error type as an interface means that if has 
an Error method that returns a string, it satisfies the error interface
,and it's an error value. That means you can define your own
type and use them anywhere an error value is required!

```go
package main

import "fmt"

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	var err error
	err = ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)
}
```
If you need an error value, but also need to track more information about
the error than just an error message string, you can create your own type
that satisfies the error interface and stores the information you want

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	err := checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}

}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

//checkTemperature uses OverheatError, it takes
//the system's actual temperature and the temperature
//that's considered safe
func checkTemperature(actual, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

```


>Important Note:<br>
> Q: How is it we’ve been using the error interface type in all these
different packages, without importing it? Its name begins with a
lowercase letter. Doesn’t that mean it’s unexported, from whatever
package it’s declared in? What package is error declared in, anyway?
> 
> A: The error type is a “predeclared identifier,” like int or string. And
so, like other predeclared identifiers, it’s not part of any package. It’s part of
the “universe block,” meaning it’s available everywhere, regardless of what
package you’re in.

---
### The Stringer interface
```go
package main

import "fmt"

type Gallons float64
type Liters float64
type Milliliters float64

func main() {
	fmt.Println(Gallons(12.09248342))//12.09248342
	fmt.Println(Liters(12.09248342)) //	12.09248342
	fmt.Println(Milliliters(12.09248342))//	12.09248342

}

```
they all look the same when printed. If there are
too many decimal places of precision on a value, that looks awkward when
printed, too.

the fmt pkg defines the Stringer interface: to allow
any type to decide how it will be displayed when printed.
just define `String()string` method on any type
```go
package main

import "fmt"

func main() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + "coffee pot"
}

```
after using the Stringer interface
```go
package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f ml", m)

}

func main() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
	//12.09 gal
	//12.09 L
	//12.09 ml
}
```
---

### The empty interface
something's been bothering me. For most of the functions
we've seen so far, we can only call them using values of
specific types. But some fmt functions like Println can
take values of _**any**_ type! How does that work?

Println signature
`func Println(a ...interface{})(n int, err error)`
<br> ... means it's variadic, can take any number of params,
<br>But what's this interface{}?<br>
But what would happen if we declared an interface type that didn’t require
any methods at all? It would be satisfied by any type! It would be satisfied
by all types! is known as the _**empty interface**_ , and it’s used to
accept values of any type. The empty interface doesn’t have any methods
that are required to satisfy it, and so every type satisfies it.

```go
package main

import "fmt"

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("string")
	AcceptAnything(true)

}

func AcceptAnything(any interface{}) {
	fmt.Println(any)

}

```