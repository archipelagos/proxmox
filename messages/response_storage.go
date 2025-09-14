package messages

import "encoding/json"

type ResponseStorage struct {
	DatastoreAllocateSpace    int `json:"Datastore.AllocateSpace"`
	DatastoreAllocateTemplate int `json:"Datastore.AllocateTemplate"`
	DatastoreAllocate         int `json:"Datastore.Allocate"`
	PermissionsModify         int `json:"Permissions.Modify"`
	DatastoreAudit            int `json:"Datastore.Audit"`
}

func ParseResponseStorage(bytes []byte) (ResponseStorage, error) {
	var responseStorage ResponseStorage
	unmarshalErr := json.Unmarshal(bytes, &responseStorage)
	if unmarshalErr != nil {
		return ResponseStorage{}, unmarshalErr
	}

	return responseStorage, nil
}
