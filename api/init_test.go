package api_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"testing"

	"github.com/archipelagos/proxmox/api"
)

func init() {
	err := api.LoadAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	api.EnvConfigProd = api.LoadEnvConfig("prod")

	// Disable security checks globally for all requests of the default client.
	// Security note: Disabling security checks is dangerous and should be avoided.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// You can use testing.T, if you want to test the code without benchmarking.
func setupSuite(_ testing.TB) func(tb testing.TB) {
	log.Println("setup suite")

	// Return a function to teardown the test.
	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

// Almost the same as the above, but this one is for single test instead of collection of tests.
func setupTest(_ testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	return func(tb testing.TB) {
		log.Println("teardown test")
	}
}
