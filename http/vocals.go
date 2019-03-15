package main

import (
	"fmt"
	"log"
	"strings"

	"git.tutorabc.com/tmc2/stark/utils"
	"golang.org/x/net/html"
)

func parseHTMLToStringArray(s string) (resultArr []string, err error) {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
		return resultArr, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					resultArr = append(resultArr, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return resultArr, err
}

func main() {

	baseURL := "http://internal.tutormeetplus.com/v2/dataware/data/talks/"

	b, err := utils.HTTPGet(baseURL, nil)
	if err != nil {
		panic(err)
	}

	dateArr, err := parseHTMLToStringArray(string(b))
	if err != nil {
		panic(err)
	}

	// 列日期
	for _, v := range dateArr {
		if v == "detail/" {
			continue
		}
		b, err := utils.HTTPGet(baseURL+v, nil)
		if err != nil {
			fmt.Println("请求错误:", v, err)
			continue
		}

		resultSessionIDs, err := parseHTMLToStringArray(string(b))
		if err != nil {
			fmt.Println("解析错误:", v, err)
			continue
		}

		// 遍历所有房间号.txt
		for _, sessionIDTxT := range resultSessionIDs {
			b, err := utils.HTTPGet(baseURL+v+sessionIDTxT, nil)
			if err != nil {
				fmt.Println("请求错误:", v, err)
				continue
			}

			lines := strings.Split(string(b), "\n")
			for _, l := range lines {
				items := strings.Split(l, "\t")

				if len(items) < 5 {
					continue
				}

				userID := items[0]
				sessionID := items[1]
				isTeacher := items[2]
				total := items[4]

				fmt.Println(userID, sessionID, isTeacher, total)
			}
		}
	}
}
