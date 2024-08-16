package tun

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToIpAndMask(t *testing.T) {
	ip, mask, _ := toIpAndMask("10.95.134.192/29")
	require.Equal(t, "10.95.134.192", ip)
	require.Equal(t, "255.255.255.248", mask)
}
