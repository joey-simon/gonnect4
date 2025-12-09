package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	board [][]int
}

func initialModel() Model {
	// 6 rows x 7 columns empty board
	board := make([][]int, 6)
	for i := range board {
		board[i] = make([]int, 7)
	}
	return Model{board: board}
}

// ---- VIEW ----
var blueCell = lipgloss.NewStyle().
	Background(lipgloss.Color("#003887")).
	Width(2).Height(1).Align(lipgloss.Left)

var (
	emptyStyle  = blueCell.Foreground(lipgloss.Color("#000000"))
	redStyle    = blueCell.Foreground(lipgloss.Color("#FF5555"))
	yellowStyle = blueCell.Foreground(lipgloss.Color("#ffff59"))
)

func renderToken(v int) string {
	switch v {
	case 1:
		return redStyle.Render("⬤")
	case 2:
		return yellowStyle.Render("⬤")
	default:
		return emptyStyle.Render("⬤")
	}
}

func renderBoard(board [][]int) string {
	var rows []string
	for _, row := range board {
		var cols []string
		cols = append(cols, lipgloss.NewStyle().Background(lipgloss.Color("#003887")).Width(1).Render())
		for _, cell := range row {
			cols = append(cols, renderToken(cell))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cols...))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m Model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m Model) View() string {
	return renderBoard(m.board)
}

// ---- UPDATE ----
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
