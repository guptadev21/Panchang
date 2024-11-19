package cmd

import (
	"bufio"
	"os"
	"time"

	Events "panchang/src/events"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var eventDate string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an event",
	Long:  `Add an event to the calendar. You will be prompted to enter the title and description of the event.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define color styles
		errorStyle := color.New(color.FgRed, color.Bold)
		successStyle := color.New(color.FgGreen, color.Bold)
		promptStyle := color.New(color.FgCyan, color.Bold)

		// Prompt for title and description
		reader := bufio.NewReader(os.Stdin)

		promptStyle.Print("Enter event title: ")
		title, err := reader.ReadString('\n')
		if err != nil {
			errorStyle.Printf("Error reading title: %v\n", err)
			return
		}
		title = title[:len(title)-1] // Remove the trailing newline

		promptStyle.Print("Enter event description: ")
		description, err := reader.ReadString('\n')
		if err != nil {
			errorStyle.Printf("Error reading description: %v\n", err)
			return
		}
		description = description[:len(description)-1] // Remove the trailing newline

		// Parse the date or use today's date as default
		var date time.Time
		if eventDate == "" {
			date = time.Now()
		} else {
			date, err = time.Parse("2006-01-02", eventDate)
			if err != nil {
				errorStyle.Println("Invalid date format. Please use YYYY-MM-DD.")
				return
			}
		}

		// Store the event
		_, err = Events.StoreEvent(title, date, description)
		if err != nil {
			errorStyle.Printf("Failed to store event: %v\n", err)
			return
		}

		successStyle.Println("Event added successfully!")
	},
}

func init() {
	eventCmd.AddCommand(addCmd)

	// Add the --date flag
	addCmd.Flags().StringVar(&eventDate, "date", "", "Date of the event in YYYY-MM-DD format (default: today's date)")
}
