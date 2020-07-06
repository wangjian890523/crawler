package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//获取URL, 对应的item,和城市的url
func Fetch(url string) ([]byte, error) {
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
	e := determineEncoding(resp.Body)
	utf8reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		//panic(err)
		log.Printf("Fetch error:%d", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
