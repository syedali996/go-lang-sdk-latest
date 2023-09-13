package models

import (
    "encoding/json"
)

// ListAddressesResponse represents a ListAddressesResponse struct.
// Response object for listing addresses
type ListAddressesResponse struct {
    // The address objects
    Data   []GetAddressResponse `json:"data"`
    // Paging object
    Paging *PagingResponse      `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListAddressesResponse.
// It customizes the JSON marshaling process for ListAddressesResponse objects.
func (l *ListAddressesResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListAddressesResponse object to a map representation for JSON marshaling.
func (l *ListAddressesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListAddressesResponse.
// It customizes the JSON unmarshaling process for ListAddressesResponse objects.
func (l *ListAddressesResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetAddressResponse `json:"data"`
        Paging *PagingResponse      `json:"paging"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Data = temp.Data
    l.Paging = temp.Paging
    return nil
}
