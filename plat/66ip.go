package plat

import (
	"fmt"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"github.com/yueqingkong/IP-Proxy-Go/util"
	"io/ioutil"
	"log"
	"net/http"
)

type IP66 struct{}

func (self *IP66) Create() {
	url := fmt.Sprintf("http://www.66ip.cn/mo.php?tqsl=%d", 100)
	self.Parser(url)
}

func (self *IP66) Parser(url string) {
	contentType := "application/x-www-form-urlencode"
	resp, err := http.Post(url, contentType, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	content := string(body)

	reg := "(((25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))):\\d+"
	arr := util.FindAllString(content, reg)
	for _, v := range arr {
		address := fmt.Sprintf("http://%s", v)
		code, speed := ProxyCheck(address)
		if code == http.StatusOK {
			log.Print(speed, v)

			db.XORM().InsertIP("IP66", v, speed)
		}
	}
}
