package tui

import (
	"fmt"
	"gameOfLife/app"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	noStyle      = lipgloss.NewStyle()
)

type Cursor int

const (
	Up Cursor = iota
	Down
)

type model struct {
	choices    []choice
	altscreen  bool
	cursor     int
	cursorMode textinput.CursorMode
	quit       chan bool
}

type choice struct {
	name  string
	input textinput.Model
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func InitialModel() model {

	input1 := textinput.New()
	input1.Placeholder = "Rows number"
	input1.Prompt = ""
	input1.CharLimit = 3
	input1.Width = 5
	input1.SetValue("50")
	input1.Focus()
	option1 := choice{name: "Height", input: input1}

	input2 := textinput.New()
	input2.Placeholder = "Columns number"
	input2.Prompt = ""
	input2.CharLimit = 3
	input2.Width = 5
	input2.SetValue("50")
	option2 := choice{name: "Width", input: input2}

	input3 := textinput.New()
	input3.Placeholder = "Number of alive cells"
	input3.Prompt = ""
	input3.CharLimit = 5
	input3.Width = 5
	input3.SetValue("2000")
	option3 := choice{name: "Population", input: input3}

	m := model{
		cursorMode: textinput.CursorBlink,
		choices:    []choice{option1, option2, option3},
		quit:       make(chan bool),
	}

	return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {

				var movement Cursor = Up
				focusOnSelectedChoice(&m, movement)

			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {

				var movement Cursor = Down
				focusOnSelectedChoice(&m, movement)

			}
		case "enter":

			var cmds []tea.Cmd

			if !m.altscreen {
				width, err := strconv.Atoi(m.choices[0].input.Value())
				height, err := strconv.Atoi(m.choices[1].input.Value())
				population, err := strconv.Atoi(m.choices[2].input.Value())

				if err != nil {
					fmt.Println("WRONG INPUT TYPE: ", err)
					return m, tea.Quit
				}

				cmds = append(cmds, tea.EnterAltScreen)

				// Start the game of life
				go app.Game{Width: width, Height: height, Population: population, Quit: m.quit}.Run()

			} else {
				// Exit game
				close(m.quit)
				cmds = append(cmds, tea.ExitAltScreen)

				// Reinitialize channel
				m.quit = make(chan bool)
			}

			m.altscreen = !m.altscreen

			return m, tea.Batch(cmds...)
		}

	}
	cmd := m.updateInputs(msg)
	return m, cmd
}

func focusOnSelectedChoice(m *model, c Cursor) {
	m.choices[m.cursor].input.Blur()
	m.choices[m.cursor].input.PromptStyle = noStyle
	m.choices[m.cursor].input.TextStyle = noStyle

	switch c {
	case Down:
		m.cursor++

	case Up:
		m.cursor--
	}

	m.choices[m.cursor].input.Focus()
	m.choices[m.cursor].input.PromptStyle = focusedStyle
	m.choices[m.cursor].input.TextStyle = focusedStyle

}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.choices))

	for i := range m.choices {
		m.choices[i].input, cmds[i] = m.choices[i].input.Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	outputString := ""
	if !m.altscreen {
		outputString = "Choose initial parameters\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = "⏺"
			}

			sep := " "
			constGap := (10 - len(choice.name))
			for i := 0; i < constGap; i++ {
				sep += " "
			}
			outputString += fmt.Sprintf("%s %s %s= %v \n", cursor, choice.name, sep, choice.input.View())
		}

		outputString += "\nPress: \n <q> to quit \n <enter> to submit parameters and start game\n"
	}

	return outputString
}
