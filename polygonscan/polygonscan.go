package polygonscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	EndpointMainnet = "https://api.polygonscan.com"
)

func GetLastPrice(token string) (*LastPrice, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api?module=stats&action=maticprice&apikey=%s", EndpointMainnet, token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	if response.Status == "1" {
		lastPrice := &LastPrice{}
		result, err := json.Marshal(response.Result)
		if err != nil {
			return nil, err
		}
		return lastPrice, json.Unmarshal(result, lastPrice)
	} else {
		return nil, errors.New(strings.ToLower(response.Result.(string)))
	}
}
