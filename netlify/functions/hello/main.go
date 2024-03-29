package main

import (
	"fmt"
	"math/rand"
	"strings"

	"example.com/m/internal"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var DaysFrom int
var DaysTill int
var MonthsTill int

func main() {
	DaysFrom, DaysTill, MonthsTill = internal.RetirementAlgorithm()

	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       genPage(DaysFrom, DaysTill, MonthsTill),
	}, nil
}

func genPage(daysFrom, days, months int) string {
	//shipIcon := "https://en.pimg.jp/102/717/473/1/102717473.jpg"
	//islandIcon := "https://icons.iconarchive.com/icons/iconarchive/seaside/512/Island-icon.png"
	// sharkIcon := "https://img.freepik.com/free-vector/aggressive-great-white-shark-cartoon_1308-102349.jpg"

	min := days - daysFrom
	max := days
	sharkDay := rand.Intn(max-min) + min
	var str strings.Builder

	str.WriteString("<html>")
	str.WriteString("<body>")
	str.WriteString("<h2>the voyage is underway, ship has sailed for " + fmt.Sprintf("%d", daysFrom) + " and ship arrives in " + fmt.Sprintf("%d", days) + " days or " + fmt.Sprintf("%d", months) + " months</h2>")
	str.WriteString("<br>")
	str.WriteString("<audio controls>")
	str.WriteString("<source src=\"https://www.silvermansound.com/wp-content/uploads/the-buccaneers-haul.mp3\" type=\"audio/mp3\">")
	str.WriteString("</audio>")
	str.WriteString("<br>")
	for i := 0; i < daysFrom; i++ {
		str.WriteString("<img src=\"https://icons.iconarchive.com/icons/iconarchive/seaside/512/Water-Wave-icon.png\" height=\"40\" width=\"40\">")
	}

	str.WriteString("<img src=\"https://en.pimg.jp/102/717/473/1/102717473.jpg\" height=\"100\" width=\"100\">")

	for i := daysFrom + 1; i < days; i++ {
		if i == sharkDay {
			str.WriteString("<img src=\"https://img.freepik.com/free-vector/aggressive-great-white-shark-cartoon_1308-102349.jpg\" height=\"40\" width=\"40\">")
		} else {
			str.WriteString("<img src=\"https://icons.iconarchive.com/icons/iconarchive/seaside/512/Water-Wave-icon.png\" height=\"40\" width=\"40\">")
		}
	}

	str.WriteString("<img src=\"https://icons.iconarchive.com/icons/iconarchive/seaside/512/Island-icon.png\" height=\"100\" width=\"100\">")
	//str.WriteString("Music: The Buccaneer's Haul by Shane Ivers - https://www.silvermansound.com")
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
