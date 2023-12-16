package main

import (
	"encoding/json"
	"log"
	"os"
)

const CONFIG_FILE = "config.json"

func main() {
	var config Config
	content, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Fatalf("Cannot read file: %s", err)
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	if len(config.Players) == 0 {
		log.Fatal("No players in the configuration")
	}

	if (config.Printer) == "" {
		log.Fatal("No printer file")
	}

	printer := Printer{PrinterOutput: config.Printer}

	config.InitAi()
	if config.Mission.MissionBreif == "" {
		config.GenerateMission()

		log.Printf("Generated mission Brief: %s\n Location: %s\n Directives: %s",
			config.Mission.MissionBreif,
			config.Mission.Location,
			config.Mission.Directives)

		err = printer.PrintMission(config.Mission)
		if err != nil {
			log.Fatal(err)
		}

		err = config.Save()
		if err != nil {
			log.Fatal(err)
		}
	}

	WithUi(func(u Ui) {

	})
}
