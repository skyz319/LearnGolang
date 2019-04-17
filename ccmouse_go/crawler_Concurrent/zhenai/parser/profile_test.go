package parser

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "普普通通", "女士", "")
	if len(result.Items) != 1 {
		t.Error("item should contain 1"+"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	fmt.Printf("%s\n", profile)
}
