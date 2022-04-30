package sqlstore_test

import (
	"os"
	"testing"

	"github.com/denisandreenko/http-rest-api/internal/app/model"
	"github.com/denisandreenko/http-rest-api/internal/app/store"
	"github.com/denisandreenko/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func getDatabaseURL() string {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	}
	return databaseURL
}

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, getDatabaseURL())
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, getDatabaseURL())
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, getDatabaseURL())
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
