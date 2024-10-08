package auction

import (
	"fmt"
	"github.com/talbx/csfloat/pkg/listing"
	"github.com/talbx/csfloat/pkg/types"
	"log"
	"math"
	"time"
)

type AuctionFilter struct{}

func (a *AuctionFilter) Filter(ofs []types.AuctionItem, config types.SearchConfig) []types.FilteredItem {
	all := make([]types.FilteredItem, 0)
	gf := listing.NewGunFilter(&config)
	for _, o := range ofs {
		gun := types.Gun{
			Name:     o.Name,
			Wear:     o.Wear,
			Stickers: o.Stickers,
		}
		if len(all) < config.Top {
			if err := gf.Filter(gun); err != nil {
				// log.Default().Println("Auction: Filtering out weapon because: ", err.Error())
				continue
			}
			now := time.Now()

			parse, err := time.Parse(time.RFC3339, o.ExpiresAt)
			if err != nil {
				log.Default().Panic(err)
			}
			sub := parse.Sub(now)

			if sub.Minutes() <= 500 {
				//log.Default().Println("Auction for", o.Name, o.Wear, "expires in", math.Floor(sub.Minutes()), "minutes")
				all = append(all, types.FilteredItem{
					Item:         o.Item,
					TerminatesIn: fmt.Sprintf("%v Minutes", math.Floor(sub.Minutes())),
				})
			}
		}
	}
	return all
}
