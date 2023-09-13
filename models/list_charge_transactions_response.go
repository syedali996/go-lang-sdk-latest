package models

import (
    "encoding/json"
)

// ListChargeTransactionsResponse represents a ListChargeTransactionsResponse struct.
// Response object for listing charge transactions
type ListChargeTransactionsResponse struct {
    // The charge transactions objects
    Data   []GetTransactionResponseInterface `json:"data"`
    // Paging object
    Paging *PagingResponse                   `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListChargeTransactionsResponse.
// It customizes the JSON marshaling process for ListChargeTransactionsResponse objects.
func (l *ListChargeTransactionsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListChargeTransactionsResponse object to a map representation for JSON marshaling.
func (l *ListChargeTransactionsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListChargeTransactionsResponse.
// It customizes the JSON unmarshaling process for ListChargeTransactionsResponse objects.
func (l *ListChargeTransactionsResponse) UnmarshalJSON(input []byte) error {
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
