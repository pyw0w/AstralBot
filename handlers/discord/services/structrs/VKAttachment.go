package structrs

type VKAttachment struct {
	Type  string  `json:"type"`
	Doc   VKDoc   `json:"doc"`
	Video VKVideo `json:"video"`
	Photo VKPhoto `json:"photo"`
	Style string  `json:"style"`
}