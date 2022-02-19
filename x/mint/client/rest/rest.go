package rest

import (
	"github.com/gorilla/mux"

	"github.com/okxtuta/go-anatha/client/context"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
}
