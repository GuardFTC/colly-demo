// Package main @Author:冯铁城 [17615007230@163.com] 2025-10-20 16:50:19
package main

import (
	"colly-demo/example/cryptocoins"
	"colly-demo/study"
)

func main() {

	//1.知识点学习测试
	studyTest()

	//2.例子测试
	//exampleTest()
}

// studyTest 知识点学习测试
func studyTest() {

	////1.基础demo测试
	//study.BaseDemoTest()
	//
	////2.异常测试
	//study.ErrorTest()
	//
	////3.登录测试
	//study.LoginTest()
	//
	////4.最大深度测试
	//study.MaxDepthTest()
	//
	////5.并发爬取测试
	//study.Parallel()
	//
	////6.代理测试
	//study.ProxyTest()
	//
	////7.队列设置
	//study.QueueTest()
	//
	////8.随机延迟测试
	//study.RandomDelayTest()
	//
	////9.redis存储测试
	//study.RedisStorageTest()
	//
	////10.请求上下文测试
	//study.RequestCtxTest()

	//11.回调测试
	study.CallbackTest()
}

// exampleTest 示例测试
func exampleTest() {

	//1.获取数字货币市场容量
	cryptocoins.TestGetCryptocoinsData()
}
