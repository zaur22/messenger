package models

import (
	"messenger/core/db"
	"strconv"
	"testing"
)

func TestUserCRU(t *testing.T) {

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

	testUserCreate(t, &user)
	testUserUpdate(user, updUser, t)
	updUser.ID = user.ID
	testUserRead(updUser, t)
	//testUserDelete(user, t)
}

func TestMessages(t *testing.T) {
	if err := db.MigrationUp(); err != nil {
		t.Fatalf(err.Error())
	}
	defer db.DropDB()

	u1, u2 := createUsers(t)
	N := 10

	for i := 0; i < N; i++ {
		err := CreateMessage(u1, u2, ([]byte)("message "+strconv.Itoa(i)))
		if err != nil {
			t.Fatal(err)
		}
	}

	list, err := GetChatsList(u1)
	if err != nil {
		t.Fatalf("Can't get chats list: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("Bad len of chats list, expected 1, got %v", len(list))
	}

	if len(list[0].Messages) != N {
		t.Fatalf("Bad len of messages list, expected %v, got %v", N, len(list[0].Messages))
	}

	err = SetStatusReaded(u1, u2, list[0].Messages[N-1].ID)
	if err != nil {
		t.Fatal("Can't set status readed: ", err)
	}
	mList, err := GetMessagesList(u1, u2)
	if err != nil {
		t.Fatal("Can't get messages list: ", err)
	}
	if len(mList) != N {
		t.Fatalf("Bad count of message: expected %v, got %v", N, len(mList))
	}

	if !mList[N-1].IsReaded {
		t.Fatalf("Bad value of 'is readed' status, expected true, got false")
	}
}

func testUserCreate(t *testing.T, user *User) {
	id, err := CreateUser(*user)
	if err != nil {
		t.Fatalf("Can't create user: %v", err)
	}
	user.ID = id
}

func testUserRead(user User, t *testing.T) {
	var newUser, err = FindUserByID(user.ID)
	if err != nil {
		t.Fatalf("Can't find user: %v", err)
	}
	if newUser.ID != 1 {
		t.Fatalf("Unexpected value of find actor id. Expected 1, got %v", newUser.ID)
	}
	user.ID = newUser.ID
	user.Password = ""
	if newUser != user {
		t.Fatalf("Unexpected finding result: expected %v, got %v", user, newUser)
	}
}

func testUserUpdate(user User, updUser User, t *testing.T) {
	err := UpdateUserByID(user.ID, updUser)
	if err != nil {
		t.Fatalf("Can't update user: %v", err)
	}
}

func createUsers(t *testing.T) (int, int) {
	user1 := User{
		Password: "sdbf39238cbds",
		Actor: Actor{
			ActorName: "user_1",
		},
	}
	user2 := User{
		Password: "sdbf39238cbds",
		Actor: Actor{
			ActorName: "user_2",
		},
	}

	id1, err := CreateUser(user1)
	if err != nil {
		t.Fatal(err)
	}
	id2, err := CreateUser(user2)
	if err != nil {
		t.Fatal(err)
	}
	return id1, id2
}
