package persist

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/engine"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {

	out := make(chan interface{})

	go func() {

		itemCount := 0

		for {
			item := <-out

			log.Printf("ItemSaver >> Got item #%d: %v", itemCount, item)
			itemCount++

			//	存储数据
			_, err := save(item)

			if err != nil {

				log.Print("Item Saver: error"+"saving item %v: %v", item, err)
			}
		}

	}()

	return out
}

//	将数据存入ElasticSearch
func save(item interface{}) (id string, err error) {

	//	ElasticSearch client
	client, err := elastic.NewClient(
		//	Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {

		return "", nil
	}
	//	存 Index -> 可创建，可修改 读 Get 找 Search
	response, err := client.Index().
		Index(engine.DataBaseName). //	数据库名
		Type(engine.TableName).     //	表名 可不设ID，由系统分配
		BodyJson(item).             //	内容
		Do(context.Background())

	if err != nil {

		return "", nil
	}

	return response.Id, nil

}
