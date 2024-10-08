package util

import (
	"encoding/json"
	"fmt"
	"github.com/talbx/csfloat/pkg/types"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetListings[R types.BuyNowItem | types.AuctionItem](filter types.SearchConfig, searchType string, returnStruct []R) ([]R, error) {
	f := []string{
		fmt.Sprintf("&category=%v", filter.Category),
		"&sort_by=highest_discount",
	}

	if filter.MaxPrice != 0 {
		f = append(f, fmt.Sprintf("&max_price=%v", filter.MaxPrice))
	}

	if filter.MinPrice != 0 {
		f = append(f, fmt.Sprintf("&min_price=%v", filter.MinPrice))
	}

	if filter.Gun != "" {
		f = append(f, fmt.Sprintf("&market_hash_name=%v", filter.Gun))
	}

	if filter.DefIndex != nil && len(filter.DefIndex) > 0 {
		f = append(f, fmt.Sprintf("&def_index=%v", strings.Join(filter.DefIndex, ",")))
	}

	filtered := strings.Join(f, "")
	l := fmt.Sprintf("https://csfloat.com/api/v1/listings?type=%v", searchType) + filtered
	request, err := http.NewRequest(http.MethodGet, l, nil)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(filter.Keyfile)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", string(file))

	c := http.Client{Timeout: time.Duration(1) * time.Second}

	do, err := c.Do(request)

	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resBody, &returnStruct)
	if err != nil {
		return nil, err
	}

	return returnStruct, nil
}
