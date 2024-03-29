package pgchat

import (
	"context"
	"hangout/entity"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (d DB) CreatePrivateChat(ctx context.Context, chat entity.PrivateChat) (*entity.PrivateChat, error) {
	const op = "ChatRepository.CreatePrivateChat"

	_, err := d.conn.Conn().ExecContext(ctx, `insert into "private_chats"("id", "name") values ($1, $2)`, chat.ID, chat.Name)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return &chat, nil
}

func (d DB) GetPrivateChatByName(ctx context.Context, name string) (*entity.PrivateChat, error) {
	const op = "ChatRepository.GetPrivateChatByName"

	var chatID string
	row := d.conn.Conn().QueryRowContext(ctx, `select "id" from "private_chats" where "name" = $1`, name)
	if err := row.Scan(&chatID); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return nil, nil
		}
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return &entity.PrivateChat{
		ID:   chatID,
		Name: name,
	}, nil
}

func (d DB) CreateGroupChat(ctx context.Context, group entity.GroupChat) (*entity.GroupChat, error) {
	const op = "ChatRepository.CreateGroupChat"

	_, err := d.conn.Conn().ExecContext(ctx, `insert into "group_chats"("id", "name") values ($1, $2)`, group.ID, group.Name)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return &group, nil
}

func (d DB) GetGroupChatByID(ctx context.Context, groupID string) (*entity.GroupChat, error) {
	const op = "ChatRepository.GetGroupChatByID"

	row := d.conn.Conn().QueryRowContext(ctx, `select "name" from "group_chats" where "id" = $1`, groupID)
	var chatName string
	if err := row.Scan(&chatName); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return nil, nil
		}
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return &entity.GroupChat{
		ID:   groupID,
		Name: chatName,
	}, nil
}

func (d DB) AddToPrivateChatParticipants(ctx context.Context, p1, p2 entity.PrivateChatParticipant) error {
	const op = "ChatRepository.AddToPrivateChatParticipants"

	_, err := d.conn.Conn().ExecContext(ctx,
		`insert into "private_chat_participants"("id", "chat_id", "user_id")
			  values ($1, $2, $3),
			         ($4, $5, $6);`, p1.ID, p1.ChatID, p1.UserID, p2.ID, p2.ChatID, p2.UserID)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}

func (d DB) GetPrivateChatMessages(ctx context.Context, chatID string) ([]entity.Message, error) {
	const op = "GetPrivateChatMessages"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from "private_messages" where "chat_id" = $1`, chatID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	messages := make([]entity.Message, 0)

	for rows.Next() {
		m := entity.Message{}
		if err := rows.Scan(&m.ID, &m.ChatID, &m.SenderID, &m.Content, &m.Type, &m.Status, &m.Timestamp); err != nil {
			return messages, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}

		messages = append(messages, m)
	}

	return messages, nil
}

func (d DB) GetUnreadPrivateChatMessages(ctx context.Context, userID string) ([]entity.Message, error) {
	const op = "ChatRepository.GetUnreadPrivateChatMessages"

	row := d.conn.Conn().QueryRowContext(ctx, `select "chat_id" from "private_chat_participants" where "user_id" = $1`, userID)
	var chatID string
	if err := row.Scan(&chatID); err != nil {
		if d.conn.IsEmptyRowError(err) {
			return nil, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
		}
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from "private_messages" where "chat_id" = $1 and status="1"`, chatID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	messages := make([]entity.Message, 0)
	for rows.Next() {
		m := entity.Message{}
		if err := rows.Scan(&m.ID, &m.ChatID, &m.SenderID, &m.Content, &m.Type, &m.Status, &m.Timestamp); err != nil {
			return messages, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
		}
		messages = append(messages, m)
	}

	return messages, nil
}
