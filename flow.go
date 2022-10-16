package dots

import (
	"context"
	"encoding/json"
	"fmt"
)

// InputCreateFlowParams
type InputCreateFlowParams struct {
	Steps  []string `json:"steps"`             // array of strings enum: 'compliance', 'manage-payements', 'manage-payouts', 'payout'
	UserID *string  `json:"user_id,omitempty"` // string <uuid>
}

// InputGetFlowParams
type InputGetFlowParams struct {
	FlowID string `json:"flow_id"` // string <uuid>
}

// Flow
type Flow struct {
	ID             string   `json:"id"`              // string <uuid>
	Created        string   `json:"created"`         // string <date-time>
	Updated        string   `json:"updated"`         // string <date-time>
	UserID         string   `json:"user_id"`         // string <uuid>
	Steps          []string `json:"steps"`           // array of strings enum: 'compliance', 'manage-payments', 'manage-payouts', 'payout'
	CompletedSteps []string `json:"completed_steps"` // array of strings enum: 'compliance', 'manage-payments', 'manage-payouts', 'payout'
	PayoutLinkID   string   `json:"payout_link_id"`  // string <uuid>
	Link           string   `json:"link"`            // string <url>
}

// FlowResponse
type FlowResponse struct {
	Success bool  `json:"success"`
	Flow    *Flow `json:"flow"`
}

// CreateFlowWithContext
func (api *API) CreateFlowWithContext(ctx context.Context, in *InputCreateFlowParams) (*Flow, error) {

	r := api.h + "/api/flow/create"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var fr FlowResponse
	if err := json.Unmarshal(b, &fr); err != nil {
		return nil, fmt.Errorf("dots api create flow json.Unmarshal err %v", err)
	}

	return fr.Flow, nil
}

// GetFlowByID
func (api *API) GetFlowByID(ctx context.Context, in *InputGetFlowParams) (*Flow, error) {

	r := api.h + "/api/flow/get/" + in.FlowID
	b, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var fr FlowResponse
	if err := json.Unmarshal(b, &fr); err != nil {
		return nil, fmt.Errorf("dots api create flow json.Unmarshal err %v", err)
	}

	return fr.Flow, nil
}
