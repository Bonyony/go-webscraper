package cmd

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"

	"github.com/spf13/cobra"
)

/*
	This is an example command just for me to understand how combra works
	It is not the point of this project
	But I guess it's an extra little feature leftover!
*/

var randomAdressCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate a random IPv4 Adress",
	Long: `Generate a random IPv4 Adress with some custom options
	Example:
	
	go-webscraper random 
	
	Could return: 123.2.45.64`,

	Run: generateAddress,
}

func init() {
	rootCmd.AddCommand(randomAdressCmd)

	randomAdressCmd.Flags().IntP("ammount", "a", 1, "Ammount of random addresses generated")

}

func generateAddress(cmd *cobra.Command, args []string) {
	ammount, _ := cmd.Flags().GetInt("ammount")

	for i := 0; i < ammount; i++ {
		newIP := ""

		for j := 0; j < 3; j++ {
			newIP += strconv.Itoa(rand.Intn(256)) + "."
		}
		newIP += strconv.Itoa(rand.Intn(256))

		// May need to rework this to be easier
		if checkIfPrivateIP(newIP) {
			fmt.Println(newIP, "[This is a private IP]")
		} else {
			fmt.Println(newIP)
		}
	}
}

// Checks if the IP is private, only made into another function as this could be a utility
func checkIfPrivateIP(ip string) bool {
	// Checks these ranges (The most common private ranges)
	// There are more private ranges than this
	// large private netwroks: 10.0.0.0 - 10.255.255.255
	// medium private networks: 172.16.0.0 – 172.31.255.255
	// home routers, LANs: 192.168.0.0 – 192.168.255.255
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil && ipCheck.IsPrivate() {
		return true
	} else {
		return false
	}
}
