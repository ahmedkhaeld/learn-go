//Package magazine
package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

//Employee track the Names and salaries of our employees
type Employee struct {
	Name   string
	Salary float64
	Address
}

//Address store the mailing addresses for both sub and emp
type Address struct {
	Street, City, State, PostCode string
}

//DefaultSubscriber return a pointer to a Subscriber
func DefaultSubscriber(Name string) *Subscriber {
	var s Subscriber
	s.Name = Name
	s.Rate = 5.99
	s.Active = true
	return &s

}

//PrintInfo takes a pointer to Subscriber
func PrintInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly Rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

//ApplyDiscount update the Rate fo a Subscriber
//it takes a pointer to Subscriber
func ApplyDiscount(s *Subscriber) {
	s.Rate = 4.99
}
