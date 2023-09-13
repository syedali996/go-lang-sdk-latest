package models

import (
    "encoding/json"
)

// ListDiscountsResponse represents a ListDiscountsResponse struct.
type ListDiscountsResponse struct {
    // The Discounts response
    Data   []GetDiscountResponse `json:"data"`
    // Paging object
    Paging *PagingResponse       `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListDiscountsResponse.
// It customizes the JSON marshaling process for ListDiscountsResponse objects.
func (l *ListDiscountsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListDiscountsResponse object to a map representation for JSON marshaling.
func (l *ListDiscountsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListDiscountsResponse.
// It customizes the JSON unmarshaling process for ListDiscountsResponse objects.
func (l *ListDiscountsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetDiscountResponse `json:"data"`
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
