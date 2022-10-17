package dots

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFlow(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputCreateFlowParams{
		UserID: "5cf4c533-93f3-485c-b9d1-31bd59954e78",
		Steps:  []string{"compliance"},
	}

	f, e := api.CreateFlow(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("successfully created flow %s", f.ID)
}

func TestGetFlowByID(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputGetFlowParams{
		FlowID: "a91e45e9-d493-4618-8bfa-90ce756eed26",
	}

	f, e := api.GetFlowByID(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, p.FlowID, f.ID)

	t.Logf("successfully retrieve flow id %s", f.ID)
}
