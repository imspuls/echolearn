package controllers

import (
	"echolearn/models"
	"log"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var user = models.User{
	UserId:   1,
	Username: "test",
	Password: "test",
}

var user2 = models.APIUser{
	UserId:   1,
	Username: "test",
}

func SetupTests() *gorm.DB { // or *gorm.DB
	mocket.Catcher.Reset()
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dialect := mysql.New(mysql.Config{
		DSN:                       "mockdb",
		DriverName:                mocket.DriverName,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialect, new(gorm.Config))
	if err != nil {
		log.Fatalf("failed to open mock database connection: %s", err)
	}

	models.Db = gormDB

	return gormDB
}

func TestGetAllUser(t *testing.T) {
	SetupTests()
	commonReply := []map[string]interface{}{{"user_id": 1, "username": "test"}}
	mocket.Catcher.Reset().NewMock().WithArgs(int64(1)).WithReply(commonReply)
	resp, _ := models.GetAllUser()

	if len(resp) < 1 {
		t.Fatalf("Returned sets is not equal to 1. Received %d", len(resp))
	}
}

func TestGetAllUserNil(t *testing.T) {
	SetupTests()
	mocket.Catcher.Reset().NewMock().WithArgs(int64(1)).WithQueryException()
	_, err := models.GetAllUser()

	if err == nil {
		t.Fatal("Error not triggered")
	}
}

func TestStoreUser(t *testing.T) {
	SetupTests()

	var mockedId int64 = 1
	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"users\"").WithID(mockedId)

	user_id, err := user.StoreUser()
	assert.Equal(t, int(mockedId), user_id)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	SetupTests()

	mocket.Catcher.Reset().NewMock().WithQuery("UPDATE \"users\"").WithExecException()

	err := user2.UpdateUser()
	assert.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	SetupTests()

	mocket.Catcher.Reset().NewMock().WithQuery("DELETE \"users\"").WithExecException()

	err := user.DeleteUser()
	assert.Error(t, err)
}
