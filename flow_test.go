package dots

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFlow(t *testing.T) {
	api := New(clientID, apiKey, true)

	p := &InputCreateFlowParams{
		Steps: []string{"compliance"},
	}

	f, e := api.CreateFlowWithContext(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("successfully created flow %s", f.ID)
}

func TestCreateFlowWithContext(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateFlowParams{
		Steps: []string{"compliance"},
	}

	f, e := api.CreateFlowWithContext(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("successfully created flow %s", f.ID)
}

func TestGetFlowByID(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputGetFlowParams{
		FlowID: "a91e45e9-d493-4618-8bfa-90ce756eed26",
	}

	f, e := api.GetFlowByID(p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, p.FlowID, f.ID)

	t.Logf("successfully retrieve flow id %s", f.ID)
}

func TestGetFlowByIDWithContext(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputGetFlowParams{
		FlowID: "a91e45e9-d493-4618-8bfa-90ce756eed26",
	}

	f, e := api.GetFlowByIdWithContext(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, p.FlowID, f.ID)

	t.Logf("successfully retrieve flow id %s", f.ID)
}
