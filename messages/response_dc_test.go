package messages_test

import (
	"reflect"
	"testing"

	"github.com/archipelagos/proxmox/messages"
)

func TestResponseDCUnmarshal(t *testing.T) {
	tests := []struct {
		name           string
		bytes          []byte
		wantResponseDC messages.ResponseDC
	}{
		{
			"odered",
			[]byte(`{"Sys.Audit":1,"SDN.Use":2,"SDN.Audit":3,"Sys.Modify":4,"SDN.Allocate":5}`),
			messages.ResponseDC{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			responseDC, parseResponseDCErr := messages.ParseResponseDC(tc.bytes)
			if parseResponseDCErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseDCErr)
			}

			if !reflect.DeepEqual(responseDC, tc.wantResponseDC) {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseDC, responseDC)
			}
		})
	}
}
