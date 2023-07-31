// Experiment to use Go Pretty for creating a table of output

package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBright)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Sacks of Grain"})
	t.AppendRows([]table.Row{
		{1, "Foo", "Bar", 4000},
		{2, "Baz", "Bill", 5000},
	})
	t.Render()
}
