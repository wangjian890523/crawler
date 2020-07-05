package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"
const text1 = `My email is ccmouse@gmail.com
email is abs@def.com
email2 is kkk@qq.com
email3 is ddd@abc.com.cn`

func main() {
	//re := regexp.MustCompile("ccmouse@gmail.com")
	re := regexp.MustCompile(`([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+\.([a-zA-Z0-9.]+)+`)
	//match := re.FindString(text)
	//match := re.FindAllString(text1, -1)
	match := re.FindAllStringSubmatch(text1, -1)
	for _, m := range match {
		fmt.Println(m)

	}

	//fmt.Println(match)

}
