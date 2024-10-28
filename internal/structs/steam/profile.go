package steam

type Profile struct {
	Response struct {
		PlayerCount int `json:"player_count"`
	} `json:"response"`
}
