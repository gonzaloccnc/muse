package checkbox

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"

	tea "github.com/charmbracelet/bubbletea"
)

func newModel() model {
	var items = []string{
		"Item 1",
		"Item 2",
		"Item 3",
		"Item 4",
		"Item 5",
		"Item 6",
		"Item 7",
		"Item 8",
		"Item 9",
		"Item 10",
		"Item 11",
		"Item 12",
		"Item 13",
		"Item 14",
		"Item 15",
		"Item 16",
		"Item 17",
		"Item 18",
		"Item 19",
	}

	p := paginator.New()
	p.Type = paginator.Dots

	p.PerPage = 4
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len(items))

	return model{
		paginator: p,
		items:     items,
		choises:   make(map[int]string),
	}
}

type model struct {
	items      []string
	choises    map[int]string
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
			v, ok := m.choises[m.realCursor]
			if ok {
				delete(m.choises, m.realCursor)
			} else {
				m.choises[m.realCursor] = v
			}

		}

	}

	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\nPaginator Example\n")

	start, end := m.paginator.GetSliceBounds(len(m.items))

	// get index = (page * items_per_page) + index_of_range
	index := (m.paginator.Page * m.paginator.PerPage)

	for i, item := range m.items[start:end] {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "

		if _, ok := m.choises[index+i]; ok {
			checked = "x"
		}

		str := fmt.Sprintf("%s [%s] %s\n", cursor, checked, item)
		b.WriteString(str)
	}

	b.WriteString("\n  " + m.paginator.View())
	b.WriteString("\n  h/l ←/→ page • enter: next\n")
	return b.String()
}

func Checkbox() {
	p := tea.NewProgram(newModel())
	m, err := p.Run()

	if err != nil {
		log.Fatal(err)
	}

	logrus.Printf("%v", m.(model).choises)
}

// goBack = (actualPage -1) * itemsperpage

// page = 1, startIndex=4 == (1-1) *4 = 0
// page = 2, startIndex=8 == (2-1) *4 = 4
// page = 3, startIndex=12 == (3-1) *4 = 8
// page = 4, startIndex=16 == (4-1) *4 = 12
// page = 5, startIndex=20 == (5-1) *4 = 16

// goTo = (actualPage +1) * itemperpage

// page = 0, startIndex=0 == (0+1) *4 = 4
// page = 1, startIndex=4 == (1+1) *4 = 8
// page = 2, startIndex=8 == (2+1) *4 = 12
// page = 3, startIndex=12 == (3+1) *4 = 16
// page = 4, startIndex=16 == (4+1) *4 = 20

// last page not must works
// page = 5, startIndex=20 == (5-1) *4 = 16
