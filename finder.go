package main

import (
	"github.com/talbx/csfloat/listing"
	auction "github.com/talbx/csfloat/listing/auction"
	"github.com/talbx/csfloat/listing/instant"
	"github.com/talbx/csfloat/types"
	"github.com/talbx/csfloat/util"
	"log"
)

type SkinFinder = func(f *types.SearchConfig)

func FindSkins(searchConfig *types.SearchConfig) {
	instants := findInstantBuys(*searchConfig)
	buynowWriter := util.BuyNowWriter{}
	buynowWriter.WriteOut(instants)

	if searchConfig.Auctions {
		auctions := findAuctions(*searchConfig)
		auctionWriter := util.AuctionWriter{}
		if len(auctions) > 0 {
			auctionWriter.WriteOut(auctions)
		} else {
			log.Default().Println("No Auctions found matching the default criteria!")
		}
	}
}

func findInstantBuys(config types.SearchConfig) []types.OutputItem {
	instantLister := instant.Instant{}
	buyNowFilter := instant.BuyNowFilter{}
	instantListings, err := instantLister.GetListings(config)
	if err != nil {
		log.Default().Fatal(err)
	}
	filteredOffers := buyNowFilter.Filter(instantListings, config)
	return listing.ProcessValidOffers(filteredOffers, config)
}

func findAuctions(searchConfig types.SearchConfig) []types.OutputItem {
	auctionLister := auction.Auction{}
	auctionFilter := auction.AuctionFilter{}
	auctionListings, err := auctionLister.GetListings(searchConfig)
	if err != nil {
		log.Default().Fatal(err)
	}
	filteredAuctions := auctionFilter.Filter(auctionListings, searchConfig)
	return listing.ProcessValidOffers(filteredAuctions, searchConfig)
}
