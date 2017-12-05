package shop

import (
	"encoding/json"
	"fmt"

	"github.com/jodi-lumbantoruan/gosample/model"
	"github.com/jodi-lumbantoruan/gosample/utils"
)

func GetShopInfo(shopID int64, shopDomain string) (*model.ShopInfoResponse, error) {
	params := map[string]string{
		"shop_id":     fmt.Sprintf("%d", shopID),
		"shop_domain": shopDomain,
		"device_id":   "b",
		"os_type":     "1",
	}

	jsonRaw, err := utils.MakeWSRequest("POST", "https://ws.tokopedia.com", "/v4/shop/get_shop_info.pl", params)
	if err != nil {
		return nil, err
	}

	result := &model.ShopInfoResponse{}

	err = json.Unmarshal(jsonRaw, result)
	if err != nil {
		return nil, err
	}

	return result, nil

}
