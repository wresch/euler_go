package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc = `

You are given the following information, but you may prefer to do some
research for yourself.

\begin{itemize}
  \item 1 Jan 1900 was a Monday
  \item Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
  \item A leap year occurs on any year evenly divisible by 4, but not
     on a century unless it is divisible by 400.
\end{itemize}

\emph{How many Sundays fell on the first of the month during the twentieth
century (1 Jan 1901 to 31 Dec 2000)?}

An approximate answer can be obtained as $1200 / 7 = 171$. That is based
on the assumption that over time about 1 out of 7 first of the months
should be a Sunday.  Naturally that's not always true.

I used modular arithmetic to solve this one. Not very elegant or
efficient, but it works.
`

// month is 1-based to make it easier to think about
func daysPerMonth(year, month int) (n int) {
	switch month {
	case 4, 6, 9, 11:
		n = 30
	case 1, 3, 5, 7, 8, 10, 12:
		n = 31
	case 2:
		if year%4 == 0 {
			if year%100 == 0 && year%400 != 0 {
				n = 28
			} else {
				n = 29
			}
		} else {
			n = 28
		}
	}
	return
}

func main() {
	// Jan 1, 1900 was a monday!
	day := 0 // 0: Mon; 1: Tue;...;6: Sun
	sunAtFirst := 0
	for year := 1900; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			//fmt.Printf("%02d/%d had %d days; ", month, year, daysPerMonth(year, month))
			//fmt.Printf(" next month starts with a %d\n", daysPerMonth(year, month) % 7 + day)
			day = (day + (daysPerMonth(year, month) % 7)) % 7
			if day == 6 {
				switch {
				case year == 2000 && month < 12:
					sunAtFirst++
				case year >= 1901 && year < 2000:
					sunAtFirst++
				}
			}
		}
	}
	eu.Output(19, desc, fmt.Sprintf("There were %d months that started with a Sunday.\n", sunAtFirst))
}
