package structrs

type VKThread struct {
	Count           int          `json:"count"`
	Items           []VKComments `json:"items"`
	CanPost         int          `json:"can_post"`
	ShowReplyButton int          `json:"show_reply_button"`
	GroupsCanPost   bool         `json:"groups_can_post"`
}
