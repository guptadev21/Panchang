package cmd

import (
	"time"

	Events "panchang/src/events"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var monthFlag string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get events",
	Long:  `Get events for a specific month. If no month is specified, the current month's events will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define color styles
		errorStyle := color.New(color.FgRed, color.Bold)
		infoStyle := color.New(color.FgBlue, color.Bold)
		titleStyle := color.New(color.FgGreen).Add(color.Underline)

		// Determine the target month
		var targetMonth time.Month
		if monthFlag == "" {
			targetMonth = time.Now().Month() // Default to the current month
		} else {
			parsedMonth, err := time.Parse("January", monthFlag)
			if err != nil {
				errorStyle.Println("Invalid month name. Please use the full month name, e.g., 'January'.")
				return
			}
			targetMonth = parsedMonth.Month()
		}

		// Fetch events for the target month
		events, err := Events.GetEvents(targetMonth)
		if err != nil {
			errorStyle.Printf("Failed to fetch events: %v\n", err)
			return
		}
		if len(events) == 0 {
			infoStyle.Printf("No events found for %s.\n", targetMonth.String())
			return
		}

		// Print the events
		titleStyle.Printf("\n\nEvents for %s:\n\n", targetMonth.String())
		for _, event := range events {
			color.Yellow("Title: %s\n", event.Title)
			color.White("Date: %s\nDescription: %s\n\n", event.Date, event.Description)
		}
	},
}

func init() {
	eventCmd.AddCommand(getCmd)

	// Add the --month flag
	getCmd.Flags().StringVarP(&monthFlag, "month", "m", "", "Month to fetch events for (default: current month, e.g., 'January')")
}
