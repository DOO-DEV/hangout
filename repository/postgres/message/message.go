package pgmessage

import (
	"context"
	"hangout/entity"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
	"time"
)

func (d DB) SavePrivateMessage(ctx context.Context, msg entity.Message) (*entity.Message, error) {
	const op = "MessageRepository.SavePrivateMessage"

	_, err := d.conn.Conn().ExecContext(ctx, `insert into "private_messages"("id", "chat_id", "sender_id", "content", "type", "status") values ($1,$2,$3,$4,$5,$6) RETURNING *`, msg.ID, msg.ChatID, msg.SenderID, msg.Content, msg.Type, msg.Status)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	row := d.conn.Conn().QueryRowContext(ctx, `select "timestamp" from "private_messages" where "id" = $1`, msg.ID)
	var timestamp time.Time
	if err := row.Scan(&timestamp); err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	msg.Timestamp = timestamp

	return &msg, nil
}
func (d DB) SaveGroupMessage(ctx context.Context, msg entity.Message) (*entity.Message, error) {
	const op = "MessageRepository.SaveGroupMessage"

	row := d.conn.Conn().QueryRowContext(ctx, `insert into "group_messages"("id", "chat_id", "sender_id", "content", "type", "status") values ($1,$2,$3,$4,$5,$6) RETURNING "timestamp"`, msg.ID, msg.ChatID, msg.SenderID, msg.Content, msg.Type, msg.Status)
	var timestamp time.Time
	if err := row.Scan(&timestamp); err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	msg.Timestamp = timestamp

	return &msg, nil
}
