package parser

import (
	"regexp"
	"strconv"
)

const ageRe = `<td><span class="label">年龄:</span>([\d]+)岁</td>`
const marriageRe = `<td><span class="label">婚况:</span>([^<]+)岁</td>`

func ParseProfile(contents []byte) engine.PareseResult {
	profile := modle.Profile{}
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)

	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			profile.Age = age
		}
	}

	re = regexp.MustCompile(marriageRe)
	match = re.FindSubmatch(contents)
	if match != nil {
		profile.Marriage = string(match[1])
	}

}

func extractString(contents []byte, re string) {

}
