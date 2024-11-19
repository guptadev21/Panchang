package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Manage events",
	Long:  `Manage events, including adding, deleting, and retrieving events.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show an error message if no subcommand is provided
		color.Red("Error: Missing subcommand. Use 'event --help' to see available subcommands.")
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)
}
