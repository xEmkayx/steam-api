package model

type PlayerStatsWrapper struct {
	PlayerStats PlayerAchievements `json:"playerstats"`
}

// PlayerStats - Second level that contains details about the player and achievements
type PlayerAchievements struct {
	SteamID      string        `json:"steamID" xml:"steamID"`
	GameName     string        `json:"gameName" xml:"gameName"`
	Achievements []Achievement `json:"achievements" xml:"achievements>achievement"`
}

// Achievement - Details each individual achievement
type Achievement struct {
	APIName    string `json:"apiname" xml:"apiname"`
	Achieved   int    `json:"achieved" xml:"achieved"`
	UnlockTime int64  `json:"unlocktime" xml:"unlocktime"` // This is usually a Unix timestamp
}
