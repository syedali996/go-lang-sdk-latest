package models

import (
    "encoding/json"
)

// ListChargesResponse represents a ListChargesResponse struct.
// Response object for listing charges
type ListChargesResponse struct {
    // The charge objects
    Data   []GetChargeResponse `json:"data"`
    // Paging object
    Paging *PagingResponse     `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListChargesResponse.
// It customizes the JSON marshaling process for ListChargesResponse objects.
func (l *ListChargesResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListChargesResponse object to a map representation for JSON marshaling.
func (l *ListChargesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListChargesResponse.
// It customizes the JSON unmarshaling process for ListChargesResponse objects.
func (l *ListChargesResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetChargeResponse `json:"data"`
        Paging *PagingResponse     `json:"paging"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Data = temp.Data
    l.Paging = temp.Paging
    return nil
}
