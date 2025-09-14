package api_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/archipelagos/proxmox/api"
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
			api.EnvConfigProd.ProxmoxHost,
			api.EnvConfigProd.ProxmoxPort,
			"username=" + url.QueryEscape(api.EnvConfigProd.ProxmoxUsername),
			"password=" + url.QueryEscape(api.EnvConfigProd.ProxmoxPassword),
			http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupTest(t)
			defer teardownTest(t)

			url := "https://" + tc.host + ":" + fmt.Sprintf("%d", tc.port) + "/api2/json/access/ticket?" + tc.username + "&" + tc.password

			req, reqErr := http.NewRequest("POST", url, nil)
			if reqErr != nil {
				t.Errorf("Failed to create request\n")
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Failed to execute request\n")
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Test fail! Response status different than expected, want: '%v', got: '%v'\n", tc.wantStatusCode, resp.StatusCode)
			}

			body, _ := io.ReadAll(resp.Body)

			//t.Log("response Body:", string(body))
			_ = body
		})
	}
}
