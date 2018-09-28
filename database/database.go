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

func SavePlayerSummary(sum *types.Summary, season string) error {
	_, err := dbClient.Exec(
		"INSERT INTO player_summary(id,name,normalizedName,teamShortName,teamLongName,position,points,goals,assists,minutes,reds,yellows,saves,pensSaved,pensMissed,bonus,cleanSheets,lastCost,season) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19);",
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
		season,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SavePlayerDetail(sum *types.Detail, season string) error {
	_, err := dbClient.Exec(
		"INSERT INTO public.player_detail (playername, opponent, elementteam, gameweek, points, goals, assists, minutes, reds,yellows, pensmissed, bonus, bps, cleansheets, cost, goalsconceded, owngoals, penaltiessaved, penaltiesmissed, saves, influence, creativity, threat, ictindex, selectedby, nettransfers, season, id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19, $20, $21, $22, $23, $24, $25, $26, $27, nextval('player_detail_id_seq'::regclass));",
		sum.PlayerName,
		sum.Opponent,
		sum.ElementTeam,
		sum.GameWeek,
		sum.Points,
		sum.Goals,
		sum.Assists,
		sum.Minutes,
		sum.Reds,
		sum.Yellows,
		sum.PensMissed,
		sum.Bonus,
		sum.Bps,
		sum.CleanSheets,
		sum.Cost,
		sum.GoalsConceded,
		sum.OwnGoals,
		sum.PenaltiesSaved,
		sum.PenaltiesMissed,
		sum.Saves,
		sum.Influence,
		sum.Creativity,
		sum.Threat,
		sum.ICTIndex,
		sum.SelectedBy,
		sum.NetTransfers,
		season,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
