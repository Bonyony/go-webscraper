package cmd

import (
	"fmt"
	"log"
	"net/http"
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

	statusCmd.Flags().BoolP("parse-url", "p", false, "Parse each part of the URL")
}

func findStatus(cmd *cobra.Command, args []string) {
	isParse, _ := cmd.Flags().GetBool("parse-url")

	if len(args) == 0 {
		fmt.Println("Please provide a URL to check the status of as an argument.\nExample: go-webscraper status https://www.google.com")
		return
	}

	for _, domain := range args {
		visitDomain(domain)
		if isParse {
			parseURL(domain)
		}
	}

}

func visitDomain(domain string) {
	// needs the Scheme to properly run, so provide one if there is not?
	res, err := http.Get(domain)
	if err != nil {
		fmt.Println("Connection failed: ", err)
		return
	}
	defer res.Body.Close()

	status := res.Status

	fmt.Println(status, " - ", domain)
}

func parseURL(domain string) {
	u, err := url.Parse(domain)
	if err != nil {
		log.Println("Unable to parse rawURL:", err)
		return
	}

	fmt.Print("Scheme:", u.Scheme, "Host:", u.Host, "Path:", u.Path, "\n\n")
}
