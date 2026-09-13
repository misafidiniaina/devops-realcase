package server

import "testing"

func TestValidate(t *testing.T) {
	if err := validate(Server{Name: "web", Hostname: "web.local", IPAddress: "10.0.0.1", Environment: "test", Provider: "aws"}); err != nil {
		t.Fatalf("valid server rejected: %v", err)
	}
	if err := validate(Server{}); err == nil {
		t.Fatal("empty server should be rejected")
	}
}
