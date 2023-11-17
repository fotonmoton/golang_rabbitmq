package warehouse

import "slices"

type WarehouseState interface {
	SaveProduct(Product) Product
	ListProducts() []Product
}

type Warehouse struct {
	state       WarehouseState
	subscribers []Observer
}

func NewWarehouse(state WarehouseState) *Warehouse {
	return &Warehouse{state: state}
}

func (w *Warehouse) AddProduct(p Product) {
	w.Notify(w.state.SaveProduct(p))
}

func (w *Warehouse) ListProducts() []Product {
	return w.state.ListProducts()
}

func (w *Warehouse) Register(listener Observer) {
	w.subscribers = append(w.subscribers, listener)
}

// Not sure it works
func (w *Warehouse) Deregister(listener Observer) {
	idx := slices.IndexFunc(w.subscribers, func(o Observer) bool { return o == listener })
	w.subscribers = slices.Delete(w.subscribers, idx, idx)
}

func (w *Warehouse) Notify(subject any) {
	for _, listener := range w.subscribers {
		listener.Observe(subject)
	}
}
