package main

import (
	"fmt"
	"github.com/swarnimcodes/go-grAIder/internal/context"
	"os"
)

func openaiConfig() {
	var apiKey string = os.Getenv("OPEN_AI_API_KEY_MS")
	if apiKey == "" {
		fmt.Println("OpenAI API Key not found!")
	} else {
		fmt.Println("API Key: ", apiKey)
	}
}

func main() {
	fmt.Println(context.Hw)
	openaiConfig()
}
