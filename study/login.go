// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-20 17:12:16
package study

import (
	"encoding/json"
	"log"

	"github.com/gocolly/colly"
)

// Response 响应结构体
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *Data  `json:"data"`
}

// Data 数据结构体
type Data struct {
	Privileges []string `json:"privileges"`
	Nickname   string   `json:"nickname"`
	Avatar     string   `json:"avatar"`
	UserID     int      `json:"userId"`
	GodMode    bool     `json:"godMode"`
	Token      string   `json:"token"`
}

// LoginTest 登录测试
func LoginTest() {

	//1.创建采集器
	c := colly.NewCollector()

	//2.登录获取Token
	token, err := login(c)
	if err != nil {
		log.Printf("login error:%v", err)
		return
	}

	//3.请求头设置Token
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("authorization", "Mon-Token "+token)
	})

	//4.响应回调
	c.OnResponse(func(r *colly.Response) {
		log.Printf("status:%d, body:%s", r.StatusCode, string(r.Body))
	})

	//5.声明请求参数,将数据转换为 JSON 格式
	jsonData, err := parsePayload(map[string]string{
		"pageNum":  "1",
		"pageSize": "15",
	})
	if err != nil {
		return
	}

	//6.访问网站
	if err = c.PostRaw("http://localhost:3000/api/project/list", jsonData); err != nil {
		log.Println("error visiting url:", err)
	}
}

func login(c *colly.Collector) (string, error) {

	//1.设置请求头为 JSON
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json")
	})

	//2.定义响应结构体
	var resp Response

	//3.登录请求之后的回调
	c.OnResponse(func(r *colly.Response) {

		//4.如果响应异常，返回
		if r.StatusCode != 200 {
			log.Println("response error")
			return
		}

		//5.解析响应体
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			log.Println("error parsing JSON:", err)
			return
		}
	})

	//6.声明用户名密码,将数据转换为 JSON 格式
	jsonData, err := parsePayload(map[string]string{
		"username": "admin",
		"password": "yuanjing@1234",
	})
	if err != nil {
		return "", err
	}

	//7.授权访问，使用 PostRaw 发送 JSON 数据
	if err = c.PostRaw("http://localhost:3000/api/user/login", jsonData); err != nil {
		return "", err
	}

	//8.判定是否登录成功
	if resp.Code != 0 {
		return "", err
	} else {
		return resp.Data.Token, nil
	}
}

// parsePayload 参数转换
func parsePayload(pagePayload map[string]string) ([]byte, error) {
	if jsonData, err := json.Marshal(pagePayload); err != nil {
		return nil, err
	} else {
		return jsonData, nil
	}
}
