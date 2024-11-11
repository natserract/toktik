package jobs

import (
	"context"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	userInterestsRepo "github.com/natserract/toktik/internal/user_interests/data/repositories"
	userInterestsEmbeddingRepo "github.com/natserract/toktik/internal/user_interests_embedding/data/repositories"
	createUserInterestEmbeddingV1 "github.com/natserract/toktik/internal/user_interests_embedding/features/create_user_interest_embedding/v1"
	"github.com/natserract/toktik/shared/store"
)

func EmbeddingJob(duration time.Duration, r *store.Store) {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
		return
	}

	j, err := s.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(
			task,
			r,
		),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Jobs start: ", j.ID())
	s.Start()
}

func task(s *store.Store) {
	cacheLen := s.UserInterests.Cache.Len()
	if cacheLen < 4 {
		log.Println("Embedding jobs running... ")

		userInterestsRepository := userInterestsRepo.NewUserInterestsRepository(s)
		userInterestsEmbeddingRepository := userInterestsEmbeddingRepo.NewUserInterestsEmbeddingRepository(s)
		userInterestsEmbeddingHandler := createUserInterestEmbeddingV1.NewCreateUserInterestEmbeddingHandler(userInterestsEmbeddingRepository)

		iterator := s.UserInterests.Cache.Iterator()
		for iterator.SetNext() {
			current, err := iterator.Value()
			if err != nil {
				log.Println(err)
			}

			userInterests, err := userInterestsRepository.GetUserInterests(current.Key())
			if err != nil {
				log.Println(err)
			}

			err = userInterestsEmbeddingHandler.Handle(context.Background(), createUserInterestEmbeddingV1.CreateUserInterestEmbedding{
				Actor:        current.Key(),
				PageContents: userInterests,
			})
			if err != nil {
				log.Println(err)
			}
		}

		log.Println("Embedding jobs finished...")
	} else {
		log.Println("Limit embedding/vectorization. Max: 3", cacheLen)
	}
}
