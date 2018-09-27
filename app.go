package main

import (
	"fmt"

	"github.com/irvana/crawler/database"

	"github.com/irvana/crawler/crawler"
)

func main() {
	database.InitDb()
	crawler.InitClient()
	sum, err := crawler.CrawlPlayerSummary("http://fplarchives.com/api/Home/SeasonDetails/LoadData?id=1")
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range sum.PlayersTable.Summaries {
		// todo crawl its detail on particular season
		database.SaveToDB(&p)
	}

}
