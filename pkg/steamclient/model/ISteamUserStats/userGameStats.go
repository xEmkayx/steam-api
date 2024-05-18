// user stats for a specific game
package model

type UserStatsResponse struct {
	PlayerStats UserStats `json:"playerstats" xml:"playerstats"`
}

type UserStats struct {
	SteamID      string                 `json:"steamID" xml:"steamID"`
	GameName     string                 `json:"gameName" xml:"gameName"`
	Stats        []Stat                 `json:"stats" xml:"stats>stat"`
	Achievements []UserStatsAchievement `json:"achievements" xml:"achievements>achievement"`
}

type Stat struct {
	Name  string `json:"name" xml:"name"`
	Value uint64 `json:"value" xml:"value"`
}

type UserStatsAchievement struct {
	Name     string `json:"name" xml:"name"`
	Achieved int    `json:"achieved" xml:"achieved"` // 1 - true, 0 - false // todo: boolean parser
}
