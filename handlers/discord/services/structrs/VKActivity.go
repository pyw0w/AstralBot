package structrs

type VKActivity struct {
	Comments      []VKComments `json:"comments"`
	PostAuthorId  int          `json:"post_author_id"`
	Type          string       `json:"type"`
	Discriminator string       `json:"discriminator"`
}
