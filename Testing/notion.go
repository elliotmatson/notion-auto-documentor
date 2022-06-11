package main

import (
	"context"
	"fmt"
	"net/http"

	notion "github.com/dstotijn/go-notion"
)

func main() {
	client := notion.NewClient(apiKey)
	//http.HandleFunc("/", handler)

	//http.ListenAndServe(":8080", nil)

	results, err := client.Search(context.Background(), &notion.SearchOpts{ /*Query: "test", */ Filter: &notion.SearchFilter{Value: "database", Property: "object"}})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for i := 0; i < len(results.Results); i++ {
		switch t := results.Results[i].(type) {
		case notion.Database:
			fmt.Printf("Page: %v\n", t.Title[0].PlainText)
			//result, _ := client.QueryDatabase(context.Background(), t.ID, &notion.DatabaseQuery{})
			initDB(client, t.ID)
			//spew.Dump(result)
		default:
			fmt.Printf("wrong type %T\n", t)
		}
	}
}

// sets up a notion db with the required parameters
func initDB(c *notion.Client, db string) {
	t := []notion.RichText{{Text: &notion.Text{Content: "test"}}}
	p := make(notion.DatabasePageProperties)
	p["Name"] = notion.DatabasePageProperty{Title: t}
	p["Poops"] = notion.DatabasePageProperty{Title: []notion.RichText{{Text: &notion.Text{Content: "poop poop"}}}}

	_, err := c.UpdateDatabase(context.Background(), db, notion.UpdateDatabaseParams{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// adds a page to a notion db
func addPage(c *notion.Client, db string, title string) {
	t := []notion.RichText{{Text: &notion.Text{Content: title}}}
	p := make(notion.DatabasePageProperties)
	p["Name"] = notion.DatabasePageProperty{Title: t}
	_, err := c.CreatePage(context.Background(), notion.CreatePageParams{ParentType: notion.ParentTypeDatabase, ParentID: db, DatabasePageProperties: &p})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}
