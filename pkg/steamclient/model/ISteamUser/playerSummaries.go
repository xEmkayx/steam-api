package model

type PlayerSummariesWrapper struct {
	PlayerSums PlayerSummaries `json:"response" xml:"response"`
}

type PlayerSummaries struct {
	PlayerSums []PlayerSummary `json:"players" xml:"player"`
}

// Player definiert die Struktur für einen Spieler.
type PlayerSummary struct {
	SteamID                  string  `json:"steamid" xml:"steamid"`
	CommunityVisibilityState int     `json:"communityvisibilitystate" xml:"communityvisibilitystate"`
	ProfileState             int     `json:"profilestate" xml:"profilestate"`
	PersonaName              string  `json:"personaname" xml:"personaname"`
	CommentPermission        int     `json:"commentpermission" xml:"commentpermission"`
	ProfileURL               string  `json:"profileurl" xml:"profileurl"`
	Avatar                   string  `json:"avatar" xml:"avatar"`
	AvatarMedium             string  `json:"avatarmedium" xml:"avatarmedium"`
	AvatarFull               string  `json:"avatarfull" xml:"avatarfull"`
	AvatarHash               string  `json:"avatarhash" xml:"avatarhash"`
	LastLogOff               int64   `json:"lastlogoff" xml:"lastlogoff"` // Unix timestamp
	PersonaState             int     `json:"personastate" xml:"personastate"`
	RealName                 string  `json:"realname" xml:"realname"`
	PrimaryClanID            string  `json:"primaryclanid" xml:"primaryclanid"`
	TimeCreated              int64   `json:"timecreated" xml:"timecreated"`             // Unix timestamp
	PersonaStateFlags        int     `json:"personastateflags" xml:"personastateflags"` // Optional integer field
	LocCountryCode           *string `json:"loccountrycode,omitempty" xml:"loccountrycode,omitempty"`
	LocStateCode             *string `json:"locstatecode,omitempty" xml:"locstatecode,omitempty"`
	LocCityId                *int    `json:"loccityid,omitempty" xml:"loccityid,omitempty"`
}
