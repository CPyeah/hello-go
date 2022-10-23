package csvParser

import (
	"strconv"
	"time"
)

type Record struct {
	Seq                 string
	Name                string
	BedNo               string
	AdmissionNo         string
	WardName            string
	SatisfactionScore   string
	Satisfaction        string
	QuestionnaireStatus string
	SendTime            string
	SubmitTime          string
	PushNurse           string
	PushWay             string
	Q1                  int    // 助产士仪表端庄，语言文明，热情接待。
	Q2                  int    //:助产士主动提供健康指导（如母乳喂养，会阴护理等）
	Q3                  int    //:分娩期间主动与您核对新生儿信息（如性别、身长、体重等）。
	Q4                  int    //:助产士主动告知产程进展情况，缓解宫缩痛的方法并注意保护隐私。
	Q5                  int    //:您对分娩期间助产士们的总体工作是否满意。
	Q6                  string //:您认为服务最好的助产士姓名：
	Q7                  string //:您认为最不满意的助产士姓名：               原因：
	Q8                  string //:您其它的意见和建议：
	Time                time.Time
	Year                int
	Month               int
}

func NewRecode(raw []string, t time.Time) Record {
	return Record{raw[0],
		raw[1],
		raw[2],
		raw[3],
		raw[4],
		raw[5],
		raw[6],
		raw[7],
		raw[8],
		raw[9],
		raw[10],
		raw[11],
		parsePoint(raw[12]),
		parsePoint(raw[13]),
		parsePoint(raw[14]),
		parsePoint(raw[15]),
		parsePoint(raw[16]),
		raw[17],
		raw[18],
		raw[19],
		t,
		t.Year(),
		int(t.Month()),
	}
}

func parsePoint(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			p := string(s[i+1])
			point, _ := strconv.Atoi(p)
			return point
		}
	}
	return -1
}

const (
	Q1 = "1:助产士仪表端庄，语言文明，热情接待。"
	Q2 = "2:助产士主动提供健康指导（如母乳喂养，会阴护理等）。"
	Q3 = "3:分娩期间主动与您核对新生儿信息（如性别、身长、体重等）。"
	Q4 = "4:助产士主动告知产程进展情况，缓解宫缩痛的方法并注意保护隐私。"
	Q5 = "5:您对分娩期间助产士们的总体工作是否满意。"
	Q6 = "6:您认为服务最好的助产士姓名："
	Q7 = "7:您认为最不满意的助产士姓名：   原因："
	Q8 = "8:您其它的意见和建议："
)
