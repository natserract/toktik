package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"strings"
	"time"

	"github.com/allegro/bigcache/v3"
)

type FeedsStore struct {
	Cache *bigcache.BigCache
}

func NewFeedsStore() (*FeedsStore, error) {
	config := bigcache.DefaultConfig(3 * time.Minute)

	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &FeedsStore{Cache: cache}, nil
}

func (c *FeedsStore) Save(key string, value interface{}) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return c.Cache.Set(key, buffer.Bytes())
}

func (c *FeedsStore) Get(key string, dest interface{}) error {
	data, err := c.Cache.Get(key)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(dest)
}

type FeedsStoreActor string

const (
	SearchFeedsActor FeedsStoreActor = "search_feeds"
	GetFeedByIdActor FeedsStoreActor = "get_feed_by_id"
)

func (c *FeedsStore) Key(actor FeedsStoreActor, params ...string) string {
	cacheKey := string(actor)

	if len(params) > 0 {
		cacheKey += "_" + strings.Join(params, "_")
	}

	return cacheKey
}
