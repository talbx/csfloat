package main

import (
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/spf13/pflag"
	"github.com/talbx/csfloat/types"
	"os"
)

func ParseFlags(flagset *pflag.FlagSet) (*types.InputConfig, error) {

	config := types.InputConfig{}
	keyfile, _ := flagset.GetString("keyfile")
	keyFileMsg := "A key file with your CSFloat API key is required.\nEither provide a file called \"key\" in the same dir from where you run float, or provide the path to your key file like \"--keyfile ../path/to/my/keyfile\""

	if keyfile == "" {
		_, err := os.ReadFile("key")
		if err != nil {
			return nil, errors.New(keyFileMsg)
		}
		config.Keyfile = "key"

	} else {
		_, err := os.ReadFile(keyfile)

		if err != nil {
			return nil, errors.New(keyFileMsg)
		}
		config.Keyfile = keyfile
	}

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

	fmt.Println("+++++++++ CSFLOAT SEARCH CONFIG +++++++++")
	fmt.Printf("--- Max Price: %v ---\n", money.New(int64(config.MaxPrice), money.USD).Display())
	fmt.Printf("--- Min Price: %v ---\n", money.New(int64(config.MinPrice), money.USD).Display())
	fmt.Printf("--- Min Discount: %v ---\n", config.MinDiscountPercentage)
	fmt.Printf("--- Min Discount Val: %v ---\n", money.New(int64(config.MinDiscountValue), money.USD).Display())
	fmt.Printf("--- Stickers: %v ---\n", config.Stickers)
	fmt.Printf("--- Category: %v ---\n", config.Category)
	fmt.Printf("--- Gun: %v ---\n", config.Gun)
	fmt.Printf("--- Top: %v ---\n", config.Top)
	fmt.Printf("+++++++++ CSFLOAT SEARCH CONFIG +++++++++\n\n")

	return &config, nil
}
