package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func (c *Config) InitAi() {
	client = openai.NewClient(c.OpenAiKey)
}

const (
	SYSTEM_CONTROLLER = "(System controller):"
)

func missionPrompt() string {
	mission := Mission{MissionBreif: "example brief",
		Location:   "example location",
		Directives: []string{"example 1", "example 2", "example 3"}}
	text, err := json.Marshal(mission)
	if err != nil {
		log.Fatalf("Cannot create prompt for mission: %s", err)
	}

	return fmt.Sprintf("%s Setup the initial mission brief and three orders that conflict. Output in JSON format such as this: %s",
		SYSTEM_CONTROLLER,
		text)
}

func (c *Config) GenerateMission() {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: missionPrompt(),
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %s\n", err)
		return
	}

  log.Print(resp.Choices[0].Message.Content)
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &c.Mission)
	if err != nil {
		log.Fatalf("Cannot generate mission: %s", err)
	}

  if c.Mission.MissionBreif == "" {
    log.Fatal("No mission brief was returned")
  }

  if c.Mission.Location == "" {
    log.Fatal("No mission location was returned")
  }

  if len(c.Mission.Directives) == 0 {
    log.Fatal("No mission directives were returned")
  }
}
