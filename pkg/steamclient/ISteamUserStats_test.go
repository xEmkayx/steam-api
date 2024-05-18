package steamclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"github.com/xemkayx/steam-api/pkg/steamclient/config"
	model "github.com/xemkayx/steam-api/pkg/steamclient/model/ISteamUserStats"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetGlobalAchievementPercentagesForApp(t *testing.T) {
	htCl := http.Client{}
	client := New("test_key", &htCl)
	tests := []struct {
		name           string
		params         GlobalAchievementPercentageParams
		mockResponse   string
		mockStatusCode int
		expectedError  bool
	}{
		{
			name: "success JSON response",
			params: GlobalAchievementPercentageParams{
				GameId: 1234,
				Format: config.Json,
			},
			mockResponse:   `{"global_achievement_percentages": {"achievements": [{"name": "ACH_1", "percent": 56.78}]}}`,
			mockStatusCode: http.StatusOK,
			expectedError:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				fmt.Fprintln(w, tt.mockResponse)
			}))
			defer server.Close()
			vals := url.Values{}
			vals.Set("gameid", strconv.FormatUint(tt.params.GameId, 10))
			vals.Set("format", tt.params.Format.String())
			if client.IsKeySet() {
				vals.Set("key", client.Key)
			}
			url := fmt.Sprintf("%s/?%s", server.URL, vals.Encode())
			fmt.Println("Mocked URL:", url)

			result, err := client.GetGlobalAchievementPercentagesForApp(tt.params)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestGetNumberOfCurrentPlayers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   NumberOfCurrentPlayersParams
		response string
		want     *model.NumberOfCurrentPlayers
		wantErr  bool
	}

	// Define test cases
	testCases := []testCase{
		{
			name: "Test Case 1",
			params: NumberOfCurrentPlayersParams{
				GameId: 12345,
				Format: config.Json,
			},
			response: `{"response":{"player_count":204,"result":1}}`,
			want: &model.NumberOfCurrentPlayers{
				PlayerCount: 204,
				Result:      1,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/ISteamUserStats/GetNumberOfCurrentPlayers/v1",
				func(req *http.Request) (*http.Response, error) {
					if req.URL.Query().Get("gameid") != strconv.FormatUint(uint64(tc.params.GameId), 10) ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}
					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				},
			)

			api := NewClientWithoutKey(&http.Client{})
			got, err := api.GetNumberOfCurrentPlayers(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetNumberOfCurrentPlayers() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetNumberOfCurrentPlayers() got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetPlayerAchievements(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   PlayerAchievementsParams
		response string
		want     *model.PlayerAchievements
		wantErr  bool
	}

	// Define test cases
	testCases := []testCase{
		{
			name: "Test Case 1",
			params: PlayerAchievementsParams{
				SteamId: 123456789,
				AppId:   730,
				Format:  config.Json,
			},
			response: `{
   				"playerstats": {
   					"steamID": "123456789",
   					"gameName": "Counter-Strike 2",
   					"achievements": [
   						{
   							"apiname": "PLAY_CS2",
   							"achieved": 1,
   							"unlocktime": 1695912022
	  					}
	  				],
	  				"success": true
	  			}
	  		}`,
			want: &model.PlayerAchievements{
				SteamID:  "123456789",
				GameName: "Counter-Strike 2",
				Achievements: []model.Achievement{
					{
						APIName:    "PLAY_CS2",
						Achieved:   1,
						UnlockTime: 1695912022,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := `https://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v1`
			httpmock.RegisterResponder("GET", url,
				func(req *http.Request) (*http.Response, error) {
					if req.URL.Query().Get("key") != "test-key" ||
						req.URL.Query().Get("steamid") != strconv.FormatUint(uint64(tc.params.SteamId), 10) ||
						req.URL.Query().Get("appid") != strconv.FormatInt(int64(tc.params.AppId), 10) ||
						req.URL.Query().Get("format") != tc.params.Format.String() {
						t.Errorf("Request parameters do not match")
					}
					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				})

			api := New("test-key", &http.Client{})
			got, err := api.GetPlayerAchievements(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetPlayerAchievements() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {

				t.Errorf("Got = %v\nWant = %v", got, tc.want)
			}
		})
	}
}

func TestGetSchemaForGame(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   SchemaForGameParams
		response string
		want     *model.GameSchemaGame
		wantErr  bool
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			params: SchemaForGameParams{
				AppId:  730,
				Format: config.Json,
			},
			response: `{
				"game": {
				  "gameName": "ValveTestApp260",
				  "gameVersion": "245",
				  "availableGameStats": {
					"stats": [
					  {
						"name": "total_kills",
						"defaultvalue": 0,
						"displayName": "Enemy players killed"
					  },
					  {
						"name": "total_deaths",
						"defaultvalue": 0,
						"displayName": "Player Deaths"
					  }
					],
					"achievements": [
					  {
						"name": "PLAY_CS2",
						"defaultvalue": 0,
						"displayName": "A New Beginning",
						"hidden": 1,
						"icon": "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/730/f75dd04fa12445a8ec43be65fa16ff1b8d2bf82e.jpg",
						"icongray": "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/730/e7d79e52002c75f027a51e8eb05ab55416678f09.jpg"
					  }
					]
				  }
				}
			  }`,
			want: &model.GameSchemaGame{
				GameName:    "ValveTestApp260",
				GameVersion: "245",
				AvailableGameStats: model.GameSchemaAvailableGameStats{
					Stats: []model.GameSchemaStat{
						{
							Name:         "total_kills",
							DefaultValue: 0,
							DisplayName:  "Enemy players killed",
						},
						{
							Name:         "total_deaths",
							DefaultValue: 0,
							DisplayName:  "Player Deaths",
						},
					},
					Achievements: []model.GameSchemaAchievement{
						{
							Name:         "PLAY_CS2",
							DefaultValue: 0,
							DisplayName:  "A New Beginning",
							Hidden:       1,
							Icon:         "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/730/f75dd04fa12445a8ec43be65fa16ff1b8d2bf82e.jpg",
							IconGray:     "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/730/e7d79e52002c75f027a51e8eb05ab55416678f09.jpg"},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			httpmock.RegisterResponder("GET", `https://api.steampowered.com/ISteamUserStats/GetSchemaForGame/v2`, func(req *http.Request) (*http.Response, error) {

				if req.URL.Query().Get("appid") != strconv.FormatUint(uint64(tc.params.AppId), 10) || req.URL.Query().Get("format") != tc.params.Format.String() {
					t.Errorf("Request parameters do not match")
				}

				resp := httpmock.NewStringResponse(200, tc.response)
				return resp, nil

			})
			api := New("test-key", &http.Client{})
			got, err := api.GetSchemaForGame(tc.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetSchemaForGame()error=%v,wantErr%v", err, tc.wantErr)
				return

			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetSchemaForGame()got=%v,want%v", got, tc.want)

			}
		})

	}
}

func TestGetUserStatsForGame(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type testCase struct {
		name     string
		params   UserStatsForGameParams
		response string
		want     *model.UserStats
		wantErr  bool
	}

	// Define test cases
	testCases := []testCase{
		{
			name: "Test Case 1",
			params: UserStatsForGameParams{
				SteamId:  123456789,
				AppId:    260,
				Format:   config.Json,
				Language: nil,
			},
			response: `{
                "playerstats": {
                    "steamID": "123456789",
                    "gameName": "ValveTestApp260",
                    "stats": [
                        {
                            "name": "total_kills",
                            "value": 33683
                        }
                    ],
                    "achievements": [
                        {
                            "name": "PLAY_CS2",
                            "achieved": 1
                        }
                    ]
                }
            }`,
			want: &model.UserStats{
				SteamID:  "123456789",
				GameName: "ValveTestApp260",
				Stats: []model.Stat{
					{
						Name:  "total_kills",
						Value: 33683,
					},
				},
				Achievements: []model.UserStatsAchievement{
					{
						Name:     "PLAY_CS2",
						Achieved: 1,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2",
				func(req *http.Request) (*http.Response, error) {
					if req.URL.Query().Get("key") != "test-key" || // Check your actual key here
						req.URL.Query().Get("steamid") != strconv.FormatInt(int64(tc.params.SteamId), 10) ||
						req.URL.Query().Get("appid") != strconv.FormatInt(int64(tc.params.AppId), 10) ||
						req.URL.Query().Get("format") != tc.params.Format.String() ||
						(tc.params.Language != nil && req.URL.Query().Get("l") != tc.params.Language.String()) {
						t.Errorf("Request parameters do not match")
					}
					resp := httpmock.NewStringResponse(200, tc.response)
					return resp, nil
				})

			api := New("test-key", &http.Client{})
			got, err := api.GetUserStatsForGame(tc.params)

			if (err != nil) != tc.wantErr {
				t.Errorf("GetUserStatsForGame() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !assert.Equal(t, tc.want, got) {
				t.Errorf("Got and Want mismatch in testcase %s", tc.name)
			}
		})
	}
}
