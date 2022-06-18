package main

import (
	"context"
	"daft-stats/db"
	"daft-stats/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Connection URI
const uri = "mongodb://localhost:27017"

func main() {
	Scrape()
}

func Scrape() {
	db := db.Connect()

	property_links := scrape_properties()

	properties := build_properties_from_links(property_links)
	var properties_interface []interface{}
	for _, v := range properties {
		properties_interface = append(properties_interface, v)
	}
	_, err := db.Collection("properties").InsertMany(context.TODO(), properties_interface)
	if err != nil {

		fmt.Print(err)
	}
}

func build_properties_from_links(links []string) []models.Property {
	properties := []models.Property{}
	var price int
	var url string
	var location string
	var id int
	var beds int
	var baths int
	var type_ string
	var entered string
	var views int

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("div[data-testid=price]", func(h *colly.HTMLElement) {
		if strings.Contains(h.Text, "€") {
			string_price := strings.Split(h.Text, " ")[0]
			removed_euro := strings.Split(string_price, "€")[1]
			clean_price := strings.Replace(removed_euro, ",", "", -1)
			price_results, err := strconv.Atoi(clean_price)
			if err != nil {
				fmt.Errorf("Error converting to int%s", err)
			}
			price = price_results
		}
	})

	c.OnHTML("h1[data-testid=address]", func(h *colly.HTMLElement) {
		location = h.Text
	})

	c.OnHTML("p[data-testid=beds]", func(h *colly.HTMLElement) {
		beds_temp, err := strconv.Atoi(strings.Split(h.Text, " ")[0])
		if err != nil {
			fmt.Errorf("Error converting to in%s", err)
		}
		beds = beds_temp
	})

	c.OnHTML("p[data-testid=baths]", func(h *colly.HTMLElement) {
		baths_temp, err := strconv.Atoi(strings.Split(h.Text, " ")[0])
		if err != nil {
			fmt.Errorf("Error converting to in%s", err)
		}
		baths = baths_temp
	})

	c.OnHTML("p[data-testid=property-type]", func(h *colly.HTMLElement) {
		type_ = strings.Split(h.Text, " ")[0]
	})
	c.OnHTML("div[data-testid=statistics]", func(h *colly.HTMLElement) {
		stats := h.ChildText("p")
		stats_split := strings.Split(stats, "Entered/Renewed")
		views_split := strings.Split(stats_split[1], "Property Views")
		views_temp, err := strconv.Atoi(strings.Replace(views_split[0], ",", "", -1))
		if err != nil {
			fmt.Errorf("Error converting to int%s", err)
		}
		views = views_temp
		entered = stats_split[0]
	})

	for _, l := range links {

		url_to_visit := fmt.Sprintf("https://www.daft.ie%s", l)
		c.Visit(url_to_visit)

		url = url_to_visit
		split_url := strings.Split(url_to_visit, "/")
		temp_id, err := strconv.Atoi(split_url[len(split_url)-1])
		if err != nil {
			fmt.Errorf("Error converting to int%s", err)
		}
		id = temp_id

		property := models.Property{}
		property.Price = price
		property.URL = url
		property.Location = location
		property.DaftId = id
		property.Bed = beds
		property.Bathroom = baths
		property.Type = type_
		property.Entered = entered
		property.Views = views

		properties = append(properties, property)
	}
	return properties
}

func scrape_properties() []string {
	from_value := 0
	c := colly.NewCollector()
	// On every a element which has href attribute call callback
	links := []string{}

	c.OnHTML("li a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if strings.Contains(link, "/for-rent/") {
			links = append(links, link)
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	for i := from_value; i < 1000; i += 20 {
		page_to_visit := fmt.Sprintf("https://www.daft.ie/property-for-rent/ireland?pageSize=20&from=%d", i)
		c.Visit(page_to_visit)
	}
	return links
}
