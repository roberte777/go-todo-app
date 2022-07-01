package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/roberte777/go-todo-app/pkg/state"
)

type model struct {
	CurrentState int
	ListState    state.State
}

func initialModel() model {
	state := &state.ListState{
		Choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		Selected: make(map[int]struct{}),
		Cursor:   0,
	}
	return model{ListState: state, CurrentState: 0}
}
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd = nil
	if m.CurrentState == 0 {
		cmd = m.ListState.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.CurrentState == 0 {
		return m.ListState.View()
	}
	return ""
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Ala, there has been an error: %v", err)
		os.Exit(1)
	}
}
