/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

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

type Product struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func scrape(cmd *cobra.Command, args []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("https://www.sweetwater.com"),
	)

	c.Visit("https://www.sweetwater.com/c590--Solidbody_Guitars")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Println(e.Request.Visit(e.Attr("href")))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
}
