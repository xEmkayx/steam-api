package steamclient

import (
	"encoding/json"
	"net/http"
	"reflect"
	"steam-api/pkg/steamclient/config"
	model "steam-api/pkg/steamclient/model/IPlayerService"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetOwnedGames(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCases := []struct {
		name     string
		params   GetOwnedGamesParams
		response string
		want     *model.OwnedGames
		wantErr  bool
	}{
		{
			name: "Valid Response",
			params: GetOwnedGamesParams{
				SteamId:                123456,
				IncludeAppinfo:         true,
				IncludePlayedFreeGames: true,
				AppIdsFilter:           nil,
				Format:                 config.Json,
				Language:               nil,
			},
			response: `{
                    "response": {
                        "game_count": 1,
                        "games": [
                            {
                                "appid": 10,
                                "playtime_forever": 69,
                                "playtime_windows_forever": 0,
                                "playtime_mac_forever": 0,
                                "playtime_linux_forever": 0,
                                "playtime_deck_forever": 0,
                                "rtime_last_played": 1503844956,
                                "playtime_disconnected": 200324100
                            }
                        ]
                    }
                }`,
			want: &model.OwnedGames{
				GameCount: 1,
				Games: []model.Game{{
					AppID:                  10,
					PlaytimeForever:        69,
					PlaytimeWindowsForever: 0,
					PlaytimeMacForever:     0,
					PlaytimeLinuxForever:   0,
					PlaytimeDeckForever:    0,
					RtimeLastPlayed:        1503844956,
					PlaytimeDisconnected:   200324100,
				}},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1",
				func(req *http.Request) (*http.Response, error) {
					inputJson := map[string]interface{}{
						"steamid":                   tc.params.SteamId,
						"include_appinfo":           tc.params.IncludeAppinfo,
						"include_played_free_games": tc.params.IncludePlayedFreeGames,
					}

					jsonBytes, err := json.Marshal(inputJson)
					if err != nil {
						t.Errorf("error marshaling JSON: %s", err)
					}

					if req.URL.Query().Get("key") != "test-key" ||
						req.URL.Query().Get("input_json") != string(jsonBytes) ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}

					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				},
			)

			api := New("test-key", &http.Client{})
			got, err := api.GetOwnedGames(tc.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetOwnedGames() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetOwnedGames() got = %v, want %v", got, tc.want)
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGetRecentlyPlayedGames(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   GetRecentlyPlayedGamesParams
		response string
		want     *model.RecentlyPlayedGames
		wantErr  bool
	}

	// Define test cases
	testCases := []testCase{
		{
			name: "Test Case 1",
			params: GetRecentlyPlayedGamesParams{
				SteamId: 123456789,
				Count:   5,
				Format:  config.Json,
			},
			response: `{"response":{"total_count":4,"games":[{"appid":553850,"name":"HELLDIVERS™ 2","playtime_2weeks":338,"playtime_forever":2534,"img_icon_url":"c3dff088e090f81d6e3d88eabbb67732647c69cf","playtime_windows_forever":2534,"playtime_mac_forever":0,"playtime_linux_forever":0,"playtime_deck_forever":0}]}}`,
			want: &model.RecentlyPlayedGames{
				TotalCount: 4,
				Games: []model.Game{
					{
						AppID:                  553850,
						Name:                   "HELLDIVERS™ 2",
						Playtime2Weeks:         338,
						PlaytimeForever:        2534,
						ImageIconURL:           "c3dff088e090f81d6e3d88eabbb67732647c69cf",
						PlaytimeWindowsForever: 2534,
						PlaytimeMacForever:     0,
						PlaytimeLinuxForever:   0,
						PlaytimeDeckForever:    0,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v1",
				func(req *http.Request) (*http.Response, error) {
					if req.URL.Query().Get("key") != "test-key" ||
						req.URL.Query().Get("steamid") != strconv.FormatInt(tc.params.SteamId, 10) ||
						req.URL.Query().Get("count") != strconv.Itoa(tc.params.Count) ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}
					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				})

			api := New("test-key", &http.Client{})
			got, err := api.GetRecentlyPlayedGames(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetRecentlyPlayedGames() error=%v , expected%v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetRecentlyplayedGames () got %v want %v ", got, tc.want)

			}
		})

	}
}
