package embedding

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/natserract/toktik/model"
	"github.com/natserract/toktik/pkg/config"
	"github.com/sashabaranov/go-openai"
)

type VectorEmbedding struct {
	model *openai.Client // OpenAI, Gemini
}

func NewVectorEmbedding() VectorEmbedding {
	cfg := config.GetConfig()
	client := model.GetOpenAiClientFromToken(cfg.OpenAIKey)
	return VectorEmbedding{
		model: client,
	}
}

func (c *VectorEmbedding) CreateVector(text string, ctx context.Context) ([]float32, error) {
	resp, err := c.model.CreateEmbeddings(
		ctx,
		openai.EmbeddingRequest{
			Input: []string{text},
			Model: openai.EmbeddingModel("text-embedding-ada-002"),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get embeddings: %w", err)
	}

	vector := resp.Data[0].Embedding
	return vector, nil
}

func (c *VectorEmbedding) CosineSimilarity(queryVec, vectors []float32) (float32, error) {
	var dotProduct, normA, normB float32
	for i := 0; i < len(queryVec); i++ {
		dotProduct += queryVec[i] * vectors[i]
		normA += queryVec[i] * queryVec[i]
		normB += vectors[i] * vectors[i]
	}

	if normA == 0 || normB == 0 {
		return 0, fmt.Errorf("one of the vectors is zero")
	}
	return dotProduct / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB)))), nil
}

type Vector struct {
	ID     string    `json:"id"`
	Values []float32 `json:"values"`
}

func (c *VectorEmbedding) LoadQueryVector(filename string) (Vector, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Vector{}, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var query Vector
	err = json.Unmarshal(byteValue, &query)
	if err != nil {
		return Vector{}, err
	}

	return query, nil
}

func (c *VectorEmbedding) LoadDataVectors(filename string) ([]Vector, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var vectors []Vector
	err = json.Unmarshal(byteValue, &vectors)
	if err != nil {
		return nil, err
	}

	return vectors, nil
}

func (c *VectorEmbedding) SaveVectorsToFile(vec Vector, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(vec)
	if err != nil {
		return err
	}

	return nil
}
