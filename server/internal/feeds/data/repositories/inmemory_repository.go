package repositories

import (
	"github.com/natserract/toktik/pkg/scraper"
	"github.com/natserract/toktik/shared/store"
)

type FeedsRepository struct {
	Store *store.Store
}

func NewFeedsRepository(s *store.Store) FeedsRepository {
	return FeedsRepository{
		Store: s,
	}
}

func (r *FeedsRepository) SaveFeeds(key string, videos *[]scraper.VideoInfo) error {
	if err := r.Store.Feeds.SetFeeds(key, videos); err != nil {
		return err
	}

	return nil
}

func (r *FeedsRepository) GetAllFeeds(key string) (*[]scraper.VideoInfo, error) {
	var results *[]scraper.VideoInfo

	err := r.Store.Feeds.GetFeeds(key, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *FeedsRepository) Instance() *store.FeedsStore {
	return r.Store.Feeds
}
