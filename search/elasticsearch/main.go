package main

import (
	"fmt"

	elastic "gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	fmt.Println("conn es succ")
	tweet := Tweet{User: "zhouwen", Message: "study es"}
	_, err = client.Index().
		Index("two").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do()
	if err != nil {
		// Handle error
		panic(err)
		return
	}
	fmt.Println("insert succ")
}
