package listing

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/talbx/csfloat/types"
)

func ProcessValidOffers(o []types.FilteredItem, config types.SearchConfig) []types.OutputItem {
	result := make([]types.OutputItem, 0)
	for _, item := range o {

		price := money.New(item.Price, money.USD)
		predicted := money.New(item.PredictedPrice, money.USD)
		minPrice := money.New(int64(config.MinPrice), money.USD)

		bigFloatPrice := ((float64(item.PredictedPrice) - float64(item.Price)) / float64(item.PredictedPrice)) * 100
		bigFloatPriceShortened := fmt.Sprintf("%.2f%%", bigFloatPrice)

		discount, err := predicted.Subtract(price)
		if err != nil {
			panic(err)
		}

		if bigFloatPrice >= config.MinDiscountPercentage && price.Amount() >= minPrice.Amount() {
			r := types.OutputItem{
				Name:         item.Name,
				Wear:         item.Wear,
				Price:        price.Display(),
				Discount:     discount.Display(),
				DiscountPerc: bigFloatPriceShortened,
				MarketPrice:  predicted.Display(),
				TerminatesIn: item.TerminatesIn,
				Link:         fmt.Sprintf("https://csfloat.com/item/%v", item.Id),
			}
			result = append(result, r)
		}
	}
	return result
}
