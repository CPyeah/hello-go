package quantityStatistic

import (
	"encoding/json"
	"time"
)

type Quantity struct {
	TotalSendCnt    int     // 发送量
	TotalRecycleCnt int     // 回收量
	RecycleRate     float32 // 回收率
	AvgSatisfaction float32 // 平均满意度
	OptionScore     int     // 选择分
	TotalScore      int     // 总分
	Time            time.Time
	Year            int
	Month           int
}

type Resp struct {
	Data    Quantity
	Success bool
	ErrMsg  string
}

// {"data":{"avgSatisfaction":0.9929,"createId":0,"optionScore":2110,"recycleRate":0.3058,"totalRecycleCnt":85,"totalScore":2125,"totalSendCnt":278,"updateId":0},"errMsg":"查询成功!","success":true}
func NewQuantity(s string, time time.Time) Quantity {
	var resp Resp
	err := json.Unmarshal([]byte(s), &resp)
	if err != nil {
		panic(err)
	}
	q := resp.Data
	q.Time = time
	q.Year = time.Year()
	q.Month = int(time.Month())
	return q
}
