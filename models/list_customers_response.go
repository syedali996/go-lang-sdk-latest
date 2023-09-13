package models

import (
    "encoding/json"
)

// ListCustomersResponse represents a ListCustomersResponse struct.
// Response for listing the customers
type ListCustomersResponse struct {
    // The customer object
    Data   []GetCustomerResponse `json:"data"`
    // Paging object
    Paging *PagingResponse       `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListCustomersResponse.
// It customizes the JSON marshaling process for ListCustomersResponse objects.
func (l *ListCustomersResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListCustomersResponse object to a map representation for JSON marshaling.
func (l *ListCustomersResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListCustomersResponse.
// It customizes the JSON unmarshaling process for ListCustomersResponse objects.
func (l *ListCustomersResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetCustomerResponse `json:"data"`
        Paging *PagingResponse       `json:"paging"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Data = temp.Data
    l.Paging = temp.Paging
    return nil
}
