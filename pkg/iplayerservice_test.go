package steam

import (
	"context"
	"os"
	"strconv"
	"testing"
)

func TestIPlayerService_GetRecentlyPlayedGames(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *GetRecentlyPlayedGamesParams
	}

	steam := NewClient(os.Getenv("STEAM_API_KEY"), nil)

	ctx := context.Background()
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)

	tests := []struct {
		name string
		args args
	}{{
		name: "ALL",
		args: args{
			ctx: ctx,
			params: &GetRecentlyPlayedGamesParams{
				SteamID: steamID,
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result, err := steam.IPlayerService.GetRecentlyPlayedGames(tt.args.ctx, tt.args.params)
			if err != nil {
				t.Error(err)
			}
			t.Logf("%+v", result)

		})
	}

}

func TestIPlayerService_GetOwnedGames(t *testing.T) {

	type args struct {
		ctx    context.Context
		params *GetOwnedGamesParams
	}

	steam := NewClient(os.Getenv("STEAM_API_KEY"), nil)

	ctx := context.Background()
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "ALL",
			args: args{
				ctx: ctx,
				params: &GetOwnedGamesParams{
					SteamID:                steamID,
					IncludeAppInfo:         true,
					IncludePlayedFreeGames: true,
				},
			},
		},
		{
			name: "IncludeAppInfo=false",
			args: args{
				ctx: ctx,
				params: &GetOwnedGamesParams{
					SteamID:                steamID,
					IncludeAppInfo:         false,
					IncludePlayedFreeGames: true,
				},
			},
		},
		{
			name: "IncludePlayedFreeGames=false",
			args: args{
				ctx: ctx,
				params: &GetOwnedGamesParams{
					SteamID:                steamID,
					IncludeAppInfo:         true,
					IncludePlayedFreeGames: false,
				},
			},
		},
		{
			name: "AppIDsFilter=[500]",
			args: args{
				ctx: ctx,
				params: &GetOwnedGamesParams{
					SteamID:                steamID,
					IncludeAppInfo:         true,
					IncludePlayedFreeGames: false,
					AppIDsFilter:           []uint32{500},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := steam.IPlayerService.GetOwnedGames(tt.args.ctx, tt.args.params)
			if err != nil {
				t.Error(err)
			}
			//t.Logf("%+v", result)
		})
	}
}

func TestIPlayerService_GetSteamLevel(t *testing.T) {

	type args struct {
		ctx    context.Context
		params *GetSteamLevelParams
	}

	steam := NewClient(os.Getenv("STEAM_API_KEY"), nil)

	ctx := context.Background()
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)
	tests := []struct {
		name string
		args args
	}{{
		name: "ALL",
		args: args{
			ctx: ctx,
			params: &GetSteamLevelParams{
				SteamID: steamID,
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := steam.IPlayerService.GetSteamLevel(tt.args.ctx, tt.args.params)
			if err != nil {
				t.Error(err)
			}
			//t.Logf("%+v", result)

		})
	}
}

func TestIPlayerService_GetBadges(t *testing.T) {

	type args struct {
		ctx    context.Context
		params *GetBadgesParams
	}

	steam := NewClient(os.Getenv("STEAM_API_KEY"), nil)

	ctx := context.Background()
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)
	tests := []struct {
		name string
		args args
	}{{
		name: "ALL",
		args: args{
			ctx: ctx,
			params: &GetBadgesParams{
				SteamID: steamID,
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := steam.IPlayerService.GetBadges(tt.args.ctx, tt.args.params)
			if err != nil {
				t.Error(err)
			}
			//t.Logf("%+v", result)

		})
	}
}

func TestIPlayerService_GetCommunityBadgeProgress(t *testing.T) {

	type args struct {
		ctx    context.Context
		params *GetCommunityBadgeProgressParams
	}

	steam := NewClient(os.Getenv("STEAM_API_KEY"), nil)

	ctx := context.Background()
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ALL",
			args: args{
				ctx: ctx,
				params: &GetCommunityBadgeProgressParams{
					SteamID: steamID,
				},
			},
		},
		{
			name: "BadgeID=7",
			args: args{
				ctx: ctx,
				params: &GetCommunityBadgeProgressParams{
					SteamID: steamID,
					BadgeID: 7,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := steam.IPlayerService.GetCommunityBadgeProgress(tt.args.ctx, tt.args.params)
			if err != nil {
				t.Error(err)
			}
			//t.Logf("%+v", result)

		})
	}
}

func TestIPlayerService_IsPlayingSharedGame(t *testing.T) {

	type args struct {
		ctx    context.Context
		params *IsPlayingSharedGameParams
	}

	steam := NewClient(os.Getenv("STEAM_API_KEY"), nil)

	ctx := context.Background()
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)
	tests := []struct {
		name string
		args args
	}{{
		name: "ALL",
		args: args{
			ctx: ctx,
			params: &IsPlayingSharedGameParams{
				SteamID:      steamID,
				AppIDPlaying: 500,
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := steam.IPlayerService.IsPlayingSharedGame(tt.args.ctx, tt.args.params)
			if err != nil {
				t.Error(err)
			}
			//t.Logf("%+v", result)

		})
	}
}
