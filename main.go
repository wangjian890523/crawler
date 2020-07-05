package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	reponse, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != http.StatusOK {
		fmt.Println("Err: status cpde",
			reponse.StatusCode)
		return
	}
	//	all, err := ioutil.ReadAll(reponse.Body)
	//	if err != nil {
	//		panic(err)
	//	}
	//fmt.Println(all)

	e := determineEncoding(reponse.Body)
	utf8Reader := transform.NewReader(reponse.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
	PrintCityList(all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func PrintCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="httpL//www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]</a>`)
	matches := re.FindAll(contents, -1)

	for _, m := range matches {
		fmt.Println("%s\n", m)
	}

	fmt.Printf("match found:%d\n", len(matches))

}
