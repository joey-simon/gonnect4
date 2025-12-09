package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const rowCount = 6
const colCount = 7

const (
	Empty = iota
	Red
	Yellow
)

const (
	R = Red
	Y = Yellow
)

type GameState struct {
	board [][]int
	turn  int
}

func (m *GameState) drop(position int) {
	for row := 0; row < rowCount; row++ {
		if m.board[row][position] == Empty {
			m.board[row][position] = m.turn
			// swap turns
			if m.turn == Red {
				m.turn = Yellow
			} else {
				m.turn = Red
			}
			return
		}
	}
}

func initialState() GameState {
	// Create empty board
	board := make([][]int, rowCount)
	for i := range board {
		board[i] = make([]int, colCount)
	}
	return GameState{board: board, turn: Red}
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
	case Red:
		return redStyle.Render("⬤")
	case Yellow:
		return yellowStyle.Render("⬤")
	default:
		return emptyStyle.Render("⬤")
	}
}

func renderBoard(board [][]int) string {
	var rows []string

	for row := len(board) - 1; row >= 0; row-- {
		var cols []string
		cols = append(cols,
			lipgloss.NewStyle().
				Background(lipgloss.Color("#003887")).
				Width(1).
				Render(),
		)

		for _, cell := range board[row] {
			cols = append(cols, renderToken(cell))
		}

		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cols...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m GameState) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m GameState) View() string {
	return renderBoard(m.board)
}

// ---- UPDATE ----
func (m GameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1", "2", "3", "4", "5", "6", "7":
			col := int(msg.String()[0] - '1') // convert "1".."7" → 0..6
			m.drop(col)
			return m, nil
		}
	}
	return m, nil
}

func main() {
	p := tea.NewProgram(initialState())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
