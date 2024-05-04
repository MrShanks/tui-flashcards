package main

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

var correctAnswerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00EE00"))
