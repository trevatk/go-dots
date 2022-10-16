package dots

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateInvoice(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateInvoiceParams{
		Amount:    100,
		ExpiresIn: 3600,
	}

	r, e := api.CreateInvoice(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	fmt.Printf("successfully created invoice %s", r.Invoice.ID)

}
