package plat

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"github.com/yueqingkong/IP-Proxy-Go/util"
	"log"
	"net/http"
)

type ProxyListPlus struct {
}

// page 6
func (self *ProxyListPlus) Create() {
	for i := 1; i <= 6; i++ {
		url := fmt.Sprintf("https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-%d", i)
		self.Parser(url)
	}
}

func (self *ProxyListPlus) Parser(url string) {
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

	doc.Find("tbody>tr.cells").Each(func(i int, sec *goquery.Selection) {
		ip := sec.Find("td").Next().First().Text()
		port := sec.Find("td").Next().First().Next().First().Text()

		address := fmt.Sprintf("%s:%s", ip, port)
		log.Printf(address)

		reg := "(((25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))):\\d+"
		if util.IsMatch(address, reg) {
			code, speed := ProxyCheck(fmt.Sprintf("http://%s:%s", ip, port))
			if code == http.StatusOK {
				log.Printf("%s %s", ip, port)

				db.XORM().InsertIP("ProxyListPlus", address, speed)
			}
		}
	})
}
