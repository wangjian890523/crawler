package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimter = time.Tick(10 * time.Millisecond )


//获取URL, 对应的item,和城市的url
func Fetch(url string) ([]byte, error) {
	<-rateLimter
	resp, err := http.Get(url)
	if err != nil {
		//panic(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error:status code ", resp.StatusCode)
		return nil,
			fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	// 如果是gbk 编码转换成utf8
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		//panic(err)
		log.Printf("Fetch error:%d", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
