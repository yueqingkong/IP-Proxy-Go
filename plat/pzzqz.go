package plat

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"log"
	"net/http"
)

// https://pzzqz.com/
type Pzzqz struct {
}

func (self *Pzzqz) Create() {
	url := "https://pzzqz.com"
	self.Parser(url)
}

func (self *Pzzqz) Parser(url string) {
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

		address := fmt.Sprintf("%s:%s", ip, port)
		log.Printf(address)

		code, speed := ProxyCheck(fmt.Sprintf("http://%s:%s", ip, port))
		if code == http.StatusOK {
			log.Printf("%s %s", ip, port)

			db.XORM().InsertIP("Pzzqz", address, speed)
		}
	})
}
