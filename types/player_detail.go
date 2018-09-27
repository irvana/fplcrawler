package types

type PlayerDetail struct {
	ClubFilter      []ClubFilter `json:"clubFilter"`
	PlayersTable    PlayersTable `json:"playersTable"`
	Title           string       `json:"title"`
	Position        string       `json:"position"`
	Team            string       `json:"team"`
	Season          string       `json:"season"`
	PreviousSeasons []string     `json:"previousSeasons"`
	NotFoundMessage interface{}  `json:"notFoundMessage"`
}

type ClubFilter struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type PlayersTable struct {
	Details                []Detail      `json:"rows"`
	Columns                []interface{} `json:"columns"`
	PlayerHasMultipleTeams bool          `json:"playerHasMultipleTeams"`
}

type Detail struct {
	Opponent        string      `json:"opponent"`
	ElementTeam     interface{} `json:"elementTeam"`
	GameWeek        int64       `json:"gameWeek"`
	Points          int64       `json:"points"`
	Goals           int64       `json:"goals"`
	Assists         int64       `json:"assists"`
	Minutes         int64       `json:"minutes"`
	Reds            int64       `json:"reds"`
	Yellows         int64       `json:"yellows"`
	PensMissed      int64       `json:"pensMissed"`
	Bonus           int64       `json:"bonus"`
	Bps             int64       `json:"bps"`
	CleanSheets     int64       `json:"cleanSheets"`
	Cost            float64     `json:"cost"`
	GoalsConceded   int64       `json:"goalsConceded"`
	OwnGoals        int64       `json:"ownGoals"`
	PenaltiesSaved  int64       `json:"penaltiesSaved"`
	PenaltiesMissed int64       `json:"penaltiesMissed"`
	Saves           int64       `json:"saves"`
	Influence       float64     `json:"influence"`
	Creativity      float64     `json:"creativity"`
	Threat          float64     `json:"threat"`
	ICTIndex        float64     `json:"ictIndex"`
	SelectedBy      int64       `json:"selectedBy"`
	NetTransfers    int64       `json:"netTransfers"`
}
