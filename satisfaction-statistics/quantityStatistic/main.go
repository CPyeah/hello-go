package quantityStatistic

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var OrmDB *gorm.DB

const url = "https://statistics.317hu.com/data-statistics-web/datastatistics/v4.1.6/satSendRecycle/getsatSendRecycle?pushHospitalId=332&pushDeptId=&pushWardId=&summaryType=1&summaryYear=%s&summaryNumber=%s&paperId=117476&questionIds="

func collectAll() {
	month, _ := time.Parse("2006-01", "2021-01")
	for month.Before(time.Now()) {
		collectAMonth(month)
		month = month.AddDate(0, 1, 0)
	}
}

func collectAMonth(time time.Time) {
	m := fmt.Sprintf("%02d", int(time.Month()))
	u := fmt.Sprintf(url, strconv.Itoa(time.Year()), m)
	resp, _ := http.Get(u)
	body, _ := ioutil.ReadAll(resp.Body)
	//Convert the body to type string
	sb := string(body)
	q := NewQuantity(sb, time)
	fmt.Println(q)
	saveToDB(q)
}

func saveToDB(q Quantity) {
	//var _ = OrmDB.AutoMigrate(&Quantity{}) // create table

	// insert date
	OrmDB.Create(q)
}

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
