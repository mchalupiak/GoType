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
var cursorStyle = gloss.NewStyle().Background(gloss.Color("#ffffff")).Inline(true).Blink(true)
var termWidth int
var termHeight int

type model struct {
	text    string
	input   string
	errText string
}

func initialModel() model {
	return model{
		text:    "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
		input:   "",
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
	var lineNum int = 0
	for _, word := range words {
		if len(strings.Split(outputText, "\n")[lineNum])+len(word) > termWidth/2 {
			outputText += "\n" + word + " "
			lineNum++
		} else {
			outputText += word + " "
		}
	}
	return outputText
}

func (m model) View() string {
	return formatText(inputStyle.Render(m.input) + cursorStyle.Render(string(m.text[0])) + textStyle.Render(m.text[1:]))
	/*
		if len(m.text) >= 1 {
			return centerStyle.Render(text)
		} else {
			return centerStyle.Render(text)
		}*/
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
