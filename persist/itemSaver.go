package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"go-spider/engine"
	"github.com/pkg/errors"
)

func ItemSaver(client *elastic.Client, c context.Context)chan engine.Item{
	saveChan := make(chan engine.Item)
	go func() {
		count := 0
		for{
			item := <- saveChan
			count++
			log.Printf("save item is %d, %v", count,item)
			id, err := saveToES(client, c, item)
			if err != nil{
				log.Printf("Item Saver error %v, %v", id, err)
			}
		}
	}()
	return saveChan
}


func saveToES(client *elastic.Client, c context.Context, item engine.Item)(id string , err error){

	if err != nil{
		return "",err
	}
	if item.Type == ""{
		return "",errors.New("must supply Type")
	}
	resp, err := client.Index().
		         Index("dating_profile").
			     Type(item.Type).Id(item.Id).BodyJson(item).Do(c)
	if err != nil{
		return "",err
	}
	return resp.Id, nil
}

