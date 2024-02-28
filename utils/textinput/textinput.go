package textinput

import (
	"errors"
	"fmt"

	ti "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var boldStyle = lipgloss.NewStyle().Bold(true)

func Run(placeholder string, label string) (string, error) {
	ti := ti.New()
	ti.Prompt = boldStyle.Render("... ")
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Validate = func(value string) error {
		if len(value) == 0 {
			return errors.New("invalid input")
		}

		return nil
	}

	m := model{
		textInput: ti,
		label:     boldStyle.Render(label),
		err:       nil,
	}

	p := tea.NewProgram(m)
	mp, err := p.Run()

	if err != nil {
		return "", err
	}

	mf := mp.(model)

	if mf.err != nil {
		return "", mf.err
	}

	return mf.textInput.Value(), nil
}

type (
	errMsg error
)

type model struct {
	textInput ti.Model
	label     string
	err       error
}

func (m model) Init() tea.Cmd {
	return ti.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	textSize := len(m.textInput.Value())

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.err = errors.New("user aborted prompt")
			return m, tea.Quit

		case tea.KeyEnter:
			if textSize != 0 {
				return m, tea.Quit
			}

		case tea.KeyTab:
			if textSize == 0 {
				m.textInput.SetValue(m.textInput.Placeholder)
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	iconIsValid := "❌"

	if len(m.textInput.Value()) != 0 {
		iconIsValid = "✔"
	}

	return fmt.Sprintf(
		"\n%s: %s %s\n",
		m.label,
		m.textInput.View(),
		iconIsValid,
	)
}
