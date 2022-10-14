package dots

import (
	"fmt"
	"testing"
)

func TestCreateInvoice(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateInvoiceParams{
		Amount:    100,
		ExpiresIn: 3600,
	}

	r, e := api.CreateInvoice(p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	fmt.Printf("successfully created invoice %s", r.Invoice.ID)

}

func TestCreateInvoiceWithContext(t *testing.T) {}
