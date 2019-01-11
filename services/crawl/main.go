package main

import (
	"github.com/stevenkitter/skills/services/crawl/server"
)

func main() {
	url := "http://udemycoupon.discountsglobal.com/coupon-category/health-fitness-course-coupons/page/444/"
	//url := "https://www.udemy.com/freestyle-101/?couponCode=OPENDOORS&awc=6554_1547094467_aa3a839aee984b655e3204f80de52965&utm_source=Growth-Affiliate&utm_medium=Affiliate-Window&utm_campaign=Campaign-Name&utm_term=13430&utm_content=Placement"
	e := server.Engine{
		Scheduler:   &server.QueuedScheduler{},
		WorkerCount: 5,
	}
	e.Run(server.Request{
		URL:       url,
		Fetcher:   server.DiscountFetcher,
		ParseFunc: server.DiscountParse,
	})
}
