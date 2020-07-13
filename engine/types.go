package engine


type ParserFunc func(
	contents []byte, url string) ParseResult


type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}


type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct{}


func (NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}
/*
func (NilParser) Serialize() (
	name string, args interface{}) {
	return config.NilParser, nil
}
*/
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(
	contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (
	name string, args interface{}) {
	return f.name, nil
}


func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
