package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	options       []string
	cursor        int
	width         int
	selectedStyle lipgloss.Style
	normalStyle   lipgloss.Style
}

func initModel() model {
	options := []string{
		"New Game",
		"Load Game",
		"Statistics",
		"Settings",
		"Exit",
	}

	return model{
		options:       options,
		cursor:        0,
		selectedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#f53c32")).Bold(true),
		normalStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("252")),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var docStyle = lipgloss.NewStyle().Margin(1,2).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63"))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "enter", " ":
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m model) View() string {
	var caveOfTimeTitle = `
 ██████╗ █████╗ ██╗   ██╗███████╗
██╔════╝██╔══██╗██║   ██║██╔════╝
██║     ███████║██║   ██║█████╗  
██║     ██╔══██║╚██╗ ██╔╝██╔══╝  
╚██████╗██║  ██║ ╚████╔╝ ███████╗
 ╚═════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝
                                 
 ██████╗ ███████╗    ████████╗██╗███╗   ███╗███████╗
██╔═══██╗██╔════╝    ╚══██╔══╝██║████╗ ████║██╔════╝
██║   ██║█████╗         ██║   ██║██╔████╔██║█████╗  
██║   ██║██╔══╝         ██║   ██║██║╚██╔╝██║██╔══╝  
╚██████╔╝██║            ██║   ██║██║ ╚═╝ ██║███████╗
 ╚═════╝ ╚═╝            ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝
 
A Choose Your Own Adventure Game by Edward Packard.
Developed by Adam`

	menu := "\n\n\n"

	for i, option := range m.options {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
			option = m.selectedStyle.Render(option)
		} else {
			option = m.normalStyle.Render(option)
		}

		menu += cursor + " " + option + "\n"
	}

	instructions := "\nUse `↑↓` to navigate, `Enter` to select, `q` to quit"

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Margin(1, 2).
		Width(100)

	content := caveOfTimeTitle + menu + instructions

	return borderStyle.Render(content)
}

func main() {
	model := initModel()
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Printf("some error: %v", err)
		os.Exit(1)
	}
}
