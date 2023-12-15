package main

import "fmt"

type Mission struct {
	MissionBreif string   `json:"missionBrief"`
	Location     string   `json:"location"`
	Directives   []string `json:"directives"`
}

func (m *Mission) DirectivesString() string {
	ret := ""
	for i, directive := range m.Directives {
		ret += fmt.Sprintf(" - %02x: %s", i+1, directive)
		if i != len(ret)-1 {
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
