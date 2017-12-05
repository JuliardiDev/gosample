package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jodi-lumbantoruan/gosample/auth"
)

func MakeWSRequest(method, domain, path string, params map[string]string) (respomseText []byte, err error) {
	reqDate := time.Now()
	authorization, contentMD5, timeStr, _ := auth.GenerateHmac(method, domain, path, "application/x-www-form-urlencoded", reqDate, params)

	h := md5.New()
	io.WriteString(h, "~b")
	hash := hex.EncodeToString(h.Sum(nil))

	params["hash"] = hash
	params["device_time"] = fmt.Sprintf("%d", reqDate.Unix())

	body := make([]string, 0)
	//build the body
	for key, val := range params {
		body = append(body, fmt.Sprintf("%s=%s", key, val))
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", domain, path), bytes.NewBuffer([]byte(strings.Join(body, "&"))))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-MD5", contentMD5)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Date", timeStr)
	req.Header.Add("X-Method", method)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		return ioutil.ReadAll(resp.Body)
	}

	return nil, errors.New("Reponse is nil")
}
