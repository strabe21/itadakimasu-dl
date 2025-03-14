package ui

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintTable(title string, header table.Row, rows []table.Row, footer table.Row, config []table.ColumnConfig) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(title)
	t.AppendHeader(header)
	t.AppendRows(rows)
	if footer != nil {
		t.AppendFooter(footer)
	}
	style := table.StyleLight
	style.Options = table.Options{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    true,
	}
	t.SetColumnConfigs(config)
	t.SetStyle(style)
	t.Render()
}
