package db

import (
	"docker/warehouse"
	"log"

	nativeMysql "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormState struct {
	db *gorm.DB
}

func NewGormConnection() *gorm.DB {
	cfg := nativeMysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		// Addr:   "127.0.0.1:8083",
		Addr:   "mysql:3306",
		DBName: "warehouse",
	}
	gorm, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return gorm
}

func (g *GormState) ListProducts() []warehouse.Product {
	products := []warehouse.Product{}
	trx := g.db.Find(&products)

	if trx.Error != nil {
		log.Println(trx.Error)
		return nil
	}
	return products
}

func (g *GormState) SaveProduct(p warehouse.Product) warehouse.Product {
	saved := g.db.Save(&p)
	if saved.Error != nil {
		log.Println(saved.Error)
		return p
	}

	return p
}

func (g *GormState) NotifyCustomers(notifications []warehouse.CustomerNotification) {
	trx := g.db.Save(notifications)
	if trx.Error != nil {
		log.Println(trx.Error)
	}
}

func (g *GormState) FindSubscriptions(productTitle string) []warehouse.CustomerSubscription {
	found := []warehouse.CustomerSubscription{}

	trx := g.db.Where("product_title LIKE ?", productTitle).Preload("Customer").Find(&found)

	if trx.Error != nil {
		log.Println(trx.Error)
	}

	return found
}

func (g *GormState) AddSubscriptions(subs []warehouse.CustomerSubscription) {
	trx := g.db.Save(subs)
	if trx.Error != nil {
		log.Println(trx.Error)
	}
}

func (g *GormState) AddCustomers(customers []warehouse.Customer) {
	trx := g.db.Create(customers)
	if trx.Error != nil {
		log.Println(trx.Error)
	}
}

func NewGormState(db *gorm.DB) *GormState {

	db.AutoMigrate(&warehouse.Product{})
	db.AutoMigrate(&warehouse.Customer{})
	db.AutoMigrate(&warehouse.CustomerSubscription{})
	db.AutoMigrate(&warehouse.CustomerNotification{})

	return &GormState{db}
}
