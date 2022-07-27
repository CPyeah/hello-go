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

var db *sql.DB

func init() {
	var linkUrl = fmt.Sprintf("%s:%s@tcp(%s)/delivery_index?charset=utf8&parseTime=True", user, password, url)
	db, _ = sql.Open("mysql", linkUrl)
	fmt.Println("db database connect success.")
}

func main() {
	defer closeMySQL()

	queryCustomerById(1)

	queryCustomers()

	var id = insertCustomer("tom")
	queryCustomerById(id)

	updateCustomer(id, "jerry")
	queryCustomerById(id)

	deleteByName("jerry")
	queryCustomers()

	transaction()
	queryCustomers()
}

func transaction() {
	var tx, err = db.Begin()
	if err != nil {
		panic(err)
	}
	var _, err1 = tx.Exec("insert into customer values (null, ?)", "Tom")
	var _, err2 = tx.Exec("insert into customer values (null, ?)", "Jerry")

	if err1 != nil || err2 != nil {
		var _ = tx.Rollback()
	} else {
		var _ = tx.Commit()
	}
}

func deleteByName(name string) {
	var result, err = db.Exec("delete from customer where name = ?", name)
	if err != nil {
		panic(err)
	}
	var count, _ = result.RowsAffected()
	fmt.Println("deleted", count, "line data.")
}

func updateCustomer(id int64, newName string) {
	var result, err = db.Exec("update customer set name = ? where id = ?", newName, id)
	if err != nil {
		panic(err)
	}
	var count, _ = result.RowsAffected()
	fmt.Println("updated", count, "line data")
}

func insertCustomer(name string) int64 {
	result, err := db.Exec("insert into customer values (null, ?)", name)
	if err != nil {
		panic(err)
	}
	var effectRowsCount, _ = result.RowsAffected()
	fmt.Println("inserted", effectRowsCount, "line data")
	var id, _ = result.LastInsertId()
	return id
}

func queryCustomers() {
	var rows, err = db.Query("select * from customer limit 10")
	if err != nil {
		panic(err)
	}
	var customers []customer
	for rows.Next() {
		var customer customer
		var _ = rows.Scan(&customer.id, &customer.name)
		customers = append(customers, customer)
	}
	fmt.Println("customers list is", customers)
}

func queryCustomerById(id int64) {
	var row = db.QueryRow("select * from customer where id = ?", id)
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
			fmt.Println("db close error.")
		}
		fmt.Println("db closed.")
	}(db)
}
