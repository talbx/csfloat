package listing

import (
	"fmt"
	"github.com/talbx/csfloat/types"
	"strings"
)

var gunsOfInterest = []string{
	"Glock-18",
	"P250",
	"P2000",
	"USP-S",
	"Five-SeveN",
	"Desert Eagle",
	"XM1014",
	"MAG-7",
	"MAC-10",
	"P90",
	"MP7",
	"Galil",
	"M4A1",
	"M4A4",
	"AK-47",
	"SSG",
	"AWP",
	"Knife",
}

type GunFilter struct {
	config *types.FilterConfig
}

func NewGunFilter(conf *types.FilterConfig) *GunFilter {
	return &GunFilter{conf}
}

func (f *GunFilter) Filter(gun types.Gun) error {
	if err := f.filterSticker(gun); err != nil {
		return err
	}
	if err := f.filterGunType(gun); err != nil {
		return err
	}
	return nil
}

func (f *GunFilter) filterSticker(gun types.Gun) error {
	if len(gun.Stickers) != 0 {
		return fmt.Errorf("gun %v contains stickers", gun.Name)
	}
	return nil
}

func (f *GunFilter) filterGunType(gun types.Gun) error {
	for _, weapon := range gunsOfInterest {
		if strings.Contains(gun.Name, weapon) {
			return nil
		}
	}
	return fmt.Errorf("weapon %v not in weapon whitelist", gun.Name)
}
