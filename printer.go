package main

import "os"

const (
  CUT_PAPER = "\ei"
)

type Printer struct {
  // i.e: /dev/usb/lp0
  PrinterOutput string
  OutputFile os.File
}

func (p *Printer) Print(body string) error {

}

