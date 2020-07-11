package parser

import (
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄:</span>([\d]+)岁</td>`)
var marriageRe = regexpp.MustCompile(`<td><span class="label">婚况:</span>([^<]+)岁</td>`)

func ParseProfile(contents []byte) engine.PareseResult {
	profile := modle.Profile{}

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
