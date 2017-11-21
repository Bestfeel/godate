package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const USAGEFORMAT = `
Usage: 日期格式化

月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
周几 Mon,Monday
时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
时区字母缩写 MST

godate -d="2017-11-10 18:22:34"  -f="2006-01-02 15:04:05"
godate -d="2017-11-10 18:22:34"  -f="2006-01-02 15:04:05" -s=1510338154
godate -s=1510338154

`

const (
	COLOR_RED = uint8(iota + 91)
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA //洋红

)

const FMT = "2006-01-02 15:04:05"

var (
	date   = flag.String("d", "", "字符串日期格式【2017-11-10 18:22:34】")
	second = flag.Int64("s", 0, "timstamp 时间格式【1510309374】")
	format = flag.String("f", "", "日期格式化[2006-01-02 15:04:05]")
)

func now() {

	ts := time.Now()

	fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_MAGENTA, "当前日期:"+ts.Format(FMT)+"\n")
	fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_BLUE, "当前时间秒:"+strconv.FormatInt(ts.Unix(), 10)+"\n")

}

func print(dateStr string, s int64, fmts string) {

	if len(dateStr) != 0 && len(fmts) != 0 {
		parse, err := time.Parse(fmts, dateStr)
		if err != nil {
			log.Fatal(err)
			now()
		}
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_RED, "格式化时间:"+strconv.FormatInt(parse.Unix(), 10)+"\n")

	} else if len(dateStr) != 0 && len(fmts) == 0 {

		parse, err := time.Parse(FMT, dateStr)
		if err != nil {
			log.Fatal(err)
			now()
		}
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_RED, "格式化时间:"+strconv.FormatInt(parse.Unix(), 10)+"\n")

	}

	if s != 0 && len(fmts) != 0 {
		tm := time.Unix(s, 0)
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_RED, "格式化时间:"+tm.Format(fmts)+"\n")
	} else if s != 0 && len(fmts) == 0 {
		tm := time.Unix(s, 0)
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_RED, "格式化时间:"+tm.Format(FMT)+"\n")
	}

}
func main() {

	flag.Usage = func() {
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_MAGENTA, USAGEFORMAT+"\n")
		flag.PrintDefaults()
	}
	args := os.Args[1:]
	if len(args) == 0 {
		now()
		return
	}
	flag.Parse()

	print(*date, *second, *format)

}
