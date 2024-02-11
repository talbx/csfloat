package auction

import (
	"github.com/talbx/csfloat/types"
	"github.com/talbx/csfloat/util"
)

type Auction struct{}

/**
{
		"id": "676412549721361235",
		"created_at": "2024-02-10T13:01:58.991561Z",
		"type": "auction",
		"price": 5400,
		"state": "listed",
		"seller": {
			"avatar": "https://avatars.steamstatic.com/e1572ab4c3e34f69b962239217fbea47f6b130b7_full.jpg",
			"away": false,
			"flags": 48,
			"has_valid_steam_api_key": true,
			"online": false,
			"stall_public": true,
			"statistics": {
				"median_trade_time": 5253,
				"total_avoided_trades": 0,
				"total_failed_trades": 6,
				"total_trades": 762,
				"total_verified_trades": 756
			},
			"steam_id": "76561198309170391",
			"username": "Konvík ❤",
			"verification_mode": "key"
		},
		"reference": {
			"base_price": 44921,
			"predicted_price": 44921,
			"quantity": 9,
			"last_updated": "2024-02-11T07:01:35.078808Z"
		},
		"item": {
			"asset_id": "35109971308",
			"def_index": 1209,
			"sticker_index": 2342,
			"icon_url": "-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXQ9QVcJY8gulRWXk3bSPP_h56EHE59IjtNs7KrFABv3_eGJGkb7t7gktSIxq7xY73Txj8B7sRy0riZpI6j2VLm_EA4Nmj2IY6TIBh-Pw9nWN9fzg",
			"rarity": 6,
			"market_hash_name": "Sticker | zehN (Gold) | Krakow 2017",
			"tradable": 0,
			"inspect_link": "steam://rungame/730/76561202255233023/+csgo_econ_action_preview%20S76561198309170391A35109971308D4611910047419060751",
			"has_screenshot": false,
			"is_commodity": true,
			"type": "sticker",
			"rarity_name": "Extraordinary",
			"type_name": "Sticker",
			"item_name": "zehN (Gold) | Krakow 2017"
		},
		"is_seller": false,
		"min_offer_price": 5346,
		"max_offer_discount": 100,
		"is_watchlisted": false,
		"watchers": 24,
		"auction_details": {
			"reserve_price": 100,
			"top_bid": {
				"id": "676713016192011508",
				"created_at": "2024-02-11T08:55:55.780131Z",
				"price": 5400,
				"contract_id": "676412549721361235",
				"state": "active",
				"obfuscated_buyer_id": "8850653525911407964"
			},
			"expires_at": "2024-02-17T13:01:58.986724Z",
			"min_next_bid": 5500
		}
	},
*/

func (i *Auction) GetListings(filter types.SearchConfig) ([]types.AuctionItem, error) {
	var ret []types.AuctionItem
	listings, err := util.GetListings(filter, "auction", ret)
	if err != nil {
		return nil, err
	}
	return listings, nil
}
