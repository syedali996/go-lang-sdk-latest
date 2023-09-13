package models

import (
    "encoding/json"
)

// ListTransferResponse represents a ListTransferResponse struct.
// List of paginated transfer objects
type ListTransferResponse struct {
    // Transfers
    Data   []GetTransferResponse `json:"data"`
    // Paging
    Paging *PagingResponse       `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListTransferResponse.
// It customizes the JSON marshaling process for ListTransferResponse objects.
func (l *ListTransferResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListTransferResponse object to a map representation for JSON marshaling.
func (l *ListTransferResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListTransferResponse.
// It customizes the JSON unmarshaling process for ListTransferResponse objects.
func (l *ListTransferResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetTransferResponse `json:"data"`
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
