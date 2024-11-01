package store

type Store struct {
	// feedsCache                *bigcache.BigCache // Cache by search keywords
	// feedsRecommendationsCache *bigcache.BigCache // Cache by user interaction/interest

	// // Private
	// _userInterestsCache *bigcache.BigCache // Collect by User Id
	Feeds *FeedsStore
}

func NewStore() (*Store, error) {
	feedsStore, err := NewFeedsStore()
	if err != nil {
		return nil, err
	}

	return &Store{ Feeds: feedsStore }, nil
}
