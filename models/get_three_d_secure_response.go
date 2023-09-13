package models

import (
    "encoding/json"
)

// GetThreeDSecureResponse represents a GetThreeDSecureResponse struct.
// 3D-S payment authentication response
type GetThreeDSecureResponse struct {
    // MPI Vendor
    Mpi           *string `json:"mpi"`
    // Electronic Commerce Indicator (ECI) (Opcional)
    Eci           *string `json:"eci"`
    // Online payment cryptogram, definido pelo 3-D Secure.
    Cavv          *string `json:"cavv"`
    // Identificador da transação (XID)
    TransactionId *string `json:"transaction_Id"`
    // Url de redirecionamento de sucessso
    SuccessUrl    *string `json:"success_url"`
}

// MarshalJSON implements the json.Marshaler interface for GetThreeDSecureResponse.
// It customizes the JSON marshaling process for GetThreeDSecureResponse objects.
func (g *GetThreeDSecureResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetThreeDSecureResponse object to a map representation for JSON marshaling.
func (g *GetThreeDSecureResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["mpi"] = g.Mpi
    structMap["eci"] = g.Eci
    structMap["cavv"] = g.Cavv
    structMap["transaction_Id"] = g.TransactionId
    structMap["success_url"] = g.SuccessUrl
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetThreeDSecureResponse.
// It customizes the JSON unmarshaling process for GetThreeDSecureResponse objects.
func (g *GetThreeDSecureResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Mpi           *string `json:"mpi"`
        Eci           *string `json:"eci"`
        Cavv          *string `json:"cavv"`
        TransactionId *string `json:"transaction_Id"`
        SuccessUrl    *string `json:"success_url"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Mpi = temp.Mpi
    g.Eci = temp.Eci
    g.Cavv = temp.Cavv
    g.TransactionId = temp.TransactionId
    g.SuccessUrl = temp.SuccessUrl
    return nil
}
