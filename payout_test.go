package dots

import (
	"context"
	"testing"
)

func TestCreatePayout(t *testing.T) {

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

	api := New(clientID, apiKey, true)

	p := &InputSendPayoutParams{}

	r, e := api.SendPayout(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("html response %s", string(r))
}
