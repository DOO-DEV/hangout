package richerror

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	Kind         Kind
	WrappedError error
	Message      string
	Operation    string
}

func New(op string) RichError {
	return RichError{
		Operation: op,
	}
}

func (r RichError) Error() string {
	return r.Message
}

func (r RichError) WithMessage(msg string) RichError {
	r.Message = msg
	return r
}

func (r RichError) WithKind(k Kind) RichError {
	r.Kind = k
	return r
}

func (r RichError) WithError(e error) RichError {
	r.WrappedError = e
	return r
}
