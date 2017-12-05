package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/google/gops/agent"

	"github.com/jodi-lumbantoruan/gosample/auth"
	"github.com/jodi-lumbantoruan/gosample/hello"
	"github.com/jodi-lumbantoruan/gosample/shop"
	"gopkg.in/tokopedia/grace.v1"
	"gopkg.in/tokopedia/logging.v1"
)

func NicePrint(val interface{}) {
	str, _ := json.MarshalIndent(val, "", "   ")
	log.Println(string(str))
}

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal(err)
	}

	hwm := hello.NewHelloWorldModule()
	// reqDate := time.Now()

	shop.GetShopSpeed(394674, 12)

	http.HandleFunc("/hello", hwm.SayHelloWorld)
	http.HandleFunc("/hello/name", hwm.HandleHelloNameWithParam)

	http.HandleFunc("/hello/login", auth.OptionalAuthorize(hwm.HandleHelloOptionalLogin))
	http.HandleFunc("/hello/mustlogin", auth.MustAuthorize(hwm.HandleHelloMustLogin))
	// go logging.StatsLog()

	log.Fatal(grace.Serve(":9000", nil))
}
