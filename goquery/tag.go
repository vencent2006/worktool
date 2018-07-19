package main

import (
	"github.com/opesun/goquery"
	"fmt"
	"flag"
	"os"
	"strings"
)


var n = flag.Int("n",1,"number of page")
var s = flag.String("s","nothing","info")

func main() {

	if nil == os.Args || len(os.Args) < 2 {
		fmt.Println("参数错误：请输入url")
		os.Exit(-1)
	}

	//var url = "https://m.ximalaya.com/28251596/album/9832179"
	var url = os.Args[1]
	if ( !strings.ContainsAny(url, "http") && !strings.ContainsAny(url, "https")) {
		//加上http
		url = "http://" + url
		//fmt.Println("补全后的url是",url)
	}
	p, err := goquery.ParseUrl(url)
	if err != nil {
		panic(err)
	} else {
		title := p.Find("title").Text()//直接提取title的内容
		fmt.Println("   title：",title)

		keywords := p.Find("meta[name=keywords]").Attr("content")
		fmt.Println("keywords：",keywords)
	}
}

