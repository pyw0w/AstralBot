package structrs

type VKResponse struct {
	Count    int      `json:"count"`
	Items    []VKItem `json:"items"`
	NextFrom int      `json:"next_from"`
}
