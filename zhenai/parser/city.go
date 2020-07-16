package parser

import (
	"fmt"
	"regexp"

	"github.com/wangjian890523/crawler/engine"
)

var(
	profileRe =regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+i)"[^>]*>([^<]+)</a>`)
	cityUrlRe =regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+"`)

	)
//const cityRe = `<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`


func ParseCity(contents []byte) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})

		limit--
		if limit == 0 {
			break
		}
		fmt.Printf("City:%s, URL:%s\n", m[2], m[1])
	}

	//fmt.Printf("match found:%d\n", len(matches))

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _,m:= range  matches{
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: ParseCity,

		})

	}

	return result
}
