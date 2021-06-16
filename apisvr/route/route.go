package route

import (
    // External dependencies
    "github.com/gorilla/mux"

    // Constants
    "github.com/juancolamendy/water-jug-riddle/lib-service/utils/constant"

    // Controllers
    "github.com/juancolamendy/water-jug-riddle/apisvr/controller/indexctl"
)

func GetRouter() *mux.Router {
	// Create router
	rtr := mux.NewRouter()

	// index_controller
	rtr.HandleFunc(constant.RTE_INDEX, indexctl.GetIndexPage).Methods(constant.HTTP_GET)

	// Return router
	return rtr
}