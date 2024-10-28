package structs

type AnilibriaTitle struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Names struct {
		Ru string `json:"ru"`
		En string `json:"en"`
	} `json:"names"`
	Announce string `json:"announce"`
	Status   struct {
		String string `json:"string"`
		Code   int    `json:"code"`
	} `json:"status"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
}
