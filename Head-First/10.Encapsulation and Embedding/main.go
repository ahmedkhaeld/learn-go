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
