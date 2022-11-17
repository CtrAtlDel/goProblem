
package testStore_test

import (
	"http-rest-api/internal/app/model"
	store "http-rest-api/internal/app/store/testStore"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s:= store.New()
	u:=model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	email := "example@example.org"
	s:= store.New()
	_, err := s.User().FindByEmail(email)
	// assert.EqualError(t, err, store.ErrRecordNotFound)
	
	u:= model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}