package messages

import "encoding/json"

type ResponseDC struct {
	SysAudit    int `json:"Sys.Audit"`
	SDNUse      int `json:"SDN.Use"`
	SDNAudit    int `json:"SDN.Audit"`
	SysModify   int `json:"Sys.Modify"`
	SDNAllocate int `json:"SDN.Allocate"`
}

func ParseResponseDC(bytes []byte) (ResponseDC, error) {
	var responseDC ResponseDC
	unmarshalErr := json.Unmarshal(bytes, &responseDC)
	if unmarshalErr != nil {
		return ResponseDC{}, unmarshalErr
	}

	return responseDC, nil
}
