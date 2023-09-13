package models

import (
    "encoding/json"
)

// PixAdditionalInformation represents a PixAdditionalInformation struct.
// Pix Additional Information
type PixAdditionalInformation struct {
    Name  string `json:"Name"`
    Value string `json:"Value"`
}

// MarshalJSON implements the json.Marshaler interface for PixAdditionalInformation.
// It customizes the JSON marshaling process for PixAdditionalInformation objects.
func (p *PixAdditionalInformation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PixAdditionalInformation object to a map representation for JSON marshaling.
func (p *PixAdditionalInformation) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["Name"] = p.Name
    structMap["Value"] = p.Value
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PixAdditionalInformation.
// It customizes the JSON unmarshaling process for PixAdditionalInformation objects.
func (p *PixAdditionalInformation) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name  string `json:"Name"`
        Value string `json:"Value"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Name = temp.Name
    p.Value = temp.Value
    return nil
}
