// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-21 10:45:43
package study

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/gocolly/redisstorage"
)

// RedisStorageTest redis存储测试
func RedisStorageTest() {

	//1.创建redis存储
	//主要功能包含如下两点
	//(1).记录和查询已访问的 URL
	//(2).存储和管理 Cookies
	storage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379", //redis地址
		Password: "",               //redis密码
		DB:       0,                //redis数据库
		Prefix:   "httpbin_test",   //redisKey前缀
	}

	//2.创建采集器
	c := colly.NewCollector()
	c.Async = true

	//3.设置存储
	err := c.SetStorage(storage)
	if err != nil {
		panic(err)
	}

	//4.每次执行之前，清空redis数据
	if err = storage.Clear(); err != nil {
		log.Fatal(err)
	}

	//5.最后关闭redis连接
	defer storage.Client.Close()

	//6.设置响应回调
	c.OnResponse(func(r *colly.Response) {
		log.Println("Cookies:", c.Cookies(r.Request.URL.String()))
	})

	//7.定义URL集合
	urls := []string{
		"http://httpbin.org/",
		"http://httpbin.org/ip",
		"http://httpbin.org/cookies/set?a=b&c=d",
		"http://httpbin.org/cookies",
	}

	//8.创建队列，并发请求数量为2，并设置存储
	q, _ := queue.New(2, storage)

	//9.将请求URL压入队列
	for _, u := range urls {
		q.AddURL(u)
	}

	//10.开始请求
	q.Run(c)

	//11.等待请求完成
	c.Wait()
}
