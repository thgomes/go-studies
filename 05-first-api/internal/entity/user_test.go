package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Thiago Gomes", "thiago@gmail.com", "teste@756")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Thiago Gomes", user.Name)
	assert.Equal(t, "thiago@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Thiago Gomes", "thiago@gmail.com", "teste@756")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("teste@756"))
	assert.False(t, user.ValidatePassword("teste123"))
	assert.NotEqual(t, user.Password, "teste@756")
}
