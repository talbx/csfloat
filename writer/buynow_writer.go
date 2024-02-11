package writer

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/talbx/csfloat/types"
	"os"
)

type BuyNowWriter struct{}

func (w *BuyNowWriter) WriteOut(result []types.OutputItem) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("BUY NOW -- Listings")
	t.AppendHeader(table.Row{"#", "Skin", "Wear", "Price", "Discount", "Base Price"})
	t.SetStyle(table.StyleColoredBright)
	for index, r := range result {
		index = index + 1
		t.AppendRow(table.Row{
			index, r.Name, r.Wear, r.Price, r.DiscountPerc, r.MarketPrice,
		})
	}

	t.AppendSeparator()
	t.Render()
}
