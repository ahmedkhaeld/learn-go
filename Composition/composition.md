# Composition
struct
composition
and that lets us build up structs from
other structs

### composition
the fields and methods of an embedded struct are promoted to the level of the embedding structure.

```go
package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair   // embedded struct
	Length int
}

func main() {
	p1 := PairWithLength{Pair{"/usr", "0xfdfe"}, 121}

	//the fields of Pair promoted to p1 to access directly  
	fmt.Println(p1.Path, p1.Length)

}
///usr 121
```

---
* promoted fields and methods 
```go
package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair   // embedded struct
	Length int
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xdead"}, 121}

	fmt.Println(p)
	fmt.Println(pl)

}
//Hash of /usr is 0xfdfe
//Hash of /usr/lib is 0xdead
```
* **when I go to print pl by itself what do I get?** <br>
pl get access to the String() method from the Pair, this prints the path and hash of pl <br>
  so the fields of `Pair` get promoted to `PairWithLength` so do the methods. <br>this String() method
of Pair does not know anything about the length field, because it did not have this field.

---
* what happens if I create a new String method for pl

```go
package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair   // embedded struct
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xdead"}, 121}

	fmt.Println(p)
	fmt.Println(pl)
	//Hash of /usr is 0xfdfe
	//Hash of /usr/lib is 0xdead; length 121

}

```
now `pl` uses its String method<br>
when it goes to resolve
these methods it says
okay is there a string method on `PairWithLength` well if there's one defined
for
the `PairWithLength` type specifically
it's going to use that, if not it's going to look to see if
there was a string method promoted into
`PairWithLength`

---
* this demo show `PairWithLength` is not a subclass of `Pair`

filename function takes in a Pair type
```go
package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair   // embedded struct
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func Filename(p Pair) string {
	return filepath.Base(p.Path)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xdead"}, 121}

	fmt.Println(p)
	fmt.Println(pl)

	fmt.Println(Filename(p))
	
	// this gets an error
	fmt.Println(Filename(pl)) 
	// cannot use pl (type PairWithLength) as type Pair in argument to Filename

}

```

but we could do this 
```go
fmt.Println(Filename(pl.Pair))
```
---
```go
package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}
func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type PairWithLength struct {
	Pair   // embedded struct
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

type Filenamer interface {
	Filename() string
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	
	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xdead"}, 121}

	fmt.Println(p)
	fmt.Println(p.Filename())

	fmt.Println(fn)
	fmt.Println(fn.Filename())

}

```
 * this assignment works <br> `
   var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xdead"}, 121}
   ` How? <br>
 because PairWithLength is a filenamer
 as the filename method on Pair promoted to PairWithLength
 Pair implements the interface method

---
* sortable interface

```go
package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

// Organs type is a slice of organ type
// has two methods Len and Swap
type Organs []Organ

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//ByName type embedding Organs, so ByName has the fields and methods of Organs
// beside them, it has Less method,
//this make ByName implements the Interface interface three methods
// ByName becomes  interface type
type ByName struct {
	Organs
}

// Less compares names
func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

//ByWeight type embedding Organs, so ByWeight has the fields and methods of Organs
// beside them, it has Less method,
//this make ByWeight implements the Interface interface three methods
// ByWeight becomes  interface type
type ByWeight struct {
	Organs
}

// Less compares weights
func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {

	s := []Organ{
		{"brain", 1235},
		{"liver", 1494}, {"spleen", 162},
		{"pancreas", 131},
		{"heart", 290},
	}

	fmt.Println("original:", s)

	// Sort takes as an input Interface type
	// that's why we made the ByWeight and ByName compatible with the Interface


	// customize Sort to sort by weight according the promoted method interface provided
	sort.Sort(ByWeight{s})
	fmt.Println("by Weight:-", s)

	// customize Sort to sort by name according the promoted method interface provided
	sort.Sort(ByName{s})
	fmt.Println("bt Name:-", s)

}


```