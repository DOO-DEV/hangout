package customerr

import "errors"

var (
	EncodePasswordErr = errors.New("can't encode password")
	DecodePasswordErr = errors.New("can't decode password")
	UserExistErr      = errors.New("user already exist")
	RecordNotFoundErr = errors.New("record not found")
)
