package shop

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/jodi-lumbantoruan/gosample/auth"
)

var base_url = `https://slicer-staging.tokopedia.com/shop-speed/cube/shop_speed_daily/aggregate`

func GetShopSpeed(shopID int64, monthRange int) {
	// ?cut=shop_id:394674|finish_date:20171027-

	year := int(monthRange / 12)
	month := monthRange % 12

	t := time.Now()
	tt := t.AddDate(-year, -month, 0)

	params := map[string]string{
		"cut": fmt.Sprintf("shop_id:%d|finish_date:%s-", shopID, tt.Format("20060102")),
	}
	u, err := url.Parse(base_url)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	u.RawQuery = q.Encode()

	authorization, contentMD5, fmtTime, _ := auth.GenerateHmac("GET", "https://slicer-staging.tokopedia.com", "/shop-speed/cube/shop_speed_daily/aggregate", "", t, params)
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", base_url, q.Encode()), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-MD5", contentMD5)
	req.Header.Add("Req-Date", fmtTime)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp == nil {
		panic("resp is null")
	}
	txt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(txt))

}
