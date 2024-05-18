package model

type RecentlyPlayedGamesWrapper struct {
	RecentlyPlayedGames RecentlyPlayedGames `json:"response" xml:"response"`
}

type RecentlyPlayedGames struct {
	TotalCount int    `json:"total_count" xml:"total_count"`
	Games      []Game `json:"games" xml:"games>message"`
}

type OwnedGamesWrapper struct {
	OwnedGames OwnedGames `json:"response" xml:"response"`
}

type OwnedGames struct {
	GameCount int    `json:"game_count" xml:"game_count"`
	Games     []Game `json:"games" xml:"games>message"`
}

// Game defines a combined struct of all attributes
// this struct is used by both responses
type Game struct {
	AppID           int    `json:"appid" xml:"appid"`
	Name            string `json:"name,omitempty" xml:"name,omitempty"` // Optional, not in every JSON
	PlaytimeForever int    `json:"playtime_forever" xml:"playtime_forever"`
	ImageIconURL    string `json:"img_icon_url,omitempty" xml:"img_icon_url,omitempty"` // Optional
	// RtimeLastPlayed        Timestamp             `json:"rtime_last_played,omitempty"` // Optional
	RtimeLastPlayed        int64    `json:"rtime_last_played,omitempty" xml:"rtime_last_played,omitempty"` // Optional
	Playtime2Weeks         int      `json:"playtime_2weeks,omitempty" xml:"playtime_2weeks,omitempty"`     // Optional
	PlaytimeWindowsForever int      `json:"playtime_windows_forever" xml:"playtime_windows_forever"`
	PlaytimeMacForever     int      `json:"playtime_mac_forever" xml:"playtime_mac_forever"`
	PlaytimeLinuxForever   int      `json:"playtime_linux_forever" xml:"playtime_linux_forever"`
	PlaytimeDeckForever    int      `json:"playtime_deck_forever" xml:"playtime_deck_forever"`
	PlaytimeDisconnected   uint32   `json:"playtime_disconnected" xml:"playtime_disconnected"`
	ContentDescriptorIds   []uint32 `json:"content_descriptorids,omitempty" xml:"content_descriptorids>uint32,omitempty"`
}

// type ContentDescriptorId struct {
// 	Ids []uint32
// }

// // Timestamp ist eine Hilfsstruktur zur Verarbeitung von Unix-Zeitstempeln.
// type Timestamp struct {
// 	time.Time
// }
