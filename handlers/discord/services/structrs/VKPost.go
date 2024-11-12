package structrs

type VKPost struct {
	Id          int            `json:"id"`
	Date        int            `json:"date"`
	Text        string         `json:"text"`
	Attachments []VKAttachment `json:"attachments"`
}
