package structrs

type VKComments struct {
	CanPost       int      `json:"can_post"`
	Count         int      `json:"count"`
	GroupsCanPost bool     `json:"groups_can_post"`
	Id            int      `json:"id"`
	FromId        int      `json:"from_id"`
	Date          int      `json:"date"`
	Text          string   `json:"text"`
	PostId        int      `json:"post_id"`
	OwnerId       int      `json:"owner_id"`
	ParentsStack  string   `json:"parents_stack"`
	Likes         VKLikes  `json:"likes"`
	Thread        VKThread `json:"thread"`
}
