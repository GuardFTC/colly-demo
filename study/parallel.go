// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 17:51:34
package study

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// Parallel 并发测试
func Parallel() {

	//1.创建采集器
	c := colly.NewCollector(

		//2.设置爬取最大深度为2
		colly.MaxDepth(2),

		//3.设置允许异步请求
		colly.Async(true),

		//4.允许重复访问同一个URL
		colly.AllowURLRevisit(),
	)

	//5.设置并发限制
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*", //域名匹配规则
		Parallelism: 4,   //最大并发数=2
	})
	if err != nil {
		log.Printf("error setting limit: %v", err)
	}

	//6.当响应为HTML，并且匹配a[href]标签时，执行回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//7.获取链接
		link := e.Attr("href")

		//8.打印链接
		fmt.Println(link)

		//9.过滤异常URL
		if link == "#" {
			return
		}

		//10.访问链接
		err = c.Visit(e.Request.AbsoluteURL(link))
		if err != nil {
			log.Printf("error visiting %s: %v", link, err)
		}
	})

	//10.范围wiki
	if err = c.Visit("https://en.wikipedia.org/"); err != nil {
		log.Printf("error visiting wikipedia: %v", err)
	}

	//11.等待所有协程爬取完毕
	c.Wait()
	log.Printf("done")
}
