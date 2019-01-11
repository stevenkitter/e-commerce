package server

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func DiscountFetcher(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	log.Printf("start url is %s", url)
	if err != nil {
		log.Printf("http.Get(url):%v", err)
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("res.Body.Close err:%v", err)
		}
	}()
	if res.StatusCode != 200 {
		log.Printf("res.StatusCode:%v", res.StatusCode)
		return nil, errors.New(fmt.Sprintf("status is wrong url is %s", url))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("goquery.NewDocumentFromReader:%v", err)
	}
	return doc, err
}
