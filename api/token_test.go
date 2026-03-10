package api_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/archipelagos/proxmox/api"
	"github.com/archipelagos/proxmox/messages"
)

func TestApi(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	tests := []struct {
		name           string
		host           string
		port           int
		username       string
		password       string
		wantStatusCode int
	}{
		{
			"prod",
			api.EnvConfigProd.ProxmoxInternalAddress,
			api.EnvConfigProd.ProxmoxInternalTCPPort,
			api.EnvConfigProd.ProxmoxUsername,
			api.EnvConfigProd.ProxmoxPassword,
			http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupTest(t)
			defer teardownTest(t)

			encodedUsername := url.QueryEscape(tc.username)
			encodedPassword := url.QueryEscape(tc.password)
			url := "https://" + tc.host + ":" + fmt.Sprintf("%d", tc.port) + "/api2/json/access/ticket?username=" + encodedUsername + "&password=" + encodedPassword

			req, reqErr := http.NewRequest("POST", url, nil)
			if reqErr != nil {
				//t.Errorf("Failed to create request: %v\n", reqErr)
				t.Errorf("Failed to create request\n")
			}

			client := &http.Client{
				Timeout: (5 * time.Second),
			}
			resp, err := client.Do(req)
			if err != nil {
				//t.Errorf("Failed to execute request: %v\n", err)
				t.Errorf("Failed to execute request\n")
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Test fail! Response status different than expected, want: '%v', got: '%v'\n", tc.wantStatusCode, resp.StatusCode)
			}

			body, _ := io.ReadAll(resp.Body)

			//t.Log("body:", string(body))

			responseData, parseResponseAccessErr := messages.ParseResponseData(body)
			if parseResponseAccessErr != nil {
				t.Errorf("Failed to parse: %v\n", parseResponseAccessErr)
			}

			//t.Log("responseData:", responseData)

			_ = responseData
		})
	}
}
