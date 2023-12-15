package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client
var chatBot openai.ChatCompletionRequest

func (c *Config) InitAi() {
	client = openai.NewClient(c.OpenAiKey)
}

const (
	SYSTEM_CONTROLLER = "(System controller):"
	USER_CONTROLLER   = "(User):"
)

func (c *Config) missionPrompt() string {
	mission := Mission{MissionBreif: "example brief",
		Location:   "example location",
		Directives: []string{"example 1", "example 2", "example 3"},
		Inventory: map[string][]InventoryItem{
			"example player": []InventoryItem{{
				Quantity:                3,
				ItemName:                "Loaf of bread",
				ItemDescriptionAndStats: "Edible and tasty",
				ItemWeight:              1,
			}}},
	}
	text, err := json.Marshal(mission)
	if err != nil {
		log.Fatalf("Cannot create prompt for mission: %s", err)
	}

	return fmt.Sprintf(`%s Setup the initial mission brief and three orders that conflict mildly.
Game context: Distopian sci-fi world, with corporate espionage very common and very little trust, 1984. As a computer you can trust no-one except your agents. One agent is a traitor.
Your players are: %s.
Allocate each player an inventory.
You are the GM for the session and should handle combat, inventory, game events and story progression. You will provide extra orders occassionally.
Output in JSON format such as this: %s`,
		SYSTEM_CONTROLLER,
		strings.Join(c.Players, ", "),
		text)
}

func (c *Config) userMessagePrompt(query string) string {
	exampleOutput := AiResponse{ResponseText: "example response"}
	text, err := json.Marshal(exampleOutput)
	if err != nil {
		log.Fatalf("Cannot create example Output: %s", err)
	}

	return fmt.Sprintf(`%s a player has sent you a message.
As the GM you will respond with story progression and describe any combats, or consequences from any actions.
Output in JSON format such as this: %s
%s Message: %s`, SYSTEM_CONTROLLER, text, USER_CONTROLLER, query)
}

func (c *Config) GenerateMission() {
	chatBot = openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: c.missionPrompt(),
			},
		},
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		chatBot)

	if err != nil {
		log.Printf("ChatCompletion error: %s\n", err)
		return
	}

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

func (c *Config) SendUserMessage(message string) (string, error) {
	return SendMessage(c.userMessagePrompt(message))
}

func SendMessage(message string) (string, error) {
	chatBot.Messages = append(chatBot.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	resp, err := client.CreateChatCompletion(context.Background(), chatBot)
	if err != nil {
		log.Printf("ChatCompletion error: %s\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
