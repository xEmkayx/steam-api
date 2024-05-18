# Steam-API for Go

![GitHub issues](https://img.shields.io/github/issues/xEmkayx/steam-api)
![code size](https://img.shields.io/github/languages/code-size/xEmkayx/steam-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Staticcheck](https://img.shields.io/badge/static%20check-passing-brightgreen) 
![GitHub last commit](https://img.shields.io/github/last-commit/xEmkayx/steam-api)

![Go Version](https://img.shields.io/github/go-mod/go-version/xEmkayx/steam-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/xEmkayx/steam-api)](https://goreportcard.com/report/github.com/xEmkayx/steam-api)
[![Coverage Status](https://coveralls.io/repos/github/xEmkayx/steam-api/badge.svg?branch=master)](https://coveralls.io/github/xEmkayx/steam-api?branch=master)
[![GoDoc](https://godoc.org/github.com/xEmkayx/steam-api?status.svg)](https://godoc.org/github.com/xEmkayx/steam-api)


This API aims to be the most comprehensive API for Steam in Go. It aims to provide all the endpoints found in https://steamapi.xpaw.me/

All API-Endpoints specified in https://developer.valvesoftware.com/wiki/Steam_Web_API are ready to use in this library. 

# Example
## Client
All endpoints are callable on a variable of the Client struct:
```go
    httpClient := http.Client{}
    client := steamclient.New("your-api-key", httpClient)
```

Here, you need to specify a (customizable) HTTP client, which is used to send the requests, and an API key. You can obtain an API key by following Valve's documentation: https://steamcommunity.com/dev

Some requests don't explicitly require a key. If you only want to call such endpoints, you can create a client without an ID:
```go
    client := steamclient.NewClientWithoutId(httpClient)
```

Calling an Endpoint
To call an endpoint, you first need to create a parameter variable for the corresponding method. Optional parameters are marked as such and are pointer types, which means you don't have to specify them when creating the parameter variable.

We will use the GetUserStatsForGame endpoint as an example:

```go
// define the output format
f := config.Json

client := steamclient.New("your-api-key", httpClient)

// define the params struct
params := steamclient.UserStatsForGameParams{
    SteamId: 76561197972495328, 
    AppId: 440, 
    Format: config.OutputFormat(f), 
    // we could also specify the language here, but chose not to because it's an optional parameter
}

// get the returned body from the request
o, err := client.GetUserStatsForGame(params)
if err != nil {
    log.Fatal("Error:", err)
}

// pretty-print the returned value
fmt.Println(format.PrettyPrint(o, f))
``` 

## Returned Values
Every endpoint returns a specific structure in the specified format (JSON and XML; VDF is currently not implemented). This library returns the responses in a directly usable object format, instead of a string. For this endpoint, the returned struct looks like this:

```go
// a variable of this struct will be returned
type UserStatsResponse struct {
	PlayerStats UserStats `json:"playerstats" xml:"playerstats"`
}

type UserStats struct {
	SteamID      string                 `json:"steamID" xml:"steamID"`
	GameName     string                 `json:"gameName" xml:"gameName"`
	Achievements []UserStatsAchievement `json:"achievements" xml:"achievements>achievement"`
}

type UserStatsAchievement struct {
	Name     string `json:"name" xml:"name"`
	Achieved int    `json:"achieved" xml:"achieved"` // 1 - true, 0 - false
}
```

# Status

Below are the status of the implementations of all the interfaces Steam provides.

## Steam-API

- [ ] IAccountCartService
- [ ] IAccountLinkingService
- [ ] IAccountPrivateAppsService
- [ ] IAuctionService
- [ ] IAuthenticationService
- [ ] IAuthenticationSupportService
- [ ] IBroadcastClientService
- [ ] IBroadcastService
- [ ] IChatRoomService
- [ ] ICheatReportingService
- [ ] ICheckoutService
- [ ] IClanFAQSService
- [ ] IClanService
- [ ] IClientCommService
- [ ] IClientMetricsService
- [ ] ICloudService
- [ ] ICommunityLinkFilterService
- [ ] ICommunityService
- [ ] IContentFilteringService
- [ ] IContentServerConfigService
- [ ] IContentServerDirectoryService
- [ ] ICredentialsService
- [ ] IDailyDealService
- [ ] IDataPublisherService
- [ ] IDeviceAuthService
- [ ] IEconMarketService
- [ ] IEconService
- [ ] IEmbeddedClientService
- [ ] IFamilyGroupsService
- [ ] IFriendMessagesService
- [ ] IFriendsListService
- [ ] IGameCoordinator
- [ ] IGameInventory
- [ ] IGameNotificationsService
- [ ] IGameRecordingClipService
- [ ] IGameServersService
- [ ] IHelpRequestLogsService
- [ ] IInventoryService
- [ ] ILobbyMatchmakingService
- [ ] ILoyaltyRewardsService
- [ ] IMarketingMessagesService
- [ ] IMobileAppService
- [ ] IMobileAuthService
- [ ] IMobileDeviceService
- [ ] IMobileNotificationService
- [ ] INewsService
- [ ] IOnlinePlayService
- [ ] IParentalService
- [ ] IPartnerMembershipInviteService
- [ ] IPartnerStoreBrowseService
- [ ] IPhoneService
- [ ] IPhysicalGoodsService
- [ ] IPlayerService
- [ ] IProductInfoService
- [ ] IPromotionEventInvitesService
- [ ] IPromotionPlanningService
- [ ] IPromotionStatsService
- [ ] IPublishedFileService
- [ ] IPublishingService
- [ ] IQuestService
- [ ] IRemoteClientService
- [ ] ISaleFeatureService
- [ ] ISaleItemRewardsService
- [ ] IShoppingCartService
- [ ] ISiteLicenseService
- [ ] ISteamApps
- [ ] ISteamAwardsService
- [ ] ISteamBitPay
- [ ] ISteamBoaCompra
- [ ] ISteamBroadcast
- [ ] ISteamCDN
- [ ] ISteamChartsService
- [ ] ISteamCloudGaming
- [ ] ISteamCommunity
- [ ] ISteamDirectory
- [ ] ISteamEconomy
- [ ] ISteamEnvoy
- [ ] ISteamGameServerStats
- [ ] ISteamLeaderboards
- [ ] ISteamLearnService
- [ ] ISteamMicroTxn
- [ ] ISteamMicroTxnSandbox
- [ ] ISteamNews
- [ ] ISteamNodwin
- [ ] ISteamNotificationService
- [ ] ISteamPayPalPaymentsHub
- [ ] ISteamPublishedItemSearch
- [ ] ISteamPublishedItemVoting
- [ ] ISteamRemoteStorage
- [ ] ISteamSpecialSurvey
- [ ] ISteamTVService
- [ ] ISteamUser
- [ ] ISteamUserAuth
- [ ] ISteamUserOAuth
- [x] ISteamUserStats
    - [ ] GetGlobalStatsForGame method
- [ ] ISteamWebAPIUtil
- [ ] ISteamWorkshop
- [ ] IStoreAppSimilarityService
- [ ] IStoreBrowseService
- [ ] IStoreMarketingService
- [ ] IStoreQueryService
- [ ] IStoreSalesService
- [ ] IStoreService
- [ ] IStoreTopSellersService
- [ ] ITestExternalPrivilegeService
- [ ] ITrustService
- [ ] ITwoFactorService
- [ ] IUserAccountService
- [ ] IUserGameNotesService
- [ ] IUserReviewsService
- [ ] IVACManagementService
- [ ] IVideoService
- [ ] IWorkshopService

## CS:GO/CS2 API
- [ ] ICSGOServers_730
- [ ] ICSGOStreamSystem_730
- [ ] ICSGOTournaments_730
- [ ] IEconItems_730
- [ ] IGCVersion_730

## Dota 2 API
- [ ] IDOTA2AutomatedTourney_570
- [ ] IDOTA2CustomGames_570
- [ ] IDOTA2Events_570
- [ ] IDOTA2Fantasy_570
- [ ] IDOTA2Guild_570
- [ ] IDOTA2League_570
- [ ] IDOTA2MatchStats_570
- [ ] IDOTA2Match_570
- [ ] IDOTA2Operations_570
- [ ] IDOTA2Plus_570
- [ ] IDOTA2StreamSystem_570
- [ ] IDOTA2Teams_570
- [ ] IDOTA2Ticket_570
- [ ] IDOTAChat_570
- [ ] IEconDOTA2_570
- [ ] IEconItems_570
- [ ] IGCVersion_570

## Team Fortress 2 API
- [ ] IEconItems_440
- [ ] IGCVersion_440
- [ ] ITFItems_440
- [ ] ITFPromos_440
- [ ] ITFSystem_440

## Portal 2 API
- [ ] IEconItems_620
- [ ] IPortal2Leaderboards_620
- [ ] ITFPromos_620

## Dota Underlords API
- [ ] IClientStats_1046930
- [ ] IEconItems_1046930
- [ ] IGCVersion_1046930

## Artifact Classic API
- [ ] IEconItems_583950
- [ ] IGCVersion_583950

## Artifact Foundry API
- [ ] IEconItems_1269260
- [ ] IGCVersion_1269260