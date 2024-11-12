package structrs

type VKDoc struct {
	Id      int       `json:"id"`
	OwnerId int       `json:"owner_id"`
	Title   string    `json:"title"`
	Size    int       `json:"size"`
	Ext     string    `json:"ext"`
	Date    int       `json:"date"`
	Type    int       `json:"type"`
	Url     string    `json:"url"`
	Preview VKPreview `json:"preview"`
}
