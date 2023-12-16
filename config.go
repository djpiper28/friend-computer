package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type InventoryItem struct {
	Quantity                float64 `json:"quantity"`
	ItemName                string  `json:"itemName"`
	ItemDescriptionAndStats string  `json:"itemDescriptionAndStats"`
	ItemWeight              float64 `json:"weight"`
}

type Mission struct {
	MissionBreif string                     `json:"missionBrief"`
	Location     string                     `json:"location"`
	Directives   []string                   `json:"directives"`
	Inventory    map[string][]InventoryItem `json:"Inventory`
}

func (m *Mission) DirectivesString() string {
	ret := ""
	for i, directive := range m.Directives {
		ret += fmt.Sprintf(" - %02x: %s", i, directive)
		if i != len(ret)-1 {
			ret += ";\n\n"
		}
	}
	return ret
}

func (m *Mission) InventoryString() string {
	ret := ""
	for player, inventory := range m.Inventory {
		ret += fmt.Sprintf("+++ %s:", player)
		ret += "\n"
		for i, item := range inventory {
			ret += fmt.Sprintf(" - %02x '%s' x %.02f", i, item.ItemName, item.Quantity)
			ret += "\n"
			ret += fmt.Sprintf("   Weight: %.02f", item.ItemWeight)
			ret += "\n"
			ret += fmt.Sprintf("   Description: %s", item.ItemDescriptionAndStats)
			ret += ";\n\n"
		}
	}

	return ret
}

type Config struct {
	Players   []string `json:"players"`
	Mission   Mission  `json:"mission"`
	OpenAiKey string   `json:"openAiKey"`
	Printer   string   `json:"printer"`
}

func (c *Config) Save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(CONFIG_FILE, data, 777)
}
