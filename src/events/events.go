package Events

import (
	"bufio"
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

var BasePath string

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Failed to get working directory: %v", err))
	}
	BasePath = dir + "/EventFiles"
	// Ensure the base path directory exists
	if _, err := os.Stat(BasePath); os.IsNotExist(err) {
		err := os.MkdirAll(BasePath, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Failed to create base directory: %v", err))
		}
	}
}

// StoreEvent stores the event details in a file named <month_name>_events.txt
func StoreEvent(title string, date time.Time, description string) {
	// Construct the filename based on the month name
	fileName := date.Month().String() + "_events.txt"
	// Create the event instance
	event := Event{
		Title:       title,
		Date:        date.Format("2006-01-02"), // Format the date as YYYY-MM-DD
		Description: description,
	}

	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(BasePath+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	defer file.Close()

	// Serialize the event as JSON
	eventData, err := json.Marshal(event)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal event: %v", err))
	}

	// Write the serialized event data to the file
	_, err = file.WriteString(string(eventData) + "\n")
	if err != nil {
		panic(fmt.Sprintf("Failed to write to file: %v", err))
	}

	fmt.Println("Event stored successfully.")
}

// GetEvents retrieves the events for the specified month
func GetEvents(month time.Month) ([]Event, error) {
	// Construct the file name based on the month
	fileName := filepath.Join(BasePath, month.String()+"_events.txt")

	// Open the file for the specified month
	file, err := os.Open(fileName)
	if err != nil {
		// Return an empty list if the file doesn't exist
		if os.IsNotExist(err) {
			return []Event{}, nil
		}
		return nil, fmt.Errorf("failed to open events file: %w", err)
	}
	defer file.Close()

	// Read and parse the file
	var events []Event
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse the JSON data into an Event object
		var event Event
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			return nil, fmt.Errorf("failed to parse event: %w", err)
		}

		// Append the event to the list
		events = append(events, event)
	}

	// Check for any error during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading events file: %w", err)
	}

	return events, nil
}
