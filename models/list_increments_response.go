package models

import (
    "encoding/json"
)

// ListIncrementsResponse represents a ListIncrementsResponse struct.
type ListIncrementsResponse struct {
    // The Increments response
    Data   []GetIncrementResponse `json:"data"`
    // Paging object
    Paging *PagingResponse        `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListIncrementsResponse.
// It customizes the JSON marshaling process for ListIncrementsResponse objects.
func (l *ListIncrementsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListIncrementsResponse object to a map representation for JSON marshaling.
func (l *ListIncrementsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListIncrementsResponse.
// It customizes the JSON unmarshaling process for ListIncrementsResponse objects.
func (l *ListIncrementsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetIncrementResponse `json:"data"`
        Paging *PagingResponse        `json:"paging"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Data = temp.Data
    l.Paging = temp.Paging
    return nil
}
