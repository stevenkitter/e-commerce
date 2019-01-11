package server

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/stevenkitter/skills/services/crawl/fetcher"
	"log"
	"os"
	"strings"
)

type Request struct {
	URL       string
	Fetcher   func(string) (*goquery.Document, error)
	ParseFunc func(string, *goquery.Document) Result
}

type Result struct {
	Requests []Request
	Items    []interface{}
}

func DiscountParse(url string, doc *goquery.Document) Result {
	result := Result{}
	doc.Find(".content-box:nth-child(2) .item-frame.row .couponAndTip .link-holder a[href]").Each(func(i int, s *goquery.Selection) {
		hr, ok := s.Attr("href")
		if !ok {
			fmt.Printf("DiscountParse not exist url is %s", url)
		}
		destUrl := hr
		result.Items = append(result.Items, destUrl)
		result.Requests = append(result.Requests, Request{
			URL:       destUrl,
			Fetcher:   fetcher.UdemyFetcher,
			ParseFunc: UdemyParse,
		})
	})

	next, exit := doc.Find(".content-box:nth-child(2) .next.page-numbers").Attr("href")
	if exit {
		result.Requests = append(result.Requests, Request{
			URL:       next,
			Fetcher:   DiscountFetcher,
			ParseFunc: DiscountParse,
		})
	} else {
		log.Printf("finished")
	}
	return result
}

func UdemyParse(url string, doc *goquery.Document) Result {
	result := Result{}
	s := doc.Find(".price-text__current[data-purpose=discount-price-text]")
	v := s.Text()
	if strings.Contains(v, "免费") {
		res := fmt.Sprintf("免费可用的链接是:%s\n", url)
		log.Println(res)
		file, err := os.OpenFile("./freeUrl.txt", os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("os.Open() err:%v", err)
		}
		n, _ := file.Seek(0, 2)
		_, err = file.WriteAt([]byte(res), n)
		if err != nil {
			log.Printf("file.WriteAt err:%v", err)
		}
	}
	return result
}
