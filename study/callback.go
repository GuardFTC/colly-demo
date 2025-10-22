// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-22 14:34:57
package study

import (
	"log"

	"github.com/gocolly/colly"
)

// CallbackTest 回调测试
func CallbackTest() {

	//1.创建采集器
	c := colly.NewCollector()

	//2.发起请求之前触发
	c.OnRequest(func(r *colly.Request) {
		log.Printf("on request is trigger. url is %v", r.URL)
	})

	//3.请求失败触发
	c.OnError(func(_ *colly.Response, err error) {
		log.Printf("on error is trigger. err is %v", err)
	})

	//4.收到响应时触发
	c.OnResponse(func(r *colly.Response) {
		log.Printf("on response is trigger. response is %+v", r)
	})

	//5.收到响应为匹配规则的HTML标签时触发
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		log.Printf("on html is trigger. element is %v", e.Text)
	})

	//6.收到响应为匹配规则的XML标签时触发
	c.OnXML("*", func(e *colly.XMLElement) {
		log.Printf("on xml is trigger. element is %v", e.Text)
	})

	//7.请求结束时触发
	c.OnScraped(func(r *colly.Response) {
		log.Printf("on scraped is trigger. request is finish")
	})

	//8.访问网站
	log.Println("-------------------------------------visit html web------------------------------------")
	if err := c.Visit("https://jsonplaceholder.typicode.com/"); err != nil {
		log.Printf("error visit url. err is %v", err)
	}

	//9.访问xml网站
	log.Println("-------------------------------------visit xml web------------------------------------")
	if err := c.Visit("https://www.w3schools.com/xml/note.xml"); err != nil {
		log.Printf("error visit url. err is %v", err)
	}

	//10.访问JSON响应网站
	log.Println("-------------------------------------visit json web------------------------------------")
	if err := c.Visit("https://jsonplaceholder.typicode.com/posts/1"); err != nil {
		log.Printf("error visit url. err is %v", err)
	}

	//11.访问异常网站
	log.Println("-------------------------------------visit error web------------------------------------")
	if err := c.Visit("123123123"); err != nil {
	}
}
