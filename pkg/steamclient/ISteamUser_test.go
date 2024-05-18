package steamclient

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/xemkayx/steam-api/pkg/steamclient/config"
	model "github.com/xemkayx/steam-api/pkg/steamclient/model/ISteamUser"

	"github.com/jarcoal/httpmock"
)

func TestGetPlayerSummaries(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		client   *Client
		name     string
		params   GetPlayerSummariesParams
		response string
		want     *model.PlayerSummaries
		wantErr  bool
	}
	httpClient := &http.Client{}

	locCountryCode := "US"
	locStateCode := "WA"
	locCityId := 3961

	testCases := []testCase{
		{
			client: New("test-key", httpClient),
			name:   "TestGetPlayerSummaries 1 - XML",
			params: GetPlayerSummariesParams{
				SteamIds: []int64{76561197960435530},
				Format:   config.Xml,
			},
			response: `<?xml version="1.0" encoding="UTF-8"?>
			<!DOCTYPE response>
			<response>
				<players>
					<player>
						<steamid>76561197960435530</steamid>
						<communityvisibilitystate>3</communityvisibilitystate>
						<profilestate>1</profilestate>
						<personaname>Robin</personaname>
						<profileurl>https://steamcommunity.com/id/robinwalker/</profileurl>
						<avatar>https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9.jpg</avatar>
						<avatarmedium>https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9_medium.jpg</avatarmedium>
						<avatarfull>https://avatars.steamstatic.com/81b5478529dce13bf24b55ac42c1af7058aaf7a9_full.jpg</avatarfull>
						<avatarhash>81b5478529dce13bf24b55ac42c1af7058aaf7a9</avatarhash>
						<personastate>0</personastate>
						<realname>Robin Walker</realname>
						<primaryclanid>103582791429521412</primaryclanid>
						<timecreated>1063407589</timecreated>
						<personastateflags>0</personastateflags>
						<loccountrycode>US</loccountrycode>
						<locstatecode>WA</locstatecode>
						<loccityid>3961</loccityid>
					</player>
				</players>
			</response>`,
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
		{
			client: New("test-key", httpClient),
			name:   "TestGetPlayerSummaries 2 - JSON",
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

					if req.URL.Query().Get("key") != tc.client.Key ||
						req.URL.Query().Get("steamids") != strings.Join(steamIds, ",") ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}
					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				})

			api := New("test-key", &http.Client{})
			// got, err := tc.client.GetPlayerSummaries(tc.params)
			got, err := api.GetPlayerSummaries(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetPlayerSummaries() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetPlayerSummaries() got = %v , expected %v", got, tc.want)
			}
		})

	}
}

func TestGetPlayerSummariesErrors(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		client   *Client
		params   GetPlayerSummariesParams
		response string
		want     error
		wantErr  bool
	}

	var sIds []int64

	for i := range 150 {
		sIds = append(sIds, int64(i))
	}

	testCases := []testCase{
		{
			name:   "TestGetPlayerSummariesErrors 1 - No Key",
			client: NewClientWithoutKey(http.DefaultClient),
			params: GetPlayerSummariesParams{
				SteamIds: sIds,
				Format:   config.Json,
			},
			response: "",
			want:     errors.New(apiKeyErrorMessage),
			wantErr:  true,
		},
		{
			name:   "TestGetPlayerSummariesErrors 1 - No Key",
			client: New("test-key", http.DefaultClient),
			params: GetPlayerSummariesParams{
				SteamIds: sIds,
				Format:   config.Json,
			},
			response: "",
			want:     errors.New("you have provided too many Steam IDs. reduce the amount to 100"),
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2",
				func(req *http.Request) (*http.Response, error) {
					return nil, nil
				})

			_, err := tc.client.GetPlayerSummaries(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetPlayerSummaries() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(err, tc.want) {
				t.Errorf("GetPlayerSummaries() got = %v , expected %v", err, tc.want)
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
			name: "TestGetFriendList - JSON Success",
			params: GetFriendListParams{
				SteamId: 12345,
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
		{
			name: "TestGetFriendList 2 - XML Success",
			params: GetFriendListParams{
				SteamId: 12345,
				Format:  config.Xml,
			},
			response: `<?xml version="1.0" encoding="UTF-8"?>
			<!DOCTYPE friendslist>
			<friendslist>
				<friends>
					<friend>
						<steamid>76561197960265731</steamid>
						<relationship>friend</relationship>
						<friend_since>0</friend_since>
					</friend>
					<friend>
						<steamid>76561199147650161</steamid>
						<relationship>friend</relationship>
						<friend_since>1614836906</friend_since>
					</friend>
					</friends>
					</friendslist>`,
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

func TestGetFriendListErrors(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		client   *Client
		params   GetFriendListParams
		response string
		want     error
		wantErr  bool
	}

	testCases := []testCase{
		{
			name:   "TestGetFriendListErrors 1 - No Key",
			client: NewClientWithoutKey(http.DefaultClient),
			params: GetFriendListParams{
				SteamId:      123,
				Relationship: config.Friend,
				Format:       config.Json,
			},
			response: "",
			want:     errors.New(apiKeyErrorMessage),
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/ISteamUser/GetFriendList/v1",
				func(req *http.Request) (*http.Response, error) {
					return nil, nil
				})

			_, err := tc.client.GetFriendList(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetPlayerSummaries() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})

	}
}
