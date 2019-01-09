package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stevenkitter/skills/services/crawl/server"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://udemycoupon.discountsglobal.com/coupon-category/health-fitness-course-coupons/"
	hf := &server.HealthFitness{Url: url}
	ch := make(chan string, 10)
	createWorker(ch)
	tasks := make(chan string, 10)
	go hf.RequestData(tasks)
	for {
		task, ok := <-tasks
		if !ok {
			log.Println("close")
			time.Sleep(50 * time.Second)
			return
		}
		ch <- task
	}
}

func createWorker(ch chan string) {
	for i := 0; i < 10; i++ {
		go worker(ch)
	}
}
func worker(ch chan string) {
	for {
		task := <-ch
		process(task)
	}
}
func process(task string) {
	log.Printf("task is %s", task)
	udemy(task)
}
func udemy(url string) {
	time.Sleep(1 * time.Second)
	agent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	rawCookies := `ud_firstvisit=2019-01-07T08:17:42.673634+00:00:1ggQ6R:9cWhkGOKLcmfVzBENEnv3FI9GLM; __udmy_2_v57r=34e20053a32b48229dcd01004fce1a84; _ga=GA1.2.651006873.1546849067; _ga=GA1.1.651006873.1546849067; IR_gbd=udemy.com; _pxvid=b76d3090-1254-11e9-a096-5580bd30fa81; _gid=GA1.2.2122561052.1546999698; _pxff_ed=1; ud_last_auth_information="{\"backend\": \"udemy-auth\"\054 \"suggested_user_name\": \"Wang\"\054 \"suggested_user_avatar\": \"https://udemy-images.udemy.com/user/50x50/anonymous_3.png\"\054 \"suggested_user_email\": \"coderxuer@gmail.com\"}:1gh3I5:Q410Zauun96KcckhFQKMY18DwS0"; dj_session_id=r7ggy17x91m906go94bc2z4grvg7terh; access_token=wkGo5nTG4Z0VP2RnOOKVThF15A4dP9lWGiwIxzag; ud_credit_last_seen=None; client_id=bd2565cb7b0c313f5e9bae44961e8db2; csrftoken=YkPyBrxbYehZnSvycTLXfdTbjNlUlUfwnKpIZsIv5syMPdb3NOAicrnoUMTlJNkB; _gid=GA1.1.2122561052.1546999698; volume=0.5; mute=0; playbackspeed=1; intercom-session-sehj53dd=Y3dENVBQNTYyTTYweSszeit0TllvOUEyTWx1SUZtR3NrUXRaMkowMnFqNjlvQm1hTjIxbjl4SXYxM0JDL3pCNS0tL3g3ZkE0aGFneSt2N29ZZENmc1NBQT09--307ccee3cd708ccd07fa37c62aa3a7ce82cc28b5; _td=340fad4f-cf2e-42e5-8837-7ef06e2cc319; ud_cache_user=da185c2cbd750dfc; ud_credit_unseen=0; ud_cache_price_country=1c2903397d883338; ud_cache_version=c4ca4238a0b92382; ud_cache_brand=fece4154b1d5d2be; ud_cache_release=89f87ceeb38a6e12; seen=1; _fbp=fb.1.1547021666142.1704085515; _aw_m_6554=6554_1547023005_ec7383de31995113a7693fba66834cd0; ki_r=; evi="SlFPIh0SV3ZKURp2X1BXdkpRXSBMWEF/AEsHYxhAT31MDgdjCAdXdkpRGnlcUFd2SlFdIExYQXoBRwdjGEBPfUwOB2MMEAIgWFEROkxTQ3kJURE6TBQUbgtHGnZaTlc6E0kaPBNOVy9dAwl7FUBEewJHCXsVQAMtE0kfclZSWW5HURF1Ex9ZbkIECXsVQER0AkUJexVAAy0TSR93XVRZbkdREXATH1luWR4JexVARHsHRQl7FUADLRNJH3VYUFluR1ERcxMfCA=="; stc111655=env:1547023017%7C20190209083657%7C20190109101637%7C11%7C1014625:20200109094637|uid:1546849068995.1101974414.5419202.111655.2128772225:20200109094637|srchist:1014625%3A1%3A20190207081749%7C1014624%3A1546999701%3A20190209020821%7C1014625%3A1547010805%3A20190209051325%7C1014624%3A1547021663%3A20190209081423%7C1014625%3A1547023017%3A20190209083657:20200109094637|tsa:-68716542:20190109101637; ud_cache_campaign_code=7815696ecbf1c96e; exaff=%7B%22start_date%22%3A%20%222019-01-09T08%3A36%3A48.375973Z%22%2C%20%22code%22%3A%20null%2C%20%22aff_type%22%3A%20%22AW%22%2C%20%22aff_id%22%3A%20124%7D:1ghARS:19RSQHbdtJtDpmCjsIXf64XeU_w; ud_rule_vars="eJx9js0KwjAQhF-l5KqW3fy0Tc6-gXgrlJiuEqwE09WD4rtbrIJ48LQwM9_O3AX7fCCmvrvGMXLKTmmSAEZ5JXe6kdL2oQcE0PtA6BvtQkrHSMIV4t6KwY_cZTpfaLq9Z2onoxUS0K4AV1AX0DisnValURbRLgAcQCuWUyr4zN0Yb2_oW-XEfnjLxXaznq1XW0iXPNK8t-N4-q20BVinKyersjaoUX9X7mOeXszsv7myrCqLxnzYh3g8ASOJVNE=:1ghARS:4_WHtI2HtBOY-w5pmYhwOPOrZds"; IR_5420=1547027198598%7C0%7C1547021664370%7C%7C; IR_PI=7163a4f1-13b3-11e9-8100-06157567e462%7C1547113598598; ki_t=1546999703037%3B1546999703037%3B1547027199935%3B1%3B197; _px2=eyJ1IjoiMjhjZmRlNTAtMTNlOS0xMWU5LWJlNDQtZmQxNzYyMDgxMjcwIiwidiI6ImI3NmQzMDkwLTEyNTQtMTFlOS1hMDk2LTU1ODBiZDMwZmE4MSIsInQiOjE1NDcwMjgxMDA4MDAsImgiOiJiMTU2ODc4NTg3NTUwYzVkNzhiN2YzMzM0NzNlZDJkY2Q3ZDM5N2QxYzNmYjA4OTFlYWZmY2JiMWMwOWM5ZGJhIn0=; _px3=4bc6be830de7ac76b7c1c53891b09ba20426a6a3d077c07f0b3668b539f27132:pprA8mhse3aKu6rWBLgWCq2b6KIKtB70ca4GOLCKWa/HwJ+nLBl/y7wL/gavSZ2nshBS5zRXJzfHhe4lQ3VyaQ==:1000:B8VCP7nqsEJ79lxHfRdxVHAIF4bHC6uhn6NK+MleR4FOcbbwOpQWMJ2m1iQDiMUJ5z0jmuiEOT5rnZ93GnOUBV1iB3ey7BCZ2xfQ5nt2hXbTCNkb+p8Pua7BMw94c3xIg5VS0U55QCdGyXPpxcdxZOAVqGQv906TNX6KQjEq0dI=`
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Cookie", rawCookies)
	request.Header.Set("User-Agent", agent)
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	if err != nil {
		log.Fatal(err)
	}
	cli := http.Client{}
	res, err := cli.Do(request)
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

	doc.Find(".ud-component--clp--price-text ").Each(func(i int, s *goquery.Selection) {
		log.Println("exist ")
		log.Printf("exist %v", s)
	})

}
