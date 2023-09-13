package models

import (
    "encoding/json"
)

// GetAutomaticAnticipationResponse represents a GetAutomaticAnticipationResponse struct.
type GetAutomaticAnticipationResponse struct {
    Enabled          *bool   `json:"enabled"`
    Type             *string `json:"type"`
    VolumePercentage *int    `json:"volume_percentage"`
    Delay            *int    `json:"delay"`
    Days             []int   `json:"days"`
}

// MarshalJSON implements the json.Marshaler interface for GetAutomaticAnticipationResponse.
// It customizes the JSON marshaling process for GetAutomaticAnticipationResponse objects.
func (g *GetAutomaticAnticipationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetAutomaticAnticipationResponse object to a map representation for JSON marshaling.
func (g *GetAutomaticAnticipationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["enabled"] = g.Enabled
    structMap["type"] = g.Type
    structMap["volume_percentage"] = g.VolumePercentage
    structMap["delay"] = g.Delay
    structMap["days"] = g.Days
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAutomaticAnticipationResponse.
// It customizes the JSON unmarshaling process for GetAutomaticAnticipationResponse objects.
func (g *GetAutomaticAnticipationResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Enabled          *bool   `json:"enabled"`
        Type             *string `json:"type"`
        VolumePercentage *int    `json:"volume_percentage"`
        Delay            *int    `json:"delay"`
        Days             []int   `json:"days"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Enabled = temp.Enabled
    g.Type = temp.Type
    g.VolumePercentage = temp.VolumePercentage
    g.Delay = temp.Delay
    g.Days = temp.Days
    return nil
}
