package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	hostname = ""
	env = ""
}

func TestBuildResponseNoEnvironment(t *testing.T) {
	expected := "none"
	t.Log("Response with no DEMO_ENV set")
	r := BuildResponse()
	if r.Environment != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, r.Environment)
	}
}

func TestBuildResponseEnvironment(t *testing.T) {
	expected := "staging"
	t.Logf("Response with DEMO_ENV set to %s", expected)

	os.Setenv("DEMO_ENV", expected)

	r := BuildResponse()
	if r.Environment != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, r.Environment)
	}
}

func TestBuildResponseHostname(t *testing.T) {
	expected, _ := os.Hostname()
	t.Log("Response with hostname")
	r := BuildResponse()
	if r.Hostname != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, r.Hostname)
	}
}

func TestBuildResponseVersion(t *testing.T) {
	t.Log("Response with version")
	r := BuildResponse()
	if r.Version != Version {
		t.Errorf("Expected: %s, Actual: %s", Version, r.Version)
	}
}
