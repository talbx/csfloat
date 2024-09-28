package instant

import (
	"github.com/talbx/csfloat/listing"
	"github.com/talbx/csfloat/types"
)

type BuyNowFilter struct{}

func (a *BuyNowFilter) Filter(ofs []types.BuyNowItem, config types.SearchConfig) []types.FilteredItem {
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
				continue
			}

			all = append(all, types.FilteredItem{
				Item: o.Item,
			})
		}
	}
	return all
}
