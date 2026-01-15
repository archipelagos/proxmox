package api_test

import (
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
		name     string
		host     string
		port     int
		username string
		password string
	}{
		{
			"miner",
			"192.168.1.5",
			8006,
			"username=" + url.QueryEscape("some_user"),
			"password=" + url.QueryEscape("some_password"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupTest(t)
			defer teardownTest(t)

			api.LoadEnvConfig("poleos")

			url := "https://" + tc.host + ":8006/api2/json/access/ticket?" + tc.username + "&" + tc.password

			req, reqErr := http.NewRequest("POST", url, nil)
			if reqErr != nil {
				t.Errorf("Failed to create request: %v\n", reqErr)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Failed to create request: %v\n", err)
			}
			defer resp.Body.Close()

			t.Log("response Status:", resp.Status)
			//t.Log("response Headers:", resp.Header)
			body, _ := io.ReadAll(resp.Body)
			t.Log("response Body:", string(body))
		})
	}
}
