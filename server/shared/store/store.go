package store

type Store struct {
	Feeds         *FeedsStore
	UserInterests *UserInterestsStore
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

	return &Store{
		Feeds:         feedsStore,
		UserInterests: userInterestsStore,
	}, nil
}
