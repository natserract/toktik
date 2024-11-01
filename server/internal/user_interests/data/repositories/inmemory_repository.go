package repositories

import "github.com/natserract/toktik/shared/store"

type userInterestsRepository struct {
	Store *store.Store
}

func NewUserInterestsRepository(s *store.Store) userInterestsRepository {
	return userInterestsRepository{
		Store: s,
	}
}

type SaveModel struct {
	ID       string
	Metadata interface{}
	Vector   float64
}

func (r *userInterestsRepository) Save(id string, metadata interface{}, vector float64) error {
	if err := r.Store.UserInterests.SetUserInterests(id, &SaveModel{
		ID:       id,
		Metadata: metadata,
		Vector:   vector,
	}); err != nil {
		return err
	}

	return nil
}
