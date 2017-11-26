package hello

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jodi-lumbantoruan/gosample/auth"
	"github.com/jodi-lumbantoruan/gosample/template"
)

func (hwm *HelloWorldModule) HandleHelloOptionalLogin(w http.ResponseWriter, r *http.Request) {
	hwm.stats.Add(1)
	serverTime := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

	loggedIn := r.Context().Value("LoggedIn").(bool)
	userData := r.Context().Value("User").(*auth.User)

	if loggedIn {
		template.ResponseOK(w, nil, []string{fmt.Sprintf("Hello %s, you've logged in", userData.FullName)}, serverTime)
		return
	}
	template.ResponseOK(w, nil, []string{"Wellcome, Anonymous"}, serverTime)
}

func (hwm *HelloWorldModule) HandleHelloMustLogin(w http.ResponseWriter, r *http.Request) {
	hwm.stats.Add(1)
	serverTime := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

	userData := r.Context().Value("User").(*auth.User)

	template.ResponseOK(w, nil, []string{fmt.Sprintf("Hello %s, you've logged in", userData.FullName)}, serverTime)
	return
}
