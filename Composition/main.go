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
