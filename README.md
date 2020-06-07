# steam-go [WIP]
🎮 Go library for accessing the [Steamworks Web API](https://partner.steamgames.com/doc/webapi)

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/YouEclipse/steam-go/pkg) ![Go](https://github.com/YouEclipse/steam-go/workflows/Go/badge.svg)



## 🏗️ Install

#### Requirments
> Go version 1.13+

#### Install
```
go get github.com/YouEclipse/steam-go
```


## 🚀 Quick start
```golang

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	steam "github.com/YouEclipse/steam-go/pkg"
)

func main() {

	steamClient := steam.NewClient(os.Getenv("STEAM_API_KEY"), nil)
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)
	ctx := context.Background()
	params := &steam.GetOwnedGamesParams{
		SteamID:                steamID,
		IncludeAppInfo:         true,
		IncludePlayedFreeGames: true,
	}

	result, err := steamClient.IPlayerService.GetOwnedGames(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}



```


## ✔️ Features v0.1.0



- [ ]  [IBroadcastService Interface](https://partner.steamgames.com/doc/webapi/IBroadcastService)
- [ ]  [ICheatReportingService Interface](https://partner.steamgames.com/doc/webapi/ICheatReportingService)
- [ ]  [ICloudService Interface](https://partner.steamgames.com/doc/webapi/icloudservice)
- [ ]  [IEconMarketService Interface](https://partner.steamgames.com/doc/webapi/IEconMarketService)
- [ ]  [IEconService Interface](https://partner.steamgames.com/doc/webapi/IEconService)
- [ ]  [IGameInventory Interface](https://partner.steamgames.com/doc/webapi/IGameInventory)
- [ ]  [IGameNotificationsService Interface](https://partner.steamgames.com/doc/webapiIGameNotificationsService)
- [ ]  [IGameServersService Interface](https://partner.steamgames.com/doc/webapi/IGameServersService)
- [ ]  [IInventoryService Interface](https://partner.steamgames.com/doc/webapi/IInventoryService)
- [ ]  [ILobbyMatchmakingService Interface](https://partner.steamgames.com/doc/webapi/ILobbyMatchmakingService)
- [x]  [IPlayerService Interface](https://partner.steamgames.com/doc/webapi/IPlayerService)
  - [ ] [GetRecentlyPlayedGames](https://partner.steamgames.com/doc/webapi/IPlayerService#GetRecentlyPlayedGames)
  - [x] [GetOwnedGames](https://partner.steamgames.com/doc/webapi/IPlayerService#GetOwnedGames)
  - [x] [GetSteamLevel](https://partner.steamgames.com/doc/webapi/IPlayerService#GetSteamLevel)
  - [x] [GetBadges](https://partner.steamgames.com/doc/webapi/IPlayerService#GetBadges)
  - [x] [GetCommunityBadgeProgress](https://partner.steamgames.com/doc/webapi/IPlayerService#GetCommunityBadgeProgress)
  - [x] [IsPlayingSharedGame](https://partner.steamgames.com/doc/webapi/IPlayerService#IsPlayingSharedGame)
- [ ]  [IPublishedFileService Interface](https://partner.steamgames.com/doc/webapi/IPublishedFileService)
- [ ]  [ISiteLicenseService Interface](https://partner.steamgames.com/doc/webapi/isitelicenseservice)
- [ ]  [ISteamApps Interface](https://partner.steamgames.com/doc/webapi/ISteamApps)
- [ ]  [ISteamCommunity Interface](https://partner.steamgames.com/doc/webapi/ISteamCommunity)
- [ ]  [ISteamEconomy Interface](https://partner.steamgames.com/doc/webapi/ISteamEconomy)
- [ ]  [ISteamGameServerStats Interface](https://partner.steamgames.com/doc/webapi/ISteamGameServerStats)
- [ ]  [ISteamLeaderboards Interface](https://partner.steamgames.com/doc/webapi/ISteamLeaderboards)
- [ ]  [ISteamMicroTxn Interface](https://partner.steamgames.com/doc/webapi/ISteamMicroTxn)
- [ ]  [ISteamMicroTxnSandbox Interface](https://partner.steamgames.com/doc/webapi/ISteamMicroTxnSandbox)
- [ ]  [ISteamNews Interface](https://partner.steamgames.com/doc/webapi/ISteamNews)
- [ ]  [ISteamPublishedItemSearch Interface](https://partner.steamgames.com/doc/webapiISteamPublishedItemSearch)
- [ ]  [ISteamPublishedItemVoting Interface](https://partner.steamgames.com/doc/webapiISteamPublishedItemVoting)
- [ ]  [ISteamRemoteStorage Interface](https://partner.steamgames.com/doc/webapi/ISteamRemoteStorage)
- [ ]  [ISteamUserAuth Interface](https://partner.steamgames.com/doc/webapi/ISteamUserAuth)
- [ ]  [ISteamUser Interface](https://partner.steamgames.com/doc/webapi/ISteamUser)
- [ ]  [ISteamUserStats Interface](https://partner.steamgames.com/doc/webapi/ISteamUserStats)
- [ ]  [ISteamWebAPIUtil Interface](https://partner.steamgames.com/doc/webapi/ISteamWebAPIUtil)
- [ ]  [IWorkshopService Interface](https://partner.steamgames.com/doc/webapi/IWorkshopService)

...


## 📄 License
This project is licensed under [Apache-2.0](./LICENSE)