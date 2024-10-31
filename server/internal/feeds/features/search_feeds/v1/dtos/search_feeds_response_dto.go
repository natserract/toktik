package dtos

import "github.com/natserract/toktik/pkg/scraper"

type SearchFeedsResponseDTO struct {
	Data []scraper.VideoInfo `json:"data"`
}
