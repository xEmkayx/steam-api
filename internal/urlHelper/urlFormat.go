package urlhelper

import (
	"fmt"
	"net/url"

	"github.com/xemkayx/steam-api/pkg/steamclient/constant"
)

type VersionedURLEndpoint struct {
	EndpointPath string // Path of the endpoint
	Version      string // version of the endpoint
}

// format the steam url based on interface, endpoint, version and queries
func RequestURLFormatter(interf string, urlEndpoint VersionedURLEndpoint, query url.Values) string {
	return fmt.Sprintf("%s/%s/%s/v%s?%s", constant.SteamWebApiBaseURL, interf, urlEndpoint.EndpointPath, urlEndpoint.Version, query.Encode())
}
