package limeanger

import (
	"log/slog"

	"github.com/pterm/pterm"
)

func test() {
	var options []string

	options = append(options, "https://raw.githubusercontent.com/taylormonacelli/navylie/master/templates/1.txtar")
	options = append(options, "https://raw.githubusercontent.com/taylormonacelli/navylie/master/templates/2.txtar")

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
	slog.Debug("selected item", "item", selectedOption)
}
