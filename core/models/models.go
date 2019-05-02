package models

import (
	"messenger/core/db"

	"golang.org/x/crypto/bcrypt"
)

type ActorType int

//Actor объект, которому можно отправлять сообщения.
//Может быть как обычным пользователем, так и чатом, и ботом, и т.п.
type Actor struct {
	ID          int
	ActorName   string
	DisplayName string
	Description string
}

//User расшрение Actora для пользователя
type User struct {
	//Password пароль
	Password string
	Actor
}

//Message хранит сообщение
type Message struct {
	ID     int
	Value  string
	Sender *Actor
	To     *Actor
}

func CreateUser(user User) (int, error) {
	var id int
	var pass, err = hashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	err = db.DB.QueryRow(`
		WITH new_actor as(
			INSERT INTO actor (actor_name, display_name, description)
			VALUES($1, $2, $3)
			RETURNING actor_id
		)
		INSERT INTO "user" (pass, actor)
		VALUES($4, (select actor_id from new_actor)) RETURNING actor`,
		user.ActorName, user.DisplayName, user.Description, pass,
	).Scan(&id)
	return id, err
}

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}
