package plat

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"log"
	"net/http"
	"strings"
)

// https://www.fanqieip.com/free/1
type FanQieIP struct {
}

// page max 1661
func (self *FanQieIP) Create() {
	for i := 1; i <= 1661; i++ {
		url := fmt.Sprintf("https://www.fanqieip.com/free/%d", i)
		self.Parser(url)
	}
}

func (self *FanQieIP) Parser(url string) {
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

		ip = strings.TrimSpace(ip)
		port = strings.TrimSpace(port)
		address:= fmt.Sprintf("%s:%s", ip, port)

		log.Printf(address)

		code, speed := ProxyCheck(fmt.Sprintf("http://%s:%s", ip, port))
		if code == http.StatusOK {
			log.Printf("%s %s", ip, port)
			db.XORM().InsertIP("FanQieIP", address, speed)
		}
	})
}
