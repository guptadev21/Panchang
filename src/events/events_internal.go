package Events

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// GetDatesWithEvents retrieves all dates that have events in the specified month
func GetDatesWithEvents(month time.Month, year int) ([]string, error) {
	// Open the centralized events file
	file, err := os.Open(EventsFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil // Return an empty list if the file doesn't exist
		}
		return nil, fmt.Errorf("failed to open events file: %w", err)
	}
	defer file.Close()

	type EventsByMonth map[string][]Event

	// Decode the JSON into the structured map
	var eventsByMonth EventsByMonth
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&eventsByMonth); err != nil {
		return nil, fmt.Errorf("failed to decode events file: %w", err)
	}

	// Get the name of the month (e.g., "November") to match the JSON key
	monthName := month.String()

	// Get events for the specified month
	events, found := eventsByMonth[monthName]
	if !found {
		return []string{}, nil // No events for this month
	}

	// Filter events for the specified year and collect unique dates
	eventDates := make(map[string]struct{})
	for _, event := range events {
		eventTime, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			return nil, fmt.Errorf("invalid event date format: %w", err)
		}
		if eventTime.Year() == year {
			eventDates[event.Date] = struct{}{}
		}
	}

	// Convert the map keys to a slice of strings
	var dates []string
	for date := range eventDates {
		dates = append(dates, date)
	}

	return dates, nil
}
