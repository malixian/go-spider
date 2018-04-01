package main

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"log"
	"go-spider/model"
)

func main(){
	item := model.Profile{
		"梦醒时分",
		"33岁",
		"162CM",
		"3001-5000元",
		"离异",
		"大专",
		"河北沧州",
	}
	client, err := elastic.NewSimpleClient()
	c := context.Background()
	if err != nil{
		panic("save error")
	}
	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").BodyJson(item).Do(c)
	if err != nil{
		panic("save error")
	}
	log.Printf("save success %v", resp)
}

