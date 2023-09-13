package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetTokenResponse represents a GetTokenResponse struct.
// Token data
type GetTokenResponse struct {
    Id        *string               `json:"id"`
    Type      *string               `json:"type"`
    CreatedAt *time.Time            `json:"created_at"`
    ExpiresAt *string               `json:"expires_at"`
    Card      *GetCardTokenResponse `json:"card"`
}

// MarshalJSON implements the json.Marshaler interface for GetTokenResponse.
// It customizes the JSON marshaling process for GetTokenResponse objects.
func (g *GetTokenResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetTokenResponse object to a map representation for JSON marshaling.
func (g *GetTokenResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["id"] = g.Id
    structMap["type"] = g.Type
    structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
    structMap["expires_at"] = g.ExpiresAt
    structMap["card"] = g.Card
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTokenResponse.
// It customizes the JSON unmarshaling process for GetTokenResponse objects.
func (g *GetTokenResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id        *string               `json:"id"`
        Type      *string               `json:"type"`
        CreatedAt *string               `json:"created_at"`
        ExpiresAt *string               `json:"expires_at"`
        Card      *GetCardTokenResponse `json:"card"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Type = temp.Type
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    }
    g.ExpiresAt = temp.ExpiresAt
    g.Card = temp.Card
    return nil
}
