package util

import (
	"context"
	"strings"

	"github.com/natserract/toktik/embedding"
	"github.com/natserract/toktik/pkg/text_processor"
)

func GenerateUserProfile() {
	const (
		USER_PROFILE = `
		User Profile

		Name: Alex Johnson
		Age: 28
		Location: London, UK

		Interests:
		America
		Barack Obama
		Favorite Book: "Dreams from My Father" by Barack Obama

		Hobbies:
		Reading about U.S. politics
		Watching documentaries on American history
		Social Media:

		Twitter: @AlexJ_AmericaFan
		`
	)

	cp := text_processor.CleanProcessor{}
	processRule := map[string]interface{}{
		"rules": map[string]interface{}{
			"pre_processing_rules": []interface{}{
				map[string]interface{}{
					"id":      "remove_extra_spaces",
					"enabled": true,
				},
				map[string]interface{}{
					"id":      "remove_urls_emails",
					"enabled": true,
				},
			},
		},
	}
	text := cp.Clean(USER_PROFILE, processRule)
	text = strings.ToLower(text)

	embed := embedding.NewVectorEmbedding()
	vec, err := embed.CreateVector(text, context.Background())
	if err == nil {
		embed.SaveVectorsToFile(embedding.Vector{
			ID:     "user_profile",
			Values: vec,
		}, "sample/user_profile.json")
	}
}
