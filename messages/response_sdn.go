package messages

import "encoding/json"

type ResponseSDN struct {
	SDNAllocate       int `json:"SDN.Allocate"`
	PermissionsModify int `json:"Permissions.Modify"`
	SDNAudit          int `json:"SDN.Audit"`
	SDNUse            int `json:"SDN.Use"`
}

func ParseResponseSDN(bytes []byte) (ResponseSDN, error) {
	var responseSDN ResponseSDN
	unmarshalErr := json.Unmarshal(bytes, &responseSDN)
	if unmarshalErr != nil {
		return ResponseSDN{}, unmarshalErr
	}

	return responseSDN, nil
}
