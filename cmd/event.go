package cmd

import (
	"fmt"
	"time"

	Events "panchang/src/events"

	"github.com/spf13/cobra"
)

var monthFlag string

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Get events for the month",
	Long:  `Get the events for the current month by default or use the --month flag to specify another month.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Determine the target month
		var targetMonth time.Month
		if monthFlag == "" {
			targetMonth = time.Now().Month() // Default to the current month
		} else {
			parsedMonth, err := time.Parse("January", monthFlag)
			if err != nil {
				fmt.Println("Invalid month name. Please use the full month name, e.g., 'January'.")
				return
			}
			targetMonth = parsedMonth.Month()
		}

		// Fetch events for the target month
		events, err := Events.GetEvents(targetMonth)
		if err != nil {
			fmt.Printf("Failed to fetch events: %v\n", err)
			return
		}
		if len(events) == 0 {
			fmt.Printf("No events found for %s.\n", targetMonth.String())
			return
		}

		// Print the events
		fmt.Printf("Events for %s:\n", targetMonth.String())
		for _, event := range events {
			fmt.Printf("Title: %s\nDate: %s\nDescription: %s\n\n", event.Title, event.Date, event.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)

	// Add the --month flag
	eventCmd.Flags().StringVar(&monthFlag, "month", "", "Month to fetch events for (default: current month, e.g., 'January')")
}
