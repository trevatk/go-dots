package dots

import (
	"context"
	"testing"
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

func TestGetFlowByID(t *testing.T) {}

func TestGetFlowByIDWithContext(t *testing.T) {}
