package main

import (
	"github.com/wangjian890523/crawler/engine"
	"github.com/wangjian890523/crawler/scheduler"
	"github.com/wangjian890523/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCity,
	})

}

/*func tMain() {*/

//reponse, err := http.Get("http://www.zhenai.com/zhenghun")
//if err != nil {
//panic(err)
//}
//defer reponse.Body.Close()

//if reponse.StatusCode != http.StatusOK {
//fmt.Println("Err: status cpde",
//reponse.StatusCode)
//return
//}
////	all, err := ioutil.ReadAll(reponse.Body)
////	if err != nil {
////		panic(err)
////	}
////fmt.Println(all)

//e := determineEncoding(reponse.Body)
//utf8Reader := transform.NewReader(reponse.Body, e.NewDecoder())
//all, err := ioutil.ReadAll(utf8Reader)
//if err != nil {
//panic(err)
//}
////fmt.Printf("%s\n", all)
//PrintCityList(all)

/*}*/

/*func determineEncoding(r io.Reader) encoding.Encoding {*/
//bytes, err := bufio.NewReader(r).Peek(1024)
//if err != nil {
//panic(err)
//}
//e, _, _ := charset.DetermineEncoding(bytes, "")
//return e
/*}*/

//const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

/*func PrintCityList(contents []byte) {*/
//re := regexp.MustCompile(cityListRe)
//matches := re.FindAllSubmatch(contents, -1)

//result := engine.ParseResult{}
//for _, m := range matches {
//result.Items = append(result.Items, m[2])
//result.Requests = append(result.Requests, engine.Requestst{
//Url:        string(m[1]),
//ParserFunc: engine.NilParser,
//})

////fmt.Printf("City:%s, URL:%s\n", m[2], m[1])
//}

//fmt.Printf("match found:%d\n", len(matches))

/*}*/
