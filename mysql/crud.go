package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	url      = "localhost:3306"
	user     = "root"
	password = "oracle"
)

var mysql *sql.DB

func init() {
	var linkUrl = fmt.Sprintf("%s:%s@tcp(%s)/delivery_index?charset=utf8&parseTime=True", user, password, url)
	var db, err = sql.Open("mysql", linkUrl)
	if err != nil {
		panic(err)
	}
	mysql = db
	fmt.Println("mysql database connect success.")
}

func main() {
	defer closeMySQL()

	queryCustomerById(1)
}

func queryCustomerById(id int) {
	var row = mysql.QueryRow("select * from customer where id = ?", id)
	var customer customer
	var err = row.Scan(&customer.id, &customer.name)
	if err == sql.ErrNoRows {
		fmt.Println("result is null")
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("result is", customer)
}

func closeMySQL() {
	func(mysql *sql.DB) {
		err := mysql.Close()
		if err != nil {
			fmt.Println("mysql close error.")
		}
		fmt.Println("mysql closed.")
	}(mysql)
}
