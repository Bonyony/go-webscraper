package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/cobra"
)

/*
	This command establishes a connection with a website and
	then checks its status (Online or Offline via TCP)
	It will also check if a domain has a valid email, etc.
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
		fmt.Println("Please provide a URL to check the status of as an argument.\nExample: go-webscraper status https://www.google.com")
		return
	}

	for _, domain := range args {
		getStatus(domain)
	}
}

func getStatus(domain string) {
	u, err := url.Parse(domain)
	if err != nil {
		log.Println("Unable to parse rawURL:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme, "Host:", u.Host, "Path:", u.Path)
}
