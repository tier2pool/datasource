package ethermine

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	apiURL = "https://api.ethermine.org"
)

func GetMinerDashboard(address string) (*MinerDashboard, error) {
	resp, err := http.Get(fmt.Sprintf("%s/miner/%s/dashboard", apiURL, address))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	if response.Status != "OK" {
		return nil, response.Error
	}
	switch data := response.Data.(type) {
	case string:
		return nil, errors.New(data)
	default:
		buffer, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		minerDashboard := MinerDashboard{}
		if err := json.Unmarshal(buffer, &minerDashboard); err != nil {
			return nil, err
		}
		return &minerDashboard, nil
	}
}

func GetMinerStats(address string) (*MinerStats, error) {
	resp, err := http.Get(fmt.Sprintf("%s/miner/%s/currentStats", apiURL, address))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	if response.Status != "OK" {
		return nil, response.Error
	}
	switch data := response.Data.(type) {
	case string:
		return nil, errors.New(data)
	default:
		buffer, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		minerStats := MinerStats{}
		if err := json.Unmarshal(buffer, &minerStats); err != nil {
			return nil, err
		}
		return &minerStats, nil
	}
}
