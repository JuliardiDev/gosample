package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jodi-lumbantoruan/gosample/model"
	oauth_client "github.com/tokopedia/oauth-client"
)

var oauthClient *oauth_client.Client

func init() {
	var err error

	oauthClient, err = oauth_client.New(&oauth_client.Options{
		BaseURL:      "http://devel-go.tkpd:8009",
		ClientID:     "UFVZgQ045uiurKJHaTBeU27J83h3yxTKmkvReJke3WJufjB4uJfxq1AJci5fle3W",
		ClientSecret: "EOlE6pq2T3fi5bpjTNs1H8lQkFZNkdbMFeGKym2EZQnLSj9uTMejH7fuWOeycRWm",
	})

	if err != nil {
		panic(err)
	}
}

func CheckOauthToken(token string) (*oauth_client.Response, error) {

	oauthResponse, err := oauthClient.GetUserInfo(token)
	if err != nil {
		return oauthResponse, err
	}

	return oauthResponse, nil
}

func GetUserInfo(token string) (*model.UserInfo, error) {
	ui := &model.UserInfo{}

	oauthUserInfo, err := CheckOauthToken(token)
	if err != nil {
		return nil, err
	}

	ui.UserID = oauthUserInfo.Data.UserID
	ui.Email = oauthUserInfo.Data.Email
	ui.Fullname = oauthUserInfo.Data.Fullname
	ui.Language = oauthUserInfo.Data.Language

	return ui, nil
}

func GetToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	splitAuth := strings.Split(auth, " ")
	if len(splitAuth) < 2 {
		return "", errors.New("token not found")
	}

	if strings.ToLower(splitAuth[0]) != "bearer" && splitAuth[1] == "" {
		return "", errors.New("token not found")
	}
	return splitAuth[1], nil
}
