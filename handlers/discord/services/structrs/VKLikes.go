package structrs

type VKLikes struct {
	CanLike        int  `json:"can_like"`
	Count          int  `json:"count"`
	UserLikes      int  `json:"user_likes"`
	CanPublish     int  `json:"can_publish"`
	CanLikeByGroup int  `json:"can_like_by_group"`
	GroupLiked     bool `json:"group_liked"`
	RepostDisabled bool `json:"repost_disabled"`
}
