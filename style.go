package main

import "github.com/charmbracelet/lipgloss"

// Define styles
var styleHeader = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#1CBAD6")).
	Padding(1, 2).
	Width(100).
	Align(lipgloss.Center)

var styleTextarea = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7D56F4")).
	Padding(1, 2).
	Margin(1, 0)

var styleInstructions = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#666666")).
	Italic(true).
	Align(lipgloss.Center).
	Margin(1, 0)

var styleSuccess = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#1CBAD6")).
	Background(lipgloss.Color("#1A1A1A")).
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#1CBAD6")).
	Width(100).
	Align(lipgloss.Center)

var tweetSuccess = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7D56F4")).
	Padding(1, 2).
	Margin(1, 0).
	Width(100)
