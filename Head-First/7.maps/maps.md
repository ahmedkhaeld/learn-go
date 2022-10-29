# Labeling data: Maps
>See? you don't have to look through the entire book to find
> the topic you want; just use the index to look up what
> page it's on!

**Throwing things in pile is fine, until you need to find
something again.** 

with arrays and slices to find a value, you'll have to start
at the beginning and look through every single value.

#### Counting votes
* There are two candidates on the ballot, **Amber Graham** and **Brian Martin**
* voters have the option to write in a candidate's name
* text files records the votes, as one vote per line(file for each District)
* process each line of the file and tally the total number of times each name occurs
* the name with most votes wins!

### Maps
is a collection where each value is accessed via a key, Keys
are an easy way to get data back out of your map. <br>
It's like having neatly labeled file folders

Whereas arrays and slices can only user integers as indexes,
a map can use any type for keys(as long as values of that type unique)
the values all have to be of the same type, and the keys all have to be
of the same type`var myMap map[string]float64` 

Just as with slices, declaring a map variable doesn't automatically
create a map; you need to call the make func <br>
`var ranks map[string]int` Declare a map variable<br>
`ranks =make(map[string]int)` actually create the map<br>

`ranks:=make(map[string]int` create a map and declare a variable to hold it

```go
package maps

import "fmt"

func maps(){
	isPrime:=make(map[int]bool)
	isPrime[4]=false
	isPrime[7]=true 
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])
}
```

#### Map literals
Just as with arrays and slices, if you know keys and values that you want
your map to have in advance, you can use a **map literal** to create it.
`myMap:=map[string]float64{"a":1.3, "b":5.6}`<br>

As with slice literals, leaving the curly braces empty creates
a map that starts empty `emptyMap :=map[string]float64{}`

#### Zero values within maps
if you access a map key that hasn't been assigned to,
you'll get a zero value back
```go
package maps

import "fmt"

func zeroValues(){
	numbers:=make(map[string]int)
	numbers["assigned"]=12
	fmt.Printf("%#v \n", numbers["assigned"]) //12
	fmt.Printf("%#v \n", numbers["not assigned"]) //0
}
```
Depending on the value type, the zero will be its zero value

#### The zero value for a map variable is nil
```go
package maps

import "fmt"

func nilMap(){
	var nilMap map[int]string
	fmt.Printf("%#v \n", nilMap) //map[int]string(nil)
	nilMap[3]="three"    //panic: assignment to entry in nil map
}
```
before attempting to add keys and values, create a map using
`make` or `map literal` and assign to it


#### How to tell zero values apart from assigned values
Zero values, although useful, can sometimes make it difficult to tell whether
a given key has been assigned the zero value, or if it has never been
assigned.
```go
package maps

import "fmt"

func status(name string){
	grades :=map[string]float64{"alma":0, "rohit":86.5}
	grade :=grades[name]
	if grade<60{
		fmt.Printf("%s is failing \n", name)
	}
}

func init()  {
	status("alma")  // alma is failing | existing key
	status("carl")  // carl is failing | non-existing key
}
```
To address situations like this, accessing a map key optionally returns a
second, Boolean value. It will be true if the returned value has actually
been assigned to the map, or false if the returned value just represents the
default zero value.

```go
package maps

import "fmt"

func counters(){
	counters :=map[string]int{"a":3, "b":0}
	value, ok:=counters["a"]
	fmt.Println(value,ok) //3 true
	
	value, ok=counters["a"]
	fmt.Println(value,ok)  //0 true
	
	value, ok=counters["a"]
	fmt.Println(value,ok) //0 false
}

```
The second (ok) return value can be used to decide whether you should treat the
value you got from the map as an assigned value that just happens to match
the zero value for that type, or as an unassigned value.
```go
package maps

import "fmt"

func status(name string) {
	grades := map[string]float64{"alma": 0, "rohit": 86.5}
	grade, ok := grades[name]
	if !ok{
		fmt.Printf("No grade recorded for %s. \n", name)
	}else if grade < 60 {
		fmt.Printf("%s is failing \n", name)
	}
}

func init() {
	status("alma") // alma is failing
	status("carl") // No grade recorded for carl.
}

```

#### Removing key/value pairs with the "delete" func
At some point after assigning a value to a key, you may want to remove it
from your map. Go provides the built-in delete function for this purpose

Just pass the delete function two things: the map you want to delete a key
from, and the key you want deleted.
```go
package maps

import "fmt"

func deleteMapKey() {
	var ok bool 
	ranks :=make(map[string]int)
	var rank int 
	ranks["bronze"]=3
	rank,ok=ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)//rank: 3, ok: true
	
	delete(ranks,"bronze")
	rank,ok=ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok) //rank: 0, ok: false
}
```
#### Using for...range loops with maps
`for key, value := range mymap{}` you can omit either key or value if you don't need to process them!

##### The for...range loop handles maps in random order!
The for...range loop processes map keys and values in a random order
because a map is an unordered collection of keys and values. When you use
a for...range loop with a map, you never know what order you’ll get the
map’s contents in!

if you need more consistent
ordering, you’ll need to write the code for that yourself.

a program that always prints the names in
alphabetical order. It does use two separate for loops. The first loops
over each key in the map, ignoring the values, and adds them to a slice of
strings. Then, the slice is passed to the sort package’s Strings function to
sort it alphabetically, in place.
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	orderedMap()
}
func orderedMap() {
	// key ,value => name,grade
	grades := map[string]float64{"alma": 74.2, "ahmed": 60.9, "carl": 48.2}
	var names []string
	// store the keys in slice of strings
	for name := range grades {
		names = append(names, name)
	}
	//sort the slice
	sort.Strings(names)
	//range the sorted slice 
	//print each name, 
	//access the grades map with this name(key) to get its grade(value
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.2f%%\n", name, grades[name])
	}
}

```
