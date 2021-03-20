package plat

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

type Plat interface {
	Create()
	Parser(url string)
}

// 测试链接ip速度
// http://124.205.155.151:9090
func ProxyCheck(address string) (statusCode int, speed int32) {
	// 测试地址 百度
	testUrl := "http://www.baidu.com"

	// 解析代理地址
	proxy, err := url.Parse(address)
	// 设置网络传输
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	// 创建连接客户端
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	begin := time.Now()

	reqest, err := http.NewRequest("GET", testUrl, nil)
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Referer", testUrl)
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	// 使用代理ip访问测试地址
	res, err := httpClient.Do(reqest)
	if err != nil {
		// log.Print(err)
		return 0, http.StatusBadRequest
	}
	defer res.Body.Close()

	speedReal := int32(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000)
	// 判断是否访问成功 成功的code 200
	if res.StatusCode != http.StatusOK {
		log.Print(res.StatusCode)
	}

	return res.StatusCode, speedReal
}
