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

	// 定时器轮训
	ticker := time.NewTicker(1 * time.Minute)
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
	// 定时器轮训
	ticker := time.NewTicker(1 * time.Minute)
	select {
	case <-ticker.C:
		log.Println("轮训 删除ip")

		ips := db.XORM().AllIps()
		for i := 0; i < len(ips); i++ {
			ip := ips[i]

			code, speed := plat.ProxyCheck(fmt.Sprintf("http://%s", ip.Address))
			if code != http.StatusOK {
				if ip.ErrorCount >= 2 {
					log.Print("[删除] ", ip.Address)
					db.XORM().Delete(ip.Id)
				} else {
					t := time.Now()

					ip.Timestamp = t.Unix()
					ip.CreateTime = t
					ip.ErrorCount = ip.ErrorCount + 1
					db.XORM().UpdateIp(ip)
				}
			} else {
				t := time.Now()

				ip.Speed = speed
				ip.Timestamp = t.Unix()
				ip.CreateTime = t
				db.XORM().UpdateIp(ip)
			}
		}
	}
}

func routers() {
	router.HttpServer()
}
