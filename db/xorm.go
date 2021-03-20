package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yueqingkong/IP-Proxy-Go/conf"
	"log"
	"time"
)

var engine *xorm.Engine
var xORM *XOrm

type XOrm struct {
}

func Load() {
	var err error

	// 获取mysql配置
	sourceName := conf.DBConnectURI()
	log.Print(sourceName)

	engine, err = xorm.NewEngine("mysql", sourceName)
	if err != nil {
		log.Fatal("[MySql] 连接失败,", err)
	}

	engine.ShowSQL(false)
	err = engine.Sync2(new(IP))
	if err != nil {
		log.Fatal("[MySql] 同步表失败", err)
	}
}

func XORM() *XOrm {
	if xORM == nil {
		xORM = new(XOrm)
	}
	return xORM
}

func (self *XOrm) Insert(i interface{}) {
	_, err := engine.Insert(i)
	if err != nil {
		log.Print("[Insert]", err, i)
	}
}

func (self *XOrm) InsertIP(plat string, address string, speed int32) {
	time := time.Now()
	ip := &IP{
		Plat:       plat,
		Address:    address,
		Speed:      speed,
		Timestamp:  time.Unix(),
		CreateTime: time,
	}

	_, err := engine.Insert(ip)
	if err != nil {
		log.Print("[InsertIP]", err, ip)
	}
}

func (self *XOrm) GetIPOne() (*IP, error) {
	ip := &IP{}
	_, err := engine.Where("id > ?", 0).Limit(1).Get(ip)
	if err != nil {
		log.Print(err)
	}

	return ip, err
}

func (self *XOrm) GetIPById(id int64) (*IP, error) {
	ip := &IP{}
	_, err := engine.Where("id > ? and speed < 5000", id).Limit(1).Get(ip)
	if err != nil {
		log.Print(err)
	}

	return ip, err
}

func (self *XOrm) GetIP(address string) (*IP, error) {
	ip := &IP{}
	_, err := engine.Where("address = ?", address).Get(ip)
	if err != nil {
		log.Print(err)
	}

	return ip, err
}

func (self *XOrm) AllIps() []IP {
	var ips []IP
	err := engine.Where("id > ?", 0).Find(&ips)
	if err != nil {
		log.Print(err)
	}

	return ips
}

func (self *XOrm) UpdateIp(ip IP) {
	_, err := engine.Id(ip.Id).Update(ip)
	if err != nil {
		log.Print(err)
	}
}

func (self *XOrm) Delete(id int64) {
	_, err := engine.Id(id).Delete(&IP{})
	if err != nil {
		log.Print(err)
	}
}

func (self *XOrm) CountIps() int32 {
	total, err := engine.Where("id >?", 0).Count(&IP{})
	if err != nil {
		return 0
	}

	return int32(total)
}
