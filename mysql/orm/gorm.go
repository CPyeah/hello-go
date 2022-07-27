package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var OrmDB *gorm.DB

func init() {
	const (
		url      = "localhost:3306"
		user     = "root"
		password = "oracle"
	)
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/delivery_index?charset=utf8mb4&parseTime=True&loc=Local", user, password, url)
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	OrmDB = db

}

func main() {

	createUserTable()

	user := createUser()
	fmt.Println(user)
}

func createUserTable() {
	var _ = OrmDB.AutoMigrate(&User{})
}

func createUser() *User {
	email := "tom@cat.com"
	user := User{
		Name:         "tom",
		Email:        &email,
		Age:          7,
		Birthday:     nil,
		MemberNumber: sql.NullString{String: "12121", Valid: true},
		ActivatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
	}
	result := OrmDB.Create(&user)
	fmt.Println("inserted", result.RowsAffected, "line data.")
	return &user
}
