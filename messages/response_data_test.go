package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseDataUnmarshal(t *testing.T) {
	tests := []struct {
		name             string
		bytes            []byte
		wantResponseData messages.ResponseData
	}{
		{
			"syntetic",
			[]byte(`{"ticket":"ticket","cap":{"sdn":{"SDN.Allocate":1,"Permissions.Modify":2,"SDN.Audit":3,"SDN.Use":4},"vms":{"VM.Config.CDROM":1,"Permissions.Modify":2,"VM.Config.CPU":3,"VM.Config.Network":4,"VM.GuestAgent.Audit":5,"VM.Clone":6,"VM.Config.Disk":7,"VM.Config.Options":8,"VM.Allocate":9,"VM.Snapshot":10,"VM.Replicate":11,"VM.Config.HWType":12,"VM.Audit":13,"VM.Console":14,"VM.Backup":15,"VM.Snapshot.Rollback":16,"VM.GuestAgent.FileWrite":17,"VM.Migrate":18,"VM.PowerMgmt":19,"VM.GuestAgent.FileRead":20,"VM.Config.Memory":21,"VM.Config.Cloudinit":22,"VM.GuestAgent.Unrestricted":23,"VM.GuestAgent.FileSystemMgmt":24},"dc":{"Sys.Audit":1,"SDN.Use":2,"SDN.Audit":3,"Sys.Modify":4,"SDN.Allocate":5},"access":{"Group.Allocate":1,"User.Modify":2,"Permissions.Modify":3},"nodes":{"Sys.Console":1,"Sys.PowerMgmt":2,"Sys.Audit":3,"Sys.Syslog":4,"Permissions.Modify":5,"Sys.Incoming":6,"Sys.Modify":7,"Sys.AccessNetwork":8},"storage":{"Datastore.AllocateSpace":1,"Datastore.AllocateTemplate":2,"Datastore.Allocate":3,"Permissions.Modify":4,"Datastore.Audit":5},"mapping":{"Mapping.Use":1,"Permissions.Modify":2,"Mapping.Audit":3,"Mapping.Modify":4}},"username":"username","CSRFPreventionToken":"CSRFPreventionToken"}`),
			messages.ResponseData{
				"ticket",
				map[string](any){
					"storage": messages.ResponseStorage{
						DatastoreAllocateSpace:    1,
						DatastoreAllocateTemplate: 2,
						DatastoreAllocate:         3,
						PermissionsModify:         4,
						DatastoreAudit:            5,
					},
					"mapping": messages.ResponseMapping{
						MappingUse:        1,
						PermissionsModify: 2,
						MappingAudit:      3,
						MappingModify:     4,
					},
					"sdn": messages.ResponseSDN{
						SDNAllocate:       1,
						PermissionsModify: 2,
						SDNAudit:          3,
						SDNUse:            4,
					},
					"vms": messages.ResponseVMs{
						VMConfigCDROM:              1,
						PermissionsModify:          2,
						VMConfigCPU:                3,
						VMConfigNetwork:            4,
						VMGuestAgentAudit:          5,
						VMClone:                    6,
						VMConfigDisk:               7,
						VMConfigOptions:            8,
						VMAllocate:                 9,
						VMSnapshot:                 10,
						VMReplicate:                11,
						VMConfigHWType:             12,
						VMAudit:                    13,
						VMConsole:                  14,
						VMBackup:                   15,
						VMSnapshotRollback:         16,
						VMGuestAgentFileWrite:      17,
						VMMigrate:                  18,
						VMPowerMgmt:                19,
						VMGuestAgentFileRead:       20,
						VMConfigMemory:             21,
						VMConfigCloudinit:          22,
						VMGuestAgentUnrestricted:   23,
						VMGuestAgentFileSystemMgmt: 24,
					},
					"dc": messages.ResponseDC{
						SysAudit:    1,
						SDNUse:      2,
						SDNAudit:    3,
						SysModify:   4,
						SDNAllocate: 5,
					},
					"access": messages.ResponseAccess{
						GroupAllocate:     1,
						UserModify:        2,
						PermissionsModify: 3,
					},
					"nodes": messages.ResponseNodes{
						SysConsole:        1,
						SysPowerMgmt:      2,
						SysAudit:          3,
						SysSyslog:         4,
						PermissionsModify: 5,
						SysIncoming:       6,
						SysModify:         7,
						SysAccessNetwork:  8,
					},
				},
				"username",
				"CSRFPreventionToken"},
		},
		{
			"real",
			[]byte(`{"ticket":"PVE:username@environemnt:487C51AE::SmQnt8iyMOFOlEwl4ifSoWz9vcZaKI521kmce15ZXoQExOkwSMJSsSp2yBz8GC4k9bP6XIPbzPYYi0agZUlYf62i+/gIsMZZYoqTu5uSSjNQva1tZ33TfvL38B2QHYqjcPuJT8FOZS6PdmHy6voOV7//kK1tQNUR4Rqs5m42HRymrGpjz0OpYl9xTnMkru4Q1tmDc9lP8/gTPpnb98G6UpKEnb5cDlH+8vMvxVjX25RCvNbViFV2HRvbRTqMx3wsqSc3mEqGQZzz1ZBBeHbUu1TcookukBZuxLgEk9WYolDC5SHytb+JZCaN/H15JqRl6CoIInvMnyFHWCr2hQJJvQ==","cap":{"sdn":{"SDN.Allocate":1,"Permissions.Modify":1,"SDN.Audit":1,"SDN.Use":1},"vms":{"VM.Config.CDROM":1,"Permissions.Modify":1,"VM.Config.CPU":1,"VM.Config.Network":1,"VM.GuestAgent.Audit":1,"VM.Clone":1,"VM.Config.Disk":1,"VM.Config.Options":1,"VM.Allocate":1,"VM.Snapshot":1,"VM.Replicate":1,"VM.Config.HWType":1,"VM.Audit":1,"VM.Console":1,"VM.Backup":1,"VM.Snapshot.Rollback":1,"VM.GuestAgent.FileWrite":1,"VM.Migrate":1,"VM.PowerMgmt":1,"VM.GuestAgent.FileRead":1,"VM.Config.Memory":1,"VM.Config.Cloudinit":1,"VM.GuestAgent.Unrestricted":1,"VM.GuestAgent.FileSystemMgmt":1},"dc":{"Sys.Audit":1,"SDN.Use":1,"SDN.Audit":1,"Sys.Modify":1,"SDN.Allocate":1},"access":{"Group.Allocate":1,"User.Modify":1,"Permissions.Modify":1},"nodes":{"Sys.Console":1,"Sys.PowerMgmt":1,"Sys.Audit":1,"Sys.Syslog":1,"Permissions.Modify":1,"Sys.Incoming":1,"Sys.Modify":1,"Sys.AccessNetwork":1},"storage":{"Datastore.AllocateSpace":1,"Datastore.AllocateTemplate":1,"Datastore.Allocate":1,"Permissions.Modify":1,"Datastore.Audit":1},"mapping":{"Mapping.Use":1,"Permissions.Modify":1,"Mapping.Audit":1,"Mapping.Modify":1}},"username":"username@environemnt","CSRFPreventionToken":"487C51AE:UBuWv9yb2G+5e7bjUUoLx4x0+Y6L9dqEOXDl8foUl3A"}`),
			messages.ResponseData{
				"PVE:username@environemnt:487C51AE::SmQnt8iyMOFOlEwl4ifSoWz9vcZaKI521kmce15ZXoQExOkwSMJSsSp2yBz8GC4k9bP6XIPbzPYYi0agZUlYf62i+/gIsMZZYoqTu5uSSjNQva1tZ33TfvL38B2QHYqjcPuJT8FOZS6PdmHy6voOV7//kK1tQNUR4Rqs5m42HRymrGpjz0OpYl9xTnMkru4Q1tmDc9lP8/gTPpnb98G6UpKEnb5cDlH+8vMvxVjX25RCvNbViFV2HRvbRTqMx3wsqSc3mEqGQZzz1ZBBeHbUu1TcookukBZuxLgEk9WYolDC5SHytb+JZCaN/H15JqRl6CoIInvMnyFHWCr2hQJJvQ==",
				map[string](any){
					"storage": messages.ResponseStorage{
						DatastoreAllocateSpace:    1,
						DatastoreAllocateTemplate: 1,
						DatastoreAllocate:         1,
						PermissionsModify:         1,
						DatastoreAudit:            1,
					},
					"mapping": messages.ResponseMapping{
						MappingUse:        1,
						PermissionsModify: 1,
						MappingAudit:      1,
						MappingModify:     1,
					},
					"sdn": messages.ResponseSDN{
						SDNAllocate:       1,
						PermissionsModify: 1,
						SDNAudit:          1,
						SDNUse:            1,
					},
					"vms": messages.ResponseVMs{
						VMConfigCDROM:              1,
						PermissionsModify:          1,
						VMConfigCPU:                1,
						VMConfigNetwork:            1,
						VMGuestAgentAudit:          1,
						VMClone:                    1,
						VMConfigDisk:               1,
						VMConfigOptions:            1,
						VMAllocate:                 1,
						VMSnapshot:                 1,
						VMReplicate:                1,
						VMConfigHWType:             1,
						VMAudit:                    1,
						VMConsole:                  1,
						VMBackup:                   1,
						VMSnapshotRollback:         1,
						VMGuestAgentFileWrite:      1,
						VMMigrate:                  1,
						VMPowerMgmt:                1,
						VMGuestAgentFileRead:       1,
						VMConfigMemory:             1,
						VMConfigCloudinit:          1,
						VMGuestAgentUnrestricted:   1,
						VMGuestAgentFileSystemMgmt: 1,
					},
					"dc": messages.ResponseDC{
						SysAudit:    1,
						SDNUse:      1,
						SDNAudit:    1,
						SysModify:   1,
						SDNAllocate: 1,
					},
					"access": messages.ResponseAccess{
						GroupAllocate:     1,
						UserModify:        1,
						PermissionsModify: 1,
					},
					"nodes": messages.ResponseNodes{
						SysConsole:        1,
						SysPowerMgmt:      1,
						SysAudit:          1,
						SysSyslog:         1,
						PermissionsModify: 1,
						SysIncoming:       1,
						SysModify:         1,
						SysAccessNetwork:  1,
					},
				},
				"username@environemnt",
				"487C51AE:UBuWv9yb2G+5e7bjUUoLx4x0+Y6L9dqEOXDl8foUl3A"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var responseData messages.ResponseData

			responseData, parseResponseAccessErr := messages.ParseResponseData(tc.bytes)
			if parseResponseAccessErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseAccessErr)
			}

			if !reflect.DeepEqual(responseData, tc.wantResponseData) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseData, responseData)
			}
		})
	}
}
