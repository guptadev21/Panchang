package cmd

import (
	"time"

	Events "panchang/src/events"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Manage events",
	Long:  `Get events for a specific date. If no date is specified, today's events will be displayed. For adding events, use the 'add' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		titleStyle := color.New(color.FgGreen).Add(color.Underline)

		var targetDate time.Time
		if eventDate == "" {
			targetDate = time.Now()
		} else {
			parsedDate, err := time.Parse("2006-01-02", eventDate)
			if err != nil {
				color.Red("Error: Invalid date format. Please use YYYY-MM-DD.")
				return
			}
			targetDate = parsedDate
		}

		events, err := Events.GetEventsByDate(targetDate)
		if err != nil {
			color.Red("Failed to fetch events: %v\n", err)
			return
		}
		if len(events) == 0 {
			color.Blue("No events found for %d-%02d-%02d.\n", targetDate.Year(), targetDate.Month(), targetDate.Day())
			return
		}

		// Print the events
		titleStyle.Printf("\n\nEvents for %d-%02d-%02d:\n\n", targetDate.Year(), targetDate.Month(), targetDate.Day())
		for _, event := range events {
			color.Yellow("Title: %s\n", event.Title)
			color.White("Date: %s\nDescription: %s\n\n", event.Date, event.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)
	eventCmd.Flags().StringVarP(&eventDate, "date", "d", "", "Event date in the format YYYY-MM-DD")
}
