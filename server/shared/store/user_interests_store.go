package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"strings"
	"time"

	"github.com/allegro/bigcache/v3"
)

type UserInterestsStore struct {
	Cache *bigcache.BigCache
}

func NewUserInterestsStore() (*UserInterestsStore, error) {
	config := bigcache.Config{
		LifeWindow: 100 * 365 * 24 * time.Hour, // 100 years
	}

	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &UserInterestsStore{Cache: cache}, nil
}

func (c *UserInterestsStore) SetUserInterests(key string, value interface{}) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return c.Cache.Set(key, buffer.Bytes())
}

func (c *UserInterestsStore) GetUserInterests(key string, dest interface{}) error {
	data, err := c.Cache.Get(key)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(dest)
}

type UserInterestsStoreActor string

const (
	SearchUserInterestsActor UserInterestsStoreActor = "search_user_interests"
	WatchUserInterestsActor  UserInterestsStoreActor = "watch_user_interests"
)

func (c *UserInterestsStore) Key(actor UserInterestsStoreActor, params ...string) string {
	cacheKey := string(actor)

	if len(params) > 0 {
		cacheKey += "_" + strings.Join(params, "_")
	}

	return cacheKey
}
