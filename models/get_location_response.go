package models

import (
    "encoding/json"
)

// GetLocationResponse represents a GetLocationResponse struct.
// Response object for geetting an order location request
type GetLocationResponse struct {
    // Latitude
    Latitude  *string `json:"latitude"`
    // Longitude
    Longitude *string `json:"longitude"`
}

// MarshalJSON implements the json.Marshaler interface for GetLocationResponse.
// It customizes the JSON marshaling process for GetLocationResponse objects.
func (g *GetLocationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetLocationResponse object to a map representation for JSON marshaling.
func (g *GetLocationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["latitude"] = g.Latitude
    structMap["longitude"] = g.Longitude
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetLocationResponse.
// It customizes the JSON unmarshaling process for GetLocationResponse objects.
func (g *GetLocationResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Latitude  *string `json:"latitude"`
        Longitude *string `json:"longitude"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Latitude = temp.Latitude
    g.Longitude = temp.Longitude
    return nil
}
