package main

import (
	"strconv"
	"time"
)

func calendar(bookings []Booking, in ...time.Month) []Month {
	res := make([]Month, 0)
	for _, m := range in {
		res = append(
			res,
			Month{
				Name: monthNames[m],
				Days: month(m, bookings),
			},
		)
	}
	return res
}

type Month struct {
	Name string
	Days []Day
}

var monthNames = map[time.Month]string{
	time.March: "März",
	time.April: "April",
	time.May:   "Mai",
}

type Day struct {
	Content string
	Classes []string
}

type Booking struct {
	From time.Time
	To   time.Time
}

func month(mon time.Month, bookings []Booking) []Day {
	res := []Day{
		Day{Content: "Sonntag", Classes: []string{"headline"}},
		Day{Content: "Montag", Classes: []string{"headline"}},
		Day{Content: "Dienstag", Classes: []string{"headline"}},
		Day{Content: "Mittwoch", Classes: []string{"headline"}},
		Day{Content: "Donnerstag", Classes: []string{"headline"}},
		Day{Content: "Freitag", Classes: []string{"headline"}},
		Day{Content: "Samstag", Classes: []string{"headline"}},
	}

	start := time.Date(2021, mon, 1, 0, 0, 0, 0, time.UTC)
	for i := time.Sunday; i < start.Weekday(); i++ {
		res = append(res, Day{Content: ""})
	}

	for ; start.Month() < mon+1; start = start.Add(time.Hour * 24) {
		classes := []string{}
		for _, b := range bookings {
			if b.From.Equal(start) {
				classes = append(classes, "event starting")
			}
			if b.To.Equal(start) {
				classes = append(classes, "event ending")
			}
			if start.After(b.From) && start.Before(b.To) {
				classes = append(classes, "event")
			}
		}
		res = append(res, Day{Content: strconv.Itoa(start.Day()), Classes: classes})
	}

	for i := start.Weekday(); i <= time.Saturday; i++ {
		res = append(res, Day{Content: ""})
	}
	return res
}
