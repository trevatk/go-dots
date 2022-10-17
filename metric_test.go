package dots

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactionVolumeMetrics(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputGetTransactionVolumeMetricsParams{}

	r, e := api.GetTransactionVolumeMetrics(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}
