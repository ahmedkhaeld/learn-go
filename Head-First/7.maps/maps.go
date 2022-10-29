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
