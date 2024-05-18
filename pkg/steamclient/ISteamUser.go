package steamclient

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	urlHelper "steam-api/internal/urlHelper"
	"steam-api/pkg/steamclient/config"
	model "steam-api/pkg/steamclient/model/ISteamUser"
	"strconv"
	"strings"
)

const (
	ISteamUser                 = "ISteamUser"
	GetPlayerSummariesEndpoint = "GetPlayerSummaries" // v0002
	GetFriendListEndpoint      = "GetFriendList"      // v0001
)

// Parameters for the GetFriendList method
type GetFriendListParams struct {
	SteamId      int64                          // SteamID of user
	Relationship config.FriendsListRelationship // 	relationship type
	Format       config.OutputFormat            // Format of the output
}

// Parameters for the GetPlayerSummaries method
type GetPlayerSummariesParams struct {
	SteamIds []int64             // Comma-delimited list of SteamIDs (max: 100)
	Format   config.OutputFormat // Format of the output
}

/*
Returns basic profile information for a list of 64-bit Steam IDs.

key required

Arguments
  - steamids
    Comma-delimited list of 64 bit Steam IDs to return profile information for. Up to 100 Steam IDs can be requested.
  - format
    Output format. json (default), xml or vdf.
*/
func (c Client) GetPlayerSummaries(params GetPlayerSummariesParams) (*model.PlayerSummaries, error) {
	if !c.IsKeySet() {
		return nil, errors.New(apiKeyErrorMessage)
	}
	version := "2"

	if len(params.SteamIds) > 100 {
		return nil, errors.New("you have provided too many Steam IDs. reduce the amount to 100")
	}

	strSlice := make([]string, len(params.SteamIds))
	for i, id := range params.SteamIds {
		strSlice[i] = strconv.FormatInt(id, 10)
	}

	vals := url.Values{}
	vals.Set("key", c.Key)
	vals.Set("steamids", strings.Join(strSlice, ","))
	vals.Set("format", params.Format.String())

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetPlayerSummariesEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamUser, versUrlEndpoint, vals)
	fmt.Println("Url:", url)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.PlayerSummariesWrapper
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.PlayerSums, nil

	case config.Xml:
		var result model.PlayerSummaries
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

/*
Returns the friend list of any Steam user, provided their Steam Community profile visibility is set to "Public".

# Key required

Arguments
  - steamid
    64 bit Steam ID to return friend list for.
  - relationship
    Relationship filter. Possibles values: all, friend.
  - format
    Output format. json (default), xml or vdf.
*/
func (c Client) GetFriendList(params GetFriendListParams) (*model.FriendList, error) {
	if !c.IsKeySet() {
		return nil, errors.New(apiKeyErrorMessage)
	}
	version := "1"

	vals := url.Values{}
	vals.Set("key", c.Key)
	vals.Set("steamid", strconv.FormatInt(params.SteamId, 10))
	vals.Set("format", params.Format.String())

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetFriendListEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamUser, versUrlEndpoint, vals)
	fmt.Println("Url:", url)

	resp, err := c.getRequest(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.FriendListWrapper
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.FriendsList, nil

	case config.Xml:
		var result model.FriendList
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

// TODO: other endpoints
