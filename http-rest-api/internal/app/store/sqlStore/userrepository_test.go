package sqlstore_test

import (
	"http-rest-api/internal/app/model"
	store "http-rest-api/internal/app/store/sqlStore"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")
	s:= store.New(db)

	u:=model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")
	email := "example@example.org"

	s:= store.New(db)
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	
	u:= model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
