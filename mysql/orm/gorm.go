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

// doc: https://gorm.io/
func main() {

	createUserTable()

	user := createUser()
	fmt.Println(user)

	user = updateUser(user)

	var allUser []User
	OrmDB.Find(&allUser)
	fmt.Println(allUser)

	selectSmall()

	selectWhere()

	count()

	deleteById(user.ID)

	count()

	nativeSql()

}

func nativeSql() {
	type result struct {
		Id   uint
		Name string
		Age  uint8
	}
	var s = "select id, name, age from users where name = ?"
	var results []result
	OrmDB.Raw(s, "tom").Scan(&results)
	fmt.Println("results is", results)
}

func deleteById(id uint) {
	OrmDB.Delete(&User{}, id)
}

func count() {
	var count int64
	OrmDB.Model(&User{}).Count(&count)
	fmt.Println("count is", count)
}

func selectWhere() {
	var users []User
	OrmDB.Where("name = ?", "tom").Order("id desc").Limit(3).Find(&users)
	fmt.Println(users)
}

func selectSmall() {
	type UserAPI struct {
		ID   uint
		Name string
	}
	var userAPIs []UserAPI
	OrmDB.Model(&User{}).Limit(10).Find(&userAPIs)
	fmt.Println(userAPIs)
}

func updateUser(user *User) *User {
	user.Name = "jerry"
	user.Age = 2
	e := "jerry@m.com"
	user.Email = &e
	OrmDB.Save(user)

	return user
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
