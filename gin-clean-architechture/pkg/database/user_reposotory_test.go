package database

import (
	"domain-driven-design/domain/e"
	"domain-driven-design/domain/entity"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

var (
	user1 = &entity.User{
		FirstName: "fn1",
		LastName:  "ln1",
		Email:     "email1@gmail.com",
		Password:  "aaa123",
	}
	user2 = &entity.User{
		FirstName: "fn2",
		LastName:  "ln2",
		Email:     "email2@gmail.com",
		Password:  "aaa123",
	}
	user3 = &entity.User{
		FirstName: "fn3",
		LastName:  "ln3",
		Email:     "email1@gmail.com",
		Password:  "aaa123",
	}
	user4 = &entity.User{
		FirstName: "fn4",
		LastName:  "ln4",
		Email:     "email4@gmail.com",
		Password:  "aaa123",
	}
)

func Test_SaveUser_Success(t *testing.T) {
	db := NewTestDBConnection()
	defer db.Close()

	userRepo := NewUserRepository(db)

	err := userRepo.Save(user1)

	require.NoError(t, err)
	require.Equal(t, "fn1", user1.FirstName)
	require.Equal(t, "ln1", user1.LastName)
	require.Equal(t, "email1@gmail.com", user1.Email)
	require.Equal(t, "aaa123", user1.Password)

	user1.FirstName = "fn-xxx"
	err = userRepo.Save(user1)

	require.Nil(t, err)
	require.Equal(t, user1.FirstName, "fn-xxx")

	db.Delete(user1)
}

func Test_SaveUser_Fail(t *testing.T) {
	db := NewTestDBConnection()
	defer db.Close()

	userRepo := NewUserRepository(db)

	err := userRepo.Save(user1)
	err = userRepo.Save(user3)
	assert.Equal(t, e.EMAIL_ALREADY_EXISTS, err)

	db.Delete(user1)
	db.Delete(user3)
}

func Test_GetByID(t *testing.T) {
	db := NewTestDBConnection()
	defer db.Close()

	userRepo := NewUserRepository(db)

	err := userRepo.Save(user1)
	require.Nil(t, err)

	u, getErr := userRepo.GetByID(user1.ID)
	require.Nil(t, getErr)
	require.EqualValues(t, user1.ID, u.ID)

	u, getErr = userRepo.GetByID(user1.ID + 1)
	require.Nil(t, u)
	require.NotNil(t, getErr)

	db.Delete(user1)
}

func Test_GetAll(t *testing.T) {
	db := NewTestDBConnection()
	defer db.Close()

	userRepo := NewUserRepository(db)

	userRepo.Save(user1)
	userRepo.Save(user2)
	userRepo.Save(user4)

	users, err := userRepo.GetAll()
	require.Nil(t, err)
	require.Equal(t, 3, len(users))

	db.Delete(user1)
	db.Delete(user2)
	db.Delete(user4)
}

func Test_GetByEmail(t *testing.T) {
	db := NewTestDBConnection()
	defer db.Close()

	userRepo := NewUserRepository(db)

	userRepo.Save(user1)

	u, err := userRepo.GetByEmail(user1.Email)
	require.Nil(t, err)
	require.Equal(t, u.Email, user1.Email)

	u, err = userRepo.GetByEmail("okk")
	require.Nil(t, u)
	require.Equal(t, e.USER_NOT_FOUND, err)

	db.Delete(user1)
}
