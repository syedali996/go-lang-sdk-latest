package models

import (
    "encoding/json"
)

// GetCardTokenResponse represents a GetCardTokenResponse struct.
// Card token data
type GetCardTokenResponse struct {
    LastFourDigits *string `json:"last_four_digits"`
    HolderName     *string `json:"holder_name"`
    HolderDocument *string `json:"holder_document"`
    ExpMonth       *string `json:"exp_month"`
    ExpYear        *string `json:"exp_year"`
    Brand          *string `json:"brand"`
    Type           *string `json:"type"`
    Label          *string `json:"label"`
}

// MarshalJSON implements the json.Marshaler interface for GetCardTokenResponse.
// It customizes the JSON marshaling process for GetCardTokenResponse objects.
func (g *GetCardTokenResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetCardTokenResponse object to a map representation for JSON marshaling.
func (g *GetCardTokenResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["last_four_digits"] = g.LastFourDigits
    structMap["holder_name"] = g.HolderName
    structMap["holder_document"] = g.HolderDocument
    structMap["exp_month"] = g.ExpMonth
    structMap["exp_year"] = g.ExpYear
    structMap["brand"] = g.Brand
    structMap["type"] = g.Type
    structMap["label"] = g.Label
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCardTokenResponse.
// It customizes the JSON unmarshaling process for GetCardTokenResponse objects.
func (g *GetCardTokenResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        LastFourDigits *string `json:"last_four_digits"`
        HolderName     *string `json:"holder_name"`
        HolderDocument *string `json:"holder_document"`
        ExpMonth       *string `json:"exp_month"`
        ExpYear        *string `json:"exp_year"`
        Brand          *string `json:"brand"`
        Type           *string `json:"type"`
        Label          *string `json:"label"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.LastFourDigits = temp.LastFourDigits
    g.HolderName = temp.HolderName
    g.HolderDocument = temp.HolderDocument
    g.ExpMonth = temp.ExpMonth
    g.ExpYear = temp.ExpYear
    g.Brand = temp.Brand
    g.Type = temp.Type
    g.Label = temp.Label
    return nil
}
