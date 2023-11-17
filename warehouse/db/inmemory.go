package db

import (
	"docker/warehouse"
	"log"
	"strings"
)

type InMemoryState struct {
	products              []warehouse.Product
	customerNotifications []warehouse.CustomerNotification
	subscriptions         []warehouse.CustomerSubscription
	customers             []warehouse.Customer
}

var productIdEpoch = 1

func NewInMemoryState() *InMemoryState {
	return &InMemoryState{}
}

func (s *InMemoryState) SaveProduct(p warehouse.Product) warehouse.Product {
	p.ID = productIdEpoch
	productIdEpoch++
	s.products = append(s.products, p)

	return p
}

func (s *InMemoryState) ListProducts() []warehouse.Product {
	return s.products
}

func (s *InMemoryState) NotifyCustomers(notifications []warehouse.CustomerNotification) {
	s.customerNotifications = append(s.customerNotifications, notifications...)
	log.Println(s.customerNotifications)
}

func (s *InMemoryState) FindSubscriptions(productTitle string) []warehouse.CustomerSubscription {
	found := []warehouse.CustomerSubscription{}

	for idx := range s.subscriptions {
		if strings.Contains(s.subscriptions[idx].ProductTitle, productTitle) {
			found = append(found, s.subscriptions[idx])
		}
	}

	return found
}

func (s *InMemoryState) AddSubscriptions(subs []warehouse.CustomerSubscription) {
	s.subscriptions = append(s.subscriptions, subs...)
}

func (s *InMemoryState) AddCustomers(customers []warehouse.Customer) {
	s.customers = append(s.customers, customers...)
}
