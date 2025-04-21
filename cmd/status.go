package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/Bonyony/go-webscraper/util"

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
	
	go-webscraper status frankfrancione.com -l -p
	
	Will return:

	200 OK  -  https://frankfrancione.com
	Scheme: https Host: frankfrancione.com Path: 

	IP addresses for frankfrancione.com:
	IPv6: 2606:4700:7::60
	IPv6: 2a06:98c1:58::60
	IPv4: 162.159.140.98
	IPv4: 172.66.0.96
	
	`,

	Run: findStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().BoolP("parse-url", "p", false, "Parse each part of the URL")
	statusCmd.Flags().BoolP("lookup-ip", "l", false, "See the IP address(es) of the URL")

}

func findStatus(cmd *cobra.Command, args []string) {
	isParse, _ := cmd.Flags().GetBool("parse-url")
	isLookup, _ := cmd.Flags().GetBool("lookup-ip")

	if len(args) == 0 {
		fmt.Println("Please provide a URL to check the status of as an argument.\nExample: go-webscraper status https://www.google.com")
		return
	}

	for _, domain := range args {
		fmt.Println()
		visitDomain(domain)
		if isParse {
			fmt.Println()
			parseURL(domain)
		}
		if isLookup {
			fmt.Println()
			lookupIP(domain)
		}
	}

}

func visitDomain(domain string) {
	cleanedURL := util.NormalizeURL(domain)

	res, err := http.Get(cleanedURL)
	if err != nil {
		fmt.Println("Connection failed: ", err)
		return
	}
	defer res.Body.Close()

	status := res.Status

	fmt.Println(status, " - ", cleanedURL)
}

func parseURL(domain string) {
	cleanedURL := util.NormalizeURL(domain)
	u, err := url.Parse(cleanedURL)
	if err != nil {
		log.Println("Unable to parse rawURL:", err)
		return
	}

	fmt.Print("Scheme: ", u.Scheme, " Host: ", u.Host, " Path: ", u.Path, "\n\n")
}

func lookupIP(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Could not find ip(s) of", domain, "Error:", err)
		return
	}

	fmt.Printf("IP addresses for %s:\n", domain)
	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Println("IPv4:", ip)
		} else {
			fmt.Println("IPv6:", ip)
		}
	}
}
