package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

var centerStyle = gloss.NewStyle()
var textStyle = gloss.NewStyle().Foreground(gloss.Color("#888888")).Inline(true)
var inputStyle = gloss.NewStyle().Foreground(gloss.Color("#ffffff")).Inline(true)
var errorStyle = gloss.NewStyle().Foreground(gloss.Color("#ff0000")).Inline(true)
var cursorStyle = gloss.NewStyle().Background(gloss.Color("#ffffff")).Foreground(gloss.Color("#111111")).Blink(false).Inline(true)
var termWidth int
var termHeight int

type model struct {
	text string
	input string
	errText string
}

func initialModel() model {
	return model {
		text: "lorem this is a lot of text, like a  lot a lot if text nlike a bunch of text",
		input: "",
		errText: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(m.text) == 0 {
		return m, tea.Quit
	}
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "ctrl+q" {
			return m, tea.Quit
		} else {
			UpdateText(msg.String(), &m)
		}

	case tea.WindowSizeMsg:
		termHeight = msg.Height
		termWidth = msg.Width
		/*textStyle = textStyle.Width(termWidth).Height(termHeight).Align(gloss.Center)
		inputStyle = inputStyle.Width(termWidth).Height(termHeight).Align(gloss.Center)
		errorStyle = errorStyle.Width(termWidth).Height(termHeight).Align(gloss.Center)
		*/
		centerStyle = centerStyle.Width(termWidth).Height(termHeight).Align(gloss.Center).MaxWidth(termWidth / 2)
	}
	
	return m, nil
}

func UpdateText(input string, m *model) {
	if input[0] == m.text[0] {
		m.input += input
		m.text = m.text[1:]
	}
}

func formatText(text string) string {
	if len(text) == 0 {
		return ""
	}
	words := strings.Fields(text)
	var outputText string
	var lineNum int
	for _, word := range words {
		if len(strings.Split(outputText, "\n")[lineNum]) + len(word) > termWidth / 2 {
			outputText += "\n" + word + " "
			lineNum++
		} else {
			outputText += word + " "
		}
	}
	return outputText
}

func (m model) View() string {
	if len(m.text) >= 1 {
		return centerStyle.Render(inputStyle.Render(m.input) + cursorStyle.Render(string(m.text[0])) + textStyle.Render(m.text[1:]))
	} else {
		return centerStyle.Render(inputStyle.Render(m.input))
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}