package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseMappingUnmarshal(t *testing.T) {
	tests := []struct {
		name                string
		bytes               []byte
		wantResponseMapping messages.ResponseMapping
	}{
		{
			"odered",
			[]byte(`{"Mapping.Use":1,"Permissions.Modify":2,"Mapping.Audit":3,"Mapping.Modify":4}`),
			messages.ResponseMapping{1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseMapping, parseResponseMappingErr := messages.ParseResponseMapping(tc.bytes)
			if parseResponseMappingErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseMappingErr)
			}

			if !reflect.DeepEqual(responseMapping, tc.wantResponseMapping) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseMapping, responseMapping)
			}
		})
	}
}
