package plat

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"github.com/yueqingkong/IP-Proxy-Go/util"
	"log"
	"net/http"
)

// https://www.my-proxy.com/free-proxy-list-2.html
type MyProxy struct {
}

func (self *MyProxy) Create() {
	for i := 1; i <= 7; i++ {
		url := fmt.Sprintf("https://www.my-proxy.com/free-proxy-list-%d.html", i)
		self.Parser(url)
	}
}

func (self *MyProxy) Parser(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Print(resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Print(err)
	}

	doc.Find("div.list").Each(func(i int, sec *goquery.Selection) {
		content := sec.Text()

		reg := "(((25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))):\\d+"
		arr := util.FindAllString(content, reg)
		for _, v := range arr {
			address := fmt.Sprintf("%s http://%s", url, v)

			code, speed := ProxyCheck(address)
			if code == http.StatusOK {
				log.Print(speed, v)

				db.XORM().InsertIP("MyProxy", address, speed)
			}
		}
	})
}
