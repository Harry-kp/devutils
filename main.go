package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/harrykp/devutils/handlers"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	initializeToolBar()
}

func initializeToolBar() {
	l := widgets.NewList()
	l.Title = " Tools"
	l.Rows = []string{
		"[0] Base64 Encoder",
		"[1] Base64 Decoder",
		"[2] JSON Formatter",
		"[3] JWT Decoder",
	}
	l.TextStyle = ui.NewStyle(ui.ColorWhite)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 25, 8)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>", "<Escape>":
			return
		case "j", "<Down>":
			if l.SelectedRow < len(l.Rows)-1 {
				l.ScrollDown()
			} else {
				l.ScrollTop()
			}
		case "k", "<Up>":
			if l.SelectedRow > 0 {
				l.ScrollUp()
			} else {
				l.ScrollBottom()
			}
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "<Tab>", "<Enter>":
			p := widgets.NewParagraph()
			p.WrapText = true
			p.SetRect(26, 0, 100, 8)
			p.BorderStyle.Fg = ui.ColorYellow
			switch l.SelectedRow {
			case 0:
				handlers.HandleBase64Encoding(uiEvents)
			case 1:
				handlers.HandleBase64Decoding(uiEvents)
			case 2:
				handlers.HandleJsonFormatter(uiEvents)
			case 3:
				handlers.HandleJwtDecoder(uiEvents)
			}
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		ui.Render(l)
	}
}
