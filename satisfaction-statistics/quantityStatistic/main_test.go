package quantityStatistic

import (
	"testing"
	"time"
)

func Test_collectAMonth(t *testing.T) {
	month, _ := time.Parse("2006-01", "2021-01")
	collectAMonth(month)

}

func Test_collectAll(t *testing.T) {
	collectAll()
}
