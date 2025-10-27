// Package study @Author:冯铁城 [17615007230@163.com] 2025-10-27 11:34:06
package study

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

// OnHtmlTest HTML响应回调测试
func OnHtmlTest() {

	//1.定义访问地址
	var url = "https://book.douban.com/tag/小说"

	//2.创建采集器
	c := colly.NewCollector()

	//3.测试Attr方法
	testAttr(c)

	//4.测试ChildAttr/ChildAttrs/ChildText方法
	testChildAttrAndChildAttrsAndChildText(c)

	//5.测试ForEach和Unmarshal方法
	testForEachAndUnmarshal(c)

	//6.测试各种选择器
	testSelector(c)

	//7.测试各种过滤器
	testFilter(c)

	//8.访问地址
	if err := c.Visit(url); err != nil {
		log.Printf("visit error:%v", err)
	}
}

// testAttr 测试Attr方法
func testAttr(c *colly.Collector) {

	//1.定义HTML回调解析方法
	c.OnHTML("div[class='nav-logo']>a", func(e *colly.HTMLElement) {

		//2.打印控制台分隔
		log.Printf("Attr方法测试:---------------------------------")

		//3.获取当前标签属性值
		attr := e.Attr("href")
		log.Printf("获取当前标签属性值:%v", attr)
		log.Println()
	})
}

// testChildAttrAndChildAttrsAndChildText 测试ChildAttr/ChildAttrs/ChildText方法
func testChildAttrAndChildAttrsAndChildText(c *colly.Collector) {

	//1.定义HTML回调解析方法
	c.OnHTML("div[class='paginator']", func(e *colly.HTMLElement) {

		//2.打印控制台分隔
		log.Printf("ChildAttr/ChildAttrs/ChildText方法测试:---------------------------------")

		//3.获取第一个子标签的属性值
		attr := e.ChildAttr("a", "href")
		log.Printf("获取第一个子标签的属性值:%v", attr)

		//4.获取所有子标签的属性值
		attrs := e.ChildAttrs("a", "href")
		log.Printf("获取所有子标签的属性值:%v", attrs)

		//5.获取所有子标签的文本
		text := e.ChildText("a")
		log.Printf("获取所有子标签的文本:%v", text)
		log.Println()
	})
}

// testForEachAndUnmarshal 测试ForEach和Unmarshal方法
func testForEachAndUnmarshal(c *colly.Collector) {

	//1.定义HTML回调解析方法
	c.OnHTML("ul[class='subject-list']", func(e *colly.HTMLElement) {

		//2.打印控制台分隔
		log.Printf("ForEach和Unmarshal方法测试:---------------------------------")

		//3.遍历li标签
		log.Printf("遍历li标签:")
		e.ForEach("li", func(i int, el *colly.HTMLElement) {

			//4.创建图书结构体实例
			var book Book

			//5.解析HTML内容到结构体
			if err := el.Unmarshal(&book); err != nil {
				log.Printf("unmarshal error:%v", err)
			} else {

				//6.填充图书信息相关字段
				infos := strings.Split(book.InfoText, "/")
				if len(infos) == 4 {
					book.Author = strings.TrimSpace(infos[0])
					book.Publisher = strings.TrimSpace(infos[1])
					book.PublishTime = strings.TrimSpace(infos[2])
					book.Price = strings.TrimSpace(infos[3])
				} else if len(infos) > 4 {
					book.Author = strings.TrimSpace(infos[0])
					book.Publisher = strings.TrimSpace(infos[1]) + "/" + strings.TrimSpace(infos[2])
					book.PublishTime = strings.TrimSpace(infos[3])
					book.Price = strings.TrimSpace(infos[4])
				}

				//7.打印图书信息
				log.Printf("反序列化图书信息:%+v", book)
			}
		})
		log.Println()
	})
}

// testSelector 测试各种选择器
func testSelector(c *colly.Collector) {

	//1.定义HTML回调解析方法
	c.OnHTML("body", func(e *colly.HTMLElement) {

		//2.打印控制台分隔
		log.Printf("选择器测试:---------------------------------")

		//3.ID选择器
		attr := e.ChildAttr("#subject_list", "id")
		log.Printf("id选择器获取id属性:%v", attr)

		//4.类选择器
		attr = e.ChildAttr("ul[class='subject-list']", "class")
		log.Printf("类选择器获取class属性-方式1:%v", attr)
		e.ChildAttr(".subject-list]", "class")
		log.Printf("类选择器获取class属性-方式2:%v", attr)

		//5.父子选择器
		text := e.ChildText("div[class='paginator'] > span[class='prev']")
		log.Printf("父子选择器获取第一个子标签的文本:%v", text)

		//6.相邻选择器
		attr = e.ChildAttr("div[class='paginator'] > span[class='thispage'] + a", "href")
		log.Printf("相邻选择器获取第一个相邻标签的属性:%v", attr)

		//7.兄弟选择器
		text = e.ChildText("div[class='paginator'] > span[class='thispage'] ~ span[class='break']")
		log.Printf("兄弟选择器获取兄弟标签的文本:%v", text)

		//8.同时选中多个
		text = e.ChildText("div[class='paginator'] > span[class='thispage'],div[class='paginator'] > span[class='break']")
		log.Printf("同时选中多个获取标签的文本:%v", text)
		log.Println()
	})
}

// testFilter 过滤器测试
func testFilter(c *colly.Collector) {

	//1.定义HTML回调解析方法
	c.OnHTML("body", func(e *colly.HTMLElement) {

		//2.打印控制台分隔
		log.Printf("过滤器测试:---------------------------------")

		//3.获取第一个子元素，不区分类型（当类型不匹配时，该值为空）
		attr := e.ChildAttr("div[class='paginator'] > a:first-child", "href")
		log.Printf("获取第一个子元素属性，不区分类型:%v", attr)

		//4.获取第一个子元素，区分类型
		attr = e.ChildAttr("div[class='paginator'] > a:first-of-type", "href")
		log.Printf("获取第一个子元素属性，区分类型:%v", attr)

		//5.获取最后一个元素，不区分类型（当类型不匹配时，该值为空）
		lastText := e.ChildText("div[class='paginator'] > a:last-child")
		log.Printf("获取最后一个子元素文本，不区分类型:%v", lastText)

		//6.获取最后一个子元素，区分类型
		lastAttr := e.ChildAttr("div[class='paginator'] > a:last-of-type", "href")
		log.Printf("获取最后一个子元素属性，区分类型:%v", lastAttr)

		//7.顺序获取第n个子元素，不区分类型（当类型不匹配时，该值为空）
		attr = e.ChildAttr("div[class='paginator'] > a:nth-child(1)", "href")
		log.Printf("顺序获取第n个子元素属性，不区分类型:%v", attr)

		//8.顺序获取第n个子元素，区分类型
		attr = e.ChildAttr("div[class='paginator'] > a:nth-of-type(1)", "href")
		log.Printf("顺序获取第n个子元素属性，区分类型:%v", attr)

		//9.倒序获取第n个子元素，不区分类型（当类型不匹配时，该值为空）
		lastText = e.ChildText("div[class='paginator'] > a:nth-last-child(1)")
		log.Printf("倒序获取第n个子元素文本，不区分类型:%v", lastText)

		//10.倒序获取第n个子元素，区分类型
		lastText = e.ChildText("div[class='paginator'] > a:nth-last-of-type(1)")
		log.Printf("倒序获取第n个子元素文本，区分类型:%v", lastText)

		//11.获取只有一个子元素，不区分类型（当类型不匹配时，该值为空）
		attr = e.ChildAttr("div[class='nav-logo'] > p:only-child", "href")
		log.Printf("获取只有一个子元素属性，不区分类型:%v", attr)

		//12.获取只有一个子元素，区分类型
		attr = e.ChildAttr("div[class='nav-logo'] > a:only-of-type", "href")
		log.Printf("获取只有一个子元素属性，区分类型:%v", attr)

		//13.获取文本包含 豆瓣 的a标签元素
		attrs := e.ChildAttrs("a:contains(豆瓣)", "href")
		log.Printf("获取文本包含 豆瓣 的a标签元素:%v", attrs)
		log.Println()
	})
}

// Book 图书信息结构体
type Book struct {
	CoverImg    string `json:"cover_img" selector:"div.pic > a.nbg > img"  attr:"src"` // 图书封面
	Title       string `json:"title" selector:"div.info > h2 > a"`                     // 图书标题
	Author      string `json:"author"`                                                 // 作者
	Publisher   string `json:"publisher"`                                              // 出版社
	PublishTime string `json:"publish_time"`                                           // 出版时间
	Price       string `json:"price"`
	Rating      string `json:"rating" selector:"div.info > div.star.clearfix > span.rating_nums"`        // 评分
	RatingCount string `json:"rating_count" selector:"div.info > div.star.clearfix > span.pl"`           // 评价数
	Description string `json:"description" selector:"div.info > p"`                                      // 描述
	EbookLink   string `json:"ebook_link" selector:"div.info > div.ft > div.ebook-link > a" attr:"href"` // 电子版链接
	InfoText    string `json:"-" selector:"div.info > div.pub"`                                          // 关键：只提取一次
}
