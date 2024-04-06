package rental

const (
	_ = iota
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Pricer interface {
	Charge(daysRented int) float64
	PriceCode() int
}

type Childrens struct {
	priceCode int
}
type NewRelease struct {
	priceCode int
}
type Regular struct {
	priceCode int
}

func CreateChildrens() Childrens {
	return Childrens{
		priceCode: CHILDRENS,
	}
}

func CreateNewRelease() NewRelease {
	return NewRelease{
		priceCode: NEW_RELEASE,
	}
}

func CreateRegular() Regular {
	return Regular{
		priceCode: REGULAR,
	}
}

func (c Regular) PriceCode() int {
	return c.priceCode
}

func (c Childrens) PriceCode() int {
	return c.priceCode
}

func (c NewRelease) PriceCode() int {
	return c.priceCode
}

func (c Childrens) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

func (c Regular) Charge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}

func (c NewRelease) Charge(daysRented int) float64 {
	return float64(daysRented) * 3.0

}

type Movie struct {
	title     string
	priceCode int
	Price     Pricer
}

func NewM(title string, charger Pricer) (m Movie) {
	return Movie{
		title:     title,
		priceCode: charger.PriceCode(),
		Price:     charger,
	}

}

func NewMovie(title string, priceCode int) (m Movie) {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}

}
func (m Movie) PriceCode() int {
	return m.priceCode
}
func (m Movie) Title() string {
	return m.title
}
