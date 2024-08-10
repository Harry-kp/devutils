package main

import (
	"encoding/base64"
	"log"

	"github.com/atotto/clipboard"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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
	l.Title = "List"
	l.Rows = []string{
		"[0] Base64 Encoder",
		"[1] Base64 Decoder",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 25, 8)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
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
		case "<Enter>":
			switch l.SelectedRow {
			case 0:
				handleBase64("encode")
			case 1:
				handleBase64("decode")
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

// handleBase64 processes both encoding and decoding based on the provided mode.
func handleBase64(mode string) {
	p := widgets.NewParagraph()
	p.WrapText = true
	p.SetRect(26, 0, 100, 8)
	p.Title = "Enter String"

	if mode == "decode" {
		p.Title = "Enter Base64 String"
	}
	ui.Render(p)

	uiEvents := ui.PollEvents()
	hasError := false

	for {
		e := <-uiEvents

		switch e.ID {
		case "<Enter>":
			if hasError {
				resetParagraph(p)
				hasError = false
				continue
			}

			if mode == "encode" {
				handleEncoding(p)
			} else if mode == "decode" {
				handleDecoding(p, &hasError)
			}
			ui.Render(p)

		case "<C-c>":
			return // Exit the function

		case "<C-v>":
			pasteFromClipboard(p)

		case "<Backspace>":
			deleteLastCharacter(p)

		default:
			if e.Type == ui.KeyboardEvent {
				p.Text += e.ID
				ui.Render(p)
			}
		}
	}
}

// handleEncoding encodes the text to Base64 and copies it to the clipboard.
func handleEncoding(p *widgets.Paragraph) {
	encoded := base64.StdEncoding.EncodeToString([]byte(p.Text))
	p.Text = encoded
	p.BorderStyle.Fg = ui.ColorGreen
	clipboard.WriteAll(encoded)
	p.Title = "Copied to clipboard!"
}

// handleDecoding decodes the Base64 string and handles errors.
func handleDecoding(p *widgets.Paragraph, hasError *bool) {
	decoded, err := base64.StdEncoding.DecodeString(p.Text)
	if err != nil {
		*hasError = true
		p.Title = "Invalid Base64 string"
		p.BorderStyle.Fg = ui.ColorRed
	} else {
		p.Text = string(decoded)
		p.BorderStyle.Fg = ui.ColorGreen
		clipboard.WriteAll(p.Text)
		p.Title = "Copied to clipboard!"
	}
}

// resetParagraph resets the paragraph widget to its default state.
func resetParagraph(p *widgets.Paragraph) {
	p.Text = ""
	p.BorderStyle.Fg = ui.ColorWhite
	ui.Render(p)
}

// pasteFromClipboard pastes the clipboard content into the paragraph widget.
func pasteFromClipboard(p *widgets.Paragraph) {
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Println("Failed to read from clipboard:", err)
		return
	}
	p.Text = text
	ui.Render(p)
}

// deleteLastCharacter removes the last character from the paragraph text.
func deleteLastCharacter(p *widgets.Paragraph) {
	if len(p.Text) > 0 {
		p.Text = p.Text[:len(p.Text)-1]
		ui.Render(p)
	}
}
