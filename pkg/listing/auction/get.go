package auction

import (
	"github.com/talbx/csfloat/pkg/types"
	"github.com/talbx/csfloat/pkg/util"
)

type Auction struct{}

func (i *Auction) GetListings(conf types.SearchConfig) ([]types.AuctionItem, error) {
	var ret []types.AuctionItem
	listings, err := util.GetListings(types.SearchConfig{Category: 1, Keyfile: conf.Keyfile}, "auction", ret)
	if err != nil {
		return nil, err
	}
	return listings, nil
}
