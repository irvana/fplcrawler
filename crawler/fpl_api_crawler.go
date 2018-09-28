package crawler

import (
	"encoding/json"
	"fmt"

	"github.com/irvana/crawler/types"
)

func RequestBootstrapStaticAPI() (*types.FPLSummary, error) {
	resp, err := DoRequest("https://fantasy.premierleague.com/drf/bootstrap-static")
	if err != nil {
		return nil, err
	}

	var res types.FPLSummary
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func RequestElementPlayerAPI(playerID int) (*types.FPLCurrentPlayer, error) {
	url := fmt.Sprintf("https://fantasy.premierleague.com/drf/element-summary/%d", playerID)
	resp, err := DoRequest(url)
	if err != nil {
		return nil, err
	}
	var res types.FPLCurrentPlayer
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
