package parser

import (
	"regexp"

	"github.com/wangjian890523/crawler/engine"
)

const cityRe = `<a href="(http://www.album.zhenhun.com/u/[0-9]+i)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: engine.NilParser,
		})

		//fmt.Printf("City:%s, URL:%s\n", m[2], m[1])
	}

	//fmt.Printf("match found:%d\n", len(matches))
	return result
}
