package dns

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alibaba/kt-connect/pkg/kt/util"
)

func Test_getAllDomainSuffixes(t *testing.T) {
	suffixes := getAllDomainSuffixes(map[string]string{
		"abc.com":   "",
		"a.b.c.net": "",
		"c.b.a.com": "",
		"xyz.net":   "",
	})
	require.True(t, util.ArrayEquals([]string{"com", "net"}, suffixes))
}
