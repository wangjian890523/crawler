package persist

import (
	"github.com/kataras/iris/core/errors"
	"github.com/lunny/log"
	"github.com/olivere/elastic/v7"
	"github.com/wangjian890523/crawler/engine"
	"golang.org/x/net/context"
)

var host="http://192.168.176.138:9200/"

func ItemServer(index string) (chan engine.Item, error){
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err !=nil{
		return nil,err
	}


	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Server: got item "+"#%d:%v", itemCount, item)
			itemCount++
			 err :=save(client,index,item)
			if err!=nil{
				log.Printf("Item server:error"+"Saving item%v:%v",
					item, err)

			}
		}
	}()
	return out, nil
}

func save( client  *elastic.Client,index string, item engine.Item)  error {

	if item.Type ==""{
		return errors.New("Must suply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != ""{
		indexService.Id(item.Id)
	}


	_, err := indexService.
		Do(context.Background())
	if err != nil{
		return  err
	}

	return nil

	}