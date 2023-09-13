package models

import (
    "encoding/json"
    "log"
    "time"
)

// GetTransactionReportFileResponse represents a GetTransactionReportFileResponse struct.
type GetTransactionReportFileResponse struct {
    Name *string    `json:"name"`
    Date *time.Time `json:"date"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransactionReportFileResponse.
// It customizes the JSON marshaling process for GetTransactionReportFileResponse objects.
func (g *GetTransactionReportFileResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetTransactionReportFileResponse object to a map representation for JSON marshaling.
func (g *GetTransactionReportFileResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["name"] = g.Name
    structMap["date"] = g.Date.Format(time.RFC3339)
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransactionReportFileResponse.
// It customizes the JSON unmarshaling process for GetTransactionReportFileResponse objects.
func (g *GetTransactionReportFileResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name *string `json:"name"`
        Date *string `json:"date"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Name = temp.Name
    if temp.Date != nil {
        DateVal, err := time.Parse(time.RFC3339, *temp.Date)
        if err != nil {
            log.Fatalf("Cannot Parse date as % s format.", time.RFC3339)
        }
        g.Date = &DateVal
    }
    return nil
}
