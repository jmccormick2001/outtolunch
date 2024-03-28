package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	daysTill, monthsTill := retirementAlgorithm()

	fmt.Printf("the ship sails in %d days or %d months \n", daysTill, monthsTill)
}

func retirementAlgorithm() (days int, months int) {
	targetDate := time.Date(2025, time.Month(4), 1, 0, 0, 0, 0, time.UTC)
	today := time.Now()
	months = diffMonths(targetDate, today)
	diff := time.Until(targetDate)
	days = roundTime(diff.Seconds() / 86400)
	return
}
func diffMonths(now time.Time, then time.Time) int {
	diffYears := now.Year() - then.Year()
	if diffYears == 0 {
		return int(now.Month() - then.Month())
	}

	if diffYears == 1 {
		return monthsTillEndOfYear(then) + int(now.Month())
	}

	yearsInMonths := (now.Year() - then.Year() - 1) * 12
	return yearsInMonths + monthsTillEndOfYear(then) + int(now.Month())
}

func monthsTillEndOfYear(then time.Time) int {
	return int(12 - then.Month())
}

func roundTime(input float64) int {
	var result float64
	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}
	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)
	return int(i)
}
