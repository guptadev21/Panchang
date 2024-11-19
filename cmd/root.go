package cmd

import (
	"os"
	"time"

	GetMonth "panchang/src"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	month int
	year  int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "panchang",
	Short: "A simple command-line calendar application",
	Long:  `A simple command-line calendar application that displays a calendar for a given month and year.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentTime := time.Now()

		// Use provided flags or default to current month/year
		if year == 0 {
			year = currentTime.Year()
		}
		if month == 0 {
			month = int(currentTime.Month())
		}

		// Validate the month range
		if month < 1 || month > 12 {
			color.New(color.FgRed).Println("Invalid month. Please specify a month between 1 and 12.")
			return
		}

		GetMonth.PrintMonth(year, time.Month(month))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add flags for specifying the month and year
	rootCmd.Flags().IntVarP(&month, "month", "m", 0, "Specify the month (1-12)")
	rootCmd.Flags().IntVarP(&year, "year", "y", 0, "Specify the year (e.g., 2024)")

	// Local flags for customization
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
