package state

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type EditState struct {
	content textinput.Model
}

func (e *EditState) Update(msg tea.Msg, m *Model) tea.Cmd {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return tea.Quit

		case "ctrl+l":
			m.AppState = m.ListState

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	e.content, cmd = e.content.Update(msg)
	m.ListState.ToDoList[m.ListState.Cursor] = e.content.Value()
	return cmd
}

func (e *EditState) View() string {
	return fmt.Sprint("testing\n", e.content.View())
}
func createEditState(content textinput.Model) *EditState {
	return &EditState{content: content}
}
