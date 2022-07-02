package state

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	AppState  State
	ListState *ListState
}

func InitialModel() *Model {
	state := createListState()
	return &Model{AppState: state, ListState: state}
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
