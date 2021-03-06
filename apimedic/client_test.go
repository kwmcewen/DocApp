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
	assert.Equal(t, "https://sandbox-authservice.priaid.ch", services[NewClient(Sandbox, nil).Mode].authURL)
}

func TestLiveClientUrlShouldBeLiveHealthUrlWhenHealthSpecified(t *testing.T) {
	assert.Equal(t, "https://healthservice.priaid.ch", services[NewClient(Live, nil).Mode].healthURL)
}

func TestGetAuthUrlShouldReturnAuthUrl(t *testing.T) {
	assert.Equal(t, "https://sandbox-authservice.priaid.ch/login", NewClient(Sandbox, nil).getAuthURL())
}

func TestComputeHashShouldComputeProperHash(t *testing.T) {
	assert.Equal(t, "xhEmHyH1FnT2K9T2DScoMw==", NewClient(Sandbox, nil).computeHash("abc"))
}
