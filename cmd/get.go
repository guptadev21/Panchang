package cmd

import (
	"time"

	Events "panchang/src/events"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var monthFlag string
var allFlag bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get events for a specific month",
	Long:  `Get events for a specific month. If no month is specified, the current month's events will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define color styles
		errorStyle := color.New(color.FgRed, color.Bold)
		infoStyle := color.New(color.FgBlue, color.Bold)
		titleStyle := color.New(color.FgGreen).Add(color.Underline)

		var events []Events.Event
		var err error
		if allFlag {
			events, err = Events.GetEvents(time.January, true)
			if err != nil {
				errorStyle.Printf("Failed to fetch events: %v\n", err)
				return
			}
		} else {
			parsedMonth, err := time.Parse("January", monthFlag)
			if err != nil {
				errorStyle.Println("Invalid month name. Please use the full month name, e.g., 'January'.")
				return
			}
			targetMonth := parsedMonth.Month()
			events, err = Events.GetEvents(targetMonth, false)
			if err != nil {
				errorStyle.Printf("Failed to fetch events for %s: %v\n", targetMonth.String(), err)
				return
			}
		}
		if err != nil {
			errorStyle.Printf("Failed to fetch events: %v\n", err)
			return
		}
		if len(events) == 0 {
			infoStyle.Printf("No events found.\n")
			return
		}

		// Print the events
		titleStyle.Printf("\n\nEvents :\n\n")
		for _, event := range events {
			color.Yellow("Title: %s\n", event.Title)
			color.White("Date: %s\nDescription: %s\n\n", event.Date, event.Description)
		}
	},
}

func init() {
	eventCmd.AddCommand(getCmd)

	// Add the --month flag
	getCmd.Flags().StringVarP(&monthFlag, "month", "m", time.Now().Format("January"), "Month to fetch events for (default: current month).\nUse the full month name, e.g., 'January'.")
	getCmd.Flags().BoolVar(&allFlag, "all", false, "Get all events")
}
