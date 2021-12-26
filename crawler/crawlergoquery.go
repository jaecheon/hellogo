package gocrawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func RunQuery(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("img").Each(func(index int, el *goquery.Selection){
		imgSrc, exists := el.Attr("src")
		if exists {
			fmt.Println(imgSrc)
		}
	})
}