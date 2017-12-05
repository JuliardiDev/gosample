package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

func GenerateHmac(method, domain, path, contentType string, timeNow time.Time, params map[string]string) (authorization, contentMD5, fmtTime string, err error) {

	//prepare the URL
	//parse the param
	u, err := url.Parse(fmt.Sprintf("%s%s", domain, path))
	if err != nil {
		return "", "", "", err
	}

	q := u.Query()
	for key, val := range params {
		q.Add(key, val)
	}

	u.RawQuery = q.Encode()

	//generate the Content MD5
	queryEncode, _ := url.QueryUnescape(u.RawQuery)
	h := md5.New()
	io.WriteString(h, queryEncode)
	contentMD5 = hex.EncodeToString(h.Sum(nil))

	key := "web_service_v4"
	fmtTime = timeNow.Format("Mon, 02 Jan 2006 15:04:05 +0700")
	hmacData := strings.Join([]string{method, contentMD5, contentType, fmtTime, path}, "\n")

	signature := hmac.New(sha1.New, []byte(key))
	signature.Write([]byte(hmacData))

	hmacKey := base64.StdEncoding.EncodeToString(signature.Sum(nil))
	authorization = "TKPD Tokopedia:" + hmacKey
	return
}
