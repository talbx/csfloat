package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/talbx/csfloat/types"
	"testing"
)

func TestGet(t *testing.T) {

	var auction []types.AuctionItem
	listings, err := GetListings(types.SearchConfig{
		MaxPrice:         500,
		MinPrice:         200,
		MinDiscountValue: 10,
		Category:         1,
		Keyfile:          "",
	}, "auction", auction)

	assert.Nil(t, listings)
	assert.NotNil(t, err)
}
