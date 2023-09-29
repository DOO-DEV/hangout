package pgchat

import (
	"context"
	"errors"
	"fmt"
	"hangout/entity"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (d DB) GetChatByUsersIds(ctx context.Context, u1, u2 string) (*entity.Chat, error) {
	const op = "ChatRepository.GetChatByUsersIds"

	row := d.conn.Conn().QueryRowContext(ctx, `select * from "chats" where (("user_1"=$1 and "user_2"=$2) or ("user_1"=$3 and "user_2"=$4))`, u1, u2, u2, u1)
	c := &entity.Chat{UsersIDs: make([]string, 2)}
	if err := row.Scan(&c.ID, &c.UsersIDs[0], &c.UsersIDs[1], &c.Type, &c.CreatedAt); err != nil {
		if d.conn.IsEmptyRowError(err) {
			// because of our service check on error. it means we want to create a chat if doesn't exist.
			return nil, nil
		}

		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return c, nil
}

func (d DB) CreateChat(ctx context.Context, c *entity.Chat) (*entity.Chat, error) {
	const op = "ChatRepository.CreateChat"

	_, err := d.conn.Conn().ExecContext(ctx, `insert into "chats"("id","user_1", "user_2", "type") values ($1, $2, $3, $4)`, c.ID, c.UsersIDs[0], c.UsersIDs[1], c.Type)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return c, nil
}

func (d DB) SaveMessage(ctx context.Context, m entity.Message) error {
	const op = "ChatService.SaveMessage"

	if _, err := d.conn.Conn().ExecContext(ctx, `insert into "messages"("chat_id", "content", "type") values ($1, $2, $3)`, m.ChatID, m.Content, m.Type); err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}

func (d DB) GetChatMessages(ctx context.Context, chatID string) ([]entity.Message, error) {
	const op = "ChatRepository.GetChatMessages"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from "messages" where "chat_id" = $1`, chatID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	defer rows.Close()

	messages := make([]entity.Message, 0)
	for rows.Next() {
		m := entity.Message{}
		if err := rows.Scan(&m.ID, &m.ChatID, &m.Content, &m.Type, &m.Timestamp); err != nil {
			if d.conn.IsEmptyRowError(err) {
				return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNoChatMessage)
			}
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}
		messages = append(messages, m)
	}
	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	if len(messages) == 0 {
		return nil, richerror.New(op).WithError(errors.New("")).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNoChatMessage)
	}

	return messages, nil
}

func (d DB) GetUserChatList(ctx context.Context, userID string) ([]entity.Chat, error) {
	const op = "GroupRepository.GetUserChatList"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from "chats" where "user_1"=$1 or "user_2"=$2`, userID, userID)
	if err != nil {
		fmt.Println("err", err)
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	defer rows.Close()
	chats := make([]entity.Chat, 0)
	for rows.Next() {
		c := entity.Chat{UsersIDs: make([]string, 2)}
		if err := rows.Scan(&c.ID, &c.UsersIDs[0], &c.UsersIDs[1], &c.Type, &c.CreatedAt); err != nil {
			if d.conn.IsEmptyRowError(err) {
				return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNoChatMessage)
			}
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}
		chats = append(chats, c)
	}
	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithError(errors.New("")).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNoChatMessage)
	}

	return chats, nil
}

func (d DB) (ctx context.Context, participants []string) (entity.Chat, error) {
	d.conn.Conn().
}
