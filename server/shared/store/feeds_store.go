package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/allegro/bigcache/v3"
)

type FeedsStore struct {
	Cache *bigcache.BigCache
}

func NewFeedsStore() (*FeedsStore, error) {
	config := bigcache.DefaultConfig(1 * time.Minute)

	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &FeedsStore{Cache: cache}, nil
}

func (c *FeedsStore) SetFeeds(key string, value interface{}) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return c.Cache.Set(key, buffer.Bytes())
}

func (c *FeedsStore) GetFeeds(key string, dest interface{}) error {
	data, err := c.Cache.Get(key)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(dest)
}
