package main

import (
	"fmt"
	"time"
)

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

type RTime struct {
	Hour   int
	Minute int
	Second int
}

type Option func(*RTime)

func TimeBefore(hour int, min int, sec int) string {
	h := HourBefore(hour)
	m := MinBefore(min)
	s := SecBefore(sec)
	ret := timeBefore(h, m, s).End()

	return ret
}

func timeBefore(opt ...Option) *RTime {

	t := new(RTime)
	for _, op := range opt {
		op(t)
	}

	return t
}

func (t *RTime) End() string {
	now := time.Now()
	t.Second += (t.Hour*60 + t.Minute) * 60
	before := now.Add(time.Duration(-t.Second) * time.Second).Format("2006-01-02 15:04:05")
	return before
}

func HourBefore(h int) Option {
	return func(t *RTime) {
		t.Hour += h
		//t.Minute += 60 * t.Hour
	}
}

func MinBefore(m int) Option {
	return func(t *RTime) {
		t.Minute += m
		//t.Second += 60 * t.Minute
	}
}

func SecBefore(s int) Option {
	return func(t *RTime) {
		t.Second += s
	}
}

func FormatTime(formatType string) string {
	return time.Now().Format(formatType)
}

func main() {

	fmt.Println("++++++++++", TimeBefore(1, 10, 0))

	//获取当前时间戳
	now := time.Now().Unix()
	fmt.Println("当前时间戳:", now)

	//***********************格式化当前时间***********************/
	nowFormat := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(nowFormat)

	minFormat := time.Now().Format("2006-01-02 15:04")
	fmt.Println(minFormat)
	/************************************************************/

	/***********************使用时间常量*************************/
	//需要什么格式，就使用什么格式的时间常量
	cnow := time.Now().Format(RFC1123)
	fmt.Println(cnow)
	fmt.Println(FormatTime("2006-01-02 15:04"))
	/***********************get Previous time***********************/
	//10min ago
	now1 := time.Now()
	tenMinsBefore := now1.Add(-10 * time.Minute).Format("2006-01-02 15:04:05")
	fmt.Println(tenMinsBefore)
}
