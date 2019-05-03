package models

import (
	"fmt"
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

func CreateUser(user User) error {
	var pass, err = hashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec(`
		WITH new_actor as(
			INSERT INTO Actors (actor_name, display_name, description)
			VALUES($1, $2, $3)
			RETURNING actor_name
		)
		INSERT INTO Users (pass, actor)
		VALUES($4, (select actor_name from new_actor))`,
		user.ActorName, user.DisplayName, user.Description, pass,
	)
	return err
}

func FindUserByActorName(actorName string) (User, error) {
	var user User
	err := db.DB.QueryRow(`
		SELECT
			a.actor_id,
			a.actor_name,
			a.display_name,
			a.description
		FROM Users u
		INNER JOIN Actors a ON u.actor = a.actor_name
		WHERE u.actor = $1`,
		actorName,
	).Scan(
		&user.ID,
		&user.ActorName,
		&user.DisplayName,
		&user.Description,
	)
	return user, err
}

func UpdateUserByActorName(actorName string, updUser User) error {
	var count int64
	res, err := db.DB.Exec(`
		UPDATE Actors
		SET
			actor_name = $1,
			display_name = $2,
			description = $3
		WHERE actor_name = $4 AND
			EXISTS (SELECT * FROM users WHERE actor = $4);
		`,
		updUser.ActorName,
		updUser.DisplayName,
		updUser.Description,
		actorName,
	)

	if err != nil {
		return err
	}

	if count, err = res.RowsAffected(); err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("User not exists")
	}

	return nil
}

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}
