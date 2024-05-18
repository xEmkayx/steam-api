package steamclient

import (
	"net/http"
	"reflect"
	"steam-api/pkg/steamclient/config"
	model "steam-api/pkg/steamclient/model/ISteamUser"
	"strconv"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetPlayerSummaries(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   GetPlayerSummariesParams
		response string
		want     *model.PlayerSummaries
		wantErr  bool
	}

	locCountryCode := "US"
	locStateCode := "WA"
	locCityId := 3961

	testCases := []testCase{
		{
			name: "Test Case 1",
			params: GetPlayerSummariesParams{
				SteamIds: []int64{76561197960435530},
				Format:   config.Json,
			},
			response: `{"response":{"players":[{"steamid":"76561197960435530","communityvisibilitystate":3,"profilestate":1,"personaname":"Robin","profileurl":"https://steamcommunity.com/id/robinwalker/","avatar":"https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9.jpg","avatarmedium":"https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9_medium.jpg","avatarfull":"https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9_full.jpg","avatarhash":"81b5478529dce13bf24b55ac42c1af7058aaf7a9","personastate":0,"realname":"Robin Walker","primaryclanid":"103582791429521412","timecreated":1063407589,"personastateflags":0,"loccountrycode":"US","locstatecode":"WA","loccityid":3961}]}}`,
			want: &model.PlayerSummaries{
				PlayerSums: []model.PlayerSummary{
					{
						SteamID:                  "76561197960435530",
						CommunityVisibilityState: 3,
						ProfileState:             1,
						PersonaName:              "Robin",
						ProfileURL:               "https://steamcommunity.com/id/robinwalker/",
						Avatar:                   "https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9.jpg",
						AvatarMedium:             "https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9_medium.jpg",
						AvatarFull:               "https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9_full.jpg",
						AvatarHash:               "81b5478529dce13bf24b55ac42c1af7058aaf7a9",
						PersonaState:             0,
						RealName:                 "Robin Walker",
						PrimaryClanID:            "103582791429521412",
						TimeCreated:              1063407589,
						PersonaStateFlags:        0,
						LocCountryCode:           &locCountryCode,
						LocStateCode:             &locStateCode,
						LocCityId:                &locCityId,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2",
				func(req *http.Request) (*http.Response, error) {
					steamIds := make([]string, len(tc.params.SteamIds))
					for i, v := range tc.params.SteamIds {
						steamIds[i] = strconv.FormatInt(v, 10)
					}

					if req.URL.Query().Get("key") != "test-key" ||
						req.URL.Query().Get("steamids") != strings.Join(steamIds, ",") ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}
					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				})

			api := New("test-key", &http.Client{})
			got, err := api.GetPlayerSummaries(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetPlayerSummaries() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetPlayerSummaries() got = %v , expected", tc.want)
			}
		})

	}
}

func TestGetFriendList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   GetFriendListParams
		response string
		want     *model.FriendList
		wantErr  bool
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			params: GetFriendListParams{
				SteamId: 76561197960265731,
				Format:  config.Json,
			},
			response: `{
                "friendslist": {
                    "friends": [
                        {
                            "steamid": "76561197960265731",
                            "relationship": "friend",
                            "friend_since": 0
                        },
                        {
                            "steamid": "76561199147650161",
                            "relationship": "friend",
                            "friend_since": 1614836906
                        }
                    ]
                }
            }`,
			want: &model.FriendList{
				Friends: []model.Friend{
					{
						SteamId:      "76561197960265731",
						Relationship: "friend",
						FriendSince:  0,
					},
					{
						SteamId:      "76561199147650161",
						Relationship: "friend",
						FriendSince:  1614836906,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/ISteamUser/GetFriendList/v1",
				func(req *http.Request) (*http.Response, error) {
					if req.URL.Query().Get("key") != "test-key" ||
						req.URL.Query().Get("steamid") != strconv.FormatInt(tc.params.SteamId, 10) ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}
					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				})

			api := New("test-key", &http.Client{})
			got, err := api.GetFriendList(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetFriendList() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetFriendList() got = %v , expected", tc.want)
			}
		})

	}
}
