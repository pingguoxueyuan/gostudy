package main

import (
	"time"
	"fmt"
)

func testTime() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day ,hour, minute, send)
	timestamp := now.Unix()
	fmt.Printf("timestamp is:%d\n", timestamp)
}

func testTimestamp(timestamp  int64) {

	timeObj := time.Unix(timestamp, 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()

	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day ,hour, minute, send)
}

func processTask() {
	fmt.Printf("do task\n")
}

func testTicker() {
	ticker := time.Tick(1*time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

func testConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micro second:%d\n", time.Microsecond)
	fmt.Printf("mili second:%d\n", time.Millisecond)
	fmt.Printf("second:%d\n", time.Second)
}

func testFormat(){

	now := time.Now()
	timeStr := now.Format("2006/01/02 15:04:05")
	fmt.Printf("time:%s\n", timeStr)
}

func testFormat2() {
	now := time.Now()
	timeStr := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d\n", 
		now.Year(), now.Month(), now.Day() ,now.Hour(), now.Minute(), now.Second())
	fmt.Printf("time:%s\n", timeStr)
}

func testCost() {
	start := time.Now().UnixNano()
	//for i := 0; i < 10; i++ {
		time.Sleep(10*time.Millisecond)
	//}
	end := time.Now().UnixNano()
	cost := (end - start)/1000
	fmt.Printf("code cost:%d us\n", cost)
}

func main() {
	//testTime()
	//timestamp := time.Now().Unix()
	//testTimestamp(timestamp)
	//testTicker()
	//testConst()

	//testFormat()
	//testFormat2()
	testCost()
}