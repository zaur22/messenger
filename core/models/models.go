package models

import (
	"fmt"
	"messenger/core/db"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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
	ID        int64
	IsReaded  bool
	Value     []byte
	Sender    int
	To        int
	CreatedAt time.Time
}

type Chat struct {
	Messages []Message
	With     int
}

func CreateMessage(sender int, to int, value []byte) error {
	_, err := db.DB.Exec(`
		INSERT INTO Messages (sender_actor, to_actor, value)
		VALUES ($1, $2, $3)`,
		sender,
		to,
		value)
	return err
}

func GetChatsList(actor_id int) ([]Chat, error) {
	var chats = []Chat{}
	var count = 10
	rows, err := db.DB.Query(`
		SELECT message_id, value, is_read, created_at,
			sender_actor, to_actor, chat_with FROM
		(
		SELECT message_id, value, is_read, created_at,
			sender_actor, to_actor,
			CASE 
				WHEN to_actor = $1 THEN sender_actor
				ELSE to_actor
			END AS chat_with,
			RANK() OVER(PARTITION BY
				least(sender_actor, to_actor),
				greatest(sender_actor, to_actor)
				ORDER BY created_at
			) num
		FROM Messages
		) X
		WHERE num <= $2 AND
			(sender_actor = $1 OR to_actor = $1)`,
		actor_id, count)
	if err != nil {
		return chats, err
	}

	var ms Message
	var chat = Chat{
		Messages: make([]Message, count),
		With:     -1,
	}
	var flagStart = true
	var chatWith int
	for rows.Next() {
		err = rows.Scan(&ms.ID, &ms.Value, &ms.IsReaded,
			&ms.CreatedAt, &ms.Sender, &ms.To, &chatWith)
		if err != nil {
			return chats, err
		}
		if chatWith != chat.With {
			if !flagStart {
				chats = append(chats, chat)
			}
			chat.With = chatWith
			chat.Messages = chat.Messages[:0]
			flagStart = false
		}
		chat.Messages = append(chat.Messages, ms)
	}
	if !flagStart {
		chats = append(chats, chat)
	}
	return chats, nil
}

func SetStatusReaded(from int, to int, lastReadedMessageID int64) error {
	_, err := db.DB.Exec(`
		UPDATE Messages
		SET is_read = TRUE
		WHERE sender_actor = $1 
			AND to_actor = $2
			AND message_id <= $3`,
		from, to, lastReadedMessageID)
	return err
}

func GetMessagesList(actor1 int, actor2 int) ([]Message, error) {
	var messages []Message
	var message Message
	rows, err := db.DB.Query(`
		SELECT message_id, value, sender_actor,
			to_actor, is_read, created_at
		FROM Messages
		WHERE (to_actor = $1 AND sender_actor = $2) OR
		(sender_actor = $1 AND to_actor = $2)
	`, actor1, actor2)
	if err != nil {
		return messages, err
	}

	for rows.Next() {
		err = rows.Scan(
			&message.ID, &message.Value, &message.Sender,
			&message.To, &message.IsReaded, &message.CreatedAt,
		)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func CreateUser(user User) (int, error) {
	var pass, err = hashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	var ID int
	err = db.DB.QueryRow(`
		WITH new_actor as(
			INSERT INTO Actors (actor_name, display_name, description)
			VALUES($1, $2, $3)
			RETURNING actor_id
		)
		INSERT INTO Users (pass, actor)
		VALUES($4, (select actor_id from new_actor))
		RETURNING actor`,
		user.ActorName, user.DisplayName, user.Description, pass,
	).Scan(&ID)
	return ID, err
}

func FindUserByID(actorID int) (User, error) {
	var user User
	err := db.DB.QueryRow(`
		SELECT
			a.actor_id,
			a.actor_name,
			a.display_name,
			a.description
		FROM Users u
		INNER JOIN Actors a ON u.actor = a.actor_id
			AND a.actor_id = $1`,
		actorID,
	).Scan(
		&user.ID,
		&user.ActorName,
		&user.DisplayName,
		&user.Description,
	)
	return user, err
}

func UpdateUserByID(actorID int, updUser User) error {
	var count int64
	res, err := db.DB.Exec(`
		UPDATE Actors
		SET
			actor_name = $1,
			display_name = $2,
			description = $3
		WHERE actor_id = $4 AND
			EXISTS (SELECT * FROM users WHERE actor = $4);
		`,
		updUser.ActorName,
		updUser.DisplayName,
		updUser.Description,
		actorID,
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

/*
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
*/
func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}
