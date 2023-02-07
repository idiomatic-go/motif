package runtime

import (
	"strings"
)

const (
	EnvVariableReference = "{env}"
)

func EnvExpansion(s string) string {
	index := strings.Index(s, EnvVariableReference)
	if index == -1 {
		return "invalid or missing environment variable reference: {env}"
	}
	t := s[:index] + GetRuntimeEnv()
	u := s[index+len(EnvVariableReference):]
	return t + u
}
