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
