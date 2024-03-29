package singleselect

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true)

	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3c2a63"))

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#0000FF"))
)

type model struct {
	label      string
	items      []string
	choice     string
	paginator  paginator.Model
	cursor     int
	realCursor int
	err        error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	start, end := m.paginator.GetSliceBounds(len(m.items))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.err = errors.New("user aborted prompt")
			return m, tea.Quit

		case "up":
			if m.realCursor > 0 {
				m.cursor--
				m.realCursor--

				if m.realCursor == start-1 {
					m.paginator.PrevPage()
					m.cursor = m.paginator.PerPage - 1
				}

			}

		case "down":
			if m.realCursor < len(m.items)-1 {
				m.realCursor++
				m.cursor++

				if m.realCursor == end && !m.paginator.OnLastPage() {
					m.paginator.NextPage()
					m.cursor = 0
				}
			}

		case "left", "h":
			if m.paginator.Page != 0 {
				m.realCursor = (m.paginator.Page - 1) * m.paginator.PerPage
				m.cursor = 0
			}

		case "right", "l":
			if !m.paginator.OnLastPage() {
				m.realCursor = (m.paginator.Page + 1) * m.paginator.PerPage
				m.cursor = 0
			}

		case "enter", " ":
			m.choice = m.items[m.realCursor]
			return m, tea.Quit

		}

	}

	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\n" + m.label + ": " + "\n\n")

	start, end := m.paginator.GetSliceBounds(len(m.items))

	for i, item := range m.items[start:end] {
		var itemRender string

		cursor := itemStyle.Render(" ")
		itemRender = itemStyle.Render(item)

		if m.cursor == i {
			itemRender = selectedItemStyle.Render(item)
			cursor = selectedItemStyle.Render("▸")
		}

		str := fmt.Sprintf("%s %s \n", cursor, itemRender)
		b.WriteString(str)
	}

	if m.paginator.TotalPages != 1 {
		b.WriteString("\n  ")
	}

	b.WriteString(m.paginator.View())

	if m.paginator.TotalPages != 1 {
		b.WriteString("\n\n  " + titleStyle.Render("⇡/⇣ ←/→ page • enter: next") + "\n")
	}

	return b.String()
}

func Run(label string, items []string, perPage int) (string, error) {
	p := paginator.New()

	p.Type = paginator.Dots
	p.PerPage = perPage
	p.SetTotalPages(len(items))

	if p.TotalPages != 1 {
		p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("●")
		p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("◌")
	} else {
		p.ActiveDot = ""
		p.InactiveDot = ""
	}

	mo := model{
		paginator: p,
		items:     items,
		label:     titleStyle.Render(label),
	}

	pgm := tea.NewProgram(&mo)
	m, err := pgm.Run()

	if err != nil {
		return "", err
	}

	mf := m.(model)

	if mf.err != nil {
		return "", mf.err
	}

	return mf.choice, nil
}
