package parser

import (
	"regexp"
	"strconv"

	"github.com/wangjian890523/crawler/modle"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄:</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况:</span>([^<]+)岁</td>`)

func ParseProfile(contents []byte, name string) engine.PareseResult {
	profile := modle.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age

	}
	profile.Marriage = extractString(contents, marriageRe)

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match = re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}