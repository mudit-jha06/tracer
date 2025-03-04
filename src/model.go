package main

// Defines the UI model for the TUI

// model struct type implements the tea.Model interface
type model struct {
	passageText        []string
	prevCorrectWords   []string
	prevIncorrectWords []string
	currentTypedWord   string
}
