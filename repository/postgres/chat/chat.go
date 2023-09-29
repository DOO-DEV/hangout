package pgchat

import (
	"context"
	"hangout/entity"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (d DB) CreatePrivateChat(ctx context.Context, chat entity.PrivateChat) (*entity.PrivateChat, error) {
	const op = "ChatRepository.CreatePrivateChat"

	_, err := d.conn.Conn().ExecContext(ctx, `insert into "private_chat"("id", "name") values ($1, $2)`, chat.ID, chat.Name)
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
