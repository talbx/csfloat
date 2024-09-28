package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/spf13/pflag"
	"github.com/talbx/csfloat/types"
	"os"
)

func ParseFlags(flagset *pflag.FlagSet) (*types.SearchConfig, error) {

	config := types.SearchConfig{}
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

	discount, err := flagset.GetFloat64("discount")
	if err != nil {
		return nil, err
	}
	config.MinDiscountPercentage = discount

	mini, err := flagset.GetInt("min")
	if err != nil {
		return nil, err
	}
	config.MinPrice = mini

	cat, err := flagset.GetInt("category")
	if err != nil {
		return nil, err
	}
	config.Category = cat

	keyword, err := flagset.GetString("keyword")
	if err != nil {
		return nil, err
	}
	config.Keyword = keyword

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

	cron, err := flagset.GetBool("cron")
	if err != nil {
		return nil, err
	}
	config.Cron = cron

	auctions, err := flagset.GetBool("auctions")
	if err != nil {
		return nil, err
	}
	config.Auctions = auctions

	if config.MinPrice == 0 {
		config.MinPrice = maxi / 10
	}

	fmt.Println(prettyPrint(config))
	return &config, nil
}

func prettyPrint(cfg types.SearchConfig) string {
	var cat string
	if cfg.Category == 1 {
		cat = "1 (Normal)"
	} else {
		cat = "2 (Normal)"
	}

	var auctions string
	if cfg.Auctions {
		auctions = "Yes"
	} else {
		auctions = "No"
	}

	var stickers string
	if cfg.Stickers {
		stickers = "Yes"
	} else {
		stickers = "No"
	}

	var cron string
	if cfg.Cron {
		cron = "Yes"
	} else {
		cron = "No"
	}
	var keyword string
	if cfg.Keyword != "" {
		keyword = cfg.Keyword
	} else {
		keyword = "n/a"
	}
	printCfg := types.PrintConfig{
		MaxPrice:              money.New(int64(cfg.MaxPrice), money.USD).Display(),
		MinPrice:              money.New(int64(cfg.MinPrice), money.USD).Display(),
		Category:              cat,
		Keyfile:               cfg.Keyfile,
		MinDiscountPercentage: fmt.Sprintf("%v%v", cfg.MinDiscountPercentage, "%"),
		Stickers:              stickers,
		Top:                   cfg.Top,
		Keyword:               keyword,
		Cron:                  cron,
		Auctions:              auctions,
	}
	s, _ := json.MarshalIndent(printCfg, "", "\t")
	return string(s)
}
