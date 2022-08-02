# date-time

This go project shows how to manage RFC3339 date/time objects. See [the spec](https://datatracker.ietf.org/doc/html/rfc3339) for details. An RFC3339 string looks like the following for the date and time when I first created this readme:

`2022-08-02T07:18:07.00Z`

Or 11:07am on Tuesday, August 2, 2022, since I live near Seattle (GMT-7).

You can get that value by entering:

`go run main.go -y 2022 -mo 8 -d 2 -h 18 -mi 7 -s 0 -ms 0`

The `go test` unit test uses these values.

You can contact me at snoopdougiedoug@gmail.com
