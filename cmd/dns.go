package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

/*
	This is an example command just for me to understand how combra works
	It is not the point of this project
	But I guess it's an extra little feature leftover!
*/

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Check the DNS records for a URL address",
	Long: `Check the DNS records for a URL address
	Example:
	
	go-webscraper dns
	
	Could return: `,

	Run: dnsCheck,
}

func init() {
	rootCmd.AddCommand(dnsCmd)
}

func dnsCheck(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("Please provide a URL to find the DNS records.")
		return
	}

	for _, domain := range args {
		fmt.Println()
		fmt.Println(strings.ToUpper(domain))
		findIP(domain)
		findMX(domain)
		findNS(domain)
		findTXT(domain)
	}
}

func findIP(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Could not find ip(s) of", domain, "Error:", err)
		return
	}

	fmt.Println("IP addresses:")
	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Println("  - IPv4:", ip)
		} else {
			fmt.Println("  - IPv6:", ip)
		}
	}
}

func findMX(domain string) {
	// mail exchange data
	mx, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("MX: Could not find mail exchange data for", domain, err)
		return
	}
	for _, record := range mx {
		fmt.Printf("MX: %s (Pref: %d)\n", record.Host, record.Pref)
	}

}

func findNS(domain string) {
	// name server
	ns, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println("NS: Could not find name server data for", domain, err)
		return
	}

	fmt.Println("NS Records:")
	for _, record := range ns {
		fmt.Println("  -", record)
	}
}

func findTXT(domain string) {
	// text records
	txt, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println("TXT: Could not find text records for", domain, err)
		return
	}
	fmt.Println("TXT Records:")
	for _, record := range txt {
		fmt.Println("  -", record)
	}
}
