package models

import (
	"messenger/core/db"
	"testing"
)

func TestUserCRUD(t *testing.T) {

	if err := db.MigrationUp(); err != nil {
		t.Fatalf(err.Error())
	}
	defer db.DropDB()

	var user = User{
		Password: "sOmeRand0m3",
		Actor: Actor{
			ActorName: "first",
		},
	}
	var updUser = user
	updUser.Description = "Abc"

	testUserCreate(t, user)
	/*testUserUpdate(user, updUser, t)
	testUserRead(updUser, t)
	testUserDelete(user, t)*/
}

func testUserCreate(t *testing.T, user User) {
	var id, err = CreateUser(user)
	if err != nil {
		t.Fatalf("Can't create user: %v", err)
	}
	if id != 1 {
		t.Fatalf("Unexpected value of created user id. Expected 1, got %v", id)
	}
}

/*
func testUserRead(user User, t *testing.T) {
	var newUser, err = FindUserByActorName(user.ActorName)
	if err != nil {
		t.Fatalf("Can't find user: %v", err)
	}
	if newUser.ID != 0 {
		t.Fatalf("Unexpected value of find actor id. Expected 1, got %v", newUser.ID)
	}
	user.ID = newUser.ID
	if newUser != user {
		t.Fatalf("Unexpected finding result: expected %v, got %v", user, newUser)
	}
}

func testUserUpdate(user User, updUser User, t *testing.T) {
	var id, err = UpdateUserByActorName(user.ActorName, updUser)
	if err != nil {
		t.Fatalf("Can't update user: %v", err)
	}
	if id != 0 {
		t.Fatalf("Unexpected value of updated actor id. Expected 1, got %v", id)
	}
}

func testUserDelete(user User, t *testing.T) {
	var id, err = DeleteUserByActorName(user.ActorName)
	if err != nil {
		t.Fatalf("Can't delete user: %v", err)
	}
	if id != 0 {
		t.Fatalf("Unexpected value of deleted actor id. Expected 1, got %v", id)
	}
}
*/

func migration_up() {

}
