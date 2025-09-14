package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseNodesUnmarshal(t *testing.T) {
	tests := []struct {
		name              string
		bytes             []byte
		wantResponseNodes messages.ResponseNodes
	}{
		{
			"odered",
			[]byte(`{"Sys.Console":1,"Sys.PowerMgmt":2,"Sys.Audit":3,"Sys.Syslog":4,"Permissions.Modify":5,"Sys.Incoming":6,"Sys.Modify":7,"Sys.AccessNetwork":8}`),
			messages.ResponseNodes{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseNodes, parseResponseNodesErr := messages.ParseResponseNodes(tc.bytes)
			if parseResponseNodesErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseNodesErr)
			}

			if !reflect.DeepEqual(responseNodes, tc.wantResponseNodes) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseNodes, responseNodes)
			}
		})
	}
}
