package dots

import (
	"context"
	"testing"
)

func TestCreatePayout(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreatePayoutParams{}

	pa, e := api.CreatePayout(p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Log("payout created successfully ", pa.Success)
}

func TestCreatePayoutWithContext(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreatePayoutParams{}

	pa, e := api.CreatePayoutWithContext(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Log("payout created successfully ", pa.Success)
}

func TestCreatePayoutLink(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreatePayoutLinkParams{}

	pl, e := api.CreatePayoutLink(p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("successfully created payout link %s", pl.PayoutLink.Link)
}

func TestCreatePayoutLinkWithContext(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreatePayoutLinkParams{}

	pl, e := api.CreatePayoutLinkWithContext(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("successfully created payout link %s", pl.PayoutLink.Link)
}
