package dots

import "testing"

const (
	clientID = "changeme"
	apiKey   = "changeme"
)

func TestNew(t *testing.T) {

	_ = New(clientID, apiKey, true)
}
