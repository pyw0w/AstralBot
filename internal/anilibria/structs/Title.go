package structs

type Title struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Names struct {
		Ru          string `json:"ru"`
		En          string `json:"en"`
		Alternative string `json:"alternative"`
	} `json:"names"`
	// ...other fields...
}
