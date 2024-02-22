package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var safetySettings = []*genai.SafetySetting{
	{
		Category:  genai.HarmCategoryHarassment,
		Threshold: genai.HarmBlockNone,
	},
	{
		Category:  genai.HarmCategoryHateSpeech,
		Threshold: genai.HarmBlockNone,
	},
	{
		Category:  genai.HarmCategorySexuallyExplicit,
		Threshold: genai.HarmBlockNone,
	},
	{
		Category:  genai.HarmCategoryDangerousContent,
		Threshold: genai.HarmBlockNone,
	},
}

func newGemini() func() *genai.GenerativeModel {
	var (
		once  sync.Once
		model *genai.GenerativeModel
	)
	return func() *genai.GenerativeModel {
		once.Do(func() {
			client, err := genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
			if err != nil {
				log.Fatal(err)
			}
			model = client.GenerativeModel("models/gemini-pro")
			model.SafetySettings = safetySettings
			if err != nil {
				log.Panicf("ChatCompletion error: %v\n", err)
			}
		})
		return model
	}
}

var gemini = newGemini()
