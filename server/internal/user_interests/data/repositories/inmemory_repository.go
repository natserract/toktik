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
	Tags        string
	TagsVector  []float32
	Title       string
	TitleVector []float32
}

func (r *UserInterestsRepository) SaveUserInterests(key string, data *[]SaveModel) error {
	if err := r.Store.UserInterests.Save(key, data); err != nil {
		return err
	}

	return nil
}

func (r *UserInterestsRepository) GetUserInterests(key string) (*[]SaveModel, error) {
	var results *[]SaveModel

	err := r.Store.UserInterests.Get(key, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *UserInterestsRepository) DB() *store.UserInterestsStore {
	return r.Store.UserInterests
}
