package structs

type SteamProfile struct {
	Response struct {
		PlayerCount int `json:"player_count"`
	} `json:"response"`
}
