package apimedic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientShouldBeSandboxWhenSandboxSpecified(t *testing.T) {
	assert.Equal(t, Sandbox, NewClient(Sandbox, nil).Mode)
}

func TestClientShouldBeLiveWhenLiveSpecified(t *testing.T) {
	assert.Equal(t, Live, NewClient(Live, nil).Mode)
}

func TestSandboxClientUrlShouldBeSandboxAuthUrlWhenAuthSpecified(t *testing.T) {
	assert.Equal(t, "https://sandbox-authservice.priaid.ch", services[NewClient(Sandbox, nil).Mode].authUrl)
}

func TestLiveClientUrlShouldBeLiveHealthUrlWhenHealthSpecified(t *testing.T) {
	assert.Equal(t, "https://healthservice.priaid.ch", services[NewClient(Live, nil).Mode].healthUrl)
}

func TestComputeHashShouldComputeProperHash(t *testing.T) {
	assert.Equal(t, "aHR0cHM6Ly9zYW5kYm94LWF1dGhzZXJ2aWNlLnByaWFpZC5jaC9sb2dpbkojquyGPxvQl01Og5ENPhc=", NewClient(Sandbox, nil).computeHash("abc"))
}
