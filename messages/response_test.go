package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

/*

 */

func TestResponseUnmarshal(t *testing.T) {
	tests := []struct {
		name         string
		bytes        []byte
		wantResponse messages.Response
	}{
		/*
			{
				"syntetic",
				[]byte(`{"data":{"username":"root@pam","cap":"cap","CSRFPreventionToken":"69B1DC3B:/LgDUXZi4lFmcW46DC3sQg+lk3lmo3vO6zq4ZVHfZqA","ticket":"PVE:root@pam:69B1DC3B::lou41PD2pLEmlID7fPstSR5nK0Itqw/KNRxmq7NemcpdotUX0pPIebsh85xuwEioRkamHZGoQ8Pnah2SgKMzOnNpw5g+En8+wI7ZKZ2daqIuow8hWZgFgH2Ybw9GzeuNvpZjF2mwUVfy79egQaJHgnMWAv3i+DVRgKMY5Plju000mNyAW0RQ/lHyEBODh7CuFJ/QyhX5ai+AUD9ByGM6espg2z4LvS2rln4VXQCgj3Tn5izR9G/Pt8mOm/augwbXe6+UDOzYLAzlnkT4BqR3//r+7WW/qXUzsxU3bDqtT/ksGmzWnI5WOmTNqVulk+zfNUC4BSwkbzKor6Zb1qLrcA=="}}`),
				messages.Response{
					map[string]([]byte){
						"username":            []byte("root@pam"),
						"cap":                 []byte("cap"),
						"CSRFPreventionToken": []byte("69B1DC3B:/LgDUXZi4lFmcW46DC3sQg+lk3lmo3vO6zq4ZVHfZqA"),
						"ticket":              []byte("PVE:root@pam:69B1DC3B::lou41PD2pLEmlID7fPstSR5nK0Itqw/KNRxmq7NemcpdotUX0pPIebsh85xuwEioRkamHZGoQ8Pnah2SgKMzOnNpw5g+En8+wI7ZKZ2daqIuow8hWZgFgH2Ybw9GzeuNvpZjF2mwUVfy79egQaJHgnMWAv3i+DVRgKMY5Plju000mNyAW0RQ/lHyEBODh7CuFJ/QyhX5ai+AUD9ByGM6espg2z4LvS2rln4VXQCgj3Tn5izR9G/Pt8mOm/augwbXe6+UDOzYLAzlnkT4BqR3//r+7WW/qXUzsxU3bDqtT/ksGmzWnI5WOmTNqVulk+zfNUC4BSwkbzKor6Zb1qLrcA=="),
					}},
			},
		*/
		//	"cap":{"sdn":{"SDN.Allocate":1,"Permissions.Modify":2,"SDN.Audit":3,"SDN.Use":4},"vms":{"VM.Config.CDROM":1,"Permissions.Modify":2,"VM.Config.CPU":3,"VM.Config.Network":4,"VM.GuestAgent.Audit":5,"VM.Clone":6,"VM.Config.Disk":7,"VM.Config.Options":8,"VM.Allocate":9,"VM.Snapshot":10,"VM.Replicate":11,"VM.Config.HWType":12,"VM.Audit":13,"VM.Console":14,"VM.Backup":15,"VM.Snapshot.Rollback":16,"VM.GuestAgent.FileWrite":17,"VM.Migrate":18,"VM.PowerMgmt":19,"VM.GuestAgent.FileRead":20,"VM.Config.Memory":21,"VM.Config.Cloudinit":22,"VM.GuestAgent.Unrestricted":23,"VM.GuestAgent.FileSystemMgmt":24},"dc":{"Sys.Audit":1,"SDN.Use":2,"SDN.Audit":3,"Sys.Modify":4,"SDN.Allocate":5},"access":{"Group.Allocate":1,"User.Modify":2,"Permissions.Modify":3},"nodes":{"Sys.Console":1,"Sys.PowerMgmt":2,"Sys.Audit":3,"Sys.Syslog":4,"Permissions.Modify":5,"Sys.Incoming":6,"Sys.Modify":7,"Sys.AccessNetwork":8},"storage":{"Datastore.AllocateSpace":1,"Datastore.AllocateTemplate":2,"Datastore.Allocate":3,"Permissions.Modify":4,"Datastore.Audit":5},"mapping":{"Mapping.Use":1,"Permissions.Modify":2,"Mapping.Audit":3,"Mapping.Modify":4}}

		{

			"syntetic",
			[]byte(`{"data":{"username":"root@pam","cap":{"sdn":{"SDN.Allocate":1,"Permissions.Modify":2,"SDN.Audit":3,"SDN.Use":4},"vms":{"VM.Config.CDROM":1,"Permissions.Modify":2,"VM.Config.CPU":3,"VM.Config.Network":4,"VM.GuestAgent.Audit":5,"VM.Clone":6,"VM.Config.Disk":7,"VM.Config.Options":8,"VM.Allocate":9,"VM.Snapshot":10,"VM.Replicate":11,"VM.Config.HWType":12,"VM.Audit":13,"VM.Console":14,"VM.Backup":15,"VM.Snapshot.Rollback":16,"VM.GuestAgent.FileWrite":17,"VM.Migrate":18,"VM.PowerMgmt":19,"VM.GuestAgent.FileRead":20,"VM.Config.Memory":21,"VM.Config.Cloudinit":22,"VM.GuestAgent.Unrestricted":23,"VM.GuestAgent.FileSystemMgmt":24},"dc":{"Sys.Audit":1,"SDN.Use":2,"SDN.Audit":3,"Sys.Modify":4,"SDN.Allocate":5},"access":{"Group.Allocate":1,"User.Modify":2,"Permissions.Modify":3},"nodes":{"Sys.Console":1,"Sys.PowerMgmt":2,"Sys.Audit":3,"Sys.Syslog":4,"Permissions.Modify":5,"Sys.Incoming":6,"Sys.Modify":7,"Sys.AccessNetwork":8},"storage":{"Datastore.AllocateSpace":1,"Datastore.AllocateTemplate":2,"Datastore.Allocate":3,"Permissions.Modify":4,"Datastore.Audit":5},"mapping":{"Mapping.Use":1,"Permissions.Modify":2,"Mapping.Audit":3,"Mapping.Modify":4}},"CSRFPreventionToken":"69B1DC3B","ticket":"PVE"}}`),
			messages.Response{
				map[string](any){
					"username": "root@pam",
					"cap": map[string](any){
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
					"CSRFPreventionToken": "69B1DC3B",
					"ticket":              "PVE",
				}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var responseData messages.Response

			responseData, parseResponseAccessErr := messages.ParseResponse(tc.bytes)
			if parseResponseAccessErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseAccessErr)
			}

			if !reflect.DeepEqual(responseData, tc.wantResponse) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponse, responseData)
			}
		})
	}
}
