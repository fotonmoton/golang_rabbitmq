package warehouse

type Customer struct {
	ID   int
	Name string
}

var customerIncrement = 1

func NewCustomer(name string) Customer {
	return Customer{
		ID:   customerIncrement,
		Name: name,
	}
}
