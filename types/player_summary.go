package types

type PlayerSummary struct {
	ClubFilter      []Filter     `json:"clubFilter"`
	PositionFilter  []Filter     `json:"positionFilter"`
	PlayersTable    SummaryTable `json:"playersTable"`
	Title           string       `json:"title"`
	NotFoundMessage interface{}  `json:"notFoundMessage"`
}

type Filter struct {
	Text  Text  `json:"text"`
	Value Value `json:"value"`
}

type SummaryTable struct {
	Summaries              []Summary     `json:"rows"`
	Columns                []interface{} `json:"columns"`
	PlayerHasMultipleTeams bool          `json:"playerHasMultipleTeams"`
}

type Summary struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	NormalizedName string   `json:"normalizedName"`
	TeamShortName  Value    `json:"teamShortName"`
	TeamLongName   Text     `json:"teamLongName"`
	Position       Position `json:"position"`
	Points         int64    `json:"points"`
	Goals          int64    `json:"goals"`
	Assists        int64    `json:"assists"`
	Minutes        int64    `json:"minutes"`
	Reds           int64    `json:"reds"`
	Yellows        int64    `json:"yellows"`
	Saves          int64    `json:"saves"`
	PensSaved      int64    `json:"pensSaved"`
	PensMissed     int64    `json:"pensMissed"`
	Bonus          int64    `json:"bonus"`
	CleanSheets    int64    `json:"cleanSheets"`
	LastCost       float64  `json:"lastCost"`
}

type Text string

const (
	Arsenal          Text = "Arsenal"
	Bournemouth      Text = "Bournemouth"
	Burnley          Text = "Burnley"
	Chelsea          Text = "Chelsea"
	CrystalPalace    Text = "Crystal Palace"
	Defenders        Text = "Defenders"
	Everton          Text = "Everton"
	Forwards         Text = "Forwards"
	Goalkeepers      Text = "Goalkeepers"
	Hull             Text = "Hull"
	Leicester        Text = "Leicester"
	Liverpool        Text = "Liverpool"
	ManchesterCity   Text = "Manchester City"
	ManchesterUnited Text = "Manchester United"
	Middlesbrough    Text = "Middlesbrough"
	Midfielders      Text = "Midfielders"
	Southampton      Text = "Southampton"
	Spurs            Text = "Spurs"
	Stoke            Text = "Stoke"
	Sunderland       Text = "Sunderland"
	Swansea          Text = "Swansea"
	Watford          Text = "Watford"
	WestBrom         Text = "West Brom"
	WestHam          Text = "West Ham"
)

type Value string

const (
	Ars      Value = "ARS"
	Bou      Value = "BOU"
	Bur      Value = "BUR"
	Che      Value = "CHE"
	Cry      Value = "CRY"
	Eve      Value = "EVE"
	Hul      Value = "HUL"
	Lei      Value = "LEI"
	Liv      Value = "LIV"
	Mci      Value = "MCI"
	Mun      Value = "MUN"
	Sou      Value = "SOU"
	Stk      Value = "STK"
	Sun      Value = "SUN"
	Swa      Value = "SWA"
	Tot      Value = "TOT"
	ValueDEF Value = "DEF"
	ValueFWD Value = "FWD"
	ValueGK  Value = "GK"
	ValueMID Value = "MID"
	Wat      Value = "WAT"
	Wba      Value = "WBA"
	Whu      Value = "WHU"
)

type Position string

const (
	PositionDEF Position = "DEF"
	PositionFWD Position = "FWD"
	PositionGK  Position = "GK"
	PositionMID Position = "MID"
)
