package state

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	CurrentState int
	AppState     State
}

func InitialModel() *Model {
	state := &ListState{
		Choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		Selected: make(map[int]struct{}),
		Cursor:   0,
	}
	return &Model{AppState: state, CurrentState: 0}
}
func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd = nil
	cmd = m.AppState.Update(msg, m)
	return m, cmd
}

func (m Model) View() string {
	return m.AppState.View()
}
