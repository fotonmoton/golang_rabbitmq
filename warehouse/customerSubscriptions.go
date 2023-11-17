package warehouse

// Products in which customer is interested in:
// 1, "blue jeans"
// 2, "gray t-shirt"
type CustomerSubscription struct {
	CustomerId   int
	Customer     Customer
	ProductTitle string
}

func NewCustomerSubscription(customer Customer, product string) CustomerSubscription {
	return CustomerSubscription{
		CustomerId:   customer.ID,
		Customer:     customer,
		ProductTitle: product,
	}
}
