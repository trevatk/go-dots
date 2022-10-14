package dots

import (
	"encoding/base64"
)

// API
type API struct {
	cl *client
}

// New
func New(clientID, apiKey string, sandbox bool) *API {

	b := []byte(clientID + ":" + apiKey)
	t := base64.StdEncoding.EncodeToString(b)

	cl := newClient(t)

	return &API{cl: cl}
}
