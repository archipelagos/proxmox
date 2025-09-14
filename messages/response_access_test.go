package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseAccessUnmarshal(t *testing.T) {
	tests := []struct {
		name               string
		bytes              []byte
		wantResponseAccess messages.ResponseAccess
	}{
		{
			"odered",
			[]byte(`{"Group.Allocate":1,"User.Modify":2,"Permissions.Modify":3}`),
			messages.ResponseAccess{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseAccess, parseResponseAccessErr := messages.ParseResponseAccess(tc.bytes)
			if parseResponseAccessErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseAccessErr)
			}

			if !reflect.DeepEqual(responseAccess, tc.wantResponseAccess) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseAccess, responseAccess)
			}
		})
	}
}
