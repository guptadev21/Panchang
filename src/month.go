package GetMonth

import (
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"

	"github.com/markkurossi/tabulate"
)

func PrintMonth(year int, month time.Month) {
	// Get the first day of the month
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// Find out which day of the week the 1st falls on
	weekday := int(firstOfMonth.Weekday())
	// Number of days in the month
	daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()

	// Create a table to store the calendar data
	table := tabulate.New(tabulate.Plain)

	// Add header for the days of the week
	table.Header(color.YellowString("Sun"))
	table.Header(color.YellowString("Mon"))
	table.Header(color.YellowString("Tue"))
	table.Header(color.YellowString("Wed"))
	table.Header(color.YellowString("Thu"))
	table.Header(color.YellowString("Fri"))
	table.Header(color.YellowString("Sat"))

	// Initialize the first row
	row := table.Row()
	// Add empty columns for the days before the 1st
	for i := 0; i < weekday; i++ {
		row.Column(color.BlackString(""))
	}
	// Add the days of the month
	for day := 1; day <= daysInMonth; day++ {
		// Check if the current day is today
		today := time.Now()
		if year == today.Year() && month == today.Month() && day == today.Day() {
			row.Column(color.BlueString(strconv.Itoa(day)))
		} else {
			row.Column(color.WhiteString(strconv.Itoa(day)))
		}
		weekday = (weekday + 1) % 7
		if weekday == 0 {
			// Start a new row for the next week
			row = table.Row()
		}
	}

	table.Print(os.Stdout)
}
