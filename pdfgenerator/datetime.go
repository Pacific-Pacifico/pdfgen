package pdfgenerator

import (
	"strconv"
	"time"
)

func GetDateInWords() string {
	currentTime := time.Now()
	// t := currentTime.String()
	// fmt.Println(t)
	day := currentTime.Day()
	month := currentTime.Month()
	year := currentTime.Year()
	// fmt.Println(reflect.TypeOf(day))
	// fmt.Println(reflect.TypeOf(month))
	// fmt.Println(reflect.TypeOf(year))

	dateStr := strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)
	// fmt.Println(dateStr)
	return dateStr
}

func GetDateTimeInWords() string {
	currentTime := time.Now()
	// t := currentTime.String()
	// fmt.Println(t)
	day := currentTime.Day()
	month := currentTime.Month()
	year := currentTime.Year()
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()
	// fmt.Println(reflect.TypeOf(day))
	// fmt.Println(reflect.TypeOf(month))
	// fmt.Println(reflect.TypeOf(year))
	// fmt.Println(reflect.TypeOf(hour))
	// fmt.Println(reflect.TypeOf(minute))
	// fmt.Println(reflect.TypeOf(second))

	dateStr := strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year) + "  " +
		strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)
	// fmt.Println(dateStr)
	return dateStr
}

// func main() {

// }
