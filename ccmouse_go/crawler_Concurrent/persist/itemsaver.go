package persist

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan engine.Item {

	out := make(chan engine.Item)

	go func() {

		itemCount := 0

		for {
			item := <-out

			log.Printf("ItemSaver >> Got item #%d: %v", itemCount, item)
			itemCount++

			//	存储数据
			err := save(item)

			if err != nil {

				log.Print("Item Saver: error"+"saving item %v: %v", item, err)
			}
		}

	}()

	return out
}

//	将数据存入ElasticSearch
func save(item engine.Item) error {

	//	ElasticSearch client
	client, err := elastic.NewClient(
		//	Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {

		return err
	}

	//	检测是否传入表名
	if item.Type == "" {

		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(engine.DataBaseName). //	数据库名
		Type(item.Type).            //	表名 可不设ID，由系统分配
		//Id(item.Id).
		BodyJson(item) //	内容

	//	检测是否传入Id
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	//	存 Index -> 可创建，可修改 读 Get 找 Search
	_, err = indexService.Do(context.Background())

	if err != nil {

		return err
	}

	return nil

}
