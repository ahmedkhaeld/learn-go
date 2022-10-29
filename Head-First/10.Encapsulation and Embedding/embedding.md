# Keep it to yourself: Encapsulation and Embedding
>I heard that Paragraph type of hers stores its data in a simple string
> field!.<br>
> And that fancy Replace method? <br>
> It's just promoted from an embedded string.Replacer!<br>
> You'd never know it from using Paragraph, though!

Mistakes happen. Sometimes, your program will receive invalid data from
user input, a file you’re reading in, or elsewhere. In this chapter, you’ll learn
about encapsulation: a way to protect your struct type’s fields from that
invalid data. That way, you’ll know your field data is safe to work with!

We’ll also show you how to embed other types within your struct type. If
your struct type needs methods that already exist on another type, you don’t
have to copy and paste the method code. You can embed the other type
within your struct type, and then use the embedded type’s methods just as if
they were defined on your own type!

#### Creating a Date struct type
A local startup called **Remind Me** is developing a calendar
application to help users remember birthdays, anniversaries, and more.

* Requirements
    * we need to be able to assign a title to each event
    * along with the year, month, and day it occurs
  
```go
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	date := Date{}
	err := date.SetYear(1995)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(8)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(28)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

}

//Date hold our year, month, and day values
type Date struct {
	Year, Month, Day int
}

//# Data Validation!
//get valid inputs from user to a Date type
//define how should be Date fields provided before accepting it
//year value set to 1 or greater
//month between 1 - 12
//day between   1 - 31
// return an error type if data is invalid with a message
// else set the value to the field and return nil error

//1.set Type fields using Setter methods(need pointer receiver)

//SetYear accepts the year to be set, year < 1 returns err
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

```

---
#### The fields can still be set to invalid values
The validation provided by your setters is great, when user
actually use them, But we've got user setting the struct
fields directly, any they're still entering invalid data!
```go
date :=Date{}
date.Year=2019
date.Month=20
date.Day=40
fmt.Println(date)
```
We need a way to protect these fields, so that users of our Date type can
only update the fields using the setter methods.

* create a pkg 
* un-export the type fields
* export the setters to update/modify fields with validation
* export the getters to retrieve the fields<br>
`date.go`
```go
package calendar

import (
	"errors"
)



//Date hold our year, month, and day values
//un-export Date fields to protect modifying them directly
// to access them use setters
// to retrieve them use getters
type Date struct {
	year, month, day int
}

//1.set Type fields using Setter methods(need pointer receiver)

//SetYear accepts the year to be set, year < 1 returns err
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

//2.getters
//By convention, a getter method’s name should be the same as the name of
//the field or variable it accesses.

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month

}

func (d *Date) Day() int {
	return d.day
}

```

```go
package main

import (
	"fmt"
	"log"
	"reminder/calendar"
)

func main() {
	date := calendar.Date{}
	err:=date.SetYear(2022)
	if err !=nil{
		log.Fatal(err)
	}
	err=date.SetMonth(10)
	if err !=nil{
		log.Fatal(err)
	}
	err=date.SetDay(25)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("whole date:",date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
	
}

```
---
## Encapsulation
The practice of hiding data in one part of a program from code
in another part is known as encapsulation

data ins encapsulated within pkgs using unexported variables, struct fileds,
functions, or methods

Go developers generally only rely on encapsulation when it’s
necessary, such as when field data needs to be validated by setter methods.
In Go, if you don’t see a need to encapsulate a field, it’s generally okay to
export it and allow direct access to it.
>Notes:<br>
> Q: Many other languages don’t allow access to encapsulated values
outside the class where they’re defined. Is it safe for Go to allow
other code in the same package to access unexported fields?<br>
> A: Generally, all the code in a package is the work of a single developer (or
group of developers). All the code in a package generally has a similar
purpose, as well. The authors of code within the same package are most
likely to need access to unexported data, and they’re also likely to only use
that data in valid ways. So, yes, sharing unexported data with the rest of the
package is generally safe.<br>
> Code outside the package is likely to be written by other developers, but
that’s okay because the unexported fields are hidden from them, so they
can’t accidentally change their values to something invalid.

###Embedding the Date type in an Event type
Now we just need to be able to assign title to our events
* **Unexported fields don’t get promoted**

  Embedding a Date in the Event type will not cause the Date fields to be
  promoted to the Event, though. The Date fields are unexported, and Go
  doesn’t promote unexported fields to the enclosing type. That makes sense;
  we made sure the fields were encapsulated so they can only be accessed
  through setter and getter methods, and we don’t want that encapsulation to
  be circumvented through field promotion.
* **Exported methods get promoted just like fields**

If you embed a type with exported methods within a struct type, its methods
will be promoted to the outer type, meaning you can call the methods as if
they were defined on the outer type.

```go
package main

import (
	"fmt"
	"log"
	"reminder/calendar"
)

func main() {

	event := calendar.Event{}
	err := event.SetTitle("My BirthDay!")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}

```
`event.go`
```go
package calendar

import (
	"errors"
	"unicode/utf8"
)

//Event include Date type
type Event struct {
	title string
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	//valid the count of title not gt 30 char
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

```