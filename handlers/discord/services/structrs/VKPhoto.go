package structrs

type VKPhoto struct {
	AlbumId      string    `json:"album_id"`
	Date         int       `json:"date"`
	Id           string    `json:"id"`
	OwnerId      string    `json:"owner_id"`
	PostId       string    `json:"post_id"`
	Sizes        []VKSizes `json:"sizes"`
	Text         string    `json:"text"`
	UserId       int       `json:"user_id"`
	WebViewToken string    `json:"web_view_token"`
	HasTags      bool      `json:"has_tags"`
	IsLicensed   int       `json:"is_licensed"`
	IsUnsafe     int       `json:"is_unsafe"`
	AccessKey    string    `json:"access_key"`
}
