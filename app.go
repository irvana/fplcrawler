package main

import (
	"fmt"
	"strings"

	"github.com/irvana/crawler/database"

	"github.com/irvana/crawler/crawler"
)

func main() {
	crawlAPI()
}

func crawlAPI() {
	crawler.InitClient()
	stat, err := crawler.RequestBootstrapStaticAPI()
	if err != nil {
		fmt.Println(err)
	}

	for _, sE := range stat.Elements {
		player, err := crawler.RequestElementPlayerAPI(int(sE.ID))
		if err != nil {
			fmt.Println(err)
		}
		if strings.ToLower(sE.WebName) == "salah" {
			fmt.Println(sE.WebName, ":", player.History[len(player.History)-1].ICTIndex)
		}
	}

}

func crawlArchive() {
	crawler.InitClient()
	database.InitDb()
	for i := 1; i <= 2; i++ {
		var season string
		switch i {
		case 2:
			season = "2015-2016"
		case 1:
			season = "2016-2017"
		}
		sum, err := crawler.CrawlPlayerSummary(fmt.Sprintf("http://fplarchives.com/api/Home/SeasonDetails/LoadData?id=%d", i))
		if err != nil {
			fmt.Println(err)
		}
		for _, p := range sum.PlayersTable.Summaries {
			// todo create player ID
			detailURL := fmt.Sprintf("http://fplarchives.com/api/Home/PlayerDetails/LoadData?id=%d&seasonId=%d", p.ID, i)
			detail, err := crawler.CrawlPlayerDetail(detailURL)
			if err != nil {
				continue
			}

			for _, d := range detail.PlayersTable.Details {
				d.PlayerName = p.Name
				err = database.SavePlayerDetail(&d, season)
				if err != nil {
					fmt.Printf("%+v", d)
					return
				}
			}
			database.SavePlayerSummary(&p, season)
		}
	}

}
