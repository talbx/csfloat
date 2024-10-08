package util

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/talbx/csfloat/pkg/types"
	"os"
)

type AuctionWriter struct{}
type BuyNowWriter struct{}

func (w *AuctionWriter) WriteOut(result []types.OutputItem) {
	header := table.Row{"#", "Skin", "Wear", "Curr. Bid", "Discount %", "Deadline", "Base Price", "Link"}
	appenderFn := func(index int, r types.OutputItem) table.Row {
		return table.Row{
			index, r.Name, r.Wear, r.Price, r.DiscountPerc, r.TerminatesIn, r.MarketPrice, r.Link,
		}
	}
	writeOut(result, "AUCTION -- Listings", header, appenderFn)
}

func (b *BuyNowWriter) WriteOut(result []types.OutputItem) {
	header := table.Row{"#", "Skin", "Wear", "Price", "Discount", "Base Price", "Link"}
	appenderFn := func(index int, r types.OutputItem) table.Row {
		return table.Row{
			index, r.Name, r.Wear, r.Price, r.DiscountPerc, r.MarketPrice, r.Link,
		}
	}
	writeOut(result, "BUY NOW -- Listings", header, appenderFn)
}

func writeOut(result []types.OutputItem, title string, header table.Row, appenderFn func(index int, item types.OutputItem) table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(title)
	t.AppendHeader(header)
	t.SetStyle(table.StyleColoredBright)
	for index, r := range result {
		index = index + 1
		t.AppendRow(appenderFn(index, r))
	}
	t.AppendSeparator()
	t.Render()
}
