package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseVMsUnmarshal(t *testing.T) {
	tests := []struct {
		name            string
		bytes           []byte
		wantResponseVMs messages.ResponseVMs
	}{
		{
			"odered",
			[]byte(`{"VM.Config.CDROM":1,"Permissions.Modify":2,"VM.Config.CPU":3,"VM.Config.Network":4,"VM.GuestAgent.Audit":5,"VM.Clone":6,"VM.Config.Disk":7,"VM.Config.Options":8,"VM.Allocate":9,"VM.Snapshot":10,"VM.Replicate":11,"VM.Config.HWType":12,"VM.Audit":13,"VM.Console":14,"VM.Backup":15,"VM.Snapshot.Rollback":16,"VM.GuestAgent.FileWrite":17,"VM.Migrate":18,"VM.PowerMgmt":19,"VM.GuestAgent.FileRead":20,"VM.Config.Memory":21,"VM.Config.Cloudinit":22,"VM.GuestAgent.Unrestricted":23,"VM.GuestAgent.FileSystemMgmt":24}`),
			messages.ResponseVMs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseVMs, parseParseResponseVMsErr := messages.ParseResponseVMs(tc.bytes)
			if parseParseResponseVMsErr != nil {
				t.Errorf("Failed to parse: %v\n", parseParseResponseVMsErr)
			}

			if !reflect.DeepEqual(responseVMs, tc.wantResponseVMs) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseVMs, responseVMs)
			}
		})
	}
}
