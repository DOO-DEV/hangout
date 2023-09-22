package pgchat

import (
	"context"
	"errors"
	"hangout/entity"
	dbparam "hangout/param/pgdb"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (d DB) SaveMessage(ctx context.Context, m entity.Message) error {
	const op = "ChatService.SaveMessage"

	if _, err := d.conn.Conn().ExecContext(ctx, `insert into "messages"("sender", "receiver", "content", "type") values ($1, $2, $3, $4)`, m.Sender, m.Receiver, m.Content, m.Type); err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}
func (d DB) GetChatMessages(ctx context.Context, sender, receiver string) ([]entity.Message, error) {
	const op = "ChatRepository.GetChatMessages"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from "messages" where "sender" = $1 and "receiver" = $2`, sender, receiver)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	defer rows.Close()

	messages := make([]entity.Message, 0)
	for rows.Next() {
		m := entity.Message{}
		if err := rows.Scan(&m.ID, &m.Sender, &m.Receiver, &m.Content, &m.Type, &m.Timestamp); err != nil {
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

func (d DB) GetUserChatList(ctx context.Context, userID string) ([]dbparam.Chat, error) {
	const op = "GroupRepository.GetUserChatList"

	rows, err := d.conn.Conn().QueryContext(ctx, `select distinct "sender", "receiver" from "messages" where "sender" = $1`, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	defer rows.Close()
	chats := make([]dbparam.Chat, 0)
	for rows.Next() {
		c := dbparam.Chat{}
		if err := rows.Scan(&c.Sender, &c.Receiver); err != nil {
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
