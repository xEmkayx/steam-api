package steamclient

import (
	"net/http"
	"reflect"
	"steam-api/pkg/steamclient/config"
	model "steam-api/pkg/steamclient/model/ISteamNews"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetNewsForApp(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   GetNewsForAppParams
		response string
		want     *model.AppNews
		wantErr  bool
	}

	// Define test cases
	testCases := []testCase{
		{
			name: "Test Case 1",
			params: GetNewsForAppParams{
				AppId:     420,
				Count:     75,
				MaxLength: 1,
				Format:    config.Json,
			},
			response: `{
				"appnews": {
					"appid":420,
					"newsitems":[
						{
							"gid":"5121196841288526580",
							"title":"Half-Life 2: Episode Two's community-made VR mod arrives on Steam today",
							"url":"https://steamstore-a.akamaihd.net/news/externalpost/Rock, Paper, Shotgun/5121196841288526580",
							"is_external_url":true,
							"author":"",
							"contents":"T...",
							"feedlabel":"Rock, Paper, Shotgun",
							"date":1680788589,
							"feedname":"Rock, Paper, Shotgun",
							"feed_type":0,
							"appid":220
						}],
						"count":75
						}
					}`,
			want: &model.AppNews{
				AppId: 420,
				Count: 75,
				NewsItemList: []model.NewsItem{
					{
						GID:           "5121196841288526580",
						Title:         "Half-Life 2: Episode Two's community-made VR mod arrives on Steam today",
						Url:           "https://steamstore-a.akamaihd.net/news/externalpost/Rock, Paper, Shotgun/5121196841288526580",
						IsExternalUrl: true,
						Author:        "",
						Contents:      "T...",
						FeedLabel:     "Rock, Paper, Shotgun",
						Date:          1680788589,
						FeedName:      "Rock, Paper, Shotgun",
						FeedType:      0,
						AppId:         220,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/ISteamNews/GetNewsForApp/v2",
				func(req *http.Request) (*http.Response, error) {
					if req.URL.Query().Get("appid") != strconv.FormatUint(uint64(tc.params.AppId), 10) ||
						req.URL.Query().Get("count") != strconv.FormatUint(uint64(tc.params.Count), 10) ||
						req.URL.Query().Get("maxlength") != strconv.FormatUint(uint64(tc.params.MaxLength), 10) ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}

					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				},
			)

			api := NewClientWithoutKey(&http.Client{})
			got, err := api.GetNewsForApp(tc.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetNewsForApp() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetNewsForApp() got = %v, want %v", got, tc.want)
			}
		})
	}
}
