// methods on the steam client
package steamclient

// ref: https://steamapi.xpaw.me/#ISteamUserStats

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"

	urlHelper "github.com/xemkayx/steam-api/internal/urlHelper"
	"github.com/xemkayx/steam-api/pkg/steamclient/config"
	model "github.com/xemkayx/steam-api/pkg/steamclient/model/ISteamUserStats"
)

const (
	ISteamUserStats                               = "ISteamUserStats"                       // v1
	GetPlayerAchievementsEndpoint                 = "GetPlayerAchievements"                 // v0001
	GetUserStatsForGameEndpoint                   = "GetUserStatsForGame"                   // v0002
	GetGlobalAchievementPercentagesForAppEndpoint = "GetGlobalAchievementPercentagesForApp" // v0002
	GetNumberOfCurrentPlayersEndpoint             = "GetNumberOfCurrentPlayers"             // v1
	GetSchemaForGameEndpoint                      = "GetSchemaForGame"                      // v2
	SetUserStatsForGameEndpoint                   = "SetUserStatsForGame"                   // v1
)

// Parameters for the GetGlobalAchievementPercentage method
type GlobalAchievementPercentageParams struct {
	GameId uint64              // GameID to retrieve the achievement percentages for
	Format config.OutputFormat // Format of the output
}

// Parameters for the GetGlobalGameStats method
type GlobalGameStatsParams struct {
	AppID     uint32              // AppID that we're getting global stats for
	Count     uint32              // Number of stats get data for
	Names     []string            // Names of stat to get data for
	Format    config.OutputFormat // Format of the output
	StartDate *uint32             // (optional) Start date for daily totals (unix epoch timestamp)
	EndDate   *uint32             // (optional) End date for daily totals (unix epoch timestamp)
}

// Parameters for the GetNumberOfCurrentPlayers method
type NumberOfCurrentPlayersParams struct {
	GameId uint64              // AppID that we're getting user count for
	Format config.OutputFormat // Format of the output
}

// Parameters for the GetPlayerAchievements method
type PlayerAchievementsParams struct {
	SteamId  uint64              // SteamID of user
	AppId    uint32              // AppID to get achievements for
	Format   config.OutputFormat // Format of the output
	Language *config.Language    // (optional) output Language
}

// Parameters for the GetSchemaForGame method
type SchemaForGameParams struct {
	AppId    uint32              // access key
	Format   config.OutputFormat // Format of the output
	Language *config.Language    // (optional) output Language
}

// Parameters for the GetUserStatsForGame method
type UserStatsForGameParams struct {
	SteamId  uint64              // SteamID of user
	AppId    uint32              // appid of game
	Format   config.OutputFormat // Format of the output
	Language *config.Language    // (optional) output Language
}

/*
Returns on global achievements overview of a specific game in percentages.

No key required.

Arguments
  - gameid
    AppID of the game you want the news of.
  - format
    Output format. json (default), xml or vdf.
*/
func (c Client) GetGlobalAchievementPercentagesForApp(params GlobalAchievementPercentageParams) (*model.GlobalAchievementPercentages, error) {
	version := "2"

	vals := url.Values{}
	vals.Set("gameid", strconv.FormatUint(params.GameId, 10))
	vals.Set("format", params.Format.String())

	if c.IsKeySet() {
		vals.Set("key", c.Key)
	}

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetGlobalAchievementPercentagesForAppEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamUserStats, versUrlEndpoint, vals)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.AchievementPercentagesResponse
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.GlobalAchievementPercentagesWrapper, nil

	case config.Xml:
		var result model.GlobalAchievementPercentages
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

// todo: create return structure
// func (c Client) GetGlobalStatsForGame(params GlobalGameStatsParams) (*model.GameStatsResponse, error) {
// 	version := "1"
// 	urlValues := url.Values{}
// 	urlValues.Add("appid", strconv.FormatUint(uint64(params.AppID), 10))
// 	urlValues.Add("count", strconv.FormatUint(uint64(params.Count), 10))
// 	urlValues.Set("format", params.Format.String())

// 	for _, name := range params.Names {
// 		urlValues.Add("name", name)
// 	}

// 	if params.StartDate != nil {
// 		urlValues.Add("startdate", strconv.FormatUint(uint64(*params.StartDate), 10))
// 	}
// 	if params.EndDate != nil {
// 		urlValues.Add("enddate", strconv.FormatUint(uint64(*params.EndDate), 10))
// 	}

// 	versUrlEndpoint := format.VersionedURLEndpoint{EndpointPath: "GetGlobalStatsForGame", Version: version}
// 	reqURL := urlHelper.RequestRequestURLFormatter(ISteamUserStats, versUrlEndpoint, urlValues)

// 	resp, err := c.HttpClient.Get(reqURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
// 	}

// 	// TODO
// 	// switch params.Format {
// 	// case config.Json:
// 	// 	var result model.AppNewsResponse
// 	// 	if _, err := decodeJSON(&result, resp.Body); err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	return &result.AppNews, nil

// 	// case config.Xml:
// 	// 	var result model.AppNews
// 	// 	if _, err := decodeXML(&result, resp.Body); err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	return &result, nil

// 	// default:
// 	// 	return nil, fmt.Errorf("unsupported format requested: %v", format)
// 	// }
// 	return nil, nil
// }

// TODO: model
func (c Client) GetNumberOfCurrentPlayers(params NumberOfCurrentPlayersParams) (*model.NumberOfCurrentPlayers, error) {
	version := "1"

	vals := url.Values{}
	vals.Set("gameid", strconv.FormatUint(params.GameId, 10))
	vals.Set("format", params.Format.String())

	if c.IsKeySet() {
		vals.Set("key", c.Key)
	}

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetNumberOfCurrentPlayersEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamUserStats, versUrlEndpoint, vals)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.NumberOfCurrentPlayersResponse
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.CurrentPlayers, nil

	case config.Xml:
		var result model.NumberOfCurrentPlayers
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

/*
Returns a list of achievements for this user by app id

# Key required

Arguments
  - steamid
    64 bit Steam ID to return friend list for.
  - appid
    The ID for the game you're requesting
  - l (Optional)
    Language. If specified, it will return language data for the requested language.
*/
func (c Client) GetPlayerAchievements(params PlayerAchievementsParams) (*model.PlayerAchievements, error) {
	if !c.IsKeySet() {
		return nil, errors.New(apiKeyErrorMessage)
	}
	version := "1"

	vals := url.Values{}
	vals.Set("key", c.Key)
	vals.Set("steamid", strconv.FormatUint(params.SteamId, 10))
	vals.Set("appid", strconv.FormatInt(int64(params.AppId), 10))
	vals.Set("format", params.Format.String())

	if params.Language != nil {
		vals.Set("l", params.Language.String())
	}

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetPlayerAchievementsEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamUserStats, versUrlEndpoint, vals)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.PlayerStatsWrapper
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.PlayerStats, nil

	case config.Xml:
		var result model.PlayerAchievements
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

// key required
func (c Client) GetSchemaForGame(params SchemaForGameParams) (*model.GameSchemaGame, error) {
	if !c.IsKeySet() {
		return nil, errors.New(apiKeyErrorMessage)
	}
	version := "2"

	vals := url.Values{}
	vals.Set("key", c.Key)
	vals.Set("appid", strconv.FormatInt(int64(params.AppId), 10))
	vals.Set("format", params.Format.String())

	if params.Language != nil {
		vals.Set("l", params.Language.String())
	}

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetSchemaForGameEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamUserStats, versUrlEndpoint, vals)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.SchemaForGameResponse
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.SchemaForGameWrapper, nil

	case config.Xml:
		var result model.GameSchemaGame
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

/*
Returns a list of achievements for this user by app id

# Key required

Arguments
  - steamid
    64 bit Steam ID to return friend list for.
  - appid
    The ID for the game you're requesting
  - l (Optional)
    Language. If specified, it will return language data for the requested language.
*/
func (c Client) GetUserStatsForGame(params UserStatsForGameParams) (*model.UserStats, error) {
	if !c.IsKeySet() {
		return nil, errors.New(apiKeyErrorMessage)
	}
	version := "2"

	vals := url.Values{}
	vals.Set("key", c.Key)
	vals.Set("steamid", strconv.FormatInt(int64(params.SteamId), 10))
	vals.Set("appid", strconv.FormatInt(int64(params.AppId), 10))
	vals.Set("format", params.Format.String())

	if params.Language != nil {
		vals.Set("l", params.Language.String())
	}

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetUserStatsForGameEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamUserStats, versUrlEndpoint, vals)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.UserStatsResponse
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.PlayerStats, nil

	case config.Xml:
		var result model.UserStats
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}
