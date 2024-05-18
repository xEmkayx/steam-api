package model

// TODO - unknown output format
type GameStatsResponse struct {
	Response struct {
		GlobalStats map[string]interface{} `json:"globalstats"`
	} `json:"response"`
}
