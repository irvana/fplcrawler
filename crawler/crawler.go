package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/irvana/crawler/types"
)

var cClient *http.Client

func InitClient() {
	if cClient == nil {
		cClient = &http.Client{
			Timeout: time.Duration(10 * time.Second),
		}
	}
}

func CrawlPlayerSummary(URL string) (*types.PlayerSummary, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json, text/plain, */*")
	resp, err := cClient.Do(req)
	if err != nil {
		fmt.Println("error get to", URL, err)
		return nil, err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error parse", err)
		return nil, err

	}

	var result types.PlayerSummary
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		fmt.Println("error unmarshal", err, string(respByte))
		return nil, err
	}

	return &result, nil
}
