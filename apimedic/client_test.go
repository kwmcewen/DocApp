package apimedic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientShouldBeSandboxWhenSandboxSpecified(t *testing.T) {
	assert.Equal(t, Sandbox, NewClient(Sandbox).Mode)
}

func TestClientShouldBeLiveWhenLiveSpecified(t *testing.T) {
	assert.Equal(t, Live, NewClient(Live).Mode)
}

func TestSandboxClientUrlShouldBeSandboxAuthUrlWhenAuthSpecified(t *testing.T) {
	assert.Equal(t, "https://sandbox-authservice.priaid.ch", services[NewClient(Sandbox).Mode].authUrl)
}

func TestLiveClientUrlShouldBeLiveHealthUrlWhenHealthSpecified(t *testing.T) {
	assert.Equal(t, "https://healthservice.priaid.ch", services[NewClient(Live).Mode].healthUrl)
}
