package database

import (
	"domain-driven-design/domain/e"
	"domain-driven-design/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	db := NewTestDBConnection()
	assert.NotNil(t, db)
	db.Close()
}

func TestSaveUser_Success(t *testing.T) {
	db := NewTestDBConnection()
	userRepo := NewUserRepository(db)

	user := &entity.User{
		FirstName: "Thanh",
		LastName:  "Dao",
		Email:     "thanhdao@gmail.com",
		Password:  "aaa123",
	}

	u, saveErr := userRepo.Save(user)

	assert.Nil(t, saveErr)
	assert.Equal(t, u.Email, "thanhdao@gmail.com")
	assert.Equal(t, u.FirstName, "Thanh")
	assert.Equal(t, u.LastName, "Dao")
	assert.Equal(t, u.Password, "aaa123")

	user.FirstName = "Thanh-1"
	u, saveErr = userRepo.Save(user)

	assert.Nil(t, saveErr)
	assert.Equal(t, u.FirstName, "Thanh-1")

	db.Close()
}

func TestSaveUser_Fail(t *testing.T) {
	db := NewTestDBConnection()
	userRepo := NewUserRepository(db)

	user1 := &entity.User{
		FirstName: "dao",
		LastName:  "thanh-1",
		Email:     "thanhdao@gmail.com",
		Password:  "aaa123",
	}

	user2 := &entity.User{
		FirstName: "dao",
		LastName:  "thanh-2",
		Email:     "thanhdao@gmail.com",
		Password:  "aaa123",
	}

	_, saveErr := userRepo.Save(user1)
	_, saveErr = userRepo.Save(user2)
	assert.Equal(t, e.EMAIL_ALREADY_EXISTS, saveErr)

	db.Close()
}

func TestGETByID_Success(t *testing.T) {
	db := NewTestDBConnection()
	userRepo := NewUserRepository(db)

	user := &entity.User{
		FirstName: "Thanh",
		LastName:  "Dao",
		Email:     "thanhdao@gmail.com",
		Password:  "aaa123",
	}

	_, saveErr := userRepo.Save(user)
	assert.Nil(t, saveErr)

	u, getErr := userRepo.GetByID(user.ID)
	assert.Nil(t, getErr)
	if assert.NotNil(t, u) {
		assert.Equal(t, u.ID, user.ID)
	}

	db.Close()
}

func TestGETByID_Fail(t *testing.T) {
	db := NewTestDBConnection()
	userRepo := NewUserRepository(db)

	u, getErr := userRepo.GetByID(1)

	assert.NotNil(t, getErr)
	assert.Nil(t, u)
	db.Close()
}
