package dots

import (
	"encoding/base64"
	"net/http"
	"time"
)

// API
type API struct {
	cl    *http.Client
	token string
}

// New
func New(clientID, apiKey string, sandbox bool) *API {

	b := []byte(clientID + ":" + apiKey)
	t := base64.StdEncoding.EncodeToString(b)

	cl := &http.Client{
		Timeout: time.Second * 3,
	}

	return &API{cl: cl, token: t}
}
