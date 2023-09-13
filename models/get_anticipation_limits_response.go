package models

import (
    "encoding/json"
)

// GetAnticipationLimitsResponse represents a GetAnticipationLimitsResponse struct.
// Anticipation limits
type GetAnticipationLimitsResponse struct {
    // Max limit
    Max *GetAnticipationLimitResponse `json:"max"`
    // Min limit
    Min *GetAnticipationLimitResponse `json:"min"`
}

// MarshalJSON implements the json.Marshaler interface for GetAnticipationLimitsResponse.
// It customizes the JSON marshaling process for GetAnticipationLimitsResponse objects.
func (g *GetAnticipationLimitsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetAnticipationLimitsResponse object to a map representation for JSON marshaling.
func (g *GetAnticipationLimitsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["max"] = g.Max
    structMap["min"] = g.Min
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAnticipationLimitsResponse.
// It customizes the JSON unmarshaling process for GetAnticipationLimitsResponse objects.
func (g *GetAnticipationLimitsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Max *GetAnticipationLimitResponse `json:"max"`
        Min *GetAnticipationLimitResponse `json:"min"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Max = temp.Max
    g.Min = temp.Min
    return nil
}
