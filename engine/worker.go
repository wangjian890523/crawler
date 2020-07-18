package engine

import (
	"github.com/lunny/log"
	"github.com/wangjian890523/crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error"+"frtching url %s: %v",r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(body, r.Url), nil
}
