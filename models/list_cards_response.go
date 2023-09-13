package models

import (
    "encoding/json"
)

// ListCardsResponse represents a ListCardsResponse struct.
// Response object for listing cards
type ListCardsResponse struct {
    // The card objects
    Data   []GetCardResponse `json:"data"`
    // Paging object
    Paging *PagingResponse   `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListCardsResponse.
// It customizes the JSON marshaling process for ListCardsResponse objects.
func (l *ListCardsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListCardsResponse object to a map representation for JSON marshaling.
func (l *ListCardsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListCardsResponse.
// It customizes the JSON unmarshaling process for ListCardsResponse objects.
func (l *ListCardsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetCardResponse `json:"data"`
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
