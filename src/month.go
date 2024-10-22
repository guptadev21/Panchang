package GetMonth

import (
	"os"
	"strconv"
	"time"

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
	table.Header("Sun")
	table.Header("Mon")
	table.Header("Tue")
	table.Header("Wed")
	table.Header("Thu")
	table.Header("Fri")
	table.Header("Sat")

	// Initialize the first row
	row := table.Row()
	// Add empty columns for the days before the 1st
	for i := 0; i < weekday; i++ {
		row.Column("")
	}
	// Add the days of the month
	for day := 1; day <= daysInMonth; day++ {
		row.Column(strconv.Itoa(day))
		weekday = (weekday + 1) % 7
		if weekday == 0 {
			// Start a new row for the next week
			row = table.Row()
		}
	}

	table.Print(os.Stdout)
}
