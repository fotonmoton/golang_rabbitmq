package controllers

import (
	"docker/warehouse"
	"docker/warehouse/http/views"
	"net/http"
	"strconv"
)

type WarehouseController struct {
	warehouse *warehouse.Warehouse
}

func NewWarehouseController(warehouse *warehouse.Warehouse) *WarehouseController {
	return &WarehouseController{warehouse}
}

func (w *WarehouseController) ListProducts(res http.ResponseWriter, req *http.Request) {
	views.HtmlProducts(res, w.warehouse.ListProducts())
}

func (w *WarehouseController) NewProduct(res http.ResponseWriter, req *http.Request) {
	views.NewHtmlProduct(res)
}

func (w *WarehouseController) SubmitProduct(res http.ResponseWriter, req *http.Request) {
	Qty, _ := strconv.ParseInt(req.PostFormValue("Qty"), 10, 64)
	Name := req.PostFormValue("Name")
	w.warehouse.AddProduct(warehouse.NewProduct(Name, int(Qty)))
	views.HtmlProducts(res, w.warehouse.ListProducts())
}
