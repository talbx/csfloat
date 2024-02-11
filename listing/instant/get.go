package instant

import (
	"github.com/talbx/csfloat/types"
	"github.com/talbx/csfloat/util"
)

type Instant struct{}

func (i *Instant) GetListings(filter types.SearchConfig) ([]types.BuyNowItem, error) {
	var ret []types.BuyNowItem
	listings, err := util.GetListings(filter, "buy_now", ret)
	if err != nil {
		return nil, err
	}
	return listings, nil
}
