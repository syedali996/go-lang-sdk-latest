package models

import (
    "encoding/json"
)

// ListPlansResponse represents a ListPlansResponse struct.
// Response object for listing plans
type ListPlansResponse struct {
    // The plan objects
    Data   []GetPlanResponse `json:"data"`
    // Paging object
    Paging *PagingResponse   `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListPlansResponse.
// It customizes the JSON marshaling process for ListPlansResponse objects.
func (l *ListPlansResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListPlansResponse object to a map representation for JSON marshaling.
func (l *ListPlansResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPlansResponse.
// It customizes the JSON unmarshaling process for ListPlansResponse objects.
func (l *ListPlansResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetPlanResponse `json:"data"`
        Paging *PagingResponse   `json:"paging"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Data = temp.Data
    l.Paging = temp.Paging
    return nil
}
