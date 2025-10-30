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

	//3.定义代理服务器IP以及端口
	proxies := []string{
		"socks5://127.0.0.1:1337",
		"socks5://127.0.0.1:1338",
	}

	//4.设置代理服务器轮询器
	rp, err := proxy.RoundRobinProxySwitcher(proxies...)
	if err != nil {
		log.Fatal(err)
	}

	//5.采集器代理设置
	c.SetProxyFunc(rp)

	//6.响应回调
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	//7.循环发送请求
	for i := 0; i < 5; i++ {
		if err := c.Visit("https://httpbin.org/ip"); err != nil {
			log.Printf("visit error: %v", err)
		}
	}
}
