// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-21 10:30:48
package study

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func RandomDelayTest() {

	//1.创建采集器，设置异步访问,开启debug模式
	c := colly.NewCollector(
		// Attach a debugger to the collector
		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
	)

	//2.设置采集器速率限制
	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",    //域名匹配规则
		Parallelism: 2,               //并发数
		RandomDelay: 5 * time.Second, //随机延迟最大seconds
	}); err != nil {
		log.Printf("error setting limit: %v", err)
	}

	//3.定义URL
	url := "https://httpbin.org/delay/2"

	//4.发送原始请求
	if err := c.Visit(url); err != nil {
		log.Printf("error visiting %s: %v", url, err)
	}

	//5.发送4个请求
	for i := 0; i < 4; i++ {
		if err := c.Visit(fmt.Sprintf("%s?n=%d", url, i)); err != nil {
			log.Printf("error visiting %s: %v", url, err)
		}
	}

	//6.等待所有请求完成
	c.Wait()
}
