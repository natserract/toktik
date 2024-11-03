package jobs

import (
	"time"

	"github.com/natserract/toktik/shared/store"
)

type Job struct {
	name string
	cron time.Duration
}

func NewJobs(s *store.Store) {
	embedding := Job{
		name: "generateEmbedding",
		cron: 10 * time.Second, // 10s
	}
	EmbeddingJob(embedding.cron, s)
}
