package messages

import (
	"encoding/json"
)

type ResponseAccess struct {
	GroupAllocate     int `json:"Group.Allocate"`
	UserModify        int `json:"User.Modify"`
	PermissionsModify int `json:"Permissions.Modify"`
}

func ParseResponseAccess(bytes []byte) (ResponseAccess, error) {
	var responseAccess ResponseAccess
	unmarshalErr := json.Unmarshal(bytes, &responseAccess)
	if unmarshalErr != nil {
		return ResponseAccess{}, unmarshalErr
	}

	return responseAccess, nil
}
