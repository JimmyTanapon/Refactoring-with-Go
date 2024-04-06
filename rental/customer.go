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

func RegularCharge(r Rental) float64 {
	result := 2.0
	if r.DaysRented() > 2 {
		result += float64(r.DaysRented()-2) * 1.5
	}
	return result
}
func NewReleaseCharge(r Rental) float64 {
	return float64(r.DaysRented()) * 3.0
}
func ChildrensCharge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

func (r Rental) Charge() float64 {
	switch r.Movie().PriceCode() {
	case REGULAR:
		return r.Movie().Charger.Charge(r.DaysRented())
	case NEW_RELEASE:
		return r.Movie().Charger.Charge(r.DaysRented())
	case CHILDRENS:
		return r.Movie().Charger.Charge(r.DaysRented())
	case 0:
		return r.Movie().Charger.Charge(r.DaysRented())
	}
	return 0

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

	for _, r := range c.rentals {
		mr := MovieRate{
			Title:  r.Movie().Title(),
			Amount: r.Charge(),
		}

		movieRates = append(movieRates, mr)

	}
	bill := Bill{
		Customer:    c,
		TotalAmount: getTotalAmount(c.rentals),
		MovieRates:  movieRates,
		Points:      getTotalPoints(c.rentals),
	}

	return renderPlenText(bill)
}
