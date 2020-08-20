package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers a7-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/a7/user", listUserHandler(cliCtx, "a7")).Methods("GET")
	r.HandleFunc("/a7/user", createUserHandler(cliCtx)).Methods("POST")
}
