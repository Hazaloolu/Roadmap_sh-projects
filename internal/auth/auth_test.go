package auth

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateSecretKey(t *testing.T) {

	key := generateSecureKey()
	decodeKey, err := base64.RawURLEncoding.DecodeString(key)
	require.NoError(t, err, "Error decoding secure key")
	assert.Equal(t, 32, len(decodeKey), "secure key length should be 32 bytes")
}

func TestHashPassword(t *testing.T) {
	password := "my_password"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err, "Error hashing password")
	assert.NotEqual(t, password, hashedPassword, "Passworld hash should not match plain password")
}
