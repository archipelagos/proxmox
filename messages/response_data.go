package messages

import (
	"encoding/json"
	"fmt"
)

type ResponseData struct {
	Ticket              string           `json:"ticket"`
	Cap                 map[string](any) `json:"cap"`
	Username            string           `json:"username"`
	CSRFPreventionToken string           `json:"CSRFPreventionToken"`
}

// TODO: Implement via Unmarshaller interface.
func ParseResponseData(bytes []byte) (ResponseData, error) {
	var responseData ResponseData
	unmarshalErr := json.Unmarshal(bytes, &responseData)
	if unmarshalErr != nil {
		return ResponseData{}, unmarshalErr
	}

	for key, value := range responseData.Cap {
		if responseData.Cap[key] == nil {
			return ResponseData{}, fmt.Errorf("Structure different than expected, no: '%v' was found\n", key)
		}

		switch key {
		case "storage":
			responseStorageMap, ok := value.(map[string]any)
			var responseStorage ResponseStorage
			if ok {
				datastoreAllocateSpace, datastoreAllocateSpaceOk := responseStorageMap["Datastore.AllocateSpace"].(float64)
				if datastoreAllocateSpaceOk {
					responseStorage.DatastoreAllocateSpace = int(datastoreAllocateSpace)
				}

				datastoreAllocateTemplate, datastoreAllocateTemplateOk := responseStorageMap["Datastore.AllocateTemplate"].(float64)
				if datastoreAllocateTemplateOk {
					responseStorage.DatastoreAllocateTemplate = int(datastoreAllocateTemplate)
				}

				datastoreAllocate, datastoreAllocateOk := responseStorageMap["Datastore.Allocate"].(float64)
				if datastoreAllocateOk {
					responseStorage.DatastoreAllocate = int(datastoreAllocate)
				}

				permissionsModify, permissionsModifyOk := responseStorageMap["Permissions.Modify"].(float64)
				if permissionsModifyOk {
					responseStorage.PermissionsModify = int(permissionsModify)
				}

				datastoreAudit, datastoreAuditOk := responseStorageMap["Datastore.Audit"].(float64)
				if datastoreAuditOk {
					responseStorage.DatastoreAudit = int(datastoreAudit)
				}
			}
			responseData.Cap[key] = responseStorage

		case "mapping":
			responseMappingMap, ok := value.(map[string]any)
			var responseMapping ResponseMapping
			if ok {
				mappingUse, mappingUseOk := responseMappingMap["Mapping.Use"].(float64)
				if mappingUseOk {
					responseMapping.MappingUse = int(mappingUse)
				}

				permissionsModify, permissionsModifyOk := responseMappingMap["Permissions.Modify"].(float64)
				if permissionsModifyOk {
					responseMapping.PermissionsModify = int(permissionsModify)
				}

				mappingAudit, mappingAuditOk := responseMappingMap["Mapping.Audit"].(float64)
				if mappingAuditOk {
					responseMapping.MappingAudit = int(mappingAudit)
				}

				mappingModify, mappingModifyOk := responseMappingMap["Mapping.Modify"].(float64)
				if mappingModifyOk {
					responseMapping.MappingModify = int(mappingModify)
				}
			}

			responseData.Cap[key] = responseMapping

		case "sdn":
			responseSDNMap, ok := value.(map[string]any)
			var responseSDN ResponseSDN
			if ok {
				SDNAllocate, SDNAllocateOk := responseSDNMap["SDN.Allocate"].(float64)
				if SDNAllocateOk {
					responseSDN.SDNAllocate = int(SDNAllocate)
				}

				permissionsModify, permissionsModifyOk := responseSDNMap["Permissions.Modify"].(float64)
				if permissionsModifyOk {
					responseSDN.PermissionsModify = int(permissionsModify)
				}

				SDNAudit, SDNAuditOk := responseSDNMap["SDN.Audit"].(float64)
				if SDNAuditOk {
					responseSDN.SDNAudit = int(SDNAudit)
				}

				SDNUse, SDNUseOk := responseSDNMap["SDN.Use"].(float64)
				if SDNUseOk {
					responseSDN.SDNUse = int(SDNUse)
				}
			}

			responseData.Cap[key] = responseSDN

		case "vms":
			responseVMsMap, ok := value.(map[string]any)
			var responseVMs ResponseVMs
			if ok {
				VMConfigCDROM, VMConfigCDROMOk := responseVMsMap["VM.Config.CDROM"].(float64)
				if VMConfigCDROMOk {
					responseVMs.VMConfigCDROM = int(VMConfigCDROM)
				}

				PermissionsModify, PermissionsModifyOk := responseVMsMap["Permissions.Modify"].(float64)
				if PermissionsModifyOk {
					responseVMs.PermissionsModify = int(PermissionsModify)
				}

				VMConfigCPU, VMConfigCPUOk := responseVMsMap["VM.Config.CPU"].(float64)
				if VMConfigCPUOk {
					responseVMs.VMConfigCPU = int(VMConfigCPU)
				}

				VMConfigNetwork, VMConfigNetworkOk := responseVMsMap["VM.Config.Network"].(float64)
				if VMConfigNetworkOk {
					responseVMs.VMConfigNetwork = int(VMConfigNetwork)
				}

				VMGuestAgentAudit, VMGuestAgentAuditOk := responseVMsMap["VM.GuestAgent.Audit"].(float64)
				if VMGuestAgentAuditOk {
					responseVMs.VMGuestAgentAudit = int(VMGuestAgentAudit)
				}

				VMClone, VMCloneOk := responseVMsMap["VM.Clone"].(float64)
				if VMCloneOk {
					responseVMs.VMClone = int(VMClone)
				}

				VMConfigDisk, VMConfigDiskOk := responseVMsMap["VM.Config.Disk"].(float64)
				if VMConfigDiskOk {
					responseVMs.VMConfigDisk = int(VMConfigDisk)
				}

				VMConfigOptions, VMConfigOptionsOk := responseVMsMap["VM.Config.Options"].(float64)
				if VMConfigOptionsOk {
					responseVMs.VMConfigOptions = int(VMConfigOptions)
				}

				VMAllocate, VMAllocateOk := responseVMsMap["VM.Allocate"].(float64)
				if VMAllocateOk {
					responseVMs.VMAllocate = int(VMAllocate)
				}

				VMSnapshot, VMSnapshotOk := responseVMsMap["VM.Snapshot"].(float64)
				if VMSnapshotOk {
					responseVMs.VMSnapshot = int(VMSnapshot)
				}

				VMReplicate, VMReplicateOk := responseVMsMap["VM.Replicate"].(float64)
				if VMReplicateOk {
					responseVMs.VMReplicate = int(VMReplicate)
				}

				VMConfigHWType, VMConfigHWTypeOk := responseVMsMap["VM.Config.HWType"].(float64)
				if VMConfigHWTypeOk {
					responseVMs.VMConfigHWType = int(VMConfigHWType)
				}

				VMAudit, VMAuditOk := responseVMsMap["VM.Audit"].(float64)
				if VMAuditOk {
					responseVMs.VMAudit = int(VMAudit)
				}

				VMConsole, VMConsoleOk := responseVMsMap["VM.Console"].(float64)
				if VMConsoleOk {
					responseVMs.VMConsole = int(VMConsole)
				}

				VMBackup, VMBackupOk := responseVMsMap["VM.Backup"].(float64)
				if VMBackupOk {
					responseVMs.VMBackup = int(VMBackup)
				}

				VMSnapshotRollback, VMSnapshotRollbackOk := responseVMsMap["VM.Snapshot.Rollback"].(float64)
				if VMSnapshotRollbackOk {
					responseVMs.VMSnapshotRollback = int(VMSnapshotRollback)
				}

				VMGuestAgentFileWrite, VMGuestAgentFileWriteOk := responseVMsMap["VM.GuestAgent.FileWrite"].(float64)
				if VMGuestAgentFileWriteOk {
					responseVMs.VMGuestAgentFileWrite = int(VMGuestAgentFileWrite)
				}

				VMMigrate, VMMigrateOk := responseVMsMap["VM.Migrate"].(float64)
				if VMMigrateOk {
					responseVMs.VMMigrate = int(VMMigrate)
				}

				VMPowerMgmt, VMPowerMgmtOk := responseVMsMap["VM.PowerMgmt"].(float64)
				if VMPowerMgmtOk {
					responseVMs.VMPowerMgmt = int(VMPowerMgmt)
				}

				VMGuestAgentFileRead, VMGuestAgentFileReadOk := responseVMsMap["VM.GuestAgent.FileRead"].(float64)
				if VMGuestAgentFileReadOk {
					responseVMs.VMGuestAgentFileRead = int(VMGuestAgentFileRead)
				}

				VMConfigMemory, VMConfigMemoryOk := responseVMsMap["VM.Config.Memory"].(float64)
				if VMConfigMemoryOk {
					responseVMs.VMConfigMemory = int(VMConfigMemory)
				}

				VMConfigCloudinit, VMConfigCloudinitOk := responseVMsMap["VM.Config.Cloudinit"].(float64)
				if VMConfigCloudinitOk {
					responseVMs.VMConfigCloudinit = int(VMConfigCloudinit)
				}

				VMGuestAgentUnrestricted, VMGuestAgentUnrestrictedOk := responseVMsMap["VM.GuestAgent.Unrestricted"].(float64)
				if VMGuestAgentUnrestrictedOk {
					responseVMs.VMGuestAgentUnrestricted = int(VMGuestAgentUnrestricted)
				}

				VMGuestAgentFileSystemMgmt, VMGuestAgentFileSystemMgmtOk := responseVMsMap["VM.GuestAgent.FileSystemMgmt"].(float64)
				if VMGuestAgentFileSystemMgmtOk {
					responseVMs.VMGuestAgentFileSystemMgmt = int(VMGuestAgentFileSystemMgmt)
				}
			}

			responseData.Cap[key] = responseVMs

		case "dc":
			responseDCMap, ok := value.(map[string](any))
			var responseDC ResponseDC
			if ok {
				sysAudit, sysAuditOk := responseDCMap["Sys.Audit"].(float64)
				if sysAuditOk {
					responseDC.SysAudit = int(sysAudit)
				}

				SDNUse, SDNUseOk := responseDCMap["SDN.Use"].(float64)
				if SDNUseOk {
					responseDC.SDNUse = int(SDNUse)
				}

				SDNAudit, SDNAuditOk := responseDCMap["SDN.Audit"].(float64)
				if SDNAuditOk {
					responseDC.SDNAudit = int(SDNAudit)
				}

				sysModify, sysModifyOk := responseDCMap["Sys.Modify"].(float64)
				if sysModifyOk {
					responseDC.SysModify = int(sysModify)
				}

				SDNAllocate, SDNAllocateOk := responseDCMap["SDN.Allocate"].(float64)
				if SDNAllocateOk {
					responseDC.SDNAllocate = int(SDNAllocate)
				}
			}

			responseData.Cap[key] = responseDC

		case "access":
			responseAccessMap, ok := value.(map[string]any)
			var responseAccess ResponseAccess
			if ok {
				groupAllocate, groupAllocateOk := responseAccessMap["Group.Allocate"].(float64)
				if groupAllocateOk {
					responseAccess.GroupAllocate = int(groupAllocate)
				}

				userModify, userModifyOk := responseAccessMap["User.Modify"].(float64)
				if userModifyOk {
					responseAccess.UserModify = int(userModify)
				}

				permissionsModify, permissionsModifyOk := responseAccessMap["Permissions.Modify"].(float64)
				if permissionsModifyOk {
					responseAccess.PermissionsModify = int(permissionsModify)
				}
			}

			responseData.Cap[key] = responseAccess

		case "nodes":
			responseNodesMap, ok := value.(map[string]any)
			var responseNodes ResponseNodes
			if ok {
				sysConsole, sysConsoleOk := responseNodesMap["Sys.Console"].(float64)
				if sysConsoleOk {
					responseNodes.SysConsole = int(sysConsole)
				}

				sysPowerMgmt, sysPowerMgmtOk := responseNodesMap["Sys.PowerMgmt"].(float64)
				if sysPowerMgmtOk {
					responseNodes.SysPowerMgmt = int(sysPowerMgmt)
				}

				sysAudit, sysAuditOk := responseNodesMap["Sys.Audit"].(float64)
				if sysAuditOk {
					responseNodes.SysAudit = int(sysAudit)
				}

				sysSyslog, sysSyslogOk := responseNodesMap["Sys.Syslog"].(float64)
				if sysSyslogOk {
					responseNodes.SysSyslog = int(sysSyslog)
				}

				permissionsModify, permissionsModifyOk := responseNodesMap["Permissions.Modify"].(float64)
				if permissionsModifyOk {
					responseNodes.PermissionsModify = int(permissionsModify)
				}

				sysIncoming, sysIncomingOk := responseNodesMap["Sys.Incoming"].(float64)
				if sysIncomingOk {
					responseNodes.SysIncoming = int(sysIncoming)
				}

				sysModify, sysModifyOk := responseNodesMap["Sys.Modify"].(float64)
				if sysModifyOk {
					responseNodes.SysModify = int(sysModify)
				}

				sysAccessNetwork, sysAccessNetworkOk := responseNodesMap["Sys.AccessNetwork"].(float64)
				if sysAccessNetworkOk {
					responseNodes.SysAccessNetwork = int(sysAccessNetwork)
				}
			}

			responseData.Cap[key] = responseNodes

		default:
		}
	}

	return responseData, nil
}
