package dtos

import "github.com/natserract/toktik/pkg/scraper"

type GetFeedByIdResponseDto struct {
	Data *scraper.VideoInfo `json:"data"`
}
