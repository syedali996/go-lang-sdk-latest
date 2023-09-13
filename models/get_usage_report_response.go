package models

import (
    "encoding/json"
)

// GetUsageReportResponse represents a GetUsageReportResponse struct.
type GetUsageReportResponse struct {
    Url              *string `json:"url"`
    UsageReportUrl   *string `json:"usage_report_url"`
    GroupedReportUrl *string `json:"grouped_report_url"`
}

// MarshalJSON implements the json.Marshaler interface for GetUsageReportResponse.
// It customizes the JSON marshaling process for GetUsageReportResponse objects.
func (g *GetUsageReportResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetUsageReportResponse object to a map representation for JSON marshaling.
func (g *GetUsageReportResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["url"] = g.Url
    structMap["usage_report_url"] = g.UsageReportUrl
    structMap["grouped_report_url"] = g.GroupedReportUrl
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetUsageReportResponse.
// It customizes the JSON unmarshaling process for GetUsageReportResponse objects.
func (g *GetUsageReportResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Url              *string `json:"url"`
        UsageReportUrl   *string `json:"usage_report_url"`
        GroupedReportUrl *string `json:"grouped_report_url"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Url = temp.Url
    g.UsageReportUrl = temp.UsageReportUrl
    g.GroupedReportUrl = temp.GroupedReportUrl
    return nil
}
