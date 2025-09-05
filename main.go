package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type errMsg error

type model struct {
	textarea textarea.Model
	done     bool
	err      error
}

func initialModel() model {
	ti := textarea.New()
	ti.Placeholder = "In today's tweet, I wanna say..."
	ti.Focus()

	ti.SetWidth(95)

	ti.Prompt = "> "

	return model{
		textarea: ti,
		err:      nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			tempModel := &m
			return tweet(tempModel), tea.Quit

		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	if m.done {
		successMessage := styleSuccess.Render("✅ Tweet posted successfully!")
		tweetContent := tweetSuccess.Render(m.textarea.Value())
		instructions := styleInstructions.Render("Press any key to exit...")

		return fmt.Sprintf("%s\n\n%s\n\n%s", successMessage, tweetContent, instructions)
	}

	header := styleHeader.Render("🐣 What's on your mind?")
	textareaView := styleTextarea.Render(m.textarea.View())
	instructions := styleInstructions.Render("(ctrl+c to quit, Enter to post)")

	return fmt.Sprintf("%s\n\n%s\n\n%s", header, textareaView, instructions)
}
