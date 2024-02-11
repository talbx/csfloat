package main

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/talbx/csfloat/types"
	"os"
)

func WriteOut(result []types.OutputItem, title string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(title)
	t.AppendHeader(table.Row{"#", "Skin", "Wear", "Price", "Discount", "Base Price"})
	t.SetStyle(table.StyleColoredBright)
	for index, r := range result {
		t.AppendRow(table.Row{
			index, r.Name, r.Wear, r.Price, r.DiscountPerc, r.MarketPrice,
		})
	}

	t.AppendSeparator()
	t.Render()
}
