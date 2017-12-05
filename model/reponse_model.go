package model

type ShopOwner struct {
	IsGoldMerchant bool   `json:"is_gold_merchant"`
	OwnerID        int64  `json:"owner_id"`
	OwnerName      string `json:"owner_name"`
}

type ShopInfo struct {
	CreateTime string `json:"date_shop_created"`
	Domain     string `json:"shop_domain"`
	ShopName   string `json:"shop_name"`
	TagLine    string `json:"shop_tagline"`
	Status     int    `json:"shop_status"`
}

type Shop struct {
	IsOpen int       `json:"is_open"`
	Owner  ShopOwner `json:"owner"`
	Info   ShopInfo  `json:"info"`
}

type ShopInfoResponse struct {
	Data              Shop   `json:"data"`
	RequestStatus     string `json:"status"`
	ServerProcessTime string `json:"server_process_time"`
}
