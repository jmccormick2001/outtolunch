package main

import (
	"math"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var DaysTill int
var MonthsTill int

func main() {
	DaysTill, MonthsTill = retirementAlgorithm()

	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       genPage(MonthsTill),
	}, nil
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

func genPage(waves int) string {
	//shipIcon := "https://en.pimg.jp/102/717/473/1/102717473.jpg"
	//islandIcon := "https://icons.iconarchive.com/icons/iconarchive/seaside/512/Island-icon.png"
	var str strings.Builder

	str.WriteString("<html>")
	str.WriteString("<body>")
	str.WriteString("<h2>the ship sails in 10 days or 5 months</h2>")
	str.WriteString("<img src=\"https://en.pimg.jp/102/717/473/1/102717473.jpg\" height=\"100\" width=\"100\">")
	for i := 0; i < waves; i++ {
		str.WriteString("<img src=\"https://icons.iconarchive.com/icons/iconarchive/seaside/512/Water-Wave-icon.png\" height=\"40\" width=\"40\">")
	}

	str.WriteString("<img src=\"https://icons.iconarchive.com/icons/iconarchive/seaside/512/Island-icon.png\" height=\"100\" width=\"100\">")
	str.WriteString("</body>")
	str.WriteString("</html>")
	return str.String()
}

/**
<html>
<body>

<h2>the ship sails in 10 days or 5 months</h2>
<img src="https://en.pimg.jp/102/717/473/1/102717473.jpg" height="100" width="100">
<img src="https://icons.iconarchive.com/icons/iconarchive/seaside/512/Water-Wave-icon.png" height="40" width="40">
<img src="https://icons.iconarchive.com/icons/iconarchive/seaside/512/Water-Wave-icon.png" height="40" width="40">

<img src="https://icons.iconarchive.com/icons/iconarchive/seaside/512/Island-icon.png" height="100" width="100">
</body>
</html>
*/
