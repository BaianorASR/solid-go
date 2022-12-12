package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("PORT", "12345")
	t.Setenv("ENV", "test")
	t.Setenv("CASE_SENSITIVE_URL", "true")

	got := GetEnv()
	want := &Environment{
		Port:             12345,
		Env:              "test",
		CaseSensitiveURL: true,
	}

	assert.Equal(t, want, got)
}
