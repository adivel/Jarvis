// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define styles for our layout
var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
	// Styles for the list, focused and unfocused
	focusedStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62"))
	unfocusedStyle = lipgloss.NewStyle().Border(lipgloss.HiddenBorder())
)

// item represents an entry in our list (file or directory)
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// model holds our application's state
type model struct {
	width, height int
	list          list.Model
	input         textinput.Model
	focusIndex    int // 0 for list, 1 for input
}

// initialModel creates the first state of our application
func initialModel() model {
	// --- START OF NEW CODE ---

	// 1. Get current directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current working directory: %v", err)
	}

	// 2. Read directory contents
	dirEntries, err := os.ReadDir(cwd)
	if err != nil {
		log.Fatalf("could not read directory: %v", err)
	}

	// 3. Create list items from directory entries
	items := []list.Item{}
	for _, entry := range dirEntries {
		var newItem item
		if entry.IsDir() {
			newItem = item{title: "📁 " + entry.Name(), desc: "Directory"}
		} else {
			newItem = item{title: "📄 " + entry.Name(), desc: "File"}
		}
		items = append(items, newItem)
	}

	// --- END OF NEW CODE ---

	// Setup the list
	fileList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	fileList.Title = "Current Directory"
	fileList.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)

	// Setup the text input
	searchInput := textinput.New()
	searchInput.Placeholder = "Search..."
	searchInput.Focus() // Start with the search input focused

	return model{
		list:       fileList,
		input:      searchInput,
		focusIndex: 1, // Start focus on search
	}
}

// Init is the first command that is run when the application starts
func (m model) Init() tea.Cmd {
	return textinput.Blink // Start the cursor blinking
}

// Update handles messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Adjust component sizes
		m.list.SetSize(m.width-2, m.height-10) // Example size, adjust as needed
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab", "shift+tab":
			// Cycle focus
			m.focusIndex = (m.focusIndex + 1) % 2
			if m.focusIndex == 1 {
				m.input.Focus()
				m.list.SetDelegate(list.NewDefaultDelegate())
			} else {
				m.input.Blur()
				// You would ideally create a custom delegate for focused state styling
			}
			return m, nil
		}
	}

	// Route updates to the focused component
	if m.focusIndex == 1 {
		m.input, cmd = m.input.Update(msg)
	} else {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

// View renders the UI
func (m model) View() string {
	var listStyle lipgloss.Style
	if m.focusIndex == 0 {
		listStyle = focusedStyle
	} else {
		listStyle = unfocusedStyle
	}

	// Using lipgloss.JoinVertical to stack components in the sidebar
	sidebar := lipgloss.JoinVertical(
		lipgloss.Left,
		m.input.View(),
		listStyle.Render(m.list.View()),
	)

	// A placeholder for the main content for now
	mainContent := "Welcome, Adiyanthy\n\nTerminal output will go here."

	// Join the panes horizontally
	view := lipgloss.JoinHorizontal(lipgloss.Top,
		docStyle.Copy().Width(m.width-32).Render(mainContent), // 32 = sidebar width + margins
		docStyle.Copy().Width(30).Render(sidebar),
	)

	return view
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}