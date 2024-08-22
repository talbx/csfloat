package main

import (
	"github.com/talbx/csfloat/listing"
	auction "github.com/talbx/csfloat/listing/auction"
	"github.com/talbx/csfloat/listing/instant"
	"github.com/talbx/csfloat/types"
	"github.com/talbx/csfloat/writer"
	"log"
)

type SkinFinder = func(f *types.InputConfig)

func FindSkins(flags *types.InputConfig, counter int) {
	searchConfig := types.SearchConfig{
		MaxPrice:         flags.MaxPrice,
		MinDiscountValue: flags.MinDiscountValue,
		Category:         flags.Category,
		MinPrice:         flags.MinPrice,
		Gun:              flags.Gun,
		Keyfile:          flags.Keyfile,
	}
	filterConfig := types.NewFilterConfig(flags)

	c := make(chan bool, 0)
	go findAuctions(searchConfig, filterConfig)
	go findInstantBuys(searchConfig, filterConfig)
	<-c

}

func findInstantBuys(searchConfig types.SearchConfig, filterConfig types.FilterConfig) {
	instantLister := instant.Instant{}
	buyNowFilter := instant.BuyNowFilter{}
	instantListings, err := instantLister.GetListings(searchConfig)
	if err != nil {
		log.Default().Fatal(err)
	}
	filteredOffers := buyNowFilter.Filter(instantListings, filterConfig)
	processedBuyNow := listing.ProcessValidOffers(filteredOffers, filterConfig)

	buynowWriter := writer.BuyNowWriter{}
	buynowWriter.WriteOut(processedBuyNow)
}

func findAuctions(searchConfig types.SearchConfig, filterConfig types.FilterConfig) {

	auctionLister := auction.Auction{}
	auctionFilter := auction.AuctionFilter{}
	auctionListings, err := auctionLister.GetListings(searchConfig)
	if err != nil {
		log.Default().Fatal(err)
	}
	filteredAuctions := auctionFilter.Filter(auctionListings, filterConfig)
	processedAuctions := listing.ProcessValidOffers(filteredAuctions, filterConfig)

	auctionWriter := writer.AuctionWriter{}
	if len(processedAuctions) > 0 {
		auctionWriter.WriteOut(processedAuctions)
	} else {
		log.Default().Println("No Auctions found matching the default criteria!")
	}
}
