package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/ratelimit"
)

type Scraper struct {
	APIHost    string
	APIKey     string
	RateLimit  ratelimit.Limiter
	HttpClient *http.Client
}

func NewScraper(apiKey string, apiHost string) *Scraper {
	return &Scraper{
		APIHost:   apiHost,
		APIKey:    apiKey,
		RateLimit: ratelimit.New(2, ratelimit.Per(time.Second)),
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type VideoInfo struct {
	ID                  string       `json:"id,omitempty"`
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

type MusicInfo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Play     string `json:"play"`
	Cover    string `json:"cover"`
	Author   string `json:"author"`
	Original bool   `json:"original"`
	Duration int    `json:"duration,omitempty"`
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

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	// ProcessedTime int    `json:"processed_time"`
	Data T `json:"data,omitempty"`
}

type TrendingsParams struct {
	Count  string
	Region string
}

func (t *Scraper) Trendings(params TrendingsParams) (*Response[[]VideoInfo], error) {
	url := fmt.Sprintf("https://%s/feed/list", t.APIHost)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("region", params.Region)
	q.Add("count", params.Count)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-RapidAPI-Key", t.APIKey)
	req.Header.Add("X-RapidAPI-Host", t.APIHost)

	t.RateLimit.Take()
	res, err := t.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response Response[[]VideoInfo]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if response.Code == -1 {
		return nil, fmt.Errorf("server returned error: %s", response.Message)
	}

	return &response, nil
}

type SearchVideosData struct {
	Videos []VideoInfo `json:"videos"`
}

type SearchVideosParams struct {
	Keywords string
	Count    string
	Region   string
}

func (t *Scraper) SearchVideos(params SearchVideosParams) (*Response[SearchVideosData], error) {
	url := fmt.Sprintf("https://%s/feed/search", t.APIHost)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("keywords", params.Keywords)
	q.Add("region", params.Region)
	q.Add("count", params.Count)
	q.Add("publish_time", "7") // latest content no more than one week old
	q.Add("sort_type", "0")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-RapidAPI-Key", t.APIKey)
	req.Header.Add("X-RapidAPI-Host", t.APIHost)

	t.RateLimit.Take()
	res, err := t.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response Response[SearchVideosData]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if response.Code == -1 {
		return nil, fmt.Errorf("server returned error: %s", response.Message)
	}

	return &response, nil
}

type GetVideoData struct {
	*VideoInfo
}

func (t *Scraper) GetVideo(videoId string) (*Response[GetVideoData], error) {
	url := fmt.Sprintf("https://%s", t.APIHost)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	// https://www.tiktok.com/@tiktok/video/${videoId}
	q.Add("url", fmt.Sprintf("https://www.tiktok.com/@tiktok/video/%s", videoId))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-RapidAPI-Key", t.APIKey)
	req.Header.Add("X-RapidAPI-Host", t.APIHost)

	t.RateLimit.Take()
	res, err := t.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response Response[GetVideoData]
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if response.Code == -1 {
		return nil, fmt.Errorf("server returned error: %s", response.Message)
	}

	return &response, nil
}
