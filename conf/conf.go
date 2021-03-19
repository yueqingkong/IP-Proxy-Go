package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Server SERVER `yaml:"server"`
	Db     DB     `yaml:"db"`
}

type SERVER struct {
	Port string `yaml:"port"`
}

type DB struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
	User string `yaml:"user"`
	PWD  string `yaml:"pwd"`
}

func init() {

}

var Cf *Config

func LoadYml(path string) {
	var fileBytes []byte
	var err error

	if fileBytes, err = ioutil.ReadFile(path); err != nil {
		log.Fatalf("ReadFile %s", path)
	}

	if err = yaml.Unmarshal(fileBytes, &Cf); err != nil {
		log.Fatalf("Unmarshal %s", path)
	}

	log.Print(Cf)
}

func Port() string {
	return Cf.Server.Port
}

// DBConnectURI 数据库连接字符串
func DBConnectURI() string {
	db := Cf.Db

	name := db.Name
	user := db.User
	host := db.Host
	port := db.Port
	password := db.PWD

	// sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s", user, password, host, port, name, url.QueryEscape("Asia/Shanghai"))
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", user, password, host, port, name)
	return sourceName
}
