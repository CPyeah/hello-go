package output

import (
	"testing"
	"time"
)

func Test_getOneMonthData(t *testing.T) {
	month, _ := time.Parse("2006-01", "2022-01")
	getOneMonthData(month)
}

func Test_outputReport(t *testing.T) {
	outputReport()
}
