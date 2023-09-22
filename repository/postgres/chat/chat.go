package pgchat

import (
	"context"
	"fmt"
	"hangout/entity"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (d DB) SaveMessage(ctx context.Context, m entity.Message) error {
	const op = "ChatService.SaveMessage"
	if _, err := d.conn.Conn().ExecContext(ctx, `insert into "messages"("sender", "receiver", "content", "type") values ($1, $2, $3, $4)`, m.Sender, m.Receiver, m.Content, m.Type); err != nil {
		fmt.Println(err)
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}
