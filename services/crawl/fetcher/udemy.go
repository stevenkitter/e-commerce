package fetcher

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var limit = time.Tick(time.Duration(rand.Intn(20)) * time.Millisecond)

func UdemyFetcher(path string) (*goquery.Document, error) {
	<-limit
	if strings.Contains(path, "awin1.com") {
		u, err := url.Parse(path)
		if err != nil {
			log.Printf("url.Parse err:%v", err)
		}
		dic, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			log.Printf("url.ParseQuery err:%v", err)
		}
		path = dic["p"][0]
	}
	agent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	cookie := `ud_firstvisit=2019-01-07T08:17:42.673634+00:00:1ggQ6R:9cWhkGOKLcmfVzBENEnv3FI9GLM; __udmy_2_v57r=34e20053a32b48229dcd01004fce1a84; _ga=GA1.2.651006873.1546849067; _ga=GA1.1.651006873.1546849067; IR_gbd=udemy.com; _pxvid=b76d3090-1254-11e9-a096-5580bd30fa81; _gid=GA1.2.2122561052.1546999698; ud_last_auth_information="{\"backend\": \"udemy-auth\"\054 \"suggested_user_name\": \"Wang\"\054 \"suggested_user_avatar\": \"https://udemy-images.udemy.com/user/50x50/anonymous_3.png\"\054 \"suggested_user_email\": \"coderxuer@gmail.com\"}:1gh3I5:Q410Zauun96KcckhFQKMY18DwS0"; dj_session_id=r7ggy17x91m906go94bc2z4grvg7terh; access_token=wkGo5nTG4Z0VP2RnOOKVThF15A4dP9lWGiwIxzag; ud_credit_last_seen=None; client_id=bd2565cb7b0c313f5e9bae44961e8db2; csrftoken=YkPyBrxbYehZnSvycTLXfdTbjNlUlUfwnKpIZsIv5syMPdb3NOAicrnoUMTlJNkB; _gid=GA1.1.2122561052.1546999698; volume=0.5; mute=0; playbackspeed=1; intercom-session-sehj53dd=Y3dENVBQNTYyTTYweSszeit0TllvOUEyTWx1SUZtR3NrUXRaMkowMnFqNjlvQm1hTjIxbjl4SXYxM0JDL3pCNS0tL3g3ZkE0aGFneSt2N29ZZENmc1NBQT09--307ccee3cd708ccd07fa37c62aa3a7ce82cc28b5; _fbp=fb.1.1547091261456.1952102127; _pxff_ed=1; ud_cache_brand=fece4154b1d5d2be; ud_credit_unseen=0; ud_cache_user=da185c2cbd750dfc; ud_cache_version=c4ca4238a0b92382; ud_cache_release=a623fbfcd1c5377d; ud_cache_price_country=1c2903397d883338; seen=1; _td=340fad4f-cf2e-42e5-8837-7ef06e2cc319; _pxff_tm=1; _gat=1; _gat_instructor=1; _gat_UA-12366301-1=1; _dc_gtm_UA-12366301-29=1; IR_5420=1547105839773%7C0%7C1547101032874%7C%7C; IR_PI=7163a4f1-13b3-11e9-8100-06157567e462%7C1547192239773; stc111655=env:1547105788%7C20190210073628%7C20190110080720%7C2%7C1014625:20200110073720|uid:1546849068995.1101974414.5419202.111655.2128772225:20200110073720|srchist:1014625%3A1%3A20190207081749%7C1014624%3A1546999701%3A20190209020821%7C1014625%3A1547010805%3A20190209051325%7C1014624%3A1547021663%3A20190209081423%7C1014625%3A1547023017%3A20190209083657%7C1014624%3A1547091259%3A20190210033419%7C1014625%3A1547091427%3A20190210033707%7C1014624%3A1547101032%3A20190210061712%7C1014625%3A1547105788%3A20190210073628:20200110073720|tsa:805708541:20190110080720; _aw_m_6554=6554_1547105826_bb01edeec929dbb963578985a9f20f09; ud_cache_campaign_code=7815696ecbf1c96e; ki_t=1546999703037%3B1547091289792%3B1547105841417%3B2%3B244; ki_r=; _px2=eyJ1IjoiOGY3Zjk2MTAtMTRhYS0xMWU5LThhMzUtMmYwY2RmYWIyOTQxIiwidiI6ImI3NmQzMDkwLTEyNTQtMTFlOS1hMDk2LTU1ODBiZDMwZmE4MSIsInQiOjE1NDcxMDYzMzA4MDIsImgiOiJiYzA3YTBkOGY4OTM4NGUxNTU2MzhiYmEyMWQ5ODIxZTY4MWEyZmVmOTdhYzZkMzJiMjg2ZTgwYWJkNjU4N2Q1In0=; _px3=ee2e96547984c1d245abac470c7d38a24d1478aded05fbe202c8be496f16ed84:StRfkphlPGAmfbkFk0rILTnHKFDa4nYTSaFQxualZTZgXNUMek16Fb5F1p3A5j/nS7c+MTgX3NO3Efr7FweGiw==:1000:T+YarjQwaNYOS4uIiJ2+WsFgOAd8Wtwg86s6mmSqhgLhO7ViE/zRDEH0YoeiwejNsz9zpXiflV7yKXklyLUV6GavYWYN985FQM9O0wy2qV97knAi6SpWQZhoz5E+25+gfdT20h96bGkv8zkYban++rO7uBXSU5yH5OBTMv6QmK4=; exaff=%7B%22start_date%22%3A%20%222019-01-10T07%3A37%3A07.170911Z%22%2C%20%22code%22%3A%20null%2C%20%22aff_type%22%3A%20%22AW%22%2C%20%22aff_id%22%3A%20124%7D:1ghUts:ghSMG6tB0I-8rwZ2cxr8md_D-LA; ud_rule_vars="eJx9j8EKwjAMhl9l9KqOpO3WtWffQLwVRu2iFJViFz0ovrtj7CAePP2Q8H3J_xIcyomYhv6RxsS5OKVJAjQqKHnQnZR2iAMggD5GwtBpF3M-JxKuEi8vYijcj-lJfhp4AV6sp7iEkftCtztNOQRethLQbgA3YCroHBqnVd0oi2hXAA4W-JjKRM3v_GVl3bYWm-abnQ_HfC8jLQZO1x8DQgXGKePA1KpFhd23YS7EmcNlaVTtd1sv3uL9AbNMVMY=:1ghUts:6dkC7nDBB92jCrLhsGXWKCgq2To"; evi="SlFILR5ATzcTQhxyWkBPNxMFSmNUVkZ0AV8JN0xYQTFMXwklDREFbgsICXBZU0duCwgJNw9AT3gCQhNtTBRXdgAOVm1MEQJuCwgJcFZRQ24LCAk3D0BPeAdAHW1MFFd2AA5WbUwKGG4LCAlwWVRDbgsICTcPQE94BUUZbUwUV3YDDlZtTAAHO10aCXsVQER6BEsJexVAAy0TSR9wWVZZbkdREXATH1luVxYJexVARHQDQQl7FUADLRNJH3deVlluR1ERcBMfCA=="`
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Printf("http.NewRequest err:%v", err)
		return nil, err
	}
	request.Header.Set("Cookie", cookie)
	request.Header.Set("User-Agent", agent)
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	cli := http.Client{}
	res, err := cli.Do(request)
	if err != nil {
		log.Printf("cli.Do(request) err:%v", err)
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("res.Body.Close err:%v", err)
		}
	}()
	if res.StatusCode != 200 {
		log.Printf("res.StatusCode:%d", res.StatusCode)
		return nil, errors.New(fmt.Sprintf("status is wrong url is %s", path))
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Printf("goquery.NewDocumentFromReader:%v", err)
	}
	return doc, err
}
