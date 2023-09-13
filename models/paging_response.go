package models

import (
    "encoding/json"
)

// PagingResponse represents a PagingResponse struct.
// Object used for returning lists of objects with pagination
type PagingResponse struct {
    // Total number of pages
    Total    *int    `json:"total"`
    // Previous page
    Previous *string `json:"previous"`
    // Next page
    Next     *string `json:"next"`
}

// MarshalJSON implements the json.Marshaler interface for PagingResponse.
// It customizes the JSON marshaling process for PagingResponse objects.
func (p *PagingResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PagingResponse object to a map representation for JSON marshaling.
func (p *PagingResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["total"] = p.Total
    structMap["previous"] = p.Previous
    structMap["next"] = p.Next
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PagingResponse.
// It customizes the JSON unmarshaling process for PagingResponse objects.
func (p *PagingResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Total    *int    `json:"total"`
        Previous *string `json:"previous"`
        Next     *string `json:"next"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Total = temp.Total
    p.Previous = temp.Previous
    p.Next = temp.Next
    return nil
}
