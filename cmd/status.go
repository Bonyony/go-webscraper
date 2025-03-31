package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/*
	This command establishes a connection with a website and
	then checks its status (Online or Offline via TCP)
*/

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "See the status of a website",
	Long: `See the status of a website and some other data
	Example:
	
	go-webscraper status 
	
	Could return: Moo-Gumbo-Gaa!`,

	Run: findStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func findStatus(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a URL to check the status of.")
		return
	}

	for _, url := range args {
		fmt.Println(url)
	}
}
