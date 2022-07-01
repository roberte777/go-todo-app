package state

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type State interface {
	Update(msg tea.Msg, m *Model) tea.Cmd
	View() string
}
type ListState struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func (l *ListState) Update(msg tea.Msg, m *Model) tea.Cmd {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return tea.Quit

		case "t":
			m.AppState = &EditState{}

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if l.Cursor > 0 {
				l.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if l.Cursor < len(l.Choices)-1 {
				l.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := l.Selected[l.Cursor]
			if ok {
				delete(l.Selected, l.Cursor)
			} else {
				l.Selected[l.Cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return nil
}
func (l ListState) GetCursor() int {
	return l.Cursor
}

func (l *ListState) View() string {
	width, height, err := term.GetSize(0)
	if err != nil {
		return ""
	}
	innerStyle := lipgloss.NewStyle()
	s := "What should we buy at the market?\n\n"

	for i, choice := range l.Choices {
		cursor := " "
		if l.GetCursor() == i {
			cursor = ">"
		}

		checked := " "

		if _, ok := l.Selected[i]; ok {

			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	mainBlock := innerStyle.Render(s)
	marginLeft := (width / 2) - (lipgloss.Width(s) / 2)
	marginTop := (height / 2) - (lipgloss.Height(s) / 2)
	outerStyle := lipgloss.NewStyle().MarginLeft(marginLeft).MarginTop(marginTop).Render(mainBlock)

	return outerStyle

}
