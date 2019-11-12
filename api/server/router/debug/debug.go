package debug

import (
	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/server/router"
	"github.com/docker/docker/pkg/signal"
	"golang.org/x/net/context"
	"net/http"
)

type debugRouter struct {
	backend interface{}
	routes  []router.Route
}

// NewRouter initializes a new troubleshooting router
func NewRouter(b interface{}) router.Router {
	r := &debugRouter{
		backend: b,
	}
	r.initRoutes()
	return r
}

func (r *debugRouter) initRoutes() {
	r.routes = []router.Route{
		// GET
		router.NewGetRoute("/debug/dump", r.getDump),
	}
}

func (r debugRouter) Routes() []router.Route {
	return r.routes
}

func (v *debugRouter) getDump(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	return httputils.WriteJSON(w, http.StatusOK, signal.GetStackDumpInformation())
}
