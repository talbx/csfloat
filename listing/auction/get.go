package auction

import (
	"github.com/talbx/csfloat/types"
	"github.com/talbx/csfloat/util"
)

type Auction struct{}

func (i *Auction) GetListings() ([]types.AuctionItem, error) {
	var ret []types.AuctionItem
	listings, err := util.GetListings(types.SearchConfig{Category: 1}, "auction", ret)
	if err != nil {
		return nil, err
	}
	return listings, nil
}
