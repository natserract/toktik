package repositories

import "github.com/natserract/toktik/shared/store"

type UserInterestsRepository struct {
	Store *store.Store
}

func NewUserInterestsRepository(s *store.Store) UserInterestsRepository {
	return UserInterestsRepository{
		Store: s,
	}
}

func (r *UserInterestsRepository) SaveUserInterests(key string, pageContents []string) error {
	if err := r.Store.UserInterests.Save(key, pageContents); err != nil {
		return err
	}

	return nil
}

func (r *UserInterestsRepository) GetUserInterests(key string) ([]string, error) {
	var results []string

	err := r.Store.UserInterests.Get(key, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *UserInterestsRepository) DB() *store.UserInterestsStore {
	return r.Store.UserInterests
}
