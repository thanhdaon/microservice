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

	db.Close()
}

func TestSaveUser_Fail(t *testing.T) {
	db := NewTestDBConnection()
	userRepo := NewUserRepository(db)

	_, err := seedUsers(db)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	user := &entity.User{
		FirstName: "dao",
		LastName:  "thanh-1",
		Email:     "thanhdao@gmail.com",
		Password:  "aaa123",
	}

	_, saveErr := userRepo.Save(user)

	assert.Equal(t, saveErr, e.EMAIL_ALREADY_EXISTS)

	db.Close()
}
