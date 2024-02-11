package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/spf13/pflag"
	"github.com/talbx/csfloat/types"
)

func ParseFlags(flagset *pflag.FlagSet) (*types.InputConfig, error) {

	config := types.InputConfig{}
	maxi, err := flagset.GetInt("max")
	if err != nil {
		return nil, err
	}
	config.MaxPrice = maxi
	config.MinPrice = maxi / 10

	discount, err := flagset.GetFloat64("discount")
	if err != nil {
		return nil, err
	}
	config.MinDiscountPercentage = discount

	dv, err := flagset.GetInt("discountValue")
	if err != nil {
		return nil, err
	}
	config.MinDiscountValue = dv

	cat, err := flagset.GetInt("category")
	if err != nil {
		return nil, err
	}
	config.Category = cat

	stickers, err := flagset.GetBool("stickers")
	if err != nil {
		return nil, err
	}
	config.Stickers = stickers

	top, err := flagset.GetInt("top")
	if err != nil {
		return nil, err
	}
	config.Top = top

	fmt.Println("+++++++++ STA CONFIG +++++++++")
	fmt.Printf("--- Max Price: %v ---\n", money.New(int64(config.MaxPrice), money.USD).Display())
	fmt.Printf("--- Min Price: %v ---\n", money.New(int64(config.MinPrice), money.USD).Display())
	fmt.Printf("--- Min Discount: %v ---\n", config.MinDiscountPercentage)
	fmt.Printf("--- Min Discount Val: %v ---\n", money.New(int64(config.MinDiscountValue), money.USD).Display())
	fmt.Printf("--- Stickers: %v ---\n", config.Stickers)
	fmt.Printf("--- Category: %v ---\n", config.Category)
	fmt.Printf("--- Top: %v ---\n", config.Top)
	fmt.Printf("+++++++++ END CONFIG +++++++++\n\n")

	return &config, nil
}
