package api

import (
	"io"
	"net/http"

	"github.com/thetogi/YReserve2/api/wrapper"
)

func (a *API) InitHealthCheck() {
	a.Router.APIRoot.Handle("/healthCheck", a.requestHandler(a.HealthCheck))
}

// HealthCheck func is used to check health check status
func (a *API) HealthCheck(c *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	// Check for status of DB or cache (Redis) in future
	io.WriteString(w, `{"alive": true}`)
}
