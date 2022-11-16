package main

import (
	"fmt"
	"strings"
)

// KeySet is a set of keys in the game.
type KeySet byte

const (
	Captain KeySet = 1 << iota // 1
	Leader                     // 2
	Striker                    // 4
	maxKey                     // 8
)

func (k KeySet) String() string {
	// check k is under the limit
	if k >= maxKey {
		return fmt.Sprintf("<unknown key: %d", k)

	}

	// if player has single key, which one?
	switch k {
	case Captain:
		return "captain"
	case Leader:
		return "leader"
	case Striker:
		return "striker"

	}

	// if player  key is not only single key, but got more than a key and under the limit
	// example   key is 5 means 4 + 1 means either a captain or striker

	var names []string
	// <<= is a lift shift and assignment operator key = key << 1 mean multiply key by 2
	for key := Captain; key < maxKey; key <<= 1 {
		if k&key != 0 {
			names = append(names, key.String())
		}
	}
	return strings.Join(names, "|")
}

//Player is a player in the game
type Player struct {
	Name string
	Keys KeySet
}

func (p *Player) AddKey(key KeySet) {
	// p.keys = p.keys | key bitwise OR and assignment
	p.Keys |= key
}

// HasKey returns true if player has a key
func (p *Player) HasKey(key KeySet) bool {

	return p.Keys&key != 0
}

//RemoveKey removes key from player
func (p *Player) RemoveKey(key KeySet) {
	// p.Keys AND with NOT key
	p.Keys &= ^key
	fmt.Printf("name is %s, key removed", p.Name)
}

//func main() {
//
//	var p1, p2, p3, p4, p5, p6, p7 Player
//	p1.Name = "hamo"
//	p1.AddKey(1)
//
//	fmt.Printf("name is %s, key is %s\n", p1.Name, p1.Keys)
//	//name is hamo, key is captain
//
//	p2.Name = "hoda"
//	p2.AddKey(2)
//	fmt.Printf("name is %s, key is %s\n", p2.Name, p2.Keys)
//	//name is hoda, key is leader
//
//	p3.Name = "pop"
//	p3.AddKey(3)
//	fmt.Printf("name is %s, key is %s\n", p3.Name, p3.Keys)
//	//name is pop, key is captain|leader
//
//	p4.Name = "ali"
//	p4.AddKey(4)
//	fmt.Printf("name is %s, key is %s\n", p4.Name, p4.Keys)
//	//name is ali, key is striker
//
//	p5.Name = "bebo"
//	p5.AddKey(5)
//	fmt.Printf("name is %s, key is %s\n", p5.Name, p5.Keys)
//	//name is bebo, key is captain|striker
//
//	p6.Name = "adel"
//	p6.AddKey(6)
//	fmt.Printf("name is %s, key is %s\n", p6.Name, p6.Keys)
//	//name is adel, key is leader|striker
//
//	p7.Name = "mom"
//	p7.AddKey(7)
//	fmt.Printf("name is %s, key is %s\n", p7.Name, p7.Keys)
//	//name is mom, key is captain|leader|striker
//
//}
