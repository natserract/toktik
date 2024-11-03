package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"strings"
	"time"

	"github.com/allegro/bigcache/v3"
)

type UserInterestsEmbeddingStore struct {
	Cache *bigcache.BigCache
}

func NewUserInterestsEmbeddingStore() (*UserInterestsEmbeddingStore, error) {
	config := bigcache.Config{
		Shards:     1024,                       // Number of cache shards
		LifeWindow: 100 * 365 * 24 * time.Hour, // 100 years
		// CleanWindow:        10 * time.Minute, // Interval for cleaning up expired entries
		// MaxEntriesInWindow: 1000 * 10 * 60,   // Max entries in life window
		// MaxEntrySize:       500,              // Max size of entry in bytes
		Verbose: true,
		// HardMaxCacheSize:   8192,             // Max size of cache in MB
		// OnRemove:           nil,              // Callback fired when an entry is removed
		// OnRemoveWithReason: nil,              // Callback fired when an entry is removed with a reason
	}

	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &UserInterestsEmbeddingStore{Cache: cache}, nil
}

func (c *UserInterestsEmbeddingStore) Save(key string, value interface{}) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return c.Cache.Set(key, buffer.Bytes())
}

func (c *UserInterestsEmbeddingStore) Add(key string, value interface{}) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return c.Cache.Append(key, buffer.Bytes())
}

func (c *UserInterestsEmbeddingStore) Get(key string, dest interface{}) error {
	data, err := c.Cache.Get(key)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(dest)
}

type UserInterestsEmbeddingStoreActor string

const (
	// Store by search (keywords)
	SearchUserInterestsEmbeddingActor UserInterestsEmbeddingStoreActor = "search_user_interests_embedding"
	// Store by watched content (video metadata)
	WatchUserInterestsEmbeddingActor UserInterestsEmbeddingStoreActor = "watch_user_interests_embedding"
)

func (c *UserInterestsEmbeddingStore) Key(actor UserInterestsEmbeddingStoreActor, params ...string) string {
	cacheKey := string(actor)

	if len(params) > 0 {
		cacheKey += "_" + strings.Join(params, "_")
	}

	return cacheKey
}
