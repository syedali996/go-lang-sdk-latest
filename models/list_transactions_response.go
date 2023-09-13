package models

import (
    "encoding/json"
)

// ListTransactionsResponse represents a ListTransactionsResponse struct.
// Response object for listing transactions
type ListTransactionsResponse struct {
    // The transaction objects
    Data   []GetTransactionResponseInterface `json:"data"`
    // Paging object
    Paging *PagingResponse                   `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListTransactionsResponse.
// It customizes the JSON marshaling process for ListTransactionsResponse objects.
func (l *ListTransactionsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListTransactionsResponse object to a map representation for JSON marshaling.
func (l *ListTransactionsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListTransactionsResponse.
// It customizes the JSON unmarshaling process for ListTransactionsResponse objects.
func (l *ListTransactionsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetTransactionResponseInterface `json:"data"`
        Paging *PagingResponse                   `json:"paging"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Data = temp.Data
    l.Paging = temp.Paging
    return nil
}
