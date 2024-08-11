package handlers

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/harrykp/devutils/utils"
)

func HandleJwtDecoder(p *widgets.Paragraph) {
	ui.Render(p)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if e.ID == "<Enter>" {
			header, payload, err := utils.DecodeJWT(p.Text)
			if err != nil {
				p.Title = err.Error()
				p.BorderStyle.Fg = ui.ColorRed
			} else {
				text := "Header: " + header + "\nPayload: " + payload
				p.Text = text
				p.BorderStyle.Fg = ui.ColorGreen
				utils.CopyToClipboard(text)
				p.Title = " Copied to clipboard!"
			}
		} else if isExitKey(e) {
			return
		} else {
			handleKeyBindings(p, e)
		}
		ui.Render(p)
	}
}

func HandleBase64Encoding(p *widgets.Paragraph) {
	ui.Render(p)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if e.ID == "<Enter>" {
			encoded := utils.EncodeBase64(p.Text)
			p.Text = encoded
			p.BorderStyle.Fg = ui.ColorGreen
			utils.CopyToClipboard(encoded)
			p.Title = " Copied to clipboard!"
		} else if isExitKey(e) {
			return
		} else {
			handleKeyBindings(p, e)
		}
		ui.Render(p)
	}
}

func HandleBase64Decoding(p *widgets.Paragraph) {
	ui.Render(p)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if e.ID == "<Enter>" {
			decoded, err := utils.DecodeBase64(p.Text)
			if err == nil {
				p.BorderStyle.Fg = ui.ColorGreen
				utils.CopyToClipboard(decoded)
				p.Text = decoded
				p.Title = " Copied to clipboard!"
			} else {
				p.Title = " Invalid base64 string"
				p.BorderStyle.Fg = ui.ColorRed
			}
		} else if isExitKey(e) {
			return
		} else {
			handleKeyBindings(p, e)
		}
		ui.Render(p)
	}
}

func HandleJsonFormatter(p *widgets.Paragraph) {
	ui.Render(p)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if e.ID == "<Enter>" {
			formatted_text, err := utils.FormatJSON(p.Text)
			if err != nil {
				p.BorderStyle.Fg = ui.ColorRed
				p.Title = err.Error()
			} else {
				p.Text = formatted_text
				p.BorderStyle.Fg = ui.ColorGreen
				utils.CopyToClipboard(formatted_text)
				p.Title = " Copied to clipboard!"
			}
		} else if isExitKey(e) {
			return
		} else {
			handleKeyBindings(p, e)
		}
		ui.Render(p)
	}
}

func handleKeyBindings(p *widgets.Paragraph, e ui.Event) {
	switch e.ID {
	case "<C-v>":
		utils.PasteFromClipboard(p)
	case "<C-u>":
		resetParagraph(p)
	case "<Backspace>":
		deleteLastCharacter(p)
	case "<Space>":
		p.Text += " "
	default:
		if e.Type == ui.KeyboardEvent {
			p.Text += e.ID
		}
	}
}

func renderDefaultState(p *widgets.Paragraph, title, text string) {
	p.Text = text
	p.Title = title
	p.BorderStyle.Fg = ui.ColorWhite
	ui.Render(p)
}

// resetParagraph resets the paragraph widget to its default state.
func resetParagraph(p *widgets.Paragraph) {
	p.Text = ""
	p.BorderStyle.Fg = ui.ColorWhite
}

// deleteLastCharacter removes the last character from the paragraph text.
func deleteLastCharacter(p *widgets.Paragraph) {
	if len(p.Text) > 0 {
		p.Text = p.Text[:len(p.Text)-1]
	}
}

func isExitKey(e ui.Event) bool {
	if e.ID == "<C-c>" || e.ID == "q" || e.ID == "<Escape>" {
		return true
	}
	return false
}
