package writer

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/talbx/csfloat/types"
	"os"
)

type AuctionWriter struct{}

func (w *AuctionWriter) WriteOut(result []types.OutputItem) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("AUCTION -- Listings")
	t.AppendHeader(table.Row{"#", "Skin", "Wear", "Curr. Bid", "Discount %", "Deadline", "Base Price", "Link"})
	t.SetStyle(table.StyleColoredBright)
	for index, r := range result {
		index = index + 1
		t.AppendRow(table.Row{
			index, r.Name, r.Wear, r.Price, r.DiscountPerc, r.TerminatesIn, r.MarketPrice, r.Link,
		})
	}

	t.AppendSeparator()
	t.Render()
}
