package auth

import (
	"errors"
	"net/http"
)

// the key that being used in cookie
// see https://phab.tokopedia.com/w/tech/accounts/#accounts-hostname-cookie
var SessionKey = map[string]string{
	"production":  "_SID_Tokopedia_",
	"staging":     "_SID_Tokopedia_Coba_",
	"alpha":       "_SID_Tokopedia_Alpha_",
	"development": "_SID_Tokopedia_",
}

//the env should be taken is environtment variables
//which is different in every stage of developement
//for now, let use the development
var env = "development"

//of course tou can add the information that you want.
type User struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"user_email"`
	FullName string `json:"full_name"`
	Lang     string `json:"lang"`
	Status   int16  `json:"status"`
	ShopID   int64  `json:"shop_id"`
}

func authorize_sites(r *http.Request) (*User, error) {

	//get the cookie
	cookie, err := r.Cookie(SessionKey[env])
	if err != nil {
		//dont return nil, it will give nil pointer derefence exception sometimes.
		return &User{}, err
	}

	//SessionHelper object is initilized in /auth/init.go
	userSession, err := SessionHelper.GetUser(cookie.Value)
	if err != nil {
		return &User{}, err
	}

	if userSession == nil {
		return &User{}, errors.New("User data not found")
	}

	return &User{
		UserID:   userSession.UserID,
		Email:    userSession.Email,
		FullName: userSession.FullName,
		Lang:     userSession.Lang,
		Status:   userSession.Status,
		ShopID:   userSession.ShopID,
	}, nil
}
