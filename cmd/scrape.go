/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

// CLI Logic

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape a website of your choice",
	Long: `Scrape a website of your choice.
	
	Example:
	
	go-webscraper scrape
	
	Could return: 'Chumbus Wumbus or Grumpus Grumbus!'`,
	Run: scrape,
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// type Product struct {
// 	Name   string `json:"name"`
// 	Price  string `json:"price"`
// 	ImgUrl string `json:"imgurl"`
// }

func scrape(cmd *cobra.Command, args []string) {
	fmt.Println("Scrape command executed!")

	c := colly.NewCollector(
		// colly.AllowedDomains("www.sweetwater.com"),
		colly.MaxDepth(1),
	)

	c.OnRequest(func(r *colly.Request) {
		// These lines pretends to be an internet browser to bypass limiting
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Referer", "https://www.google.com/")
		r.Headers.Set("DNT", "1") // Do Not Track
		r.Headers.Set("Connection", "keep-alive")
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Status Code %v Error on %s: %s\n", r.StatusCode, r.Request.URL, err)
		fmt.Println("Response Headers")
		for key, value := range *r.Headers {
			fmt.Printf(" %s: %s\n", key, value)
		}
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print Link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// err := c.Visit("https://www.sweetwater.com/")
	err := c.Visit("https://en.wikipedia.org/")
	if err != nil {
		fmt.Println("Colly error:", err)
	}
}
