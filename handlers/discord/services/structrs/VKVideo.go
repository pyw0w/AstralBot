package structrs

type VKVideo struct {
	ResponseType  string         `json:"response_type"`
	AccessKey     string         `json:"access_key"`
	CanComment    int            `json:"can_comment"`
	CanLike       int            `json:"can_like"`
	CanRepost     int            `json:"can_repost"`
	CanSubscribe  int            `json:"can_subscribe"`
	CanAddToFaves int            `json:"can_add_to_faves"`
	CanAdd        int            `json:"can_add"`
	Comments      int            `json:"comments"`
	Date          int            `json:"date"`
	Description   string         `json:"description"`
	Duration      int            `json:"duration"`
	Image         []VKImage      `json:"image"`
	FirstFrame    []VKFirstFrame `json:"first_frame"`
	Width         int            `json:"width"`
	Height        int            `json:"height"`
	Id            int            `json:"id"`
	OwnerId       int            `json:"owner_id"`
	Title         string         `json:"title"`
	IsFavorite    bool           `json:"is_favorite"`
	TrackCode     string         `json:"track_code"`
	Type          string         `json:"type"`
	Views         int            `json:"views"`
	LocalViews    int            `json:"local_views"`
	CanDislike    int            `json:"can_dislike"`
}
