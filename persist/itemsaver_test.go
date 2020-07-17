package persist

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/wangjian890523/crawler/engine"
	"github.com/wangjian890523/crawler/model"
	"golang.org/x/net/context"
	"testing"
)

func TestItemServer(t *testing.T) {

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	err := save(expected)
	if err !=nil{
		panic(err)

	}
	client,err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err !=nil{
		panic(err)
	}
	resp,err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err!=nil{
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal([]byte (resp.Source), &actual)
	if err !=nil{
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual !=expected{
		t.Errorf("got %v; expected %v", actual, expected)
	}
}