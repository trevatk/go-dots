package dots

import (
	"context"
	"encoding/json"
	"fmt"
)

// InputCreatePayoutParams create new payout input
type InputCreatePayoutParams struct {
	UserID        string `json:"user_id"`
	Amount        int    `json:"amount"`
	PayoutMethod  string `json:"payout_method"`
	ACHAccountID  string `json:"ach_account_id"`
	IntlAccountID string `json:"intl_account_id"`
	Fund          bool   `json:"fund"`
}

// InputCreatePayoutLinkParams create payout link input
type InputCreatePayoutLinkParams struct {
	Delivery  Delivery `json:"delivery"`
	Amount    int      `json:"amount"`
	Notes     string   `json:"notes"`
	TaxExempt bool     `json:"tax_exempt"`
	Payee     *Payee   `json:"payee,omitempty"`
}

// InputSendPayoutParams send payout input
type InputSendPayoutParams struct {
	Amount                            int       `json:"amount"`
	UserID                            string    `json:"user_id"`
	Payee                             *Payee    `json:"payee"`
	Delivery                          *Delivery `json:"delivery"`
	Notes                             string    `json:"notes"`                                // custom data that will be attached to the transaction when the recipient claims the link
	ForceCollectComplianceInformation bool      `json:"force_collect_compliance_information"` // Require the recipient to fill out compliance information (i.e. form 1099) when below the payout limit.
}

// InputDirectPayoutParams create direct payout input
type InputDirectPayoutParams struct {
	PayoutType     string `json:"payout_type"`     // enum 'paypal', 'venmo', 'visa', 'amazon'
	DeliveryMethod string `json:"delivery_method"` // phone number for venmo and email address for all other types
	International  bool   `json:"international"`   // set to true if international
}

// CreatePayoutResponse create payout response
type CreatePayoutResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

// CreatePayoutLinkResponse create payout link response
type CreatePayoutLinkResponse struct {
	Success    bool        `json:"success"`
	PayoutLink *PayoutLink `json:"payout_link"`
}

// DirectPayoutResponse create direct payout response object
type DirectPayoutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// CreatePayout create new payout
func (api *API) CreatePayout(ctx context.Context, in *InputCreatePayoutParams) (*CreatePayoutResponse, error) {

	r := api.h + "/api/payouts/create"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var pr CreatePayoutResponse
	if e := json.Unmarshal(b, &pr); e != nil {
		return nil, fmt.Errorf("dots api create payout json.Unmarshal err %v", e)
	}

	return &pr, nil
}

// CreatePayoutLink create new payout link
func (api *API) CreatePayoutLink(ctx context.Context, in *InputCreatePayoutLinkParams) (*CreatePayoutLinkResponse, error) {

	r := api.h + "/api/payouts/create_payout_link"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var pl CreatePayoutLinkResponse
	if e := json.Unmarshal(b, &pl); e != nil {
		return nil, fmt.Errorf("dots api create payout link json.Unmarshal err %v", e)
	}

	return &pl, nil
}

// SendPayout send a payout
func (api *API) SendPayout(ctx context.Context, in *InputSendPayoutParams) ([]byte, error) {

	r := api.h + "/api/v2/payouts/send_payout"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return []byte{}, e
	}

	return b, nil
}

// CreateDirectPayout creata direct payout
func (api *API) CreateDirectPayout(ctx context.Context, in *InputDirectPayoutParams) (*DirectPayoutResponse, error) {

	r := api.h + "/api/payouts/create_direct_payout"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var p DirectPayoutResponse
	if e := json.Unmarshal(b, &p); e != nil {
		return nil, fmt.Errorf("dots api create direct payout json.Unmarshal err %v html response %s", e, string(b))
	}

	return &p, nil
}
