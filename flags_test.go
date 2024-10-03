package main

import (
	"github.com/spf13/pflag"
	"github.com/talbx/csfloat/types"
	"os"
	"reflect"
	"testing"
)

var def = types.SearchConfig{
	MaxPrice:              1000,
	MinPrice:              100,
	Category:              1,
	Gun:                   "",
	Keyfile:               "key-temp",
	MinDiscountPercentage: 10,
	Stickers:              false,
	Top:                   10,
	Keyword:               "",
	Cron:                  false,
	Auctions:              false,
}

func TestParseFlags(t *testing.T) {
	_, _ = os.Create("key-temp")
	defer os.Remove("key-temp")
	var maxPriceEnhanced = def
	maxPriceEnhanced.MaxPrice = 2500

	tests := []struct {
		name    string
		flagset *pflag.FlagSet
		want    *types.SearchConfig
		wantErr bool
	}{
		{name: "default", flagset: createFlagSet(&def), want: &def, wantErr: false},
		{name: "maxPriceEnhanced", flagset: createFlagSet(&maxPriceEnhanced), want: &maxPriceEnhanced, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFlags(tt.flagset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlags() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func createFlagSet(cfg *types.SearchConfig) *pflag.FlagSet {
	set := pflag.NewFlagSet("default", 0)
	set.Int("max", cfg.MaxPrice, "")
	set.Int("min", cfg.MinPrice, "")
	set.Int("category", cfg.Category, "")
	set.Int("top", cfg.Top, "")
	set.Float64("discount", cfg.MinDiscountPercentage, "")
	set.String("keyword", cfg.Keyword, "")
	set.String("keyfile", cfg.Keyfile, "")
	set.Bool("auctions", cfg.Auctions, "")
	set.Bool("stickers", cfg.Stickers, "")
	set.Bool("cron", cfg.Cron, "")
	return set
}
