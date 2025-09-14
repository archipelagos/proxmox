package messages_test

import (
	"encoding/json"
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
			"odered",
			[]byte(`{"ticket":"ticket","cap":{"sdn":{},"vms":{},"dc":{},"access":{},"nodes":{},"storage":{},"mapping":{}},"username":"username","CSRFPreventionToken":"CSRFPreventionToken"}`),
			messages.ResponseData{"ticket", map[string](any){}, "username", "CSRFPreventionToken"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var responseData messages.ResponseData
			unmarshalErr := json.Unmarshal(tc.bytes, &responseData)
			if unmarshalErr != nil {
				t.Errorf("Failed to create request: %v\n", unmarshalErr)
			}

			if responseData.Ticket != tc.wantResponseData.Ticket {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseData.Ticket, responseData.Ticket)
			}

			if len(responseData.Cap) != 7 {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", 7, len(responseData.Cap))
			}

			if responseData.Cap["sdn"] == nil {
				t.Errorf("Test fail! Structure different than expected, no: '%v' was found\n", "sdn")
			}

			if responseData.Cap["vms"] == nil {
				t.Errorf("Test fail! Structure different than expected, no: '%v' was found\n", "vms")
			}

			if responseData.Cap["dc"] == nil {
				t.Errorf("Test fail! Structure different than expected, no: '%v' was found\n", "dc")
			}

			if responseData.Cap["access"] == nil {
				t.Errorf("Test fail! Structure different than expected, no: '%v' was found\n", "access")
			}

			if responseData.Cap["nodes"] == nil {
				t.Errorf("Test fail! Structure different than expected, no: '%v' was found\n", "nodes")
			}

			if responseData.Cap["storage"] == nil {
				t.Errorf("Test fail! Structure different than expected, no: '%v' was found\n", "storage")
			}

			if responseData.Cap["mapping"] == nil {
				t.Errorf("Test fail! Structure different than expected, no: '%v' was found\n", "mapping")
			}

			if responseData.Username != tc.wantResponseData.Username {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseData.Username, responseData.Username)

			}

			if responseData.CSRFPreventionToken != tc.wantResponseData.CSRFPreventionToken {
				t.Errorf("Test fail! Structure different than expected, want: '%v', got: '%v'\n", tc.wantResponseData.CSRFPreventionToken, responseData.CSRFPreventionToken)

			}
		})
	}
}
