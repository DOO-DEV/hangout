package chatservice

import (
	"context"
	"errors"
	param "hangout/param/http"
	"hangout/pkg/errmsg"
	"hangout/pkg/richerror"
)

func (s Service) GetChatHistory(ctx context.Context, _ param.GetChatHistoryRequest, sender, receiver string) (*param.GetChatHistoryResponse, error) {
	const op = "ChatService.GetChatHistory"

	chat, err := s.chatRepo.GetChatByUsersIds(ctx, sender, receiver)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	if chat == nil {
		wErr := richerror.New(op).WithError(errors.New("")).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgUserNotFound)
		return nil, richerror.New(op).WithError(wErr)
	}

	chatMsgs, err := s.chatRepo.GetChatMessages(ctx, chat.ID)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	list := make([]param.ChatMsg, 0)
	for _, v := range chatMsgs {
		m := param.ChatMsg{
			Content:   v.Content,
			Timestamp: v.Timestamp,
		}
		list = append(list, m)
	}

	res := &param.GetChatHistoryResponse{
		Sender:   sender,
		Receiver: receiver,
		Data:     list,
	}

	return res, nil
}
