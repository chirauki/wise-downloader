package wise

import "github.com/jedib0t/go-pretty/v6/table"

var (
	defaultTableStyle = table.Style{
		Name:   "StyleDefault",
		Box:    table.StyleBoxDefault,
		Color:  table.ColorOptionsDefault,
		Format: table.FormatOptionsDefault,
		HTML:   table.DefaultHTMLOptions,
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateHeader:  false,
		},
		Title: table.TitleOptionsDefault,
	}
)
