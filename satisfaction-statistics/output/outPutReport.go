package output

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hello-go/satisfaction-statistics/csvParser"
	"hello-go/satisfaction-statistics/quantityStatistic"
	"time"
)

var OrmDB *gorm.DB

func outputReport() {
	allData := getAllData()
	fmt.Println(len(allData))
	export2Excel(allData)
}

func export2Excel(reports []Report) {
	f := excelize.NewFile() // 设置单元格的值
	// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "月份")
	f.SetCellValue("Sheet1", "B1", "发送量")
	f.SetCellValue("Sheet1", "C1", "去年同月发送量")
	f.SetCellValue("Sheet1", "D1", "同比增长发送量")
	f.SetCellValue("Sheet1", "E1", "回收量")
	f.SetCellValue("Sheet1", "F1", "去年同月回收量")
	f.SetCellValue("Sheet1", "G1", "同比增长回收量")
	f.SetCellValue("Sheet1", "H1", "回收率")
	f.SetCellValue("Sheet1", "I1", "去年同月回收率")
	f.SetCellValue("Sheet1", "J1", "同比增长回收率")
	f.SetCellValue("Sheet1", "K1", "平均满意度")
	f.SetCellValue("Sheet1", "L1", "去年同月平均满意度")
	f.SetCellValue("Sheet1", "M1", "同比增长平均满意度")
	f.SetCellValue("Sheet1", "N1", "表扬的护士")
	f.SetCellValue("Sheet1", "O1", "主要问题")

	line := 1

	// 循环写入数据
	for _, v := range reports {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.month)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.TotalSendCnt)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.LastYearTotalSendCnt)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.IncreaseLastYearTotalSendCnt)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", line), v.TotalRecycleCnt)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", line), v.LastYearTotalRecycleCnt)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", line), v.IncreaseLastYearTotalRecycleCnt)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", line), v.RecycleRate)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", line), v.LastYearRecycleRate)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", line), v.IncreaseLastYearRecycleRate)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", line), v.AvgSatisfaction)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", line), v.LastYearAvgSatisfaction)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", line), v.IncreaseLastYearAvgSatisfaction)
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", line), v.PraisingNurse)
		f.SetCellValue("Sheet1", fmt.Sprintf("O%d", line), v.MainProblem)
	}

	// 保存文件
	if err := f.SaveAs("report.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func getAllData() []Report {
	all := make([]Report, 0)
	month, _ := time.Parse("2006-01", "2022-01")
	for month.Before(time.Now()) {
		r := getOneMonthData(month)
		all = append(all, r)
		month = month.AddDate(0, 1, 0)
	}
	return all
}

func getOneMonthData(month time.Time) Report {
	// select quantityStatistic
	var thisYear quantityStatistic.Quantity
	OrmDB.Where("time = ?", month).Limit(1).Find(&thisYear)
	// select last year quantityStatistic
	var lastYear quantityStatistic.Quantity
	OrmDB.Where("time = ?", month.AddDate(-1, 0, 0)).Limit(1).Find(&lastYear)
	// select all detail records
	var records []csvParser.Record
	OrmDB.Where("time = ?", month).Order("seq").Find(&records)
	return NewReport(thisYear, lastYear, records)
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
