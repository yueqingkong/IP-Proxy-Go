package main

import (
	"flag"
	"fmt"
	"github.com/yueqingkong/IP-Proxy-Go/conf"
	"github.com/yueqingkong/IP-Proxy-Go/db"
	"github.com/yueqingkong/IP-Proxy-Go/plat"
	"github.com/yueqingkong/IP-Proxy-Go/router"
	"log"
	"net/http"
	"time"
)

var (
	configPath = flag.String("configpath", "./conf/conf.yml", "configpath")
)

func init() {
	flag.Parse()

	conf.LoadYml(*configPath)
	db.Load()
}

func main() {
	log.Print("[IP Prxoy] start...")

	go Crawl()
	go ConnectCheck()
	routers()
}

func Crawl() {
	plats := []plat.Plat{&plat.IP66{}, &plat.FanQieIP{},
		&plat.IP3366{}, &plat.KuaiDaiLi{},
		&plat.ProxyListPlus{}, &plat.Pzzqz{}}

	for _, p := range plats {
		go p.Create()
	}

	// 定时器轮训
	ticker := time.NewTicker(30 * time.Minute)
	select {
	case <-ticker.C:
		log.Print("轮循 ip")
		for _, p := range plats {
			go p.Create()
		}
		break
	}
}

func ConnectCheck() {
	start := int64(0)

	// 定时器轮训
	ticker := time.NewTicker(30 * time.Minute)
	select {
	case <-ticker.C:
		ips := db.XORM().AllIps(start)
		if len(ips) == 0 {
			start = 0
		} else {
			for i := 0; i < len(ips); i++ {
				ip := ips[i]
				if i == len(ips)-1 {
					start = ip.Id
				}

				code, _ := plat.ProxyCheck(fmt.Sprintf("http://%s", ip.Address))
				if code != http.StatusOK {
					if ip.ErrorCount >= 5 {
						db.XORM().Delete(ip.Id)
					} else {
						ip.ErrorCount = ip.ErrorCount + 1
						db.XORM().UpdateIp(ip)
					}
				}
			}
		}
		break
	}
}

func routers() {
	router.HttpServer()
}
