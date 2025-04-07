/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace an IP address of your choice",
	Long: `Trace an IP of your choice. You can enter as many IPs as you want. 

go-webscraper trace 1.1.1.1 2.2.2.2 5.5.5.5

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

	traceCmd.Flags().BoolP("scan-port", "s", false, "Scan the IP address for open ports")
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
	isScan, _ := cmd.Flags().GetBool("scan-port")

	if len(args) == 0 {
		fmt.Println("Please provide an IP to trace.")
		return
	}

	// setup for the tabwriter
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "\nIP\tCITY\tCOUNTRY\tREGION\tLOCATION\tTIMEZONE\tPOSTAL")

	for _, ip := range args {
		showData(w, ip)
		if isScan {
			scanPorts(w, ip)
		}
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
		log.Println("Bad response:", res.StatusCode, res.Request.URL)
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Unable to read response")
	}

	return resByte
}

func scanPorts(w *tabwriter.Writer, ip string) {
	topPorts := []int{80, 22, 443, 3306, 3389, 21, 23, 8080, 8443, 53, 25}
	services := map[int]string{
		22:   "SSH",
		80:   "HTTP",
		443:  "HTTPS",
		3306: "MySQL",
		3389: "RDP",
		21:   "FTP",
		23:   "Telnet",
		8080: "Alt HTTP",
		8443: "Alt HTTPS",
		53:   "DNS",
		25:   "SMTP",
	}

	fmt.Fprintln(w, "\nPORT\tSERVICE\tSTATUS")
	for _, port := range topPorts {
		target := fmt.Sprintf("%s:%d", ip, port)

		conn, err := net.DialTimeout("tcp", target, 2*time.Second)
		if err != nil {
			fmt.Fprintf(w, "%d\t%s\tClosed\n", port, services[port])
			continue
		}

		fmt.Fprintf(w, "%d\t%s\tOpen\n", port, services[port])

		conn.Close()
	}
	w.Flush()
}
