package database

import (
	"fmt"

	"github.com/irvana/crawler/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbClient *sqlx.DB

func InitDb() error {
	if dbClient == nil {
		var err error
		dbClient, err = sqlx.Connect("postgres", "user=postgres password=postgres dbname=postgres host=localhost port=5432")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func SaveToDB(sum *types.Summary) error {
	_, err := dbClient.Exec(
		"INSERT INTO player_summary(id,name,normalizedName,teamShortName,teamLongName,position,points,goals,assists,minutes,reds,yellows,saves,pensSaved,pensMissed,bonus,cleanSheets,lastCost) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18);",
		sum.ID,
		sum.Name,
		sum.NormalizedName,
		sum.TeamShortName,
		sum.TeamLongName,
		sum.Position,
		sum.Points,
		sum.Goals,
		sum.Assists,
		sum.Minutes,
		sum.Reds,
		sum.Yellows,
		sum.Saves,
		sum.PensSaved,
		sum.PensMissed,
		sum.Bonus,
		sum.CleanSheets,
		sum.LastCost,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
