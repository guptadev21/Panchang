package Events

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Event Struct
type Event struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

// EventsWrapper represents a JSON wrapper containing an array of events
type EventsWrapper struct {
	Events []Event `json:"events"`
}

var (
	BasePath       string
	EventsFilePath string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("Failed to get user home directory: %v", err))
	}
	BasePath = filepath.Join(homeDir, "Documents", "EventFiles")
	// Ensure the base path directory exists
	if _, err := os.Stat(BasePath); os.IsNotExist(err) {
		err := os.MkdirAll(BasePath, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Failed to create base directory: %v", err))
		}
	}
	EventsFilePath = filepath.Join(BasePath, "events.json")
}

// StoreEvent stores the event details in the centralized events file
func StoreEvent(title string, date time.Time, description string) (bool, error) {
	// Read existing events from the centralized file
	allEvents := make(map[string][]Event)
	file, err := os.Open(EventsFilePath)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&allEvents); err != nil {
			return false, fmt.Errorf("failed to decode existing events: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return false, fmt.Errorf("failed to open events file: %w", err)
	}

	// Create a new event instance
	event := Event{
		Title:       title,
		Date:        date.Format("2006-01-02"),
		Description: description,
	}

	// Add the event to the correct month
	monthName := date.Month().String()
	allEvents[monthName] = append(allEvents[monthName], event)

	// Write the updated events back to the centralized file
	file, err = os.Create(EventsFilePath)
	if err != nil {
		return false, fmt.Errorf("failed to open events file for writing: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(allEvents); err != nil {
		return false, fmt.Errorf("failed to write events: %w", err)
	}

	return true, nil
}

// GetEvents retrieves the events for the specified month or all events if `all` is true
func GetEvents(month time.Month, all bool) ([]Event, error) {
	// Open the centralized events file
	file, err := os.Open(EventsFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return an empty list if the file doesn't exist
			return []Event{}, nil
		}
		return nil, fmt.Errorf("failed to open events file: %w", err)
	}
	defer file.Close()

	// Parse the JSON data from the file
	var allEvents map[string][]Event
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&allEvents); err != nil {
		return nil, fmt.Errorf("failed to decode events file: %w", err)
	}

	// If `all` is true, collect all events across months
	if all {
		var events []Event
		for _, monthEvents := range allEvents {
			events = append(events, monthEvents...)
		}
		return events, nil
	}

	// Retrieve events for the specified month
	monthName := month.String()
	if events, exists := allEvents[monthName]; exists {
		return events, nil
	}

	return []Event{}, nil
}

func GetEventsByDate(date time.Time) ([]Event, error) {
	events, err := GetEvents(date.Month(), false)
	if err != nil {
		return nil, err
	}

	var filteredEvents []Event
	for _, event := range events {
		if event.Date == date.Format("2006-01-02") {
			filteredEvents = append(filteredEvents, event)
		}
	}

	return filteredEvents, nil
}
