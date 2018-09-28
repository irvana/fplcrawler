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

func DoRequest(URL string) ([]byte, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.89 Safari/537.36")
	resp, err := cClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respByte, nil
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

func CrawlPlayerDetail(URL string) (*types.PlayerDetail, error) {
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

	var result types.PlayerDetail
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		fmt.Println("error unmarshal", err, string(respByte))
		return nil, err
	}

	return &result, nil
}
