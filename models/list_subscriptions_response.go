package models

import (
    "encoding/json"
)

// ListSubscriptionsResponse represents a ListSubscriptionsResponse struct.
// Response object for listing subscriptions
type ListSubscriptionsResponse struct {
    // The subscription objects
    Data   []GetSubscriptionResponse `json:"data"`
    // Paging object
    Paging *PagingResponse           `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionsResponse.
// It customizes the JSON marshaling process for ListSubscriptionsResponse objects.
func (l *ListSubscriptionsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionsResponse object to a map representation for JSON marshaling.
func (l *ListSubscriptionsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["data"] = l.Data
    structMap["paging"] = l.Paging
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionsResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionsResponse objects.
func (l *ListSubscriptionsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Data   []GetSubscriptionResponse `json:"data"`
        Paging *PagingResponse           `json:"paging"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Data = temp.Data
    l.Paging = temp.Paging
    return nil
}
