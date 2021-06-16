package indexctl

import (
	// External dependencies
	"net/http"

	"github.com/juancolamendy/water-jug-riddle/apisvr/server/wsserver"
)

func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	svr := wsserver.NewWsServer()
	svr.Serve(w, r)
}
