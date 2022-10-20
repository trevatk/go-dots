package dots

import (
	"context"
	"os"
	"testing"
)

func TestCreatePayout(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputCreatePayoutParams{}

	pa, e := api.CreatePayout(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Log("payout created successfully ", pa.Success)
}

func TestCreatePayoutLink(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputCreatePayoutLinkParams{
		Delivery: Delivery{
			Method: "",
		},
	}

	pl, e := api.CreatePayoutLink(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("successfully created payout link %s", pl.PayoutLink.Link)
}

func TestSendPayout(t *testing.T) {

	// TODO:
	// finish response model

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputSendPayoutParams{}

	r, e := api.SendPayout(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("html response %s", string(r))
}

func TestCreateDirectPayout(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputDirectPayoutParams{}

	r, e := api.CreateDirectPayout(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Log("payout created ", r.Success)
}
