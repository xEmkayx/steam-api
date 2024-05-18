package steamclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"

	urlHelper "github.com/xemkayx/steam-api/internal/urlHelper"
	"github.com/xemkayx/steam-api/pkg/steamclient/config"
	model "github.com/xemkayx/steam-api/pkg/steamclient/model/IPlayerService"
)

const (
	IPlayerService                 = "IPlayerService"
	GetOwnedGamesEndpoint          = "GetOwnedGames"          // v0001
	GetRecentlyPlayedGamesEndpoint = "GetRecentlyPlayedGames" // v0001
)

// Parameters for the GetOwnedGames method
type GetOwnedGamesParams struct {
	SteamId                int64               // The player we're asking about
	IncludeAppinfo         bool                // true if we want additional details (name, icon) about each game
	IncludePlayedFreeGames bool                // Free games are excluded by default. If this is set, free games the user has played will be returned.
	Format                 config.OutputFormat // Format of the output
	AppIdsFilter           []int64             // (optional) if set, restricts result set to the passed in apps
	Language               *config.Language    // (optional) output Language
}

// Parameters for the GetRecentlyPlayedGames method
type GetRecentlyPlayedGamesParams struct {
	SteamId int64               // The player we're asking about
	Count   int                 // The number of games to return (0/unset: all)
	Format  config.OutputFormat // Format of the output
}

/*
GetOwnedGames returns a list of games a player owns along with some playtime information, if the profile is publicly visible.
Private, friends-only, and other privacy settings are not supported unless you are asking for your own personal details
(ie the WebAPI key you are using is linked to the steamid you are requesting).

# Key required

Arguments
  - steamid
    The SteamID of the account.
  - include_appinfo
    Include game name and logo information in the output. The default is to return appids only.
  - include_played_free_games
    By default, free games like Team Fortress 2 are excluded (as technically everyone owns them).
    If include_played_free_games is set, they will be returned if the player has played them at some point.
    This is the same behavior as the games list on the Steam Community.
  - format
    Output format. json (default), xml or vdf.
  - appids_filter
    You can optionally filter the list to a set of appids.
*/
func (c Client) GetOwnedGames(params GetOwnedGamesParams) (*model.OwnedGames, error) {
	if !c.IsKeySet() {
		return nil, errors.New(apiKeyErrorMessage)
	}
	version := "1"

	inputJson := map[string]interface{}{
		"steamid":                   params.SteamId,
		"include_appinfo":           params.IncludeAppinfo,
		"include_played_free_games": params.IncludePlayedFreeGames,
	}

	if len(params.AppIdsFilter) > 0 {
		inputJson["appids_filter"] = params.AppIdsFilter
	}

	jsonBytes, err := json.Marshal(inputJson)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %s", err)
	}

	vals := url.Values{}
	vals.Set("key", c.Key)
	vals.Set("format", params.Format.String())
	vals.Set("input_json", string(jsonBytes))

	if params.Language != nil {
		vals.Set("l", params.Language.String())
	}

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetOwnedGamesEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(IPlayerService, versUrlEndpoint, vals)

	fmt.Println("URL:", url)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.OwnedGamesWrapper
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.OwnedGames, nil

	case config.Xml:
		var result model.OwnedGames
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

/*
GetRecentlyPlayedGames returns a list of games a player has played in the last two weeks, if the profile is publicly visible.
Private, friends-only, and other privacy settings are not supported unless you are asking for your own personal details
(ie the WebAPI key you are using is linked to the steamid you are requesting).

# Key required

Arguments
  - steamid
    The SteamID of the account.
  - count
    Optionally limit to a certain number of games (the number of games a person has played in the last 2 weeks is typically very small)
  - format
    Output format. json (default), xml or vdf.
*/
func (c Client) GetRecentlyPlayedGames(params GetRecentlyPlayedGamesParams) (*model.RecentlyPlayedGames, error) {
	if !c.IsKeySet() {
		return nil, errors.New(apiKeyErrorMessage)
	}
	version := "1"

	vals := url.Values{}
	vals.Set("key", c.Key)
	vals.Set("steamid", strconv.FormatInt(params.SteamId, 10))
	vals.Set("count", strconv.Itoa(params.Count))
	vals.Set("format", params.Format.String())

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetRecentlyPlayedGamesEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(IPlayerService, versUrlEndpoint, vals)
	fmt.Println("Url:", url)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.RecentlyPlayedGamesWrapper
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.RecentlyPlayedGames, nil

	case config.Xml:
		var result model.RecentlyPlayedGames
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

// TODO: other endpoints
