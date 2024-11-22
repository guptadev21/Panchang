package GetMonth

import (
	"os"
	"strconv"
	"time"

	Events "panchang/src/events"

	"github.com/fatih/color"
	"github.com/markkurossi/tabulate"
)

// Constants for readability
const (
	dateFormat  = "2006-01-02"
	daysInAWeek = 7
)

// PrintMonth prints the calendar for the specified month and year
func PrintMonth(year int, month time.Month) {
	// Get the first day of the month
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	weekday := int(firstOfMonth.Weekday())
	daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()

	// Retrieve all dates with events for the month
	datesWithEvents, err := Events.GetDatesWithEvents(month, year)
	if err != nil {
		color.Red("Error retrieving event dates: %v\n", err)
		return
	}

	// print dates with events
	for _, date := range datesWithEvents {
		print(date)
	}

	// Convert datesWithEvents to a map for faster lookups
	eventDates := make(map[string]struct{})
	for _, date := range datesWithEvents {
		eventDates[date] = struct{}{}
	}

	// Get today's date
	today := time.Now()

	// Create a table for the calendar
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
	for i := 0; i < weekday; i++ {
		row.Column(color.BlackString("")) // Fill empty slots before the first day
	}

	// Populate the days of the month
	for day := 1; day <= daysInMonth; day++ {
		// Create the full date string for the current day
		currentDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Format(dateFormat)

		// Add the day to the row with appropriate formatting
		addDayToRow(row, day, currentDate, eventDates, today)

		// Start a new row for the next week
		weekday = (weekday + 1) % 7
		if weekday == 0 {
			row = table.Row() // Create a new row
		}
	}

	// Fill the remaining columns in the last row, if necessary
	for i := weekday; i < daysInAWeek; i++ {
		row.Column("")
	}

	// Print the calendar
	table.Print(os.Stdout)
}

// addDayToRow adds a day to the current row with appropriate color formatting
func addDayToRow(row *tabulate.Row, day int, currentDate string, eventDates map[string]struct{}, today time.Time) {
	_, hasEvent := eventDates[currentDate]
	// print(eventDates[currentDate])
	switch {
	case hasEvent && currentDate == today.Format(dateFormat):
		// Today's date with events
		row.Column(color.MagentaString(strconv.Itoa(day)))
	case hasEvent:
		// Event date
		row.Column(color.GreenString(strconv.Itoa(day)))
	case currentDate == today.Format(dateFormat):
		// Today's date without events
		row.Column(color.CyanString(strconv.Itoa(day)))
	default:
		// Regular date
		row.Column(color.WhiteString(strconv.Itoa(day)))
	}
}
