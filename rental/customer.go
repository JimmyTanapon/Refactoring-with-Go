package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) (rcvr Customer) {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}

}
func (rcvr Customer) AddRental(arg Rental) {
	rcvr.rentals = append(rcvr.rentals, arg)
}
func (rcvr Customer) Name() string {
	return rcvr.name
}

func (r Rental) Charge() float64 {
	result := 0.0
	switch r.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if r.DaysRented() > 2 {
			result += float64(r.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if r.DaysRented() > 3 {
			result += float64(r.DaysRented()-3) * 1.5
		}
	}
	return result
}

func (r Rental) GetPoints() int {

	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func getTotalAmount(rentals []Rental) float64 {
	result := 0.0
	for _, r := range rentals {

		result += r.Charge()

	}
	return result
}
func getTotalPoints(rentals []Rental) int {
	result := 0
	for _, r := range rentals {

		result += r.GetPoints()

	}
	return result
}

type Bill struct {
	Customer    Customer
	TotalAmount float64
	MovieRates  []MovieRate
	Points      int
}
type MovieRate struct {
	Title  string
	Amount float64
}

func renderPlenText(b Bill) string {
	result := fmt.Sprintf("Rental Record for %v\n", b.Customer.Name())
	for _, r := range b.MovieRates {

		result += fmt.Sprintf("\t%v\t%.1f\n", r.Title, r.Amount)

	}

	result += fmt.Sprintf("Amount owed is %.1f\n", b.TotalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", b.Points)
	return result
}

func (c Customer) Statement() string {
	movieRates := []MovieRate{}

	// result := fmt.Sprintf("Rental Record for %v\n", c.Name())

	for _, r := range c.rentals {
		mr := MovieRate{
			Title:  r.Movie().Title(),
			Amount: r.Charge(),
		}

		movieRates = append(movieRates, mr)
		// result += fmt.Sprintf("\t%v\t%.1f\n", mr.Title, mr.Amount)

	}
	bill := Bill{
		Customer:    c,
		TotalAmount: getTotalAmount(c.rentals),
		MovieRates:  movieRates,
		Points:      getTotalPoints(c.rentals),
	}

	// result += fmt.Sprintf("Amount owed is %.1f\n", bill.TotalAmount)
	// result += fmt.Sprintf("You earned %v frequent renter points", bill.Points)

	return renderPlenText(bill)
}
