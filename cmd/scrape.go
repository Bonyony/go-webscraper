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

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

// CLI Logic

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape a website of your choice",
	Long: `Scrape a website of your choice.

	Current entrypoints are:
		"https://www.musiciansfriend.com/sitemap"
		"https://reverb.com/sitemaps/sitemap.xml.gz"

	
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
	scrapeCmd.Flags().BoolP("reverb", "r", false, "Include Reverb to be scraped")
}

type Product struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

// maps key value pairs of category names and links
// var categories = make(map[string]string)

func scrape(cmd *cobra.Command, args []string) {
	fmt.Println("Scrape command executed!")

	isMusiciansFriend, _ := cmd.Flags().GetBool("musicians-friend")
	isReverb, _ := cmd.Flags().GetBool("reverb")

	// need to rework logic here so that flag order is respected
	if isMusiciansFriend {
		scrapeMusiciansFriendSitemap()
	}
	if isReverb {
		scrapeReverbSitemap()
	}
}

func scrapeMusiciansFriendSitemap() {
	c := colly.NewCollector()

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

	var categoriesList []string

	c.OnHTML("h2.sitemap-heading", func(e *colly.HTMLElement) {
		categoryName := e.Text
		if categoryName != "" {
			categoriesList = append(categoriesList, categoryName)
		}
	})

	err := c.Visit("https://www.musiciansfriend.com/sitemap")
	if err != nil {
		log.Fatal("Failed to scrape sitemap:", err)
	}

	fmt.Println("\nChoose a category:")
	chooseOptionFromList(categoriesList, scrapeMusiciansFriendCategories)
}

func scrapeMusiciansFriendCategories(category string) {
	c := colly.NewCollector()

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

	// maps key value pairs of category names and links
	var categories = make(map[string]string)

	c.OnHTML("h2.sitemap-heading", func(e *colly.HTMLElement) {

		sectionTitle := e.Text
		if sectionTitle == category {
			e.DOM.Next().Find("li a").Each(func(i int, s *goquery.Selection) {
				subCategoryName := s.Text()
				subCategoryUrl, _ := s.Attr("href")

				if subCategoryName != "" && subCategoryUrl != "" {
					categories[subCategoryName] = e.Request.AbsoluteURL(subCategoryUrl)
				}
			})
		}
	})

	err := c.Visit("https://www.musiciansfriend.com/sitemap")
	if err != nil {
		log.Fatal("Failed to scrape sitemap:", err)
	}

	fmt.Println("\nChoose a category:")
	chooseOptionFromMap(categories, scrapeMusiciansFriendProducts)
}

func scrapeMusiciansFriendProducts(subCategoryUrl string) {
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

	err := c.Visit(subCategoryUrl)
	if err != nil {
		fmt.Println("Colly error:", err)
	}

	jsonData, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	if len(products) == 0 {
		fmt.Println("No product cards on this page. Visit this URL to see what is offered here:", subCategoryUrl)
	} else {
		fmt.Println(string(jsonData))
	}

}

// Prompts the user to input a number corresponding to available choices (takes in a map of strings)
func chooseOptionFromMap(options map[string]string, callback func(string)) {
	// loop through and display the options
	keys := make([]string, 0, len(options))
	for k := range options {
		keys = append(keys, k)
	}
	for i, name := range keys {
		fmt.Printf("[%d] %s\n", i+1, name)
	}

	// get user input (number for category)
	var choice int
	fmt.Print("Enter the number of your category: ")
	fmt.Scan(&choice)

	if choice > 0 && choice <= len(keys) {
		callback(options[keys[choice-1]])
	} else {
		fmt.Println("Invalid choice.")
	}
}

// Prompts the user to input a number corresponding to available choices (takes in a list of strings)
func chooseOptionFromList(options []string, callback func(string)) {
	for i, name := range options {
		fmt.Printf("[%d] %s\n", i+1, name)
	}

	var choice int
	fmt.Print("Enter the number of your category: ")
	fmt.Scan(&choice)

	if choice > 0 && choice <= len(options) {
		callback(options[choice-1])
	} else {
		fmt.Println("Invalid choice.")
	}
}

func scrapeReverbSitemap() {
	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
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

	var urlList []string

	c.OnXML("//url/loc", func(x *colly.XMLElement) {
		url := x.Text
		if url != "" {
			urlList = append(urlList, url)
		}
	})

	err := c.Visit("https://reverb.com/sitemaps/sitemap.xml.gz")
	if err != nil {
		log.Fatal("Error scraping sitemap:", err)
	}

	fmt.Println("\nChoose a link to visit:")
	chooseOptionFromList(urlList, func(s string) {})

}
