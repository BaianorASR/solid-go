package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("PORT", "12345")

	got := GetEnv()
	want := 12345

	assert.Equal(t, want, got.Port)
}
