package state

import tea "github.com/charmbracelet/bubbletea"

// type State interface {
// 	Update(msg tea.Msg, m *Model) tea.Cmd
// 	View() string
// }

type EditState struct{}

func (e *EditState) Update(msg tea.Msg, m *Model) tea.Cmd {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return tea.Quit

		case "t":
			m.AppState = &ListState{}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return nil
}

func (e *EditState) View() string {
	return ""
}
