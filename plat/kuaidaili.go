package plat

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"log"
	"net/http"
)

// https://www.kuaidaili.com
type KuaiDaiLi struct {
}

// page 10
func (self *KuaiDaiLi) Create() {
	styles := []string{"inha", "intr"} // 国内高匿 国内普通
	for i := 1; i <= 10; i++ {
		for j := 0; j < 2; j++ {
			url := fmt.Sprintf("https://www.kuaidaili.com/free/%s/%d", styles[j], i)
			self.Parser(url)
		}
	}
}

func (self *KuaiDaiLi) Parser(url string) {
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
		log.Printf(address)

		code, speed:= ProxyCheck(fmt.Sprintf("http://%s:%s", ip, port))
		if code == http.StatusOK {
			log.Printf("%s %s", ip, port)

			db.XORM().InsertIP("KuaiDaiLi", address, speed)
		}
	})
}
