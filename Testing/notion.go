package main

import (
	"context"
	"fmt"

	//"net/http"
	"os"

	//"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"

	notion "github.com/dstotijn/go-notion"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading environment variables: %v\n", err)
	}

	client := notion.NewClient(os.Getenv("NOTION_API_KEY"))
	results, err := client.Search(context.Background(), &notion.SearchOpts{ /*Query: "test", */ Filter: &notion.SearchFilter{Value: "database", Property: "object"}})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for i := 0; i < len(results.Results); i++ {
		switch t := results.Results[i].(type) {
		case notion.Database:
			fmt.Printf("Page: %v\n", t.Title[0].PlainText)
			//result, _ := client.QueryDatabase(context.Background(), t.ID, &notion.DatabaseQuery{})
			InitDB(client, t.ID)
			//spew.Dump(result)
		default:
			fmt.Printf("wrong type %T\n", t)
		}
	}
}

// sets up a notion db with the required parameters
func InitDB(c *notion.Client, db string) {
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
func AddPage(c *notion.Client, db string, title string) {
	t := []notion.RichText{{Text: &notion.Text{Content: title}}}
	p := make(notion.DatabasePageProperties)
	p["Name"] = notion.DatabasePageProperty{Title: t}
	_, err := c.CreatePage(context.Background(), notion.CreatePageParams{ParentType: notion.ParentTypeDatabase, ParentID: db, DatabasePageProperties: &p})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
