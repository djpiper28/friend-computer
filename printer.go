package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	CUT_PAPER    = "\x1bi"
	SCRATCH_FILE = "print_scratch_file.txt"
	EOF_PADDING  = "\n\n\n\n\n\n"
	NEW_LINES    = "\n\n"
)

type Printer struct {
	// i.e: /dev/usb/lp0
	PrinterOutput string
	OutputFile    os.File
}

func PaperHeader(title string) string {
	const HEADER_MARKER = "***"
	return fmt.Sprintf("%s %s %s", HEADER_MARKER, title, HEADER_MARKER)
}

func PaperSectionHeader(title string) string {
	const SECTION_MARKER = "====="
	return fmt.Sprintf("%s %s %s", SECTION_MARKER, title, SECTION_MARKER)
}

func GenerateTicket(title, description, sectionTitle, sectionBody string) string {
	const begin = "BEGIN"
	const end = "END"
	text := PaperHeader(title)
	text += NEW_LINES
	text += time.Now().Format(time.Stamp)
	text += NEW_LINES
	text += description
	text += NEW_LINES
	text += PaperSectionHeader(begin + " " + sectionTitle)
	text += NEW_LINES
	text += sectionBody
	text += NEW_LINES
	text += PaperSectionHeader(end + " " + sectionTitle)
	text += NEW_LINES
	text += "Messages sent by Friend Computer are TOP SECRET. Friend Computer requests you inform on developments."
	text += EOF_PADDING
	return text
}

func (p *Printer) PrintMission(mission Mission) error {
	return p.Print(GenerateTicket("MISSION BRIEF",
		"Friend computer requires agents to execute a mission. Best of luck.",
		"BRIEF",
		"MISSION BREIF:"+mission.MissionBreif+NEW_LINES+
			"LOCATION: "+mission.Location+NEW_LINES+
			"DIRECTIVES:"+NEW_LINES+mission.DirectivesString()+
			"INVENTORIES: "+NEW_LINES+mission.InventoryString()))
}

func (p *Printer) Print(body string) error {
	text := body + CUT_PAPER
	err := os.WriteFile(SCRATCH_FILE, []byte(text), 0777)
	if err != nil {
		return err
	}

	command := fmt.Sprintf("cat %s > \"%s\"", SCRATCH_FILE, p.PrinterOutput)
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()

	if err != nil {
		err = errors.New(fmt.Sprintf("%s: %s", err, output))
		log.Print("Cannot print: %s", err)
		return err
	}
	return nil
}
