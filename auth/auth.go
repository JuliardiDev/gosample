package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jodi-lumbantoruan/gosample/template"
)

func OptionalAuthorize(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userData, err := authorize_sites(r)
		log.Println(userData)

		//we saved it to the context
		r = r.WithContext(context.WithValue(r.Context(), "LoggedIn", err == nil))
		r = r.WithContext(context.WithValue(r.Context(), "User", userData))

		h(w, r)
	}
}

func MustAuthorize(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		userData, err := authorize_sites(r)
		log.Println(userData)

		if err != nil {
			template.ResponseError(w, http.StatusUnauthorized, []string{"Invalid session"}, t)
			return
		}
		//we saved it to the context
		r = r.WithContext(context.WithValue(r.Context(), "User", userData))

		h(w, r)
	}
}
