package options

import (
	"strings"

	"github.com/alibaba/kt-connect/hack"
)

func GetCustomizeKubeConfig() (string, bool) {
	if len(hack.CustomizeKubeConfig) > 50 {
		return hack.CustomizeKubeConfig, true
	}
	return "", false
}

func GetCustomizeKtConfig() (string, bool) {
	if strings.Contains(hack.CustomizeKtConfig, ":") {
		return hack.CustomizeKtConfig, true
	}
	return "", false
}
