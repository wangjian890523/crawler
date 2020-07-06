package parser

import (
	"fmt"
	"testing"

	"github.com/wangjian890523/crawler/fetcher"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch(
		"http://www.zhenai.com/zhenghun")
	//TestParseCityList()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", contents)
	//ParseCityList(Contents)

	//verify result
}
