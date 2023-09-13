package models

import (
    "encoding/json"
)

// GetAntifraudResponse represents a GetAntifraudResponse struct.
type GetAntifraudResponse struct {
    Status        *string `json:"status"`
    ReturnCode    *string `json:"return_code"`
    ReturnMessage *string `json:"return_message"`
    ProviderName  *string `json:"provider_name"`
    Score         *string `json:"score"`
}

// MarshalJSON implements the json.Marshaler interface for GetAntifraudResponse.
// It customizes the JSON marshaling process for GetAntifraudResponse objects.
func (g *GetAntifraudResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetAntifraudResponse object to a map representation for JSON marshaling.
func (g *GetAntifraudResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["status"] = g.Status
    structMap["return_code"] = g.ReturnCode
    structMap["return_message"] = g.ReturnMessage
    structMap["provider_name"] = g.ProviderName
    structMap["score"] = g.Score
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAntifraudResponse.
// It customizes the JSON unmarshaling process for GetAntifraudResponse objects.
func (g *GetAntifraudResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Status        *string `json:"status"`
        ReturnCode    *string `json:"return_code"`
        ReturnMessage *string `json:"return_message"`
        ProviderName  *string `json:"provider_name"`
        Score         *string `json:"score"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Status = temp.Status
    g.ReturnCode = temp.ReturnCode
    g.ReturnMessage = temp.ReturnMessage
    g.ProviderName = temp.ProviderName
    g.Score = temp.Score
    return nil
}
