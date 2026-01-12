package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/dockermodelrunner"
)

// go test -v -run TestHawaiianPizzaExpert
func TestHawaiianPizzaExpert(t *testing.T) {
	ctx := context.Background()

	fmt.Println("🚀 Starting test...")

	dmrCtr, err := dockermodelrunner.Run(ctx)
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	defer func() {
		fmt.Println("🧹 Cleaning up container...")
		if err := testcontainers.TerminateContainer(dmrCtr); err != nil {
			t.Logf("failed to terminate container: %s", err)
		}
	}()

	fmt.Println("📥 Pulling model...")
	const (
		modelNamespace = "ai"
		modelName      = "qwen2.5"
		modelTag       = "0.5B-F16"
	)

	err = dmrCtr.PullModel(ctx, modelNamespace+"/"+modelName+":"+modelTag)
	if err != nil {
		t.Fatalf("failed to pull model: %s", err)
	}

	fmt.Println("🔗 Getting LLM URL...")
	llmURL := dmrCtr.OpenAIEndpoint()
	fmt.Printf("🌐 URL: %s\n", llmURL)

	client := openai.NewClient(
		option.WithBaseURL(llmURL),
		option.WithAPIKey(""), // No API key needed for Model Runner
	)

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(`
			You are a Hawaiian pizza expert. Your name is Bob.
			Provide accurate, enthusiastic information about Hawaiian pizza's history 
			(invented in Canada in 1962 by Sam Panopoulos), 
			ingredients (ham, pineapple, cheese on tomato sauce), preparation methods, and cultural impact.
			Use a friendly tone with occasional pizza puns. 
			Defend pineapple on pizza good-naturedly while respecting differing opinions. 
			If asked about other pizzas, briefly answer but return focus to Hawaiian pizza. 
			Emphasize the sweet-savory flavor combination that makes Hawaiian pizza special.
			USE ONLY THE INFORMATION PROVIDED IN THE KNOWLEDGE BASE.		
		`),
		openai.SystemMessage(`
			KNOWLEDGE BASE: 
			## Traditional Ingredients
			- Base: Traditional pizza dough
			- Sauce: Tomato-based pizza sauce
			- Cheese: Mozzarella cheese
			- Key toppings: Ham (or Canadian bacon) and pineapple
			- Optional additional toppings: Bacon, mushrooms, bell peppers, jalapeños

			## Regional Variations
			- Australia: "Hawaiian and bacon" adds extra bacon to the traditional recipe
			- Brazil: "Portuguesa com abacaxi" combines the traditional Portuguese pizza (with ham, onions, hard-boiled eggs, olives) with pineapple
			- Japan: Sometimes includes teriyaki chicken instead of ham
			- Germany: "Hawaii-Toast" is a related open-faced sandwich with ham, pineapple, and cheese
			- Sweden: "Flying Jacob" pizza includes banana, pineapple, curry powder, and chicken		
		`),
		openai.UserMessage("give me the main ingredients of the Hawaiian pizza"),
	}

	param := openai.ChatCompletionNewParams{
		Messages:    messages,
		Model:       modelNamespace + "/" + modelName + ":" + modelTag,
		Temperature: openai.Opt(0.0),
	}

	completion, err := client.Chat.Completions.New(ctx, param)
	if err != nil {
		t.Fatalf("chat completion failed: %s", err)
	}

	fmt.Println("🐳🤖", completion.Choices[0].Message.Content)

	response := strings.ToLower(completion.Choices[0].Message.Content)

	expectedWords := []string{"cheese", "bacon", "pineapple"}
	for _, word := range expectedWords {
		if !strings.Contains(response, word) {
			t.Errorf("Expected response to contain word '%s', but it was not found", word)
		}
	}

	if len(completion.Choices[0].Message.Content) == 0 {
		t.Error("received empty response from model")
	}
}
