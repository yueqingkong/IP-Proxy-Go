package db

import "time"

type IP struct {
	Id         int64
	Plat       string    `xorm:"varchar(255) unique(name,symbol,plat,section,create_time)"` // 平台名称
	Address    string    `xorm:"varchar(255) unique(name,symbol,plat,section,create_time)"` // 地址 xxx.xx.x.x:port
	Status     int32     `xorm:"int"`                                                       // 0: 没验证过 1: 验证通过
	Speed      int32     `xorm:"int"`                                                       // 网络连接速度
	ErrorCount int32     `xorm:"int"`                                                       // 验证失败次数
	Timestamp  int64     `xorm:"bigint"`                                                    // 最后一次check时间戳
	CreateTime time.Time `xorm:"DATETIME unique(name,symbol,plat,section,create_time))"`    // 最后一次check时间
}
