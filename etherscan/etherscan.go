package etherscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	EndpointMainnet = "https://api.etherscan.io"
)

var (
	ErrorUnsupportedDataFormat = errors.New("unsupported data format")
)

func GetGasTrackerEstimate(token string) (*GasTrackerEstimate, error) {
	fmt.Println(fmt.Sprintf("%s/api?module=gastracker&action=gasestimate&apikey=%s", EndpointMainnet, token))
	resp, err := http.Get(fmt.Sprintf("%s/api?module=gastracker&action=gasestimate&apikey=%s", EndpointMainnet, token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{
		Result: &GasTrackerEstimate{},
	}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	if response.Status != "1" {
		return nil, errors.New("status error")
	}
	switch data := response.Result.(type) {
	case *GasTrackerEstimate:
		return data, nil
	default:
		return nil, ErrorUnsupportedDataFormat
	}
}
