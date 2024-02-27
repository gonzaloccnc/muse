package multiselect

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	titleStyle = lipgloss.NewStyle().
			SetString(fmt.Sprintln()).
			Bold(true)

	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3c2a63"))

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#0000FF"))

	isBeOnChoices = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#fa0cc3"))
)

type model struct {
	label      string
	items      []string
	choices    map[int]string
	paginator  paginator.Model
	cursor     int
	realCursor int
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
		case "enter", "ctrl+c":
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

		case " ":
			_, ok := m.choices[m.realCursor]

			if ok {
				delete(m.choices, m.realCursor)
			} else {
				m.choices[m.realCursor] = m.items[m.realCursor]
			}

		}
	}

	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(m.label + ": " + "\n\n")

	start, end := m.paginator.GetSliceBounds(len(m.items))

	index := m.paginator.Page * m.paginator.PerPage

	for i, item := range m.items[start:end] {
		var itemRender, checkedRender string
		cursor := itemStyle.Render(" ")
		checkedRender = itemStyle.Render("[ ]")
		itemRender = itemStyle.Render(item)

		if m.cursor == i {
			itemRender = selectedItemStyle.Render(item)
			cursor = selectedItemStyle.Render("▸")
			checkedRender = selectedItemStyle.Render("[ ]")
		}

		if _, ok := m.choices[index+i]; ok {
			itemRender = isBeOnChoices.Render(item)
			checkedRender = isBeOnChoices.Render("[✔]")
		}

		str := fmt.Sprintf("%s %s %s\n", cursor, checkedRender, itemRender)
		b.WriteString(str)
	}

	b.WriteString("\n  " + m.paginator.View())
	b.WriteString("\n  " + titleStyle.Render("⇡/⇣ ←/→ page • enter: next") + "\n")
	return b.String()
}

func Run(label string, items []string, perPage int) []string {
	p := paginator.New()
	p.Type = paginator.Dots

	p.PerPage = perPage
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("●")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("◌")
	p.SetTotalPages(len(items))

	mo := model{
		paginator: p,
		items:     items,
		choices:   make(map[int]string),
		label:     titleStyle.Render(label),
	}

	pgm := tea.NewProgram(&mo)
	m, err := pgm.Run()

	if err != nil {
		log.Fatal(err)
	}

	mChoices := m.(model).choices

	choices := make([]string, len(mChoices))

	idx := 0
	for _, v := range choices {
		choices[idx] = v
		idx++
	}

	return choices
}
