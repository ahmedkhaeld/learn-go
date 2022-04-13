package main

import "fmt"

func do(m1 *map[int]int) {
	// add a new key to the map
	(*m1)[4] = 400

	// make a new map descriptor
	*m1 = make(map[int]int)
	// here m1 as it has control on the passed descriptor
	// this step change the descriptor direction to point to the new created map
	(*m1)[4] = 4

	fmt.Println("m1", *m1) // m1 map[4:4]

}

func main() {
	// declare an map
	m := map[int]int{1: 100, 2: 200, 3: 300}
	fmt.Println("m", m) // m map[1:100 2:200 3:300]

	// when calling do() we pass the reference
	// which inside do() the reference get to point to different map
	// and will have the value in the new map only
	do(&m)
	fmt.Println("m", m) //m map[4:4]

}
