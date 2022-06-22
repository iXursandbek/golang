package main

import (
	"fmt"
	"time"
)

func main() {

	dobStr := "14.10.1996" // Replace this date with your birthday
	givenDate, err := time.Parse("02.01.2006", dobStr)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s is %s", givenDate.Format("02-01-2006"), FindWeekday(givenDate))
	fmt.Println(FindWeekday((givenDate)))
}

func FindWeekday(date time.Time) (weekday string) {
	return date.Weekday().String()
}
