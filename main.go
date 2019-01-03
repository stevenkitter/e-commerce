package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func main() {
	for {
		sess := sessionId()
		log.Printf("sessionId %s", sess)
		vote(sess)
	}
	//vote("5d991fd1fd3805e46b3ea112c3265a30")
	//requClient()
}

func vote(sessID string) {
	url := "http://dp.sngzzc.com/app/index.php?i=7&c=entry&rid=822632&id=38367&do=vote&m=tyzm_diamondvote"
	agent := "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16C101 MicroMessenger/7.0.1(0x17000120) NetType/WIFI Language/zh_CN"
	postData := map[string]interface{}{
		"latitude":  0,
		"longitude": 0,
		"verify":    0,
	}
	jsonValue, _ := json.Marshal(postData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("http.NewRequest err:%v", err)
	}
	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	req.Header.Set("User-Agent", agent)
	//"5d991fd1fd3805e46b3ea112c3265a30"
	req.Header.Set("Cookie", "PHPSESSID="+sessID+"; Hm_lpvt_08c6f5e17c0761a968c5658ccf6ff5ad=1546498961; Hm_lvt_08c6f5e17c0761a968c5658ccf6ff5ad=1546472301")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err:%v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll err:%v", err)
	}

	dict := map[string]interface{}{}
	err = json.Unmarshal(body, &dict)
	if err != nil {
		return
		//log.Printf("json.Unmarshal err:%v",err)
	}
	if dict["status"].(string) == "1" {
		log.Printf("dict is %v", dict)
		log.Printf("sessionId is %s", sessID)
	}

}

//appId ：wxeaa826389ad94106
//nonceStr p489Hw4LQK4lL6S4
//timestamp 1546503020
//signature 70928606aebff35d0a1c783c21e339bf13288c6d
//http://dp.sngzzc.com/app/index.php?i=7&c=entry&id=39461&rid=822632&do=view&m=tyzm_diamondvote&from=groupmessage

func requClient() {
	url := "http://dp.sngzzc.com/app/index.php?i=7&c=entry&id=39461&rid=822632&do=view&m=tyzm_diamondvote&from=groupmessage"
	host := "dp.sngzzc.com"
	userAgent := "Mozilla/5.0 (Linux; Android 8.0; MI 6 Build/OPR1.170623.027; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/6.2 TBS/044455 Mobile Safari/537.36 MicroMessenger/6.6.6.1300(0x26060638) NetType/WIFI Language/zh_CN"
	acc := "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,image/wxpic,image/sharpp,image/apng,image/tpg,*/*;q=0.8"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("err:%v", err)
	}
	req.Header.Set("Accept", acc)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Host", host)
	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	req.Header.Set("Location", "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxeaa826389ad94106&redirect_uri=http%3A%2F%2Fdp.sngzzc.com%2Fapp%2Findex.php%3Fi%3D7%26c%3Dauth%26a%3Doauth%26scope%3Dsnsapi_base&response_type=code&scope=snsapi_base&state=we7sid-5f0e18cadd9d668179d70255ced0ce79#wechat_redirect")
	req.Header.Set("Set-Cookie", "PHPSESSID=5f0e18cadd9d668179d70255ced0ce79; path=/; HttpOnly")

	cl := http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		log.Printf("client.Do err:%v", err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll err:%v", err)
	}
	//fmt.Println(string(body))
}

func sessionId() string {
	str := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,0,1,2,3,4,5,6,7,8,9"
	arr := strings.Split(str, ",")

	res := ""
	for i := 0; i < 32; i++ {

		rd := rand.Intn(len(arr))

		res += arr[rd]
	}
	return res
}

//id	String	333805
//tid	String	38351
//ptid	String	201901031340067784066766
//rid	String	822632
//uniacid	String	7
//uniontid	String	2019010313400000002062684462
//paytype	String
//oauth_openid	String	oNju2wYc_TAl67G9WcJLAILwx1iI
//openid	String	oNju2wYc_TAl67G9WcJLAILwx1iI
//avatar	String	http://thirdwx.qlogo.cn/mmopen/vi_32/dQZH4oEsCUia0gn7FD1CgT3rNIj8piah2tGOE5qqiar1Rx6RAyJ4Ibe3QwN9UmpQ5BicD9bedtRVMogRgqTdw01tgQ/132
//nickname	String	王进
//user_ip	String	180.126.43.65
//gifttitle	String	豪华邮轮
//giftcount	String	1
//gifticon	String	images/7/2018/10/i4m6Kw83J89qC94Wkgq5NkQq7V6948.jpg
//fee	String	100.00
//giftvote	String	300
//ispay	String	1
//isdeal	String	1
//status	String	0
//createtime	String	2019-01-03 13:40:00
//cont	String	王进，给TA送了1份豪华邮轮！
