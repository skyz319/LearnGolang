package persist

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/model"
	"LearnGolang/ccmouse_go/crawler_Concurrent/zhenai/parser"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"testing"
)

func TestSaver(t *testing.T) {

	contents, err := ioutil.ReadFile("profile1.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(contents, "普普通通", "女士")
	if len(result.Items) != 1 {
		t.Error("item should contain 1"+"element; but was %v", result.Items)
	}

	t.Logf("%+v", result.Items[0])

	id, err := save(result.Items[0])

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	resp, err := client.Get().Index("dating_profile_test").Type("zhenai_test").Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}
	t.Log("===========")
	//t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	t.Logf("%+v", actual)
}
