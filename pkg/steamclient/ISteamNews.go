package steamclient

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	urlHelper "github.com/xemkayx/steam-api/internal/urlHelper"
	"github.com/xemkayx/steam-api/pkg/steamclient/config"
	model "github.com/xemkayx/steam-api/pkg/steamclient/model/ISteamNews"
)

const (
	ISteamNews            = "ISteamNews"
	GetNewsForAppEndpoint = "GetNewsForApp" // v0002
)

type GetNewsForAppParams struct {
	AppId     uint32              // AppID to retrieve news for
	Count     uint32              // # of posts to retrieve (default 20)
	MaxLength uint32              // Maximum length for the content to return, if this is 0 the full content is returned, if it's less then a blurb is generated to fit.
	Format    config.OutputFormat // Format of the output
}

/*
GetNewsForApp returns the latest of a game specified by its appID.

Arguments:
  - appid
    AppID of the game you want the news of.
  - count
    How many news enties you want to get returned.
  - maxlength
    Maximum length of each news entry.
  - format
    Output format. json (default), xml or vdf.

returns: output in the specified format
*/
func (c Client) GetNewsForApp(params GetNewsForAppParams) (*model.AppNews, error) {
	version := "2"

	vals := url.Values{}
	vals.Set("appid", strconv.FormatUint(uint64(params.AppId), 10))
	vals.Set("count", strconv.FormatUint(uint64(params.Count), 10))
	vals.Set("maxlength", strconv.FormatUint(uint64(params.MaxLength), 10))
	vals.Set("format", params.Format.String())

	if c.IsKeySet() {
		vals.Set("key", c.Key)
	}

	versUrlEndpoint := urlHelper.VersionedURLEndpoint{EndpointPath: GetNewsForAppEndpoint, Version: version}
	url := urlHelper.RequestURLFormatter(ISteamNews, versUrlEndpoint, vals)

	resp, err := c.getRequest(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code was " + fmt.Sprint(resp.StatusCode))
	}
	defer resp.Body.Close()

	switch params.Format {
	case config.Json:
		var result model.AppNewsResponse
		if _, err := decodeJSON(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result.AppNews, nil

	case config.Xml:
		var result model.AppNews
		if _, err := decodeXML(&result, resp.Body); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, fmt.Errorf("unsupported format requested: %v", params.Format)
	}
}

// TODO: implement other endpoint
