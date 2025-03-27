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
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace an IP address of your choice",
	Long: `Trace an IP of your choice. You can enter as many IPs as you want. 

go-webscraper trace 1.1.1.1, 2.2.2.2, 5.5.5.5

Would output:

IP        |CITY      |COUNTRY   |REGION      |LOCATION           |TIMEZONE            |POSTAL
1.1.1.1   |Brisbane  |AU        |Queensland  |-27.4820,153.0136  |Australia/Brisbane  |4101
2.2.2.2   |Latham    |US        |New York    |42.7470,-73.7590   |America/New_York    |12110
5.5.5.5   |Delhi     |IN        |Delhi       |28.6519,77.2315    |Asia/Kolkata        |110001
`,
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
	if len(args) == 0 {
		fmt.Println("Please provide an IP to trace.")
		return
	}

	// setup for the tabwriter
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "\nIP\tCITY\tCOUNTRY\tREGION\tLOCATION\tTIMEZONE\tPOSTAL")

	for _, ip := range args {
		showData(w, ip)
	}

	// Outputs the writer
	w.Flush()
}

func showData(w *tabwriter.Writer, ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := IP{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
		return
	}

	fmt.Fprint(w, data.IP, "\t", data.City, "\t", data.Country, "\t", data.Region, "\t", data.Loc, "\t", data.Timezone, "\t", data.Postal, "\n")
}

func getData(url string) []byte {
	// Sends a GET request to the URL
	res, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	// 200 should be the only response wanted
	if res.StatusCode != 200 {
		log.Println("Bad response:", res.StatusCode)
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Unable to read response")
	}

	return resByte
}
