package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	TEXT                    = "Raindrops tapped against the window as the wind howled outside. A candle flickered, casting warm light over scattered books and forgotten notes. The clock ticked steadily, marking time’s quiet passage. Somewhere, a distant train rumbled, its whistle fading into the night, leaving only echoes in the silence."
	DEFAULT_TIME_IN_SECONDS = 10
)

var (
	// Define some styles using Lipgloss
	//36 is green, 40 is bright green, 196 is red, 185 is yellow
	//titleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("185")).PaddingLeft(1).PaddingRight(1)
	//highlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Bold(true)
)

func main() {
	//Get initial state of the TUI:
	initState := getInititalState()
	//timeoutCtx, _ := context.WithTimeout(context.Background(), DEFAULT_TIME_IN_SECONDS*time.Second)
	p := tea.NewProgram(initState)
	// Start timer in new goroutine
	go func() {
		<-time.After(DEFAULT_TIME_IN_SECONDS * time.Second)
		//Call Interrupt() method of Program
		//fmt.Println("Time over, sending interrupt msg")
		p.Send(tea.SuspendMsg{})
	}()
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

func getInititalState() *model {
	passageText := getRandomPassage()
	return &model{
		passageText:             passageText,
		prevCorrectWords:        []string{},
		prevIncorrectWords:      []string{},
		currentPassageTextIndex: 0,
		currentTypedWord:        "",
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.SuspendMsg:
		fmt.Println("TIME OVER ---")
		fmt.Println("*********************")

	// Is it a key press?
	case tea.KeyMsg:
		if msg.Type != tea.KeyRunes {
			switch msg.String() {
			//How to trigger UI update on other key presses
			// These keys should exit the program.
			case "ctrl+c":
				fmt.Println("Quitting the program")
				fmt.Println("Your stats:")
				fmt.Println("Correctly typed words: ", len(m.prevCorrectWords))
				fmt.Println("Incorrectly typed words:", len(m.prevIncorrectWords))
				return m, tea.Quit
			case "backspace":
				//Update currently typed word
				length := len(m.currentTypedWord)
				if length > 0 {
					m.currentTypedWord = m.currentTypedWord[:length-1]
				}
			case " ":
				//Pressing space should update stats
				if len(m.currentTypedWord) == len(m.passageText[0]) &&
					m.currentTypedWord == m.passageText[0] {
					m.prevCorrectWords = append(m.prevCorrectWords, m.currentTypedWord)
				} else {
					m.prevIncorrectWords = append(m.prevIncorrectWords, m.currentTypedWord)
				}
				m.currentTypedWord = ""
				// Re-render
				m.passageText = m.passageText[1:]

				return m, nil
			}
		} else {
			// Key pressed is one of the letters on the keyboard - update model state accordingly
			m.currentTypedWord += string(msg.Runes)
			//fmt.Println(m.currentTypedWord)
			return m, nil

		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

// View gives a stringified representation of the state of the TUI
// Here, that is the passage text followed by the current typed word in the next line

// View() is called every time the model is updated
// TODO: Add padding around first word
func (m model) View() string {
	passageText := strings.Join(m.passageText[1:], " ")
	lineBreak := "\n==============================================\n"
	enterTextHerePrompt := "Type text here >"
	//return passageText + "\n" + lineBreak + m.currentTypedWord
	return fmt.Sprintf("%s%s\n%s\n%s %s",
		textStyle.Render(m.passageText[0]), passageText, lineBreak, enterTextHerePrompt, m.currentTypedWord)
}

/*
func (m model) View() string {
	return fmt.Sprintf("%s\n%s\n%s",
		titleStyle.Render("Bubble Tea TUI App"),
		textStyle.Render("Welcome to the colorful world of Bubble Tea!"),
		highlightStyle.Render("Press 'q' to quit."),
	)
}
*/
