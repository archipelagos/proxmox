package messages

import "encoding/json"

type ResponseVMs struct {
	VMConfigCDROM              int `json:"VM.Config.CDROM"`
	PermissionsModify          int `json:"Permissions.Modify"`
	VMConfigCPU                int `json:"VM.Config.CPU"`
	VMConfigNetwork            int `json:"VM.Config.Network"`
	VMGuestAgentAudit          int `json:"VM.GuestAgent.Audit"`
	VMClone                    int `json:"VM.Clone"`
	VMConfigDisk               int `json:"VM.Config.Disk"`
	VMConfigOptions            int `json:"VM.Config.Options"`
	VMAllocate                 int `json:"VM.Allocate"`
	VMSnapshot                 int `json:"VM.Snapshot"`
	VMReplicate                int `json:"VM.Replicate"`
	VMConfigHWType             int `json:"VM.Config.HWType"`
	VMAudit                    int `json:"VM.Audit"`
	VMConsole                  int `json:"VM.Console"`
	VMBackup                   int `json:"VM.Backup"`
	VMSnapshotRollback         int `json:"VM.Snapshot.Rollback"`
	VMGuestAgentFileWrite      int `json:"VM.GuestAgent.FileWrite"`
	VMMigrate                  int `json:"VM.Migrate"`
	VMPowerMgmt                int `json:"VM.PowerMgmt"`
	VMGuestAgentFileRead       int `json:"VM.GuestAgent.FileRead"`
	VMConfigMemory             int `json:"VM.Config.Memory"`
	VMConfigCloudinit          int `json:"VM.Config.Cloudinit"`
	VMGuestAgentUnrestricted   int `json:"VM.GuestAgent.Unrestricted"`
	VMGuestAgentFileSystemMgmt int `json:"VM.GuestAgent.FileSystemMgmt"`
}

func ParseResponseVMs(bytes []byte) (ResponseVMs, error) {
	var responseVMs ResponseVMs
	unmarshalErr := json.Unmarshal(bytes, &responseVMs)
	if unmarshalErr != nil {
		return ResponseVMs{}, unmarshalErr
	}

	return responseVMs, nil
}
