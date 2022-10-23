package output

import (
	"hello-go/satisfaction-statistics/csvParser"
	"hello-go/satisfaction-statistics/quantityStatistic"
)

type Report struct {
	month                           string
	TotalSendCnt                    int     // 发送量
	LastYearTotalSendCnt            int     // 去年发送量
	IncreaseLastYearTotalSendCnt    float32 // 同比增长发送量
	TotalRecycleCnt                 int     // 回收量
	LastYearTotalRecycleCnt         int     // 去年回收量
	IncreaseLastYearTotalRecycleCnt float32 // 同比增长回收量
	RecycleRate                     float32 // 回收率
	LastYearRecycleRate             float32 // 去年回收率
	IncreaseLastYearRecycleRate     float32 // 同比增长回收率
	AvgSatisfaction                 float32 // 平均满意度
	LastYearAvgSatisfaction         float32 // 去年平均满意度
	IncreaseLastYearAvgSatisfaction float32 // 同比增长平均满意度
	//OptionScore                     int     // 选择分
	//TotalScore                      int     // 总分
	PraisingNurse string // 表扬的护士
	MainProblem   string // 主要问题
}

func NewReport(thisYear quantityStatistic.Quantity, lastYear quantityStatistic.Quantity, records []csvParser.Record) Report {
	report := Report{
		month: thisYear.Time.Format("2006-01"),

		TotalSendCnt:                 thisYear.TotalSendCnt,
		LastYearTotalSendCnt:         lastYear.TotalSendCnt,
		IncreaseLastYearTotalSendCnt: CalculateYearOnYear(thisYear.TotalSendCnt, lastYear.TotalSendCnt),

		TotalRecycleCnt:                 thisYear.TotalRecycleCnt,
		LastYearTotalRecycleCnt:         lastYear.TotalRecycleCnt,
		IncreaseLastYearTotalRecycleCnt: CalculateYearOnYear(thisYear.TotalRecycleCnt, lastYear.TotalRecycleCnt),

		RecycleRate:                 thisYear.RecycleRate,
		LastYearRecycleRate:         lastYear.RecycleRate,
		IncreaseLastYearRecycleRate: CalculateYearOnYear0(thisYear.RecycleRate, lastYear.RecycleRate),

		AvgSatisfaction:                 thisYear.AvgSatisfaction,
		LastYearAvgSatisfaction:         lastYear.AvgSatisfaction,
		IncreaseLastYearAvgSatisfaction: CalculateYearOnYear0(thisYear.AvgSatisfaction, lastYear.AvgSatisfaction),

		PraisingNurse: buildPraisingNurse(records),
		MainProblem:   buildMainProblem(records),
	}

	return report
}

func buildMainProblem(records []csvParser.Record) string {
	return ""
}

func buildPraisingNurse(records []csvParser.Record) string {
	return ""
}

func CalculateYearOnYear(thisYear int, lastYear int) float32 {
	t := float32(thisYear)
	l := float32(lastYear)
	return (t - l) / l
}

func CalculateYearOnYear0(thisYear float32, lastYear float32) float32 {
	return (thisYear - lastYear) / lastYear
}
