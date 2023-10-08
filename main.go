package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
	// InsertIndexEs()
	GetEsDoc()
}

func GetEsDoc() {
	ctx := context.Background()
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Search for documents
	query := `{"query": {"match": {"title": "học code"}}}`
	req := esapi.SearchRequest{
		Index: []string{"learninginternation"},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		log.Fatalf("Error searching documents: %s", err)
	}
	defer res.Body.Close()

	// Đọc và in ra nội dung chi tiết của các tài liệu từ trường "_source"
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error decoding response body: %s", err)
	}

	hits, _ := response["hits"].(map[string]interface{})
	hitsArray, _ := hits["hits"].([]interface{})

	for _, hit := range hitsArray {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		title := source["title"].(string)
		content := source["content"].(string)
		// Các trường dữ liệu khác ở đây
		fmt.Printf("title: %s, content: %s\n", title, content)
	}
}

func InsertIndexEs() {
	// Create a context object for the API calls
	ctx := context.Background()

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Index a document
	doc := `{"title": "Indonesia học code", "content": "Tôi đang học NextJs và Vuejs"}`
	req := esapi.IndexRequest{
		Index:      "learninginternation",
		DocumentID: "4",
		Body:       strings.NewReader(doc),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res)
}
