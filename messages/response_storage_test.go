package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseStorageUnmarshal(t *testing.T) {
	tests := []struct {
		name                string
		bytes               []byte
		wantResponseStorage messages.ResponseStorage
	}{
		{
			"odered",
			[]byte(`{"Datastore.AllocateSpace":1,"Datastore.AllocateTemplate":2,"Datastore.Allocate":3,"Permissions.Modify":4,"Datastore.Audit":5}`),
			messages.ResponseStorage{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseStorage, parseResponseStorageErr := messages.ParseResponseStorage(tc.bytes)
			if parseResponseStorageErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseStorageErr)
			}

			if !reflect.DeepEqual(responseStorage, tc.wantResponseStorage) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseStorage, responseStorage)
			}
		})
	}
}
