/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace an IP address of your choice",
	Long: `Trace an IP of your choice. A potential return would look like: 

Chumble Wumble`,
	Run: traceIP,
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

// sample response for 1.1.1.1
// {
//     "ip": "1.1.1.1",
//     "hostname": "one.one.one.one",
//     "city": "Brisbane",
//     "region": "Queensland",
//     "country": "AU",
//     "loc": "-27.4820,153.0136",
//     "org": "AS13335 Cloudflare, Inc.",
//     "postal": "4101",
//     "timezone": "Australia/Brisbane",
//     "readme": "https://ipinfo.io/missingauth"
// }

type IP struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func traceIP(cmd *cobra.Command, args []string) {
	fmt.Println("trace command called!")

	if len(args) > 0 {
		for _, ip := range args {
			showData(ip)
		}
	} else {
		fmt.Println("Please provide an IP to trace.")
	}
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := IP{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	fmt.Println("DATA FOUND: ")
	fmt.Printf("IP: %s\nCITY: %s\nCOUNTRY: %s\nREGION: %s\nLOCATION: %s\nTIMEZONE: %s\nPOSTAL: %s\n", data.IP, data.City, data.Country, data.Region, data.Loc, data.Timezone, data.Postal)

}

func getData(url string) []byte {
	// Sends a GET request to the URL
	res, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Unable to read response")
	}

	return resByte
}
