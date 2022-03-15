package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Ui struct {
	spinner spinner.Model
	loading bool
}

func (u Ui) Init() tea.Cmd {
	return u.spinner.Tick
}

func (u Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	default:
		var cmd tea.Cmd
		u.spinner, cmd = u.spinner.Update(msg)
		return u, cmd
	}
}

func (u Ui) View() string {
	str := fmt.Sprintf("\n\n   %s Loading forever...press q to quit\n\n", u.spinner.View())

	return str
}

func (u Ui) SetLoading(loading bool) {
	u.loading = loading

}

func New() Ui {
	return Ui{
		spinner: spinner.New(),
		loading: false,
	}
}

func (u Ui) Start() {
	program := tea.NewProgram(u)
	if err := program.Start(); err != nil {
		panic(err)
	}
}
