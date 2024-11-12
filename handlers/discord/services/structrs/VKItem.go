package structrs

type VKItem struct {
	InnerType                   string          `json:"inner_type"`
	AdsEasyPromote              VKAEP           `json:"ads_easy_promote"`
	Donut                       VKDonut         `json:"donut"`
	Comments                    VKComments      `json:"comments"`
	MarkedAsAds                 int             `json:"marked_as_ads"`
	Activity                    VKActivity      `json:"activity"`
	ShortTextRate               float32         `json:"short_text_rate"`
	CompactAttachmentsBeforeCut int             `json:"compact_attachments_before_cut"`
	Hash                        string          `json:"hash"`
	Type                        string          `json:"type"`
	Attachments                 []VKAttachments `json:"attachments"`
	Date                        int             `json:"date"`
	FromId                      int             `json:"from_id"`
	Id                          int             `json:"id"`
	IsFavorite                  bool            `json:"is_favorite"`
	Likes                       VKLikes         `json:"likes"`
	OwnerId                     int             `json:"owner_id"`
	PostSource                  VKPostSource    `json:"post_source"`
	PostType                    string          `json:"post_type"`
	Reposts                     VKReposts       `json:"reposts"`
	Text                        string          `json:"text"`
	Views                       VKViews         `json:"views"`
}
