package structrs

type VKAEP struct {
	Type        int    `json:"type"`
	Text        string `json:"text"`
	LabelText   string `json:"label_text"`
	ButtonText  string `json:"button_text"`
	IsAdNotEasy bool   `json:"is_ad_not_easy"`
}
