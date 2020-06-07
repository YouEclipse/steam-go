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
