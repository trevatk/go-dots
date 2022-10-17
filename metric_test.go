package dots

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactionVolumeMetrics(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputGetTransactionVolumeMetricsParams{}

	r, e := api.GetTransactionVolumeMetrics(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}
