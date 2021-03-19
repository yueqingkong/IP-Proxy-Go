package plat

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"log"
	"net/http"
)

// http://www.ip3366.net

type IP3366 struct {

}

// page 10
func (self *IP3366) Create() {
	for i := 1; i <= 7; i++ {
		for j := 1; j <= 2; j++ { // 国内高匿 国内普通
			url := fmt.Sprintf("http://www.ip3366.net/free/?stype=%d&page=%d", j, i)
			self.Parser(url)
		}
	}
}

func (self *IP3366) Parser(url string) {
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

	doc.Find("tbody>tr").Each(func(i int, sec *goquery.Selection) {
		ip := sec.Find("td").First().Text()
		port := sec.Find("td").Next().First().Text()

		address:= fmt.Sprintf("%s:%s", ip, port)

		code, speed := ProxyCheck(fmt.Sprintf("http://%s:%s", ip, port))
		if code == http.StatusOK {
			log.Printf("%s %s", ip, port)

			db.XORM().InsertIP("IP3366", address, speed)
		}
	})
}