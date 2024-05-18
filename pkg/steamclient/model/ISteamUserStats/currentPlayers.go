package model

type NumberOfCurrentPlayersResponse struct {
	CurrentPlayers NumberOfCurrentPlayers `json:"response" xml:"response"`
}

type NumberOfCurrentPlayers struct {
	PlayerCount uint64 `json:"player_count" xml:"player_count"`
	Result      int    `json:"result" xml:"result"`
}
