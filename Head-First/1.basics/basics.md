#Basics

#### Strings
A string is a series of bytes that usually represent text characters. You can define strings 
directly withing your code using **string literals**: text between double quotation marks
that GO will treat as a string
`"Hello, Go!`

#### Runes
Whereas strings are usually used to represent a whole series of text characters,
**Go's runes** are used to represent single characters. rune literals are written
with single quotation marks.

Go programs can user almost any character from almost any language on earth, 
because Go uses the unicode standard for storing runes,
Runes are kept as numeric codes, not characters themselves, and if you pass
a rune to `fmt.Println('A')`, you will see that numeric code in the output `65` not the 
original character.

just as with string literals, escape sequences can be used in a rune literal to 
represent characters that would be hard to include in program code `'\t'` => `9`

#### Types

> Go is statically typed, which means that it knows what the types of your
values are even before your program runs. Functions expect their arguments
to be of particular types, and their return values have types as well (which
may or may not be the same as the argument types). If you accidentally use
the wrong type of value in the wrong place, Go will give you an error
message. This is a good thing: it lets you find out there’s a problem before
your users do!


#### Declaring Variables
In Go, a variable is a piece of storage containing a value. You can give a
variable a name by using a variable declaration.
``` 
var quantity int
var length, width float64
var customerName string
```
Once you declare a variable, you can assign any value of that type to it with
= (that’s a single equals sign):
```
quantity=3
```

#### Zero values
if you declare a variable without assigning it a value, that variable will
contain the `zero value` for its type 











