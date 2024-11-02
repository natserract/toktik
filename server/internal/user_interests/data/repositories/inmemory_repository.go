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

type SaveModel struct {
	ID          string
	PageContent string
	Metadata    interface{}
}

func (r *UserInterestsRepository) SaveUserInterest(id string, pageContent string, metadata interface{}) error {
	if err := r.Store.UserInterests.SetUserInterests(id, &SaveModel{
		ID:          id,
		PageContent: pageContent,
		Metadata:    metadata,
	}); err != nil {
		return err
	}

	return nil
}

func (r *UserInterestsRepository) GetUserInterests(key string) (*SaveModel, error) {
	var results *SaveModel

	err := r.Store.UserInterests.GetUserInterests(key, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *UserInterestsRepository) DB() *store.UserInterestsStore {
	return r.Store.UserInterests
}
