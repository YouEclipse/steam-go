package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// IPlayerService provides additional methods for interacting with Steam Users.
// See https://partner.steamgames.com/doc/webapi/IPlayerService for more information.
type IPlayerService service

// GetRecentlyPlayedGamesParams defines the parameters to  GetRecentlyPlayedGames.
type GetRecentlyPlayedGamesParams struct {
	// The player we're asking about
	SteamID uint64 `json:"steamid"`
	// The number of games to return (0/unset: all)
	Count uint32 `json:"count"`
}

// GetRecentlyPlayedGamesResult defines the result of GetRecentlyPlayedGames.
type GetRecentlyPlayedGamesResult struct {
	//TODO
}

// GetRecentlyPlayedGames gets information about a player's recently played games.
func (s *IPlayerService) GetRecentlyPlayedGames(ctx context.Context, params *GetRecentlyPlayedGamesParams) (*GetRecentlyPlayedGamesResult, error) {
	endpoint := "IPlayerService/GetRecentlyPlayedGames/v1/"

	ret := &Result{
		Response: &GetRecentlyPlayedGamesResult{},
	}
	err := s.doGet(ctx, endpoint, params, ret)
	if err != nil {
		return nil, err
	}

	return ret.Response.(*GetRecentlyPlayedGamesResult), nil

}

// GetOwnedGamesParams defines the parameters to GetRecentlyPlayedGames.
type GetOwnedGamesParams struct {
	// The player we're asking about
	SteamID uint64 `json:"steamid"`
	// true if we want additional details (name, icon) about each game
	IncludeAppInfo bool `json:"include_appinfo"`
	// Free games are excluded by default. If this is set, free games the user has played will be returned
	IncludePlayedFreeGames bool `json:"include_played_free_games"`
	// if set, restricts result set to the passed in apps,
	// actually it should passed like appids_filter[0]=8930&appids_filter[1]=500
	AppIDsFilter []uint32 `json:"appids_filter"`
}

// GetOwnedGamesResult defines the result of GetRecentlyPlayedGames.
type GetOwnedGamesResult struct {
	GameCount int     `json:"game_count"`
	Games     []*Game `json:"games"`
}

// Game defines the game in GetOwnedGamesResult.
type Game struct {
	Appid                    int    `json:"appid"`
	Name                     string `json:"name"`
	PlaytimeForever          int    `json:"playtime_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	ImgLogoURL               string `json:"img_logo_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
	PlaytimeMacForever       int    `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
}

// GetOwnedGames return a list of games owned by the player.
func (s *IPlayerService) GetOwnedGames(ctx context.Context, params *GetOwnedGamesParams) (*GetOwnedGamesResult, error) {
	endpoint := "IPlayerService/GetOwnedGames/v1/"

	ret := &Result{
		Response: &GetOwnedGamesResult{},
	}
	err := s.doGet(ctx, endpoint, params, ret)
	if err != nil {
		return nil, err
	}

	return ret.Response.(*GetOwnedGamesResult), nil
}

// GetSteamLevelParams defines the parameters to  GetRecentlyPlayedGames.
type GetSteamLevelParams struct {
	// The player we're asking about
	SteamID uint64 `json:"steamid"`
}

// GetSteamLevelResult defines the result of GetRecentlyPlayedGames.
type GetSteamLevelResult struct {
	PlayerLevel uint64 `json:"player_level"`
}

// GetSteamLevel returns the Steam Level of a user.
func (s *IPlayerService) GetSteamLevel(ctx context.Context, params *GetSteamLevelParams) (*GetSteamLevelResult, error) {
	endpoint := "IPlayerService/GetSteamLevel/v1/"

	var ret = &Result{
		Response: &GetSteamLevelResult{},
	}
	err := s.doGet(ctx, endpoint, params, ret)
	if err != nil {
		return nil, err
	}

	return ret.Response.(*GetSteamLevelResult), nil
}

// GetBadgesParams defines the parameters to  GetRecentlyPlayedGames.
type GetBadgesParams struct {
	// The player we're asking about
	SteamID uint64 `json:"steamid"`
}

// GetBadgesResult defines the result of GetRecentlyPlayedGames.
type GetBadgesResult struct {
	Badges                     []*Badge `json:"badges"`
	PlayerXp                   int      `json:"player_xp"`
	PlayerLevel                int      `json:"player_level"`
	PlayerXpNeededToLevelUp    int      `json:"player_xp_needed_to_level_up"`
	PlayerXpNeededCurrentLevel int      `json:"player_xp_needed_current_level"`
}

// Badge defines the badge in GetBadgesResult.
type Badge struct {
	Badgeid        int32 `json:"badgeid"`
	Level          int   `json:"level"`
	CompletionTime int   `json:"completion_time"`
	Xp             int   `json:"xp"`
	Scarcity       int   `json:"scarcity"`
}

// GetBadges gets badges that are owned by a specific user.
func (s *IPlayerService) GetBadges(ctx context.Context, params *GetBadgesParams) (*GetBadgesResult, error) {
	endpoint := "IPlayerService/GetBadges/v1/"

	var ret = &Result{
		Response: &GetBadgesResult{},
	}
	err := s.doGet(ctx, endpoint, params, ret)
	if err != nil {
		return nil, err
	}

	return ret.Response.(*GetBadgesResult), nil
}

// GetCommunityBadgeProgressParams defines the parameters to GetCommunityBadgeProgress.
type GetCommunityBadgeProgressParams struct {
	// The player we're asking about
	SteamID uint64 `json:"steamid"`
	// The badge we're asking about
	BadgeID int32 `json:"badgeid"`
}

// GetCommunityBadgeProgressResult defines the result of GetCommunityBadgeProgress.
type GetCommunityBadgeProgressResult struct {
	Quests []BadgeQuests `json:"quests"`
}

// BadgeQuests defines the quests of GetCommunityBadgeProgressResult.
type BadgeQuests struct {
	Questid   int  `json:"questid"`
	Completed bool `json:"completed"`
}

// GetCommunityBadgeProgress gets all the quests needed to get the specified badge, and which are completed.
func (s *IPlayerService) GetCommunityBadgeProgress(ctx context.Context, params *GetCommunityBadgeProgressParams) (
	*GetCommunityBadgeProgressResult, error) {
	endpoint := "IPlayerService/GetCommunityBadgeProgress/v1/"

	var ret = &Result{
		Response: &GetCommunityBadgeProgressResult{},
	}
	err := s.doGet(ctx, endpoint, params, ret)
	if err != nil {
		return nil, err
	}

	return ret.Response.(*GetCommunityBadgeProgressResult), nil
}

// IsPlayingSharedGameParams defines the parameters to  GetRecentlyPlayedGames.
type IsPlayingSharedGameParams struct {
	// The player we're asking about
	SteamID uint64 `json:"steamid"`
	// The game player is currently playing
	AppIDPlaying uint32 `json:"appid_playing"`
}

// IsPlayingSharedGameResult defines the result of IsPlayingSharedGame.
type IsPlayingSharedGameResult struct {
	LenderSteamID string `json:"lender_steamid"`
}

// IsPlayingSharedGame returns valid lender SteamID if game currently played is borrowed.
func (s *IPlayerService) IsPlayingSharedGame(ctx context.Context, params *IsPlayingSharedGameParams) (
	*IsPlayingSharedGameResult, error) {
	endpoint := "IPlayerService/IsPlayingSharedGame/v1/"

	var ret = &Result{
		Response: &IsPlayingSharedGameResult{},
	}
	err := s.doGet(ctx, endpoint, params, ret)
	if err != nil {
		return nil, err
	}

	return ret.Response.(*IsPlayingSharedGameResult), nil
}

func (s *IPlayerService) doGet(ctx context.Context, endpoint string, params, result interface{}) error {
	q := url.Values{}
	inputJSON, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("marshal params err :%w", err)
	}
	q.Set("input_json", string(inputJSON))

	respBytes, err := (*service)(s).get(ctx, endpoint, q)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return fmt.Errorf("unmarshal error  :%w", err)
	}
	return nil
}
