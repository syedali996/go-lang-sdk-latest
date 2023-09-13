package models

import (
    "encoding/json"
)

// GetSplitResponse represents a GetSplitResponse struct.
// Split response
type GetSplitResponse struct {
    // Type
    Type      *string                           `json:"type"`
    // Amount
    Amount    *int                              `json:"amount"`
    // Recipient
    Recipient Optional[GetRecipientResponse]    `json:"recipient"`
    // The split rule gateway id
    GatewayId *string                           `json:"gateway_id"`
    Options   Optional[GetSplitOptionsResponse] `json:"options"`
    Id        *string                           `json:"id"`
}

// MarshalJSON implements the json.Marshaler interface for GetSplitResponse.
// It customizes the JSON marshaling process for GetSplitResponse objects.
func (g *GetSplitResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetSplitResponse object to a map representation for JSON marshaling.
func (g *GetSplitResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["type"] = g.Type
    structMap["amount"] = g.Amount
    if g.Recipient.IsValueSet() {
        structMap["recipient"] = g.Recipient.Value()
    }
    structMap["gateway_id"] = g.GatewayId
    if g.Options.IsValueSet() {
        structMap["options"] = g.Options.Value()
    }
    structMap["id"] = g.Id
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSplitResponse.
// It customizes the JSON unmarshaling process for GetSplitResponse objects.
func (g *GetSplitResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Type      *string                           `json:"type"`
        Amount    *int                              `json:"amount"`
        Recipient Optional[GetRecipientResponse]    `json:"recipient"`
        GatewayId *string                           `json:"gateway_id"`
        Options   Optional[GetSplitOptionsResponse] `json:"options"`
        Id        *string                           `json:"id"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Type = temp.Type
    g.Amount = temp.Amount
    g.Recipient = temp.Recipient
    g.GatewayId = temp.GatewayId
    g.Options = temp.Options
    g.Id = temp.Id
    return nil
}
