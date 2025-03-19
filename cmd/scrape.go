/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

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
	
	Could return: 'Chumbus Wumbus or Grumpus Grumbus!'
	{
		"name": "Fender Player II Stratocaster HSS Rosewood Fingerboard...",
		"price": "$649.99",
		"imgurl": "https://media.musiciansfriend.com/is/image/MMGS7/Player-II-Stratocaster-HSS-Rosewood-Fingerboard-Limited-Edition-Electric-Guitar-Candy-Red-Burst/M11732000001000-00-400x400.jpg"
	},`,

	Run: scrape,
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	scrapeCmd.Flags().BoolP("musicians-friend", "m", false, "Include Musiciansfriend to be scraped")
}

type Product struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func scrape(cmd *cobra.Command, args []string) {
	fmt.Println("Scrape command executed!")

	isMusiciansFriend, _ := cmd.Flags().GetBool("musicians-friend")

	if isMusiciansFriend {
		scrapeMusiciansFriend()
	}
}

func scrapeMusiciansFriend() {
	products := make([]Product, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("www.musiciansfriend.com"),
		colly.MaxDepth(1),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*musiciansfriend.com",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		// These lines pretends to be an internet browser to bypass limiting
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
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
	c.OnHTML("div.product-card", func(e *colly.HTMLElement) {
		fmt.Println("Found a product-card div!")
		card, err := e.DOM.Html()
		if err != nil {
			fmt.Println("Error extracting the HTML")
			return
		}
		fmt.Println("Scraped card HTML:\n", card)

		product := Product{
			Name:   e.ChildText("a.ui-link"),
			Price:  e.ChildText(".sale-price"),
			ImgUrl: e.ChildAttr("img", "src"),
		}

		// combats lazy loading of product images
		if !strings.Contains(product.ImgUrl, "https") {
			product.ImgUrl = e.ChildAttr("img", "data-src")
		}
		products = append(products, product)
	})

	err := c.Visit("https://www.musiciansfriend.com/electric-guitars")
	if err != nil {
		fmt.Println("Colly error:", err)
	}

	jsonData, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	fmt.Println(string(jsonData))
}
