package db

import (
	"database/sql"
	"docker/warehouse"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MysqlState struct {
	db *sql.DB
}

func NewMysqlConnection() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:8083",
		DBName: "warehouse",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(1)

	return db
}

func NewMysqlState() *MysqlState {
	return &MysqlState{
		NewMysqlConnection(),
	}
}

func (w *MysqlState) List() []warehouse.Product {
	products := []warehouse.Product{}
	rows, err := w.db.Query("SELECT * FROM products")

	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var p warehouse.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Qty); err != nil {
			log.Println(err)

			return nil
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)

		return nil
	}

	return products
}

func (w *MysqlState) Save(p warehouse.Product) int {
	result, err := w.db.Exec("INSERT INTO products (name, qty) VALUES (?, ?)", p.Name, p.Qty)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return int(id)
}
