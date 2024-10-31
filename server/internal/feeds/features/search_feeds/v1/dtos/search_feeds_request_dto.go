package dtos

type SearchFeedsRequestDTO struct {
	// keywords, region, count, publish_time, sort_type
	Keywords string `query:"keywords" json:"keywords" validate:"required"`
}
