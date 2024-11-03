package store

type Store struct {
	Feeds                  *FeedsStore
	UserInterests          *UserInterestsStore
	UserInterestsEmbedding *UserInterestsEmbeddingStore
}

func NewStore() (*Store, error) {
	feedsStore, err := NewFeedsStore()
	if err != nil {
		return nil, err
	}

	userInterestsStore, err := NewUserInterestsStore()
	if err != nil {
		return nil, err
	}

	userInterestsEmbeddingStore, err := NewUserInterestsEmbeddingStore()
	if err != nil {
		return nil, err
	}

	return &Store{
		Feeds:                  feedsStore,
		UserInterests:          userInterestsStore,
		UserInterestsEmbedding: userInterestsEmbeddingStore,
	}, nil
}
