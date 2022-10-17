package dots

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	_ = New(clientID, apiKey, true)
}
