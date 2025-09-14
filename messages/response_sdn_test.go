package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseSDNUnmarshal(t *testing.T) {
	tests := []struct {
		name            string
		bytes           []byte
		wantResponseSDN messages.ResponseSDN
	}{
		{
			"odered",
			[]byte(`{"SDN.Allocate":1,"Permissions.Modify":2,"SDN.Audit":3,"SDN.Use":4}`),
			messages.ResponseSDN{1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseSDN, parseResponseSDNErr := messages.ParseResponseSDN(tc.bytes)
			if parseResponseSDNErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseSDNErr)
			}

			if !reflect.DeepEqual(responseSDN, tc.wantResponseSDN) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseSDN, responseSDN)
			}
		})
	}
}
