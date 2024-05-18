package model

type AppNewsResponse struct {
	AppNews AppNews `json:"appnews"`
}

type AppNews struct {
	AppId        int64      `json:"appid" xml:"appid"`
	NewsItemList []NewsItem `json:"newsitems" xml:"newsitems>newsitem"`
	Count        int        `json:"count" xml:"count"`
}

type NewsItem struct {
	GID           string `json:"gid" xml:"gid"`
	Title         string `json:"title" xml:"title"`
	Url           string `json:"url" xml:"url"`
	IsExternalUrl bool   `json:"is_external_url" xml:"is_external_url"`
	Author        string `json:"author,omitempty" xml:"author,omitempty"`
	Contents      string `json:"contents" xml:"contents"`
	FeedLabel     string `json:"feedlabel" xml:"feedlabel"`
	Date          uint64 `json:"date" xml:"date"`
	FeedName      string `json:"feedname" xml:"feedname"`
	FeedType      int    `json:"feed_type" xml:"feed_type"`
	AppId         int64  `json:"appid" xml:"appid"`
}
