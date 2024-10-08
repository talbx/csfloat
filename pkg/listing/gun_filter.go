package listing

import (
	"fmt"
	"github.com/talbx/csfloat/pkg/types"
	"strings"
)

type GunFilter struct {
	config *types.SearchConfig
}

func NewGunFilter(conf *types.SearchConfig) *GunFilter {
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
	if f.config.Keyword != "" {
		if strings.Contains(strings.ToLower(gun.Name), strings.ToLower(f.config.Keyword)) {
			return nil
		}
		return fmt.Errorf("weapon %v not in weapon whitelist", gun.Name)
	}
	return nil
}
