package repositories

import (
	"github.com/natserract/toktik/shared/store"
)

type UserInterestsEmbeddingRepository struct {
	Store *store.Store
}

func NewUserInterestsEmbeddingRepository(s *store.Store) UserInterestsEmbeddingRepository {
	return UserInterestsEmbeddingRepository{
		Store: s,
	}
}

type SaveUserInterestsEmbeddingModel struct {
	ID          string
	Tag         string
	TagsVector  []float32
	Title       string
	TitleVector []float32
}

func (r *UserInterestsEmbeddingRepository) SaveUserInterestsEmbedding(key string, data *[]SaveUserInterestsEmbeddingModel) error {
	if err := r.Store.UserInterestsEmbedding.Save(key, data); err != nil {
		return err
	}

	return nil
}

func (r *UserInterestsEmbeddingRepository) GetUserInterestsEmbedding(key string) (*[]SaveUserInterestsEmbeddingModel, error) {
	var results *[]SaveUserInterestsEmbeddingModel

	err := r.Store.UserInterestsEmbedding.Get(key, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *UserInterestsEmbeddingRepository) GetAllUserInterestsEmbedding() ([]SaveUserInterestsEmbeddingModel, error) {
	var results []SaveUserInterestsEmbeddingModel

	iterator := r.Store.UserInterestsEmbedding.Cache.Iterator()
	for iterator.SetNext() {
		current, err := iterator.Value()
		if err != nil {
			return nil, err
		}

		var embeddings *[]SaveUserInterestsEmbeddingModel
		err = r.Store.UserInterestsEmbedding.Get(current.Key(), &embeddings)
		if err != nil {
			return nil, err
		}

		results = append(results, *embeddings...)
	}

	return results, nil
}

func (r *UserInterestsEmbeddingRepository) DB() *store.UserInterestsEmbeddingStore {
	return r.Store.UserInterestsEmbedding
}
