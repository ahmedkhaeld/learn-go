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
