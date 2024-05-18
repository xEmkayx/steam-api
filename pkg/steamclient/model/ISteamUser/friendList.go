package model

type FriendListWrapper struct {
	FriendsList FriendList `json:"friendslist" xml:"friendslist"`
}

type FriendList struct {
	Friends []Friend `json:"friends" xml:"friends>friend"`
}

type Friend struct {
	SteamId      string `json:"steamid" xml:"steamid"`
	Relationship string `json:"relationship" xml:"relationship"`
	FriendSince  int64  `json:"friend_since" xml:"friend_since"`
}
