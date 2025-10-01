package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main(){
	fmt.Println("This will work to take scraping data and return json to client")
	collector := colly.NewCollector(colly.AllowedDomains("www.scrapingcourse.com"),)


	

	collector.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting: ", r.URL)
    })

	collector.OnResponse(func(r *colly.Response){
		fmt.Println("Page Visited: ", r.Request.URL )
	})

	collector.Visit("https://www.scrapingcourse.com/ecommerce")
}
