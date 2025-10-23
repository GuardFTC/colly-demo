// Package douban @Author:冯铁城 [17615007230@163.com] 2025-10-23 16:04:30
package douban

import (
	"strings"

	"github.com/gocolly/colly"
)

// Book 图书信息结构体
type Book struct {
	CoverImg    string `json:"cover_img"`    // 图书封面
	Title       string `json:"title"`        // 图书标题
	Author      string `json:"author"`       // 作者
	Publisher   string `json:"publisher"`    // 出版社
	PublishTime string `json:"publish_time"` // 出版时间
	Price       string `json:"price"`        // 单价
	Rating      string `json:"rating"`       // 评分
	RatingCount string `json:"rating_count"` // 评价数
	Description string `json:"description"`  // 描述
	EbookLink   string `json:"ebook_link"`   // 电子版链接
}

// newBook 创建图书结构体实例
func newBook(el *colly.HTMLElement) Book {

	//1.创建图书结构体实例
	book := Book{}

	//2.获取图书封面
	book.CoverImg = el.ChildAttr("div[class='pic'] > a[class='nbg'] > img", "src")

	//3.获取图书标题
	book.Title = el.ChildText("div[class='info'] > h2 > a")

	//4.获取作者、出版社、出版时间、单价
	infoText := el.ChildText("div[class='info'] > div[class='pub']")
	infos := strings.Split(infoText, "/")
	if len(infos) >= 4 {
		book.Author = strings.TrimSpace(infos[0])
		book.Publisher = strings.TrimSpace(infos[1])
		book.PublishTime = strings.TrimSpace(infos[2])
		book.Price = strings.TrimSpace(infos[3])
	}

	//5.获取评价数以及评分
	book.Rating = el.ChildText("div[class='info'] > div[class='star clearfix'] > span[class='rating_nums']")
	book.RatingCount = el.ChildText("div[class='info'] > div[class='star clearfix'] > span[class='pl']")

	//6.获取描述
	book.Description = el.ChildText("div[class='info'] > p")

	//7.获取电子版链接
	book.EbookLink = el.ChildAttr("div[class='info'] > div[class='ft'] > div[class='ebook-link'] > a", "href")

	//8.返回图书
	return book
}
