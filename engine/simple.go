package engine

import (
	"log"

	"github.com/wangjian890523/crawler/fetcher"
)


type SimpleEngine struct {}
func (SimpleEngine)Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//log.Printf("Fetching %s", r.Url)
		//body, err := fetcher.Fetch(r.Url)
		//if err != nil {
		//	log.Printf("Fetcher:error"+r.Url, err)
		//	continue
		//}
		//parseResult := r.ParseFunc(body)

		parseResult, err := worker(r)
		if err != nil{
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {

			log.Printf("Got item %v", item)
		}
	}
}


func worker(r Request) (ParseResult, error){
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error"+r.Url, err)
		return   ParseResult{},err
	}

	return r.ParseFunc(body),nil
}