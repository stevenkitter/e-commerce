package server

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type HealthFitness struct {
	Url string
}

func (s *HealthFitness) RequestData(ch chan string) {
	res, err := http.Get(s.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".content-box:nth-child(2) .item-frame.row .couponAndTip .link-holder a[href]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		hr, ok := s.Attr("href")
		if !ok {
			fmt.Println("not exist")
		}
		ch <- hr
	})
}
