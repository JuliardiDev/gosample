package hello

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jodi-lumbantoruan/gosample/template"
)

func (hwm *HelloWorldModule) HandleHelloNameWithParam(w http.ResponseWriter, r *http.Request) {
	hwm.stats.Add(1)
	serverTime := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

	name := r.FormValue("name")

	// this is HTTP, it's text
	// there is no null value for string, just empty string
	//  no, there is is_set like PHP
	if name == "" {
		template.ResponseError(w, http.StatusBadRequest, []string{"Who are you? Name Please"}, serverTime)
		return
	}

	data := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	template.ResponseOK(w, data, []string{fmt.Sprintf("Why, Hello there %s", name)}, serverTime)

}
