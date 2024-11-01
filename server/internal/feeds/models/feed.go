package models

type MusicInfo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Play     string `json:"play"`
	Cover    string `json:"cover"`
	Author   string `json:"author"`
	Original bool   `json:"original"`
	Duration int    `json:"duration"`
	Album    string `json:"album"`
}

type CommerceInfo struct {
	AdvPromotable          bool `json:"adv_promotable"`
	AuctionAdInvited       bool `json:"auction_ad_invited"`
	BrandedContentType     int  `json:"branded_content_type"`
	WithCommentFilterWords bool `json:"with_comment_filter_words"`
}

type Author struct {
	ID       string `json:"id"`
	UniqueID string `json:"unique_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type Feed struct {
	AwemeID             string       `json:"aweme_id"`
	VideoID             string       `json:"video_id"`
	Region              string       `json:"region"`
	Title               string       `json:"title"`
	Cover               string       `json:"cover"`
	AIDynamicCover      string       `json:"ai_dynamic_cover"`
	OriginCover         string       `json:"origin_cover"`
	Duration            int          `json:"duration"`
	Play                string       `json:"play"`
	WMPlay              string       `json:"wmplay"`
	Size                int64        `json:"size"`
	WMSize              int64        `json:"wm_size"`
	Music               string       `json:"music"`
	MusicInfo           MusicInfo    `json:"music_info"`
	PlayCount           int          `json:"play_count"`
	DiggCount           int          `json:"digg_count"`
	CommentCount        int          `json:"comment_count"`
	ShareCount          int          `json:"share_count"`
	DownloadCount       int          `json:"download_count"`
	CreateTime          int64        `json:"create_time"`
	Anchors             interface{}  `json:"anchors"`
	AnchorsExtras       string       `json:"anchors_extras"`
	IsAd                bool         `json:"is_ad"`
	CommerceInfo        CommerceInfo `json:"commerce_info"`
	CommercialVideoInfo string       `json:"commercial_video_info"`
	ItemCommentSettings int          `json:"item_comment_settings"`
	MentionedUsers      string       `json:"mentioned_users"`
	Author              Author       `json:"author"`
	IsTop               int          `json:"is_top"`
}
