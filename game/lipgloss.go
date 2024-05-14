package game

import "github.com/charmbracelet/lipgloss"

var normal = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder(), true, false).
	BorderForeground(lipgloss.Color("#3C3C3C")).
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#3C3C3C")).
	PaddingTop(1).
	PaddingBottom(1).
	Align(lipgloss.Center).
	Width(22)

var wrong = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder(), true, false).
	BorderForeground(lipgloss.Color("#3C3C3C")).
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#E7625F")).
	PaddingTop(1).
	PaddingBottom(1).
	Align(lipgloss.Center).
	Width(22)

var correct = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder(), true, false).
	BorderForeground(lipgloss.Color("#3C3C3C")).
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#8BCD50")).
	PaddingTop(1).
	PaddingBottom(1).
	Align(lipgloss.Center).
	Width(22)

// defaultAnswerStyle is rendered White
var defaultAnswerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEEEEE"))

// correctAnswerStyle is rendered Green
var correctAnswerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00EE00"))

// warningAnswerStyle is rendered Orange
var warningAnswerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEEE00"))

// wrongAnswerStyle is rendered Red
var wrongAnswerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EE0000"))

// exampleStyle is rendered Blue
var exampleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000EE"))
