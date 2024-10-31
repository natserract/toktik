package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/allegro/bigcache/v3"
)

type Store struct {
	feedsCache                *bigcache.BigCache // Cache by search keywords
	feedsRecommendationsCache *bigcache.BigCache // Cache by user interaction/interest

	// Private
	_userInterestsCache *bigcache.BigCache
}

func NewStore() (*Store, error) {
	config := bigcache.DefaultConfig(10 * time.Minute)

	userInterestsCache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	feedsCache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	feedsRecommendationsCache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		feedsCache:                feedsCache,
		feedsRecommendationsCache: feedsRecommendationsCache,
		_userInterestsCache:       userInterestsCache,
	}, nil
}

func (cm *Store) SetFeeds(key string, value interface{}) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return cm.feedsCache.Set(key, buffer.Bytes())
}

func (cm *Store) GetFeeds(key string, dest interface{}) error {
	data, err := cm.feedsCache.Get(key)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(dest)
}
