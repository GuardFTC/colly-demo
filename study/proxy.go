// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 19:22:14
package study

import (
	"bytes"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

// ProxyTest 代理测试
func ProxyTest() {

	//1.创建采集器
	c := colly.NewCollector(

		//2.设置允许访问重复URL
		colly.AllowURLRevisit(),
	)

	//3.设置代理服务器协议、IP、端口
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	if err != nil {
		log.Fatal(err)
	}

	//4.采集器代理设置
	c.SetProxyFunc(rp)

	//5.响应回调
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	//6.循环发送请求
	for i := 0; i < 5; i++ {
		if err := c.Visit("https://httpbin.org/ip"); err != nil {
			log.Printf("visit error: %v", err)
		}
	}
}
