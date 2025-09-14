package messages

import "encoding/json"

type ResponseMapping struct {
	MappingUse        int `json:"Mapping.Use"`
	PermissionsModify int `json:"Permissions.Modify"`
	MappingAudit      int `json:"Mapping.Audit"`
	MappingModify     int `json:"Mapping.Modify"`
}

func ParseResponseMapping(bytes []byte) (ResponseMapping, error) {
	var responseMapping ResponseMapping
	unmarshalErr := json.Unmarshal(bytes, &responseMapping)
	if unmarshalErr != nil {
		return ResponseMapping{}, unmarshalErr
	}

	return responseMapping, nil
}
