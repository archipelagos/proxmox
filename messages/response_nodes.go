package messages

import "encoding/json"

type ResponseNodes struct {
	SysConsole        int `json:"Sys.Console"`
	SysPowerMgmt      int `json:"Sys.PowerMgmt"`
	SysAudit          int `json:"Sys.Audit"`
	SysSyslog         int `json:"Sys.Syslog"`
	PermissionsModify int `json:"Permissions.Modify"`
	SysIncoming       int `json:"Sys.Incoming"`
	SysModify         int `json:"Sys.Modify"`
	SysAccessNetwork  int `json:"Sys.AccessNetwork"`
}

func ParseResponseNodes(bytes []byte) (ResponseNodes, error) {
	var responseNodes ResponseNodes
	unmarshalErr := json.Unmarshal(bytes, &responseNodes)
	if unmarshalErr != nil {
		return ResponseNodes{}, unmarshalErr
	}

	return responseNodes, nil
}
