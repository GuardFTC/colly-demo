// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 17:12:16
package study

import (
	"encoding/json"
	"log"

	"github.com/gocolly/colly"
)

// LoginTest 登录测试
func LoginTest() {

	//1.创建采集器
	c := colly.NewCollector()

	//2.声明用户名密码
	userAuth := map[string]string{
		"username": "admin",
		"password": "yuanjing@1234",
	}

	//3.将数据转换为 JSON 格式
	jsonData, err := json.Marshal(userAuth)
	if err != nil {
		log.Printf("error marshaling JSON: %v", err)
		return
	}

	//4.设置请求头为 JSON
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json")
	})

	//5.授权访问，使用 PostRaw 发送 JSON 数据
	err = c.PostRaw("http://localhost:3000/api/user/login", jsonData)
	if err != nil {
		log.Printf("error posting JSON: %v", err)
	}

	//6.登录请求之后的回调
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	//7.访问网站
	if err = c.Visit("http://localhost:3000/"); err != nil {
		log.Println("error visiting url:", err)
	}
}
