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

// func TestGenerateJwt(t *testing.T) {
// 	user := &model.User{
// 		ID:       1,
// 		Username: "john_doe",
// 	}

// 	// Generate token
// 	token, err := GenerateJwt(user, "test_secret_key")
// 	require.NoError(t, err, "Error generatimg token")
// 	assert.NotEmpty(t, token, "token should not be empty")

// 	claims, err := ValidateJwt(token)
// 	require.NoError(t, err, "Error validating token")
// 	assert.NotNil(t, claims, "claims should not be nil")
// 	assert.Equal(t, user.ID, claims.UserID, "UserID in claims does not match")
// 	assert.Equal(t, user.Username, claims.Username, "Username in claim does not match")

// }
