package dots

import (
	"context"
	"encoding/json"
	"fmt"
)

// InputCreateInvoiceParams
type InputCreateInvoiceParams struct {
	Amount               int        `json:"amount"`
	ExpiresIn            int        `json:"expires_int"`
	Items                []*Item    `json:"items"`
	Breakdown            *Breakdown `json:"breakdown"`
	RequestedInformation []string   `json:"requested_information"`
	Metadata             []string   `json:"metadata"`
}

// Invoice
type Invoice struct {
	ID                   string              `json:"id"`                    // string <uuid> UUID id of the invoice
	Payer                *Payer              `json:"payer"`                 // object
	APIApp               *APIApp             `json:"api_app"`               // object
	Amount               float64             `json:"amount"`                // float total amount to charge the customer (sum of all breakdown fields)
	Expiry               string              `json:"expiry"`                // string (date-time)
	Status               []InvoiceStatusEnum `json:"status"`                // string enm 'created', 'completed', 'expired'
	Link                 string              `json:"link"`                  // string (uri)
	Items                []*Item             `json:"items"`                 // list object
	Breakdown            *Breakdown          `json:"breakdown"`             // object
	RequestedInformation []string            `json:"requested_information"` // information requested from the user that will populate on completion of transaction, (shipping_address)
	Metadata             []string            `json:"metadata"`
}

// Payer
type Payer struct {
	FirstName string  `json:"first_name"` // string
	LastName  string  `json:"last_name"`  // string
	Username  string  `json:"username"`   // string
	Wallet    *Wallet `json:"wallet"`     // object
}

// APIApp
type APIApp struct {
	ID     string `json:"id"`      // string <uuid>
	Name   string `json:"name"`    // string
	UserID string `json:"user_id"` // string <uuid>
}

// Item
type Item struct {
	Name        string `json:"name"`        // name of the item
	UnitAmount  int    `json:"unit_amount"` // cost of 1 unit of the item
	Quantity    int    `json:"quantity"`    // quantity of this item
	Description string `json:"description"` // description of the item
}

// Breakdown
type Breakdown struct {
	ItemsTotal int `json:"intems_total"` // total of the items (unit_amount * quantity)
	Shipping   int `json:"shipping"`     // cost of shipping
	Tax        int `json:"tax"`          // tax
}

// CreateInvoiceResponse
type CreateInvoiceResponse struct {
	Success bool     `json:"success"`
	Invoice *Invoice `json:"invoice"`
}

// GetInvoiceResponse
type GetInvoiceResponse struct {
	Success bool     `json:"success"`
	Invoice *Invoice `json:"invoice"`
}

// CreateInvoice
func (api *API) CreateInvoice(ctx context.Context, in *InputCreateInvoiceParams) (*CreateInvoiceResponse, error) {

	r := api.h + "/api/invoice/create"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var cir CreateInvoiceResponse
	if e := json.Unmarshal(b, &cir); e != nil {
		return nil, fmt.Errorf("dots api create invoice json.Unmarshal err %v", e)
	}

	return &cir, nil
}

// GetInvoiceByID
func (api *API) GetInvoiceByID(ctx context.Context, ID string) (*GetInvoiceResponse, error) {

	r := api.h + "/api/invoice/get/" + ID
	b, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var i GetInvoiceResponse
	if e := json.Unmarshal(b, &i); e != nil {
		return nil, fmt.Errorf("dots create get invoice by id json.Unmarshal err %v", e)
	}

	return &i, nil
}
