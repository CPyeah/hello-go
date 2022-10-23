package csvParser

import (
	"encoding/csv"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

const fileRootPath string = "/Users/chengpeng/zhujing/环比同比统计/"

var OrmDB *gorm.DB

func main() {
	month, _ := time.Parse("2006-01", "2021-01")
	for month.Before(time.Now()) {
		filePath := fileRootPath + month.Format("2006-01") + ".csv"
		parseFile(filePath, month)
		month = month.AddDate(0, 1, 0)
	}
}

func parseFile(filePath string, month time.Time) {
	rawRecords := readCsvFile(filePath)
	records := make([]Record, 0)
	for i := 1; i < len(rawRecords); i++ {
		r := NewRecode(rawRecords[i], month)
		records = append(records, r)
	}
	fmt.Println(len(records))
	save2DB(records, month)
}

func save2DB(records []Record, month time.Time) {
	//createTable()

	// delete history data
	//var s = "delete from records where year = ? and month = ?"
	OrmDB.Where("year = ? and month = ?", month.Year(), int(month.Month())).Delete(Record{})

	// insert new records
	OrmDB.CreateInBatches(records, 30)
}

func createTable() {
	var _ = OrmDB.AutoMigrate(&Record{})
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	//csvReader := csv.NewReader(transform.NewReader(f, simplifiedchinese.GBK.NewDecoder()))
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
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
