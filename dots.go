package dots

import (
	"encoding/base64"
)

// API
type API struct {
	cl *client
	h  string
}

// New
func New(clientID, apiKey string, debug bool) *API {

	var host string

	if debug {
		host = sandbox
	} else {
		host = production
	}

	b := []byte(clientID + ":" + apiKey)
	t := base64.StdEncoding.EncodeToString(b)

	cl := newClient(t)

	return &API{cl: cl, h: host}
}
