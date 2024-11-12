package structrs

type VKImage struct {
	Url         string `json:"url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	WithPadding int    `json:"with_padding"`
}
