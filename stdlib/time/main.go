package main

import (
	"fmt"
	"time"
)

func main() {
	durationOperation()
	locationOperation()
	timeOperation()
	tickerOperation()
	timerOperation()

}

func timerOperation() {
	var timer = time.NewTimer(time.Second)

	select {
	case t := <-timer.C:
		fmt.Println("time is", t)
	}

	time.AfterFunc(time.Second, func() {
		fmt.Println("delay one second function")
	})

	time.Sleep(time.Second * 2)
}

func tickerOperation() {
	var tic = time.NewTicker(time.Second)
	var after5seconds = time.Now().Add(time.Second * 5)

ticLab:
	for true {
		select {
		case t := <-tic.C:
			if time.Now().After(after5seconds) {
				break ticLab
			}
			fmt.Println(t)
		}
	}

	tic.Reset(time.Hour)
	tic.Stop()

}

func timeOperation() {
	var now = time.Now()
	var thousandDay = time.Hour * 24 * 1000
	var fallInLoveTime, _ = time.Parse("2006-01-02 15:04:05", "2019-10-13 20:00:00")
	fmt.Println(now, thousandDay, fallInLoveTime)

	fmt.Println("thousand day:", fallInLoveTime.Add(thousandDay))
	fmt.Println("still now:", now.Sub(fallInLoveTime).Hours()/24)
	fmt.Println("now is", now.Format("2006-01-02 15:04:05"))

}

func locationOperation() {
	var l, err = time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)

	l, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
}

func durationOperation() {
	var d = time.Second * 1000
	fmt.Println(d)
	time.Sleep(time.Second)

	var d1, err = time.ParseDuration("5m40s")
	if err != nil {
		panic(err)
	}
	fmt.Println(d1)

	var t1, err2 = time.Parse("2006年1月2日", "2019年10月13日")
	if err2 != nil {
		panic(err)
	}
	fmt.Println(t1)

	fmt.Println(time.Since(t1).Hours())
}
