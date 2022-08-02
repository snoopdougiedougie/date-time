package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

/*
2022-08-02T18:07:00.00Z

Or 11:07am on Tuesday, August 2, 2022, since I live near Seattle (GMT-7)
*/

type datetime struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
	msec   int
	//	offset int
}

func newDateTime(year int, month int, day int, hour int, minute int, second int, msec int /*, offset int*/) *datetime {
	dt := datetime{year: year}
	dt.month = month
	dt.day = day
	dt.hour = hour
	dt.minute = minute
	dt.second = second
	dt.msec = msec
	// dt.offset = offset
	return &dt
}

func validateInt(value int, min int, max int, errMsg string) {
	if value < min || value > max {
		fmt.Println(errMsg)
		os.Exit(1)
	}
}

func myIntToString(i int, minDigits int) string {
	s := strconv.Itoa(i)

	for len(s) < minDigits {
		s = "0" + s
	}

	return s
}

// 2022-08-02T18:07:00.00Z
func getDateTimeAsString(dt datetime) string {
	str := myIntToString(dt.year, 4) + "-" +
		myIntToString(dt.month, 2) + "-" +
		myIntToString(dt.day, 2) + "T" +
		myIntToString(dt.hour, 2) + ":" +
		myIntToString(dt.minute, 2) + ":" +
		myIntToString(dt.second, 2) + ":" +
		myIntToString(dt.msec, 2) + "Z"

	return str
}

/* go run main.go -y 2022 -mo 8 -d 2 -h 18 -mi 7 -s 0 -ms 0
   should result in:
   2022-08-02T18:07:00:00Z
*/

func main() {
	yearPtr := flag.Int("y", -1, "The year")
	monthPtr := flag.Int("mo", -1, "The month")
	dayPtr := flag.Int("d", -1, "The day")
	hourPtr := flag.Int("h", -1, "The hour")
	minutePtr := flag.Int("mi", -1, "The minute")
	secondPtr := flag.Int("s", -1, "The second")
	msecPtr := flag.Int("ms", 0, "The millisecond")

	flag.Parse()

	validateInt(*yearPtr, 0, 2100, "The year is out of range (0-2100)")
	validateInt(*monthPtr, 1, 12, "The month is out of range (1-12)")
	validateInt(*dayPtr, 1, 31, "The day is out of range (1-31)")
	validateInt(*hourPtr, 0, 24, "The hour is out of range (0-24)")
	validateInt(*minutePtr, 0, 59, "The minute is out of range (0-59)")
	validateInt(*secondPtr, 0, 59, "The second is out of range _0-59)")
	validateInt(*msecPtr, 0, 99, "The millisecond is out of range (0, 99)")

	now := newDateTime(*yearPtr, *monthPtr, *dayPtr, *hourPtr, *minutePtr, *secondPtr, *msecPtr)

	str := getDateTimeAsString(*now)

	fmt.Println(str)
}
