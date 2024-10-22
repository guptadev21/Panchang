/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	GetMonth "panchang/src"

	"github.com/spf13/cobra"
)

var month int
var year int

// monthCmd represents the month command
var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Get Panchang for a month",
	Long:  `Get Panchang details for a specific month and year.`,
	Run: func(cmd *cobra.Command, args []string) {
		if month < 1 || month > 12 {
			fmt.Println("Invalid month. Please enter a value between 1 and 12.")
			return
		}

		// If no year is passed, use the current year
		if year == 0 {
			year = time.Now().Year()
		}

		// Fetch and display Panchang for the specified month and year
		GetMonth.PrintMonth(year, time.Month(month))
	},
}

func init() {
	rootCmd.AddCommand(monthCmd)

	// Here you will define your flags and configuration settings.

	// Add the --year and --month flags to the month command
	monthCmd.Flags().IntVarP(&year, "year", "y", 0, "Year for the Panchang (default is current year)")
	monthCmd.Flags().IntVarP(&month, "month", "m", int(time.Now().Month()), "Month for the Panchang")
}
